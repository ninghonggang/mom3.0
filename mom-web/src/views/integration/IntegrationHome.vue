<template>
  <div class="module-home integration-home">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Setting /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">12</div>
              <div class="stat-label">接口配置</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><List /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">1,456</div>
              <div class="stat-label">执行日志</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><SuccessFilled /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">98.5%</div>
              <div class="stat-label">成功率</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Clock /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">156ms</div>
              <div class="stat-label">平均耗时</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-4">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>集成管理模块说明</span>
          </template>
          <div class="module-intro">
            <p><strong>集成管理</strong>模块提供以下功能：</p>
            <ul>
              <li>接口配置 - 管理系统间集成接口配置</li>
              <li>执行日志 - 查看接口调用执行日志</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { Setting, List, SuccessFilled, Clock } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.integration-home {
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
    background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
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
