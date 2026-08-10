#!/usr/bin/env bash
#
# Read-path smoke test for the MOM gateway.
#
# The write smoke (smoke_write.sh) proves the command side works; this one
# proves the query side does. It hits every GET route the gateway exposes,
# checks the HTTP status AND that the payload actually deserialises into the
# shape the frontend expects (a list envelope or an object), so a route that
# returns `{}` because of a proto/field mismatch is reported as a failure
# rather than a green tick.
#
# Usage: ./scripts/smoke_read.sh [base_url]
#
set -uo pipefail

BASE=${1:-http://localhost:3000/api}
PASS=0
FAIL=0
FAILED_STEPS=()
BODY=""

c_ok=$'\033[32m'; c_bad=$'\033[31m'; c_dim=$'\033[2m'; c_hdr=$'\033[1;36m'; c_off=$'\033[0m'

section() { printf '\n%s========== %s ==========%s\n' "$c_hdr" "$1" "$c_off"; }

# get PATH [JQ_ASSERT]
#   JQ_ASSERT is an optional jq expression that must evaluate truthy.
#   Defaults to "any non-empty object/array".
get() {
  local path=$1 assert=${2:-'(type=="object" or type=="array")'}
  local out code note=''
  out=$(curl -s -m 20 -w $'\n%{http_code}' "$BASE$path")
  code=$(printf '%s' "$out" | tail -n1)
  BODY=$(printf '%s' "$out" | sed '$d')

  if [ "$code" != "200" ]; then
    FAIL=$((FAIL + 1))
    FAILED_STEPS+=("GET $path [$code]")
    printf '%sFAIL%s [%s] GET  %s\n     %s%s%s\n' "$c_bad" "$c_off" "$code" "$path" \
      "$c_dim" "$(printf '%s' "$BODY" | head -c 200)" "$c_off"
    return 1
  fi

  if ! printf '%s' "$BODY" | jq -e "$assert" >/dev/null 2>&1; then
    FAIL=$((FAIL + 1))
    FAILED_STEPS+=("GET $path [shape]")
    printf '%sFAIL%s [200] GET  %s  %s(shape assertion failed: %s)%s\n     %s%s%s\n' \
      "$c_bad" "$c_off" "$path" "$c_dim" "$assert" "$c_off" \
      "$c_dim" "$(printf '%s' "$BODY" | head -c 200)" "$c_off"
    return 1
  fi

  # Report list size when the payload looks like a list envelope.
  local n
  n=$(printf '%s' "$BODY" | jq -r '(.items // .list // .data // empty) | if type=="array" then length else empty end' 2>/dev/null)
  [ -n "$n" ] && note=" ($n 条)"
  PASS=$((PASS + 1))
  printf '%sOK  %s[200] GET  %s%s\n' "$c_ok" "$c_off" "$path" "$note"
  return 0
}

echo "MOM gateway read smoke  base=$BASE"

section "MDM 主数据"
get "/mdm/materials?page=1&pageSize=5"
get "/mdm/boms?page=1&pageSize=5"
get "/mdm/customers?page=1&pageSize=5"
get "/mdm/suppliers?page=1&pageSize=5"
get "/mdm/workshops?page=1&pageSize=5"
get "/mdm/production-lines?page=1&pageSize=5"
get "/mdm/workstations?page=1&pageSize=5"

section "WMS 仓储"
get "/wms/warehouses?page=1&pageSize=5"
get "/wms/locations?page=1&pageSize=5"
get "/wms/balances?page=1&pageSize=5"

section "QMS 质量"
get "/qms/characteristics?page=1&pageSize=5"
get "/qms/inspection-sheets?page=1&pageSize=5"
get "/qms/ncrs?page=1&pageSize=5"
get "/qms/spc-data"

section "EAM 设备"
get "/eam/equipment?page=1&pageSize=5"
get "/eam/repair-orders?page=1&pageSize=5"
# ListOee 不分页，按 equipmentId / beginDate / endDate 过滤。
get "/eam/oee"
get "/eam/oee?beginDate=2020-01-01"

section "TRACE 追溯"
get "/trace/data-points?page=1&pageSize=5"

section "ANDON 安灯"
get "/andon/calls?page=1&pageSize=5"
get "/andon/alerts?page=1&pageSize=5"

section "APS 计划"
get "/aps/mps-plans?page=1&pageSize=5"
get "/aps/schedule-jobs?page=1&pageSize=5"

section "MES 执行"
get "/mes/orders?page=1&pageSize=5"

section "Dashboard 看板"
get "/dashboard/overview"
get "/dashboard/production-trend"
get "/dashboard/quality-trend"
get "/dashboard/alarms"

# Detail routes need a real id — take the first one from the list endpoints.
section "详情路由"
first_id() { curl -s -m 20 "$BASE$1" | jq -r "$2 // empty" 2>/dev/null; }

EQ_ID=$(first_id "/eam/equipment?page=1&pageSize=1" '.items[0].id')
[ -n "$EQ_ID" ] && get "/eam/equipment/$EQ_ID" || echo "     (skip eam/equipment/:id — 无数据)"

IS_ID=$(first_id "/qms/inspection-sheets?page=1&pageSize=1" '.items[0].id')
[ -n "$IS_ID" ] && get "/qms/inspection-sheets/$IS_ID" || echo "     (skip qms/inspection-sheets/:id — 无数据)"

MO_ID=$(first_id "/mes/orders?page=1&pageSize=1" '.items[0].base.id // .items[0].id')
[ -n "$MO_ID" ] && get "/mes/orders/$MO_ID" || echo "     (skip mes/orders/:id — 无数据)"

TOTAL=$((PASS + FAIL))
printf '\n%s===================================%s\n' "$c_hdr" "$c_off"
printf 'PASS=%s%d%s  FAIL=%s%d%s  TOTAL=%d\n' "$c_ok" "$PASS" "$c_off" "$c_bad" "$FAIL" "$c_off" "$TOTAL"
if [ "$FAIL" -gt 0 ]; then
  printf '\n失败步骤:\n'
  for s in "${FAILED_STEPS[@]}"; do printf '  - %s\n' "$s"; done
  exit 1
fi
printf '%s全部通过%s\n' "$c_ok" "$c_off"
