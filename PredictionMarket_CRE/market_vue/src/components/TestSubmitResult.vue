<template>
  <div class="test-panel">
    <h2>🧪 测试 submitResult 功能</h2>
    <p class="warning">⚠️ 此工具用于测试，会尝试提交结果（不检查市场状态）</p>
    
    <div class="section">
      <h3>当前配置</h3>
      <div class="info-grid">
        <div class="info-item">
          <span class="label">合约地址:</span>
          <span class="value mono">{{ CONTRACT_ADDRESS }}</span>
        </div>
        <div class="info-item">
          <span class="label">信任签名者:</span>
          <span class="value mono">{{ trustedSigner || '加载中...' }}</span>
        </div>
        <div class="info-item">
          <span class="label">当前钱包:</span>
          <span class="value mono">{{ account || '未连接' }}</span>
        </div>
        <div class="info-item">
          <span class="label">匹配状态:</span>
          <span class="value" :class="isSignerMatch ? 'success' : 'error'">
            {{ isSignerMatch ? '✅ 匹配' : '❌ 不匹配' }}
          </span>
        </div>
      </div>
    </div>

    <div class="section">
      <h3>提交结果</h3>
      
      <div class="form-group">
        <label>市场 ID</label>
        <input
          v-model.number="marketId"
          type="number"
          min="1"
          placeholder="输入市场 ID (例如: 3)"
        />
      </div>

      <div class="form-group">
        <label>结果 (Outcome)</label>
        <div class="radio-group">
          <label class="radio-label">
            <input type="radio" v-model.number="outcome" :value="0" />
            <span>0 - 否</span>
          </label>
          <label class="radio-label">
            <input type="radio" v-model.number="outcome" :value="1" />
            <span>1 - 是</span>
          </label>
        </div>
      </div>

      <button
        @click="testSubmitResult"
        :disabled="isSubmitting || !account || !marketId"
        class="btn-submit"
      >
        {{ isSubmitting ? '提交中...' : '🚀 测试提交结果' }}
      </button>

      <div v-if="txHash" class="result success">
        <h4>✅ 提交成功！</h4>
        <p>交易哈希: 
          <a :href="`${EXPLORER_URL}/tx/${txHash}`" target="_blank">
            {{ txHash }}
          </a>
        </p>
      </div>

      <div v-if="error" class="result error">
        <h4>❌ 提交失败</h4>
        <p>{{ error }}</p>
      </div>
    </div>

    <div class="section">
      <h3>📝 签名过程说明</h3>
      <ol class="steps">
        <li>构造消息哈希: <code>keccak256(abi.encodePacked(marketId, outcome))</code></li>
        <li>使用钱包签名（自动添加以太坊前缀）</li>
        <li>调用合约: <code>submitResult(marketId, outcome, signature)</code></li>
        <li>合约验证签名并更新市场状态</li>
      </ol>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ethers } from 'ethers'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS || '0x3f2E7c68Ec5fC4E6419C5922a444C0e20Db13b60'
const RPC_URL = import.meta.env.VITE_RPC_URL || 'https://ethereum-sepolia-rpc.publicnode.com'
const EXPLORER_URL = import.meta.env.VITE_EXPLORER_URL || 'https://sepolia.etherscan.io'

const CONTRACT_ABI = [
  'function trustedSigner() view returns (address)',
  'function submitResult(uint256 marketId, uint8 outcome, bytes signature)',
]

const props = defineProps<{
  account: string | null
}>()

const marketId = ref<number>(3)
const outcome = ref<number>(1)
const isSubmitting = ref(false)
const txHash = ref<string | null>(null)
const error = ref<string | null>(null)
const trustedSigner = ref<string | null>(null)

const isSignerMatch = computed(() => {
  if (!props.account || !trustedSigner.value) return false
  return props.account.toLowerCase() === trustedSigner.value.toLowerCase()
})

async function loadTrustedSigner() {
  try {
    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)
    
    const signer = await contract.trustedSigner()
    trustedSigner.value = signer
    
    if (signer === ethers.ZeroAddress) {
      trustedSigner.value = '未设置'
    }
  } catch (err) {
    console.error('加载信任签名者失败:', err)
  }
}

async function testSubmitResult() {
  if (!props.account) {
    alert('请先连接钱包')
    return
  }

  if (!marketId.value) {
    alert('请输入市场 ID')
    return
  }

  try {
    isSubmitting.value = true
    txHash.value = null
    error.value = null

    console.log('=== 开始测试 submitResult ===')
    console.log('市场 ID:', marketId.value)
    console.log('结果 (outcome):', outcome.value)
    console.log('当前钱包:', props.account)
    console.log('信任签名者:', trustedSigner.value)

    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    // 步骤 1: 构造消息哈希
    const messageHash = ethers.solidityPackedKeccak256(
      ['uint256', 'uint8'],
      [marketId.value, outcome.value]
    )
    console.log('步骤 1 - 消息哈希:', messageHash)

    // 步骤 2: 签名
    const signature = await signer.signMessage(ethers.getBytes(messageHash))
    console.log('步骤 2 - 签名:', signature)

    // 步骤 3: 调用合约
    console.log('步骤 3 - 调用 submitResult...')
    const tx = await contract.submitResult(marketId.value, outcome.value, signature)
    console.log('交易已提交:', tx.hash)
    txHash.value = tx.hash
    
    console.log('等待交易确认...')
    const receipt = await tx.wait()
    console.log('交易已确认！区块:', receipt.blockNumber)
    
    console.log('=== 测试成功！===')
    alert('✅ 提交成功！市场状态应该已更新为 Resolved。')
  } catch (err: any) {
    console.error('=== 测试失败 ===')
    console.error('完整错误:', err)
    
    let errorMessage = '提交失败'
    
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Invalid signature')) {
      errorMessage = '签名无效：当前钱包地址不是信任签名者'
    } else if (err.message?.includes('Not resolving')) {
      errorMessage = '市场状态不是 Resolving（这是预期的错误，说明合约验证正常工作）'
    } else if (err.reason) {
      errorMessage = `合约错误: ${err.reason}`
    } else if (err.data?.message) {
      errorMessage = `错误: ${err.data.message}`
    } else {
      errorMessage = `错误: ${err.message || '未知错误'}`
    }
    
    error.value = errorMessage
    alert(errorMessage)
  } finally {
    isSubmitting.value = false
  }
}

onMounted(() => {
  loadTrustedSigner()
})
</script>

<style scoped>
.test-panel {
  max-width: 800px;
  margin: 40px auto;
  padding: 24px;
}

.test-panel h2 {
  font-size: 28px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 12px;
  text-align: center;
}

.warning {
  background: #fef3c7;
  color: #92400e;
  padding: 12px 16px;
  border-radius: 8px;
  font-size: 14px;
  margin-bottom: 24px;
  text-align: center;
}

.section {
  background: white;
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.section h3 {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 16px;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 12px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.label {
  font-size: 13px;
  color: #6b7280;
  font-weight: 500;
}

.value {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.value.mono {
  font-family: monospace;
  font-size: 12px;
  word-break: break-all;
}

.value.success {
  color: #059669;
}

.value.error {
  color: #dc2626;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: #374151;
  margin-bottom: 8px;
}

.form-group input[type="number"] {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  font-size: 16px;
}

.form-group input[type="number"]:focus {
  outline: none;
  border-color: #667eea;
}

.radio-group {
  display: flex;
  gap: 24px;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  font-size: 14px;
  color: #374151;
}

.radio-label input[type="radio"] {
  width: 18px;
  height: 18px;
  cursor: pointer;
}

.btn-submit {
  width: 100%;
  padding: 16px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-submit:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(16, 185, 129, 0.4);
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.result {
  margin-top: 20px;
  padding: 16px;
  border-radius: 12px;
}

.result.success {
  background: #d1fae5;
  color: #065f46;
}

.result.error {
  background: #fee2e2;
  color: #991b1b;
}

.result h4 {
  font-size: 16px;
  font-weight: 700;
  margin-bottom: 8px;
}

.result p {
  font-size: 14px;
  word-break: break-all;
}

.result a {
  color: #667eea;
  text-decoration: underline;
}

.steps {
  padding-left: 24px;
  color: #374151;
}

.steps li {
  margin-bottom: 8px;
  font-size: 14px;
  line-height: 1.6;
}

.steps code {
  background: #f3f4f6;
  padding: 2px 6px;
  border-radius: 4px;
  font-family: monospace;
  font-size: 13px;
}

@media (max-width: 768px) {
  .info-grid {
    grid-template-columns: 1fr;
  }
  
  .radio-group {
    flex-direction: column;
    gap: 12px;
  }
}
</style>
