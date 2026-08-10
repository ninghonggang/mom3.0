#!/usr/bin/env bash
#
# Chained end-to-end smoke test for the MOM gateway write paths.
#
# Design notes
#   * Every run stamps its own RUN_ID, so business codes are unique and the
#     script is re-runnable without tripping unique constraints.
#   * IDs are threaded from one step to the next (warehouse -> location ->
#     receive order -> ...), instead of assuming "1". A failure therefore means
#     a real bug, not stale fixture data.
#   * Prerequisites the domain services legitimately require (a storage bin, an
#     alert config, a data point definition) are created up front through the
#     public API rather than seeded behind the gateway's back.
#   * State machines are driven through their legal transitions
#     (repair: REPORTED -> ASSIGNED -> IN_PROGRESS -> COMPLETED -> VERIFIED).
#
# Usage: ./scripts/smoke_write.sh [base_url]
#
set -uo pipefail

BASE=${1:-http://localhost:3000/api}
RUN_ID=$(date +%y%m%d%H%M%S)
PASS=0
FAIL=0
FAILED_STEPS=()

# Last response body, set by req()
BODY=""

c_ok=$'\033[32m'; c_bad=$'\033[31m'; c_dim=$'\033[2m'; c_hdr=$'\033[1;36m'; c_off=$'\033[0m'

section() { printf '\n%s========== %s ==========%s\n' "$c_hdr" "$1" "$c_off"; }

# req METHOD PATH [JSON_BODY]
# Populates $BODY, updates counters, returns 0 on 2xx.
req() {
  local method=$1 path=$2 data=${3:-}
  local out code
  if [ -n "$data" ]; then
    out=$(curl -s -m 20 -w $'\n%{http_code}' -X "$method" "$BASE$path" \
      -H 'Content-Type: application/json' -d "$data")
  else
    out=$(curl -s -m 20 -w $'\n%{http_code}' -X "$method" "$BASE$path")
  fi
  code=$(printf '%s' "$out" | tail -n1)
  BODY=$(printf '%s' "$out" | sed '$d')

  case "$code" in
    200|201)
      PASS=$((PASS + 1))
      printf '%sOK  %s[%s] %-6s %s\n' "$c_ok" "$c_off" "$code" "$method" "$path"
      return 0
      ;;
    *)
      FAIL=$((FAIL + 1))
      FAILED_STEPS+=("$method $path [$code]")
      printf '%sFAIL%s [%s] %-6s %s\n     %s%s%s\n' \
        "$c_bad" "$c_off" "$code" "$method" "$path" \
        "$c_dim" "$(printf '%s' "$BODY" | head -c 240)" "$c_off"
      return 1
      ;;
  esac
}

# pick JQ_FILTER [FALLBACK] -- read a value out of the last response
pick() {
  local val
  val=$(printf '%s' "$BODY" | jq -r "$1 // empty" 2>/dev/null)
  if [ -z "$val" ] || [ "$val" = "null" ]; then
    printf '%s' "${2:-}"
  else
    printf '%s' "$val"
  fi
}

echo "MOM gateway smoke test  base=$BASE  run=$RUN_ID"

# ---------------------------------------------------------------- MDM ------
# Everything downstream references a material, so create a fresh one and use
# its code throughout. This also removes the dependency on pre-seeded data.
section "MDM 主数据"
MAT_CODE="SMK-$RUN_ID"
req POST /mdm/materials "{\"materialCode\":\"$MAT_CODE\",\"name\":\"冒烟测试物料\",\"specification\":\"Q235\",\"unit\":\"kg\",\"category\":\"RAW\",\"materialType\":\"raw\"}"
MAT_ID=$(pick '.material.base.id // .base.id // .id')
echo "     material: $MAT_CODE (id=$MAT_ID)"

# ---------------------------------------------------------------- WMS ------
section "WMS 仓储 (入库→上架→出库→拣货→发运→盘点)"
WH_CODE="WH-$RUN_ID"
req POST /wms/warehouses "{\"warehouseCode\":\"$WH_CODE\",\"name\":\"冒烟仓库\",\"type\":\"RAW\",\"address\":\"1号厂房\"}"
WH_ID=$(pick '.id' '0')
echo "     warehouse: $WH_CODE (id=$WH_ID)"

LOC_CODE="LOC-$RUN_ID"
req POST /wms/locations "{\"warehouseId\":\"$WH_ID\",\"locationCode\":\"$LOC_CODE\",\"locationType\":\"STORAGE\",\"capacity\":1000}"
LOC_ID=$(pick '.id' '0')
echo "     location: $LOC_CODE (id=$LOC_ID)"

# unitPrice 必须带上：上架时它是库存台账 unit_cost 的唯一来源（移动加权平均），
# 缺失会让看板 inventoryValue 静默停留在 0。
req POST /wms/receive-orders "{\"receiveNo\":\"RO-$RUN_ID\",\"warehouseId\":\"$WH_ID\",\"supplierId\":\"1\",\"poNo\":\"PO-$RUN_ID\",\"lines\":[{\"materialCode\":\"$MAT_CODE\",\"quantity\":100,\"unit\":\"kg\",\"batchNo\":\"B-$RUN_ID\",\"unitPrice\":15.5}]}"
RO_ID=$(pick '.id' '0')
echo "     receive order id=$RO_ID"

req POST "/wms/receive-orders/$RO_ID/confirm" '{"remark":"全量收货"}'
req POST "/wms/receive-orders/$RO_ID/putaway" "{\"locationCode\":\"$LOC_CODE\",\"operatorId\":\"1\"}"

req POST /wms/delivery-orders "{\"deliveryNo\":\"DO-$RUN_ID\",\"warehouseId\":\"$WH_ID\",\"customerId\":\"1\",\"soNo\":\"SO-$RUN_ID\",\"lines\":[{\"materialCode\":\"$MAT_CODE\",\"quantity\":10,\"unit\":\"kg\",\"batchNo\":\"B-$RUN_ID\"}]}"
DO_ID=$(pick '.id' '0')
echo "     delivery order id=$DO_ID"

req POST "/wms/delivery-orders/$DO_ID/pick" "{\"locationCodes\":[\"$LOC_CODE\"],\"operatorId\":\"1\"}"
req POST "/wms/delivery-orders/$DO_ID/ship" '{"carrier":"顺丰","trackingNo":"SF'"$RUN_ID"'","operatorId":"1"}'

req POST /wms/count-plans "{\"planNo\":\"CP-$RUN_ID\",\"warehouseId\":\"$WH_ID\",\"countType\":\"FULL\",\"plannedDate\":\"2026-08-10T08:00:00Z\"}"
CP_ID=$(pick '.id' '0')
req POST /wms/count-records "{\"planId\":\"$CP_ID\",\"locationCode\":\"$LOC_CODE\",\"materialCode\":\"$MAT_CODE\",\"countedQuantity\":88,\"operatorId\":\"1\"}"

# ---------------------------------------------------------------- QMS ------
section "QMS 质量 (检验特性→检验单→录结果→判退→NCR→处置)"
# 检验结果必须挂在已存在的检验特性上（USL/LSL 决定自动判定）。
req POST /qms/characteristics "{\"charCode\":\"CH-$RUN_ID\",\"charName\":\"外径\",\"dataType\":\"NUMBER\",\"usl\":\"10.05\",\"lsl\":\"9.95\",\"target\":\"10.00\",\"unit\":\"mm\"}"
CH_ID=$(pick '.id' '0')
echo "     characteristic id=$CH_ID"

req POST /qms/inspection-sheets "{\"sheetNo\":\"IS-$RUN_ID\",\"inspectionType\":\"IQC\",\"orderId\":\"1\",\"materialCode\":\"$MAT_CODE\",\"sampleSize\":5,\"inspectorId\":\"1\"}"
IS_ID=$(pick '.id' '0')
echo "     inspection sheet id=$IS_ID"

# 录一条超差值 -> 检验单自动进入 IN_PROGRESS
req POST "/qms/inspection-sheets/$IS_ID/results" "{\"results\":[{\"itemCode\":\"$CH_ID\",\"result\":\"FAIL\",\"measuredValue\":\"10.42\",\"remark\":\"超上差\"}],\"inspectorId\":\"1\"}"

# NCR 只能建立在 FAILED 的检验单上，先做判退（PENDING→IN_PROGRESS→FAILED）
req PATCH "/qms/inspection-sheets/$IS_ID" '{"status":"FAILED","defectCount":2,"inspectorId":"1"}'

req POST /qms/ncrs "{\"ncrNo\":\"NCR-$RUN_ID\",\"sheetId\":\"$IS_ID\",\"materialCode\":\"$MAT_CODE\",\"quantity\":2,\"description\":\"尺寸超差\",\"defectCategory\":\"DIM\",\"severity\":\"MAJOR\"}"
NCR_ID=$(pick '.id' '0')
echo "     ncr id=$NCR_ID"

req POST "/qms/ncrs/$NCR_ID/actions" '{"actionType":"rework","description":"返修处理","handlerId":"1"}'
# NCR 闭环需逐级流转：OPEN → INVESTIGATING → DISPOSITIONED → VERIFIED → CLOSED
req PATCH "/qms/ncrs/$NCR_ID" '{"status":"INVESTIGATING"}'
req PATCH "/qms/ncrs/$NCR_ID" '{"status":"DISPOSITIONED"}'
req PATCH "/qms/ncrs/$NCR_ID" '{"status":"VERIFIED"}'
req PATCH "/qms/ncrs/$NCR_ID" '{"status":"CLOSED","description":"已闭环"}'

# ---------------------------------------------------------------- EAM ------
section "EAM 设备 (台账→报修→派工→完修→停机)"
EQ_CODE="EQ-$RUN_ID"
req POST /eam/equipment "{\"equipmentCode\":\"$EQ_CODE\",\"name\":\"注塑机\",\"category\":\"MACHINE\",\"workshopId\":\"1\",\"lineId\":\"1\",\"model\":\"HT-200\",\"targetOee\":\"0.85\"}"
EQ_ID=$(pick '.id' '0')
echo "     equipment: $EQ_CODE (id=$EQ_ID)"

req POST /eam/repair-orders "{\"equipmentId\":\"$EQ_ID\",\"faultDescription\":\"液压异常\",\"faultCategory\":\"HYDRAULIC\",\"reporterId\":\"1\",\"priority\":\"URGENT\"}"
RP_ID=$(pick '.id' '0')
echo "     repair order id=$RP_ID"

# Walk the state machine in order — jumping straight to IN_PROGRESS is a
# legitimate 409, so the script must assign first.
req PATCH "/eam/repair-orders/$RP_ID" '{"status":"ASSIGNED","technicianId":"1"}'
req PATCH "/eam/repair-orders/$RP_ID" '{"status":"IN_PROGRESS","repairDescription":"更换密封件","technicianId":"1"}'
req PATCH "/eam/repair-orders/$RP_ID" '{"status":"COMPLETED","repairDescription":"已更换密封件并试机","technicianId":"1"}'
req PATCH "/eam/repair-orders/$RP_ID" '{"status":"VERIFIED","technicianId":"1"}'

req POST /eam/downtime/start "{\"equipmentId\":\"$EQ_ID\",\"reason\":\"故障停机\",\"type\":\"UNPLANNED\",\"reporterId\":\"1\"}"
DT_ID=$(pick '.id' '0')
req POST "/eam/downtime/$DT_ID/resolve" '{"resolution":"已修复","resolverId":"1"}'

# OEE 上报：服务端按 A×P×Q 计算，(设备,日期) 唯一，重复上报走 upsert。
# 连报两次是刻意的——第二次必须仍返回 200/201 且覆盖旧值，用来守住幂等语义。
TODAY=$(date +%F)
req POST /eam/oee "{\"equipmentId\":\"$EQ_ID\",\"calcDate\":\"$TODAY\",\"availability\":0.92,\"performance\":0.88,\"quality\":0.97}"
req POST /eam/oee "{\"equipmentId\":\"$EQ_ID\",\"calcDate\":\"$TODAY\",\"availability\":0.94,\"performance\":0.89,\"quality\":0.98}"

# -------------------------------------------------------------- TRACE ------
section "TRACE 追溯 (记录→序列号→数据点→采集→扫码)"
SN="SN-$RUN_ID"
req POST /trace/records "{\"serialNo\":\"$SN\",\"productCode\":\"$MAT_CODE\",\"orderId\":\"1\",\"workstationId\":\"1\",\"batchNo\":\"B-$RUN_ID\"}"
req POST /trace/serials/generate "{\"productCode\":\"$MAT_CODE\",\"count\":3,\"prefix\":\"SN$RUN_ID\"}"

# A data point is a *definition*; collection references it by id.
req POST /trace/data-points "{\"pointCode\":\"TEMP-$RUN_ID\",\"pointName\":\"炉温\",\"equipmentId\":\"$EQ_ID\",\"dataType\":\"NUMBER\",\"upperLimit\":\"120\",\"lowerLimit\":\"20\",\"collectIntervalSeconds\":30}"
DP_ID=$(pick '.id' '0')
echo "     data point id=$DP_ID"

req POST /trace/collect "{\"dataPointId\":\"$DP_ID\",\"value\":\"85.5\",\"quality\":\"GOOD\"}"
req POST /trace/scan-logs "{\"serialNo\":\"$SN\",\"scanPoint\":\"WS-1\",\"operatorId\":\"1\",\"result\":\"OK\"}"

# -------------------------------------------------------------- ANDON ------
section "ANDON 安灯 (呼叫→响应→关闭 / 告警配置→触发)"
req POST /andon/calls '{"callPoint":"WS-1","workstationId":"1","type":"QUALITY","description":"质量异常","callerId":"1","priority":"HIGH"}'
CALL_ID=$(pick '.id' '0')
echo "     andon call id=$CALL_ID"

req POST "/andon/calls/$CALL_ID/acknowledge" '{"responderId":"1","remark":"处理中"}'
req POST "/andon/calls/$CALL_ID/resolve" '{"resolution":"已解决","resolverId":"1"}'

# Alerts are rule-driven: the config must exist before one can fire.
req POST /andon/alert-configs "{\"configCode\":\"CFG-$RUN_ID\",\"configName\":\"温度过高\",\"triggerType\":\"THRESHOLD\",\"severity\":\"P1\",\"triggerCondition\":\"value > 100\",\"notifyChannels\":\"wecom,sms\"}"
CFG_ID=$(pick '.id' '0')
echo "     alert config id=$CFG_ID"

req POST /andon/alerts "{\"configId\":\"$CFG_ID\",\"targetId\":\"$EQ_ID\",\"targetType\":\"EQUIPMENT\"}"

# ---------------------------------------------------------------- APS ------
section "APS 计划 (MPS→MRP→排产)"
req POST /aps/mps-plans "{\"planNo\":\"MPS-$RUN_ID\",\"productCode\":\"$MAT_CODE\",\"plannedQuantity\":200,\"planMonth\":\"2026-08\",\"remark\":\"冒烟\"}"
MPS_ID=$(pick '.mps.base.id // .mps.id // .id' '0')
echo "     mps id=$MPS_ID"

req POST /aps/mrp/generate "{\"mpsId\":\"$MPS_ID\",\"operator\":\"smoke\"}"
req POST /aps/schedule-jobs "{\"planNo\":\"SCH-$RUN_ID\",\"mpsId\":\"$MPS_ID\",\"planType\":\"FINE\",\"algorithm\":\"EDD\",\"workcenterId\":\"1\",\"plannedStart\":\"2026-08-10T08:00:00Z\",\"plannedEnd\":\"2026-08-20T18:00:00Z\"}"

# ---------------------------------------------------------------- MES ------
section "MES 执行 (工单→派工→报工→完工)"
req POST /mes/orders "{\"orderNo\":\"MO-$RUN_ID\",\"productCode\":\"$MAT_CODE\",\"plannedQuantity\":100,\"workshopId\":\"1\",\"lineId\":\"2\",\"workstationId\":\"1\",\"plannedStartTime\":\"2026-08-10T08:00:00Z\",\"plannedEndTime\":\"2026-08-12T18:00:00Z\"}"
MO_ID=$(pick '.order.base.id // .base.id // .id' '0')
echo "     production order id=$MO_ID"

req POST "/mes/orders/$MO_ID/dispatch" '{"workshopId":"1","lineId":"2","workstationId":"1","operatorId":"1","remark":"派工"}'
req POST "/mes/orders/$MO_ID/report" '{"goodQuantity":98,"scrapQuantity":2,"operatorId":"1","workstationId":"1","shift":"DAY","remark":"报工"}'
req POST "/mes/orders/$MO_ID/complete" "{\"actualQuantity\":98,\"warehouseId\":\"$WH_ID\",\"locationCode\":\"$LOC_CODE\",\"batchNo\":\"B-$RUN_ID\",\"actualEndTime\":\"2026-08-12T18:00:00Z\",\"remark\":\"完工入库\"}"

# --------------------------------------------------------------- 汇总 ------
TOTAL=$((PASS + FAIL))
printf '\n%s===================================%s\n' "$c_hdr" "$c_off"
printf 'PASS=%s%d%s  FAIL=%s%d%s  TOTAL=%d\n' "$c_ok" "$PASS" "$c_off" "$c_bad" "$FAIL" "$c_off" "$TOTAL"
if [ "$FAIL" -gt 0 ]; then
  printf '\n失败步骤:\n'
  for s in "${FAILED_STEPS[@]}"; do printf '  - %s\n' "$s"; done
  exit 1
fi
printf '%s全部通过%s\n' "$c_ok" "$c_off"
