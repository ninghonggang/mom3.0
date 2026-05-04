<template>
  <div class="module-home mes-home">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><User /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">12</div>
              <div class="stat-label">班组数量</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Guide /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">48</div>
              <div class="stat-label">工艺路线</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Star /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">156</div>
              <div class="stat-label">人员技能</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon"><el-icon><Warning /></el-icon></div>
            <div class="stat-info">
              <div class="stat-value">3</div>
              <div class="stat-label">离线产品</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-4">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>MES作业模块说明</span>
          </template>
          <div class="module-intro">
            <p><strong>MES作业</strong>模块提供以下功能：</p>
            <ul>
              <li>班组管理 - 管理生产班组和人员配置</li>
              <li>工艺路线 - 定义产品生产工序流程</li>
              <li>人员技能 - 维护员工技能矩阵和能力评估</li>
              <li>产品离线 - 处理生产异常离线产品</li>
              <li>物料追溯 - 全流程物料追溯查询</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { User, Guide, Star, Warning } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.mes-home {
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
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
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
