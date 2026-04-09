<template>
  <div class="ai-config">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>AI 助手配置</span>
          <el-button type="primary" @click="handleTest" :loading="testing" :disabled="!form.endpoint">
            测试连接
          </el-button>
        </div>
      </template>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="140px"
      >
        <el-divider content-position="left">基本设置</el-divider>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="配置名称">
              <el-input v-model="form.config_name" placeholder="例如：我的OpenAI配置" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="启用状态">
              <el-switch v-model="form.enable" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">AI模型设置</el-divider>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="AI提供商" prop="provider">
              <el-select v-model="form.provider" placeholder="选择AI提供商" style="width: 100%">
                <el-option value="openai" label="OpenAI" />
                <el-option value="azure" label="Azure OpenAI" />
                <el-option value="ollama" label="Ollama (本地)" />
                <el-option value="custom" label="自定义" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="模型名称" prop="model_name">
              <el-input v-model="form.model_name" placeholder="例如：gpt-4o" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="API Endpoint" prop="endpoint">
          <el-input
            v-model="form.endpoint"
            placeholder="例如：https://api.openai.com/v1/chat/completions"
          />
        </el-form-item>

        <el-form-item v-if="form.provider === 'azure'" label="API版本" prop="api_version">
          <el-input v-model="form.api_version" placeholder="例如：2024-02-15-preview" />
        </el-form-item>

        <el-form-item label="API Key" prop="api_key">
          <el-input
            v-model="form.api_key"
            type="password"
            show-password
            placeholder="输入API Key（不会保存明文）"
          />
          <div class="form-tip">API Key仅保存在服务器端，不会泄露到前端</div>
        </el-form-item>

        <el-divider content-position="left">生成参数</el-divider>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="Temperature">
              <el-slider
                v-model="form.temperature"
                :min="0"
                :max="2"
                :step="0.1"
                show-stops
                :marks="{ 0: '0', 1: '1', 2: '2' }"
              />
              <div class="slider-tip">较低的值使输出更确定性，较高的值使输出更有创造性</div>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="最大Token数">
              <el-input-number
                v-model="form.max_tokens"
                :min="100"
                :max="32000"
                :step="100"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="请求超时(秒)">
              <el-input-number
                v-model="form.timeout"
                :min="10"
                :max="300"
                :step="10"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-divider content-position="left">操作</el-divider>

        <el-form-item>
          <el-button type="primary" @click="handleSave" :loading="saving">
            保存配置
          </el-button>
          <el-button @click="handleReset">
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 测试结果对话框 -->
    <el-dialog v-model="testResultVisible" title="测试结果" width="400px">
      <div v-if="testSuccess" class="test-result success">
        <el-icon color="#67c23a" :size="32"><CircleCheck /></el-icon>
        <p>连接成功！</p>
        <p class="tip">AI模型响应正常，可以开始使用</p>
      </div>
      <div v-else class="test-result error">
        <el-icon color="#f56c6c" :size="32"><CircleClose /></el-icon>
        <p>连接失败</p>
        <p class="tip">{{ testErrorMessage }}</p>
      </div>
      <template #footer>
        <el-button @click="testResultVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { CircleCheck, CircleClose } from '@element-plus/icons-vue'
import { getAiConfig, updateAiConfig, testAiConfig, type AIConfig } from '@/api/ai-chat'

const formRef = ref()
const saving = ref(false)
const testing = ref(false)
const testResultVisible = ref(false)
const testSuccess = ref(false)
const testErrorMessage = ref('')

const form = reactive<AIConfig>({
  config_name: '',
  provider: 'openai',
  endpoint: '',
  api_version: '',
  model_name: 'gpt-4o',
  api_key: '',
  temperature: 0.7,
  max_tokens: 2048,
  timeout: 60,
  enable: true
})

const rules = {
  provider: [{ required: true, message: '请选择AI提供商', trigger: 'change' }],
  model_name: [{ required: true, message: '请输入模型名称', trigger: 'blur' }],
  endpoint: [{ required: true, message: '请输入API Endpoint', trigger: 'blur' }]
}

onMounted(async () => {
  await loadConfig()
})

async function loadConfig() {
  try {
    const res = await getAiConfig()
    if (res.code === 200 && res.data) {
      const data = res.data
      Object.assign(form, {
        id: data.id,
        tenant_id: data.tenant_id,
        config_name: data.config_name || '',
        provider: data.provider || 'openai',
        endpoint: data.endpoint || '',
        api_version: data.api_version || '',
        model_name: data.model_name || 'gpt-4o',
        api_key: '', // 不显示已保存的API Key
        temperature: data.temperature || 0.7,
        max_tokens: data.max_tokens || 2048,
        timeout: data.timeout || 60,
        enable: data.enable !== false
      })
    }
  } catch (error) {
    console.error('加载AI配置失败', error)
  }
}

async function handleSave() {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  saving.value = true
  try {
    const res = await updateAiConfig(form)
    if (res.code === 200) {
      ElMessage.success('保存成功')
      await loadConfig()
    } else {
      ElMessage.error(res.message || '保存失败')
    }
  } catch (error: any) {
    ElMessage.error(error.message || '保存失败')
  } finally {
    saving.value = false
  }
}

async function handleTest() {
  testing.value = true
  testResultVisible.value = true
  testSuccess.value = false
  testErrorMessage.value = ''

  try {
    const res = await testAiConfig(form)
    testSuccess.value = res.code === 200
    if (!testSuccess.value) {
      testErrorMessage.value = res.message || '未知错误'
    }
  } catch (error: any) {
    testSuccess.value = false
    testErrorMessage.value = error?.message || '连接失败，请检查配置'
  } finally {
    testing.value = false
  }
}

function handleReset() {
  Object.assign(form, {
    config_name: '',
    provider: 'openai',
    endpoint: '',
    api_version: '',
    model_name: 'gpt-4o',
    api_key: '',
    temperature: 0.7,
    max_tokens: 2048,
    timeout: 60,
    enable: true
  })
  formRef.value?.clearValidate()
}
</script>

<style scoped lang="scss">
.ai-config {
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}

.slider-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

.test-result {
  text-align: center;
  padding: 20px;

  p {
    margin: 12px 0 0;
    font-size: 16px;
  }

  .tip {
    font-size: 13px;
    color: #909399;
  }

  &.success {
    color: #67c23a;
  }

  &.error {
    color: #f56c6c;
  }
}
</style>
