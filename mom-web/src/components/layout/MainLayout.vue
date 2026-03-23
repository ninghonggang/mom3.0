<template>
  <el-container class="main-layout">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <div class="logo">
        <img v-if="!isCollapse" src="/favicon.svg" alt="logo" class="logo-img">
        <span v-if="!isCollapse" class="logo-text">MOM3.0</span>
      </div>

      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :router="true"
        class="sidebar-menu"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
      >
        <el-menu-item index="/dashboard">
          <el-icon><House /></el-icon>
          <template #title>首页</template>
        </el-menu-item>

        <el-sub-menu index="/system">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="/system/user">用户管理</el-menu-item>
          <el-menu-item index="/system/role">角色管理</el-menu-item>
          <el-menu-item index="/system/menu">菜单管理</el-menu-item>
          <el-menu-item index="/system/dept">部门管理</el-menu-item>
          <el-menu-item index="/system/dict">字典管理</el-menu-item>
          <el-menu-item index="/system/post">岗位管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/mdm">
          <template #title>
            <el-icon><Box /></el-icon>
            <span>主数据</span>
          </template>
          <el-menu-item index="/mdm/material">物料管理</el-menu-item>
          <el-menu-item index="/mdm/workshop">车间管理</el-menu-item>
          <el-menu-item index="/mdm/line">生产线管理</el-menu-item>
          <el-menu-item index="/mdm/workstation">工位管理</el-menu-item>
          <el-menu-item index="/mdm/shift">班次管理</el-menu-item>
          <el-menu-item index="/mdm/bom">BOM管理</el-menu-item>
          <el-menu-item index="/mdm/operation">工序管理</el-menu-item>
          <el-menu-item index="/mdm/mdm-shift">班次定义</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/production">
          <template #title>
            <el-icon><List /></el-icon>
            <span>生产执行</span>
          </template>
          <el-menu-item index="/production/order">生产工单</el-menu-item>
          <el-menu-item index="/production/sales-order">销售订单</el-menu-item>
          <el-menu-item index="/production/report">生产报工</el-menu-item>
          <el-menu-item index="/production/dispatch">派工</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/equipment">
          <template #title>
            <el-icon><Monitor /></el-icon>
            <span>设备管理</span>
          </template>
          <el-menu-item index="/equipment">设备台账</el-menu-item>
          <el-menu-item index="/equipment/check">设备点检</el-menu-item>
          <el-menu-item index="/equipment/maintenance">设备保养</el-menu-item>
          <el-menu-item index="/equipment/repair">设备维修</el-menu-item>
          <el-menu-item index="/equipment/spare">备件管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/wms">
          <template #title>
            <el-icon><House /></el-icon>
            <span>仓储管理</span>
          </template>
          <el-menu-item index="/wms/warehouse">仓库管理</el-menu-item>
          <el-menu-item index="/wms/location">库位管理</el-menu-item>
          <el-menu-item index="/wms/inventory">库存管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/quality">
          <template #title>
            <el-icon><CircleCheck /></el-icon>
            <span>质量管理</span>
          </template>
          <el-menu-item index="/quality/iqc">IQC检验</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/aps">
          <template #title>
            <el-icon><Calendar /></el-icon>
            <span>APS计划</span>
          </template>
          <el-menu-item index="/aps/mps">MPS计划</el-menu-item>
          <el-menu-item index="/aps/mrp">MRP计划</el-menu-item>
          <el-menu-item index="/aps/schedule">排程计划</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="/trace">
          <template #title>
            <el-icon><Search /></el-icon>
            <span>追溯管理</span>
          </template>
          <el-menu-item index="/trace/query">追溯查询</el-menu-item>
          <el-menu-item index="/trace/andon">安东呼叫</el-menu-item>
        </el-sub-menu>

        <el-menu-item index="/energy/monitor">
          <el-icon><Lightning /></el-icon>
          <template #title>能源监控</template>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <el-container>
      <!-- 头部 -->
      <el-header class="header">
        <div class="header-left">
          <el-icon class="collapse-icon" @click="isCollapse = !isCollapse">
            <Fold v-if="!isCollapse" />
            <Expand v-else />
          </el-icon>
        </div>

        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" :src="userInfo?.avatar">
                {{ userInfo?.nickname?.charAt(0) || 'U' }}
              </el-avatar>
              <span class="username">{{ userInfo?.nickname || userInfo?.username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人中心</el-dropdown-item>
                <el-dropdown-item command="password">修改密码</el-dropdown-item>
                <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区 -->
      <el-main class="main-content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessageBox } from 'element-plus'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const isCollapse = ref(false)
const userInfo = computed(() => authStore.userInfo)

const activeMenu = computed(() => route.path)

const handleCommand = async (command: string) => {
  switch (command) {
    case 'logout':
      await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
        type: 'warning'
      })
      authStore.logoutAction()
      break
    case 'profile':
      router.push('/profile')
      break
    case 'password':
      router.push('/password')
      break
  }
}
</script>

<style scoped lang="scss">
.main-layout {
  height: 100vh;
}

.sidebar {
  background-color: #304156;
  transition: width 0.3s;

  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 60px;
    background: #2b3a4b;

    .logo-img {
      width: 32px;
      height: 32px;
    }

    .logo-text {
      margin-left: 8px;
      font-size: 18px;
      font-weight: 600;
      color: #fff;
    }
  }

  .sidebar-menu {
    border-right: none;
  }
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  border-bottom: 1px solid #e4e7ed;
  padding: 0 16px;

  .header-left {
    .collapse-icon {
      font-size: 20px;
      cursor: pointer;
      &:hover {
        color: #409eff;
      }
    }
  }

  .header-right {
    .user-info {
      display: flex;
      align-items: center;
      cursor: pointer;

      .username {
        margin: 0 8px;
      }
    }
  }
}

.main-content {
  background: #f5f7fa;
  padding: 16px;
}
</style>
