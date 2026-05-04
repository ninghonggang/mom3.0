<template>
  <div class="module-home agv-home">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Document /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">128</div>
              <div class="stat-label">AGV任务</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Processor /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">15</div>
              <div class="stat-label">AGV设备</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Location /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">42</div>
              <div class="stat-label">库位映射</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><VideoPlay /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">8</div>
              <div class="stat-label">运行中</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-4">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>AGV管理模块说明</span>
          </template>
          <div class="module-intro">
            <p><strong>AGV管理</strong>模块提供以下功能：</p>
            <ul>
              <li>AGV任务 - 管理自动导引车任务调度</li>
              <li>AGV设备 - 维护AGV设备台账</li>
              <li>库位映射 - 管理AGV与仓储库位映射</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { Document, Processor, Location, VideoPlay } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.agv-home {
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
    background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
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
