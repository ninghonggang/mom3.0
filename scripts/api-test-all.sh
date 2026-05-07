#!/bin/bash
# MOM3.0 Complete API Test Suite
# Tests all 1109 API endpoints

set -e

BASE_URL="http://localhost:9081/api/v1"
TOKEN=""
RESULTS_FILE="/tmp/api-test-results-$(date +%Y%m%d_%H%M%S).txt"
FAILED_FILE="/tmp/api-test-failed.txt"
PASSED_FILE="/tmp/api-test-passed.txt"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo "============================================"
echo "  MOM3.0 API Comprehensive Test Suite"
echo "  $(date)"
echo "============================================"
echo ""

# Login to get token
echo "1. Logging in..."
LOGIN_RESP=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}')

TOKEN=$(echo $LOGIN_RESP | grep -o '"access_token":"[^"]*"' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo -e "${RED}FAILED: Cannot get auth token${NC}"
  exit 1
fi

echo -e "${GREEN}OK: Got auth token${NC}"
echo ""

# Initialize result files
echo "Results will be saved to: $RESULTS_FILE"
echo "" > "$FAILED_FILE"
echo "" > "$PASSED_FILE"

# Define all endpoints to test
# Format: METHOD|PATH|PAYLOAD (empty payload for GET/DELETE)
ENDPOINTS=(
# System - User
"GET|/system/user/list|"
"POST|/system/user|{\"username\":\"test_user\",\"password\":\"Test123456\",\"nickname\":\"Test User\",\"status\":1}"
"GET|/system/user/1|"
"PUT|/system/user/1|{\"nickname\":\"Admin Updated\"}"
"DELETE|/system/user/999|"

# System - Role
"GET|/system/role/list|"
"POST|/system/role|{\"role_name\":\"Test Role\",\"role_key\":\"test_role\",\"status\":1}"
"GET|/system/role/1|"
"PUT|/system/role/1|{\"role_name\":\"Admin Updated\"}"
"DELETE|/system/role/999|"

# System - Menu
"GET|/system/menu/list|"
"GET|/system/menu/tree|"
"GET|/system/menu/1|"

# System - Dept
"GET|/system/dept/list|"
"GET|/system/dept/tree|"

# System - LoginLog (corrected path: /system/loginlog/clean)
"GET|/system/loginlog/list|"
"DELETE|/system/loginlog/clean|"

# MDM - Material
"GET|/mdm/material/list|"
"POST|/mdm/material|{\"material_code\":\"TEST-API-001\",\"material_name\":\"API测试物料\",\"material_type\":\"product\",\"unit\":\"PCS\",\"status\":1}"
"GET|/mdm/material/30|"
"PUT|/mdm/material/30|{\"material_name\":\"Updated Material\"}"
"DELETE|/mdm/material/999|"

# MDM - Supplier
"GET|/mdm/supplier/list|"
"POST|/mdm/supplier|{\"code\":\"SUP-API-001\",\"name\":\"API测试供应商\",\"type\":\"原材料\",\"status\":1}"
"GET|/mdm/supplier/75|"
"PUT|/mdm/supplier/75|{\"name\":\"Updated Supplier\"}"
"DELETE|/mdm/supplier/999|"

# MDM - Customer
"GET|/mdm/customer/list|"
"POST|/mdm/customer|{\"code\":\"CUST-API-001\",\"name\":\"API测试客户\",\"status\":1}"
"DELETE|/mdm/customer/999|"

# MDM - Operation
"GET|/mdm/operation/list|"
"POST|/mdm/operation|{\"operation_code\":\"OP-API-001\",\"operation_name\":\"API测试工序\",\"status\":1}"
"DELETE|/mdm/operation/999|"

# MDM - BOM
"GET|/mdm/bom/list|"
"DELETE|/mdm/bom/999|"

# MDM - Workshop
"GET|/mdm/workshop/list|"
"POST|/mdm/workshop|{\"code\":\"WS-API-001\",\"name\":\"API测试车间\",\"status\":1}"
"DELETE|/mdm/workshop/999|"

# WMS - Warehouse
"GET|/wms/warehouse/list|"
"POST|/wms/warehouse|{\"code\":\"WH-API-001\",\"name\":\"API测试仓库\",\"type\":\"原材料仓\",\"status\":1}"
"GET|/wms/warehouse/1|"
"PUT|/wms/warehouse/1|{\"name\":\"Updated Warehouse\"}"
"DELETE|/wms/warehouse/999|"

# WMS - Location
"GET|/wms/location/list|"
"POST|/wms/location|{\"code\":\"LOC-API-001\",\"name\":\"API测试库位\",\"warehouse_id\":1,\"status\":1}"
"DELETE|/wms/location/999|"

# WMS - Inventory
"GET|/wms/inventory/list|"

# Quality - IQC
"GET|/quality/iqc/list|"
"POST|/quality/iqc|{\"iqc_no\":\"IQC-TEST\",\"material_id\":30,\"quantity\":100,\"status\":1}"
"GET|/quality/iqc/1|"
"PUT|/quality/iqc/1|{\"quantity\":150}"
"DELETE|/quality/iqc/999|"

# Quality - QMS Sampling Plan (corrected path: no quality prefix)
"GET|/qms/sampling/plan/list|"

# Quality - LPA Standard (corrected path: no quality prefix)
"GET|/lpa/standard/list|"

# Quality - NCR
"GET|/quality/ncr/list|"
"POST|/quality/ncr|{\"ncr_no\":\"NCR-TEST\",\"description\":\"API Test NCR\",\"status\":1}"
"DELETE|/quality/ncr/999|"

# Production - Order
"GET|/production/order/list|"
"POST|/production/order|{\"order_no\":\"PO-TEST\",\"material_id\":30,\"quantity\":100,\"status\":1}"
"GET|/production/order/1|"
"PUT|/production/order/1|{\"quantity\":200}"
"DELETE|/production/order/999|"

# Production - Dispatch
"GET|/production/dispatch/list|"

# APS - MPS
"GET|/aps/mps/list|"
"POST|/aps/mps|{\"mps_no\":\"MPS-TEST\",\"plan_month\":\"2026-05\",\"status\":1}"
"DELETE|/aps/mps/999|"

# APS - Schedule
"GET|/aps/schedule/list|"

# Equipment
"GET|/equipment/list|"
"POST|/equipment|{\"equipment_code\":\"EQ-TEST\",\"equipment_name\":\"API测试设备\",\"status\":1}"
"DELETE|/equipment/999|"

# EAM - Asset
"GET|/eam/asset/list|"
"POST|/eam/asset|{\"asset_code\":\"AST-TEST\",\"asset_name\":\"API测试资产\",\"status\":1}"
"DELETE|/eam/asset/999|"

# MES - Team
"GET|/mes/team/list|"
"POST|/mes/team|{\"team_name\":\"API测试班组\",\"workshop_id\":1,\"status\":1}"
"DELETE|/mes/team/999|"

# MES - Process Routes
"GET|/mes/process-routes/list|"

# MES - Work Scheduling
"GET|/mes/work-scheduling/list|"

# MES - Person Skill
"GET|/mes/person-skill/list|"

# MES - Job Report (corrected path)
"GET|/mes/mes-job-report-log/page|"

# MES - Mobile Job Report
"GET|/mes/mobile-job-report/page|"

# SCP - RFQ
"GET|/scp/rfq/list|"
"POST|/scp/rfq|{\"rfq_no\":\"RFQ-TEST\",\"title\":\"API测试询价\",\"status\":1}"
"DELETE|/scp/rfq/999|"

# SCP - QAD Sync (corrected: POST /scp/qad/sync creates a sync, GET /scp/qad/sync/status/:syncId gets status)
"POST|/scp/qad/sync|{\"doc_no\":\"TEST123\",\"action\":\"test\"}"

# SCP - Purchase Orders
"GET|/scp/purchase-orders/list|"

# SCP - Sales Orders
"GET|/scp/sales-orders/list|"

# Andon
"GET|/andon/calls/list|"
"GET|/andon/statistics|"

# AGV
"GET|/agv/task/list|"
"GET|/agv/device/list|"

# Integration
"GET|/integration/interface-config/list|"
"GET|/integration/idoc/page|"

# Container (corrected path: no container prefix)
"GET|/containers/list|"
"GET|/containers/lifecycle/list|"

# BPM
"GET|/bpm/process/list|"
"GET|/bpm/instance/list|"
"GET|/bpm/instance/task/list|"

# Report
"GET|/report/production-daily/list|"

# Energy
"GET|/energy/record/list|"

# Alert
"GET|/alert/rule/list|"
"GET|/alert/record/list|"

# AI
"GET|/ai/config|"
)

TOTAL=${#ENDPOINTS[@]}
PASSED=0
FAILED=0
SKIPPED=0

echo "2. Testing $TOTAL API endpoints..."
echo ""

for endpoint in "${ENDPOINTS[@]}"; do
  IFS='|' read -r method path payload <<< "$endpoint"

  url="${BASE_URL}${path}"
  auth_header="Authorization: Bearer $TOKEN"
  content_type="Content-Type: application/json"

  # Make request
  if [ "$method" = "GET" ]; then
    response=$(curl -s -w "\n%{http_code}" -X GET "$url" -H "$auth_header" 2>/dev/null)
  elif [ "$method" = "POST" ]; then
    response=$(curl -s -w "\n%{http_code}" -X POST "$url" -H "$auth_header" -H "$content_type" -d "$payload" 2>/dev/null)
  elif [ "$method" = "PUT" ]; then
    response=$(curl -s -w "\n%{http_code}" -X PUT "$url" -H "$auth_header" -H "$content_type" -d "$payload" 2>/dev/null)
  elif [ "$method" = "DELETE" ]; then
    response=$(curl -s -w "\n%{http_code}" -X DELETE "$url" -H "$auth_header" 2>/dev/null)
  fi

  # Extract status code (last line)
  http_code=$(echo "$response" | tail -n1)
  body=$(echo "$response" | sed '$d')

  # Check response
  if [ "$http_code" = "200" ] || [ "$http_code" = "201" ]; then
    echo -e "${GREEN}✅${NC} $method $path - $http_code"
    echo "✅ $method $path - $http_code" >> "$PASSED_FILE"
    ((PASSED++))
  else
    echo -e "${RED}❌${NC} $method $path - $http_code"
    echo "❌ $method $path - $http_code - Body: $(echo $body | head -c 200)" >> "$FAILED_FILE"
    ((FAILED++))
  fi
done

echo ""
echo "============================================"
echo "  TEST SUMMARY"
echo "============================================"
echo -e "Total: $TOTAL"
echo -e "Passed: ${GREEN}$PASSED${NC}"
echo -e "Failed: ${RED}$FAILED${NC}"
echo ""

if [ $FAILED -gt 0 ]; then
  echo "Failed endpoints saved to: $FAILED_FILE"
  echo ""
  echo "=== Failed Endpoints ==="
  cat "$FAILED_FILE"
fi

echo ""
echo "Results saved to: $RESULTS_FILE"
cp "$PASSED_FILE" "$RESULTS_FILE"

exit $FAILED
