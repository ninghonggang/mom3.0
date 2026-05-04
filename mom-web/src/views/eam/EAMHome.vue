<template>
  <div class="module-home eam-home">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><OfficeBuilding /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">5</div>
              <div class="stat-label">厂区数量</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Grid /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">128</div>
              <div class="stat-label">设备层级</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Switch /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">23</div>
              <div class="stat-label">停机记录</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Calendar /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">365</div>
              <div class="stat-label">工厂日历</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-4">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>EAM资产管理模块说明</span>
          </template>
          <div class="module-intro">
            <p><strong>EAM资产管理</strong>模块提供以下功能：</p>
            <ul>
              <li>厂区管理 - 管理工厂厂区和建筑设施</li>
              <li>设备层级 - 维护设备层级结构树</li>
              <li>停机记录 - 记录设备异常停机事件</li>
              <li>工厂日历 - 定义工作日历和节假日</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { OfficeBuilding, Grid, Switch, Calendar } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.eam-home {
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
    background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
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
