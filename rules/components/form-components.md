# 表单组件规范

> 本文件适用于 MOM3.0 Vue3 前端项目。

## 1. 表单设计原则

### 1.1 布局原则

| 场景 | 布局方式 | 说明 |
|------|----------|------|
| 登录/注册 | 单列布局 | 简单直接 |
| 搜索表单 | 行内表单 | 节省空间 |
| 弹窗表单 | 两列栅格 | 充分利用空间 |
| 详情页 | 标签页 | 信息分组 |

### 1.2 表单栅格

```vue
<el-form label-width="100px">
  <el-row :gutter="20">
    <el-col :span="12">
      <el-form-item label="名称">
        <el-input v-model="form.name" />
      </el-form-item>
    </el-col>
    <el-col :span="12">
      <el-form-item label="编码">
        <el-input v-model="form.code" />
      </el-form-item>
    </el-col>
  </el-row>
</el-form>
```

## 2. 输入组件

### 2.1 文本输入

```vue
<el-form-item label="名称" prop="name">
  <el-input
    v-model="form.name"
    placeholder="请输入名称"
    clearable
    :maxlength="50"
    show-word-limit
  />
</el-form-item>
```

### 2.2 数字输入

```vue
<el-form-item label="数量" prop="quantity">
  <el-input-number
    v-model="form.quantity"
    :min="1"
    :max="999999"
    :precision="0"
    controls-position="right"
  />
</el-form-item>
```

### 2.3 下拉选择

```vue
<el-form-item label="状态" prop="status">
  <el-select
    v-model="form.status"
    placeholder="请选择状态"
    clearable
  >
    <el-option label="启用" :value="1" />
    <el-option label="禁用" :value="0" />
  </el-select>
</el-form-item>
```

### 2.4 日期选择

```vue
<!-- 日期 -->
<el-form-item label="日期" prop="date">
  <el-date-picker
    v-model="form.date"
    type="date"
    placeholder="选择日期"
    value-format="YYYY-MM-DD"
  />
</el-form-item>

<!-- 日期时间 -->
<el-form-item label="时间" prop="datetime">
  <el-date-picker
    v-model="form.datetime"
    type="datetime"
    placeholder="选择时间"
    value-format="YYYY-MM-DD HH:mm:ss"
  />
</el-form-item>

<!-- 日期范围 -->
<el-form-item label="日期范围" prop="dateRange">
  <el-date-picker
    v-model="form.dateRange"
    type="daterange"
    start-placeholder="开始日期"
    end-placeholder="结束日期"
    value-format="YYYY-MM-DD"
  />
</el-form-item>
```

### 2.5 树形选择

```vue
<el-form-item label="部门" prop="deptId">
  <el-tree-select
    v-model="form.deptId"
    :data="deptTree"
    :props="{ value: 'id', label: 'name', children: 'children' }"
    placeholder="请选择部门"
    check-strictly
  />
</el-form-item>
```

### 2.6 用户/物料选择器

```vue
<!-- 通用选择器组件 -->
<el-form-item label="负责人" prop="leaderId">
  <UserSelect
    v-model="form.leaderId"
    placeholder="请选择负责人"
  />
</el-form-item>

<el-form-item label="物料" prop="materialId">
  <MaterialSelect
    v-model="form.materialId"
    placeholder="请选择物料"
  />
</el-form-item>

<el-form-item label="设备" prop="equipmentId">
  <EquipmentSelect
    v-model="form.equipmentId"
    placeholder="请选择设备"
  />
</el-form-item>
```

## 3. 表单验证

### 3.1 常用验证规则

```javascript
const rules = {
  // 必填
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' }
  ],

  // 长度限制
  code: [
    { required: true, message: '请输入编码', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],

  // 数字范围
  quantity: [
    { required: true, message: '请输入数量', trigger: 'blur' },
    { type: 'number', min: 1, max: 999999, message: '数量在 1 到 999999 之间', trigger: 'blur' }
  ],

  // 邮箱
  email: [
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],

  // 手机号
  phone: [
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }
  ],

  // 自定义验证
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { validator: validatePassword, trigger: 'blur' }
  ],

  // 异步验证
  email: [
    { validator: validateEmailUnique, trigger: 'blur' }
  ]
}

// 自定义验证函数
const validatePassword = (rule, value, callback) => {
  if (value.length < 8) {
    callback(new Error('密码长度不能少于 8 位'))
  } else if (!/[A-Z]/.test(value)) {
    callback(new Error('密码必须包含大写字母'))
  } else if (!/[a-z]/.test(value)) {
    callback(new Error('密码必须包含小写字母'))
  } else if (!/[0-9]/.test(value)) {
    callback(new Error('密码必须包含数字'))
  } else {
    callback()
  }
}

// 异步验证
const validateEmailUnique = async (rule, value, callback) => {
  if (!value) {
    callback()
    return
  }
  const exists = await checkEmailExists(value)
  if (exists) {
    callback(new Error('邮箱已被使用'))
  } else {
    callback()
  }
}
```

### 3.2 动态验证规则

```javascript
const rules = reactive({
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
})

// 根据条件添加规则
watch(() => form.type, (newType) => {
  if (newType === 'special') {
    rules.specialCode = [
      { required: true, message: '请输入特殊编码', trigger: 'blur' }
    ]
  } else {
    delete rules.specialCode
  }
})
```

## 4. 表单操作

### 4.1 弹窗表单

```vue
<template>
  <el-dialog
    v-model="visible"
    :title="title"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item label="名称" prop="name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-switch v-model="form.status" />
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" :loading="loading" @click="handleSubmit">
        确定
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  data: { type: Object, default: () => ({}) }
})

const emit = defineEmits(['update:modelValue', 'success'])

// 对话框状态
const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const title = computed(() => props.data?.id ? '编辑' : '新增')

// 表单数据
const formRef = ref()
const loading = ref(false)
const form = reactive({
  id: '',
  name: '',
  status: true
})

// 表单规则
const rules = {
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
}

// 初始化
watch(() => props.data, (val) => {
  if (val) {
    Object.assign(form, val)
  }
}, { immediate: true })

// 提交
const handleSubmit = async () => {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    await saveData(form)
    ElMessage.success('保存成功')
    emit('success')
    handleClose()
  } finally {
    loading.value = false
  }
}

// 关闭
const handleClose = () => {
  formRef.value?.resetFields()
  visible.value = false
}
</script>
```

### 4.2 抽屉表单

```vue
<template>
  <el-drawer
    v-model="visible"
    :title="title"
    size="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
      <!-- 表单项 -->
    </el-form>

    <template #footer>
      <div class="drawer-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" :loading="loading" @click="handleSubmit">
          确定
        </el-button>
      </div>
    </template>
  </el-drawer>
</template>
```

## 5. 特殊表单组件

### 5.1 富文本编辑器

```vue
<el-form-item label="详细描述" prop="description">
  <RichEditor v-model="form.description" :height="300" />
</el-form-item>
```

### 5.2 文件上传

```vue
<el-form-item label="附件" prop="attachments">
  <UploadFile
    v-model="form.attachments"
    :limit="5"
    :accept="['.pdf', '.doc', '.docx', '.xls', '.xlsx']"
    :max-size="10"
    list-type="picture-card"
  />
</el-form-item>
```

### 5.3 代码编辑器

```vue
<el-form-item label="SQL" prop="sql">
  <CodeEditor
    v-model="form.sql"
    language="sql"
    :height="200"
  />
</el-form-item>
```

### 5.4 JSON 编辑器

```vue
<el-form-item label="配置" prop="config">
  <JsonEditor v-model="form.config" />
</el-form-item>
```

## 6. 表单布局模式

### 6.1 基础两列布局

```vue
<template>
  <el-form :model="form" label-width="120px">
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="字段1">
          <el-input v-model="form.field1" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="字段2">
          <el-input v-model="form.field2" />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
```

### 6.2 标签页布局

```vue
<el-tabs v-model="activeTab">
  <el-tab-pane label="基本信息" name="basic">
    <!-- 基本信息表单项 -->
  </el-tab-pane>
  <el-tab-pane label="扩展信息" name="extend">
    <!-- 扩展信息表单项 -->
  </el-tab-pane>
  <el-tab-pane label="其他" name="other">
    <!-- 其他表单项 -->
  </el-tab-pane>
</el-tabs>
```

### 6.3 分组布局

```vue
<el-form>
  <div class="form-section">
    <div class="section-title">基本信息</div>
    <!-- 基本信息表单项 -->
  </div>

  <el-divider />

  <div class="form-section">
    <div class="section-title">扩展信息</div>
    <!-- 扩展信息表单项 -->
  </div>
</el-form>

<style scoped>
.form-section {
  margin-bottom: 16px;
}
.section-title {
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 16px;
  color: #303133;
}
</style>
```

## 7. 表单辅助函数

### 7.1 表单重置

```javascript
const handleReset = () => {
  formRef.value?.resetFields()
  // 手动清除无法自动清除的字段
  form.field = null
  form.files = []
}
```

### 7.2 表单克隆

```javascript
const handleCopy = () => {
  const clonedForm = JSON.parse(JSON.stringify(form))
  // 清除 ID 等唯一字段
  clonedForm.id = ''
  clonedForm.code = ''
  Object.assign(form, clonedForm)
}
```

### 7.3 表单回显

```vue
<script setup>
// 编辑时加载数据
const handleEdit = async (row) => {
  const res = await getById(row.id)
  Object.assign(form, res.data)
  visible.value = true
}
</script>
```

## 8. 无障碍设计

### 8.1 必填标记

```vue
<el-form-item label="名称" prop="name">
  <template #label>
    <span>
      名称
      <span class="required-mark">*</span>
    </span>
  </template>
  <el-input v-model="form.name" />
</el-form-item>

<style scoped>
.required-mark {
  color: #F56C6C;
  margin-left: 2px;
}
</style>
```

### 8.2 错误提示

```vue
<el-form-item label="名称" prop="name" :error="formErrors.name">
  <el-input v-model="form.name" />
</el-form-item>
```
