<template>
  <div class="module-home fin-home">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Money /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">18</div>
              <div class="stat-label">待付款申请</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><ShoppingCart /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">¥1.2M</div>
              <div class="stat-label">采购结算</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Sell /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">¥2.8M</div>
              <div class="stat-label">销售结算</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Document /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">45</div>
              <div class="stat-label">本月付款</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-4">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>财务管理模块说明</span>
          </template>
          <div class="module-intro">
            <p><strong>财务管理</strong>模块提供以下功能：</p>
            <ul>
              <li>付款申请 - 处理采购付款申请流程</li>
              <li>采购结算 - 管理供应商采购结算</li>
              <li>销售结算 - 处理客户销售结算</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { Money, ShoppingCart, Sell, Document } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.fin-home {
  padding: 20px;
}
.stat-card {
  .stat-content {
    display: flex;
    align-items: center;
    gap: 16px;
  }
  .stat-icon {
    width: 60px;
    height: 60px;
    border-radius: 8px;
    background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    .el-icon {
      font-size: 28px;
      color: #fff;
    }
  }
  .stat-info {
    .stat-value {
      font-size: 28px;
      font-weight: 600;
      color: #303133;
    }
    .stat-label {
      font-size: 14px;
      color: #909399;
    }
  }
}
.mt-4 {
  margin-top: 16px;
}
.module-intro {
  ul {
    margin: 12px 0;
    padding-left: 20px;
    li {
      margin: 8px 0;
      color: #606266;
    }
  }
}
</style>
