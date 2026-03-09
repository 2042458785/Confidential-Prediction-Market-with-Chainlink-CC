<template>
  <div class="admin-panel">
    <h2>管理员面板 - 提交结算结果</h2>
    
    <div class="section">
      <h3>1. 设置信任签名者</h3>
      <p class="info">首次使用前，需要设置信任的签名者地址（只有 owner 可以操作）</p>
      
      <div class="input-group">
        <input
          v-model="trustedSignerAddress"
          type="text"
          placeholder="输入信任的签名者地址 (0x...)"
          class="address-input"
        />
        <button
          @click="setTrustedSigner"
          :disabled="isSettingSigner || !trustedSignerAddress"
          class="btn-primary"
        >
          {{ isSettingSigner ? '设置中...' : '设置信任签名者' }}
        </button>
      </div>
      
      <p v-if="currentTrustedSigner" class="current-signer">
        当前信任签名者: {{ currentTrustedSigner }}
      </p>
    </div>

    <div class="section">
      <h3>2. 提交结算结果</h3>
      <p class="info">手动输入市场 ID 并提交结果（outcome 固定为 1）</p>
      
      <div class="input-group">
        <input
          v-model.number="manualMarketId"
          type="number"
          min="1"
          placeholder="输入市场 ID (例如: 3)"
          class="market-id-input"
        />
        <button
          @click="submitManualResult"
          :disabled="isSubmitting || !manualMarketId"
          class="btn-submit"
        >
          {{ isSubmitting ? '提交中...' : '提交结果 (outcome=1)' }}
        </button>
      </div>

      <p v-if="submitTxHash" class="tx-link">
        交易哈希: 
        <a :href="`${EXPLORER_URL}/tx/${submitTxHash}`" target="_blank">
          {{ submitTxHash.slice(0, 10) }}...{{ submitTxHash.slice(-8) }}
        </a>
      </p>
      
      <details class="advanced-section">
        <summary>高级：自动查找待结算市场</summary>
        
        <button @click="loadResolvingMarkets" class="btn-secondary" :disabled="isLoading">
          {{ isLoading ? '加载中...' : '刷新待结算市场' }}
        </button>

        <div v-if="resolvingMarkets.length === 0" class="empty-message">
          暂无待结算的市场（状态为 Resolving）
        </div>

        <div v-else class="markets-list">
          <div v-for="market in resolvingMarkets" :key="market.id" class="market-item">
            <div class="market-info">
              <span class="market-id">市场 #{{ market.id }}</span>
              <span class="market-desc">{{ market.description }}</span>
            </div>
            
            <button
              @click="submitResult(market.id)"
              :disabled="isSubmitting"
              class="btn-submit-small"
            >
              {{ isSubmitting ? '提交中...' : '提交结果 (outcome=1)' }}
            </button>
          </div>
        </div>
      </details>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ethers } from 'ethers'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS || '0x3f2E7c68Ec5fC4E6419C5922a444C0e20Db13b60'
const RPC_URL = import.meta.env.VITE_RPC_URL || 'https://ethereum-sepolia-rpc.publicnode.com'
const EXPLORER_URL = import.meta.env.VITE_EXPLORER_URL || 'https://sepolia.etherscan.io'

const CONTRACT_ABI = [
  'function setTrustedSigner(address _signer)',
  'function trustedSigner() view returns (address)',
  'function submitResult(uint256 marketId, uint8 outcome, bytes signature)',
  'function markets(uint256 marketId) view returns (uint256 id, string description, uint8 state, uint48 createTimeStamp, uint48 closeTime, uint48 resolveWindow, uint48 resolveTimeStamp, uint8 outcome, bool resolved)',
  'function OrderMarketid() view returns (uint256)',
]

const props = defineProps<{
  account: string | null
}>()

const trustedSignerAddress = ref('')
const currentTrustedSigner = ref<string | null>(null)
const isSettingSigner = ref(false)
const isLoading = ref(false)
const isSubmitting = ref(false)
const submitTxHash = ref<string | null>(null)
const manualMarketId = ref<number | null>(null)

interface Market {
  id: number
  description: string
  state: number
}

const resolvingMarkets = ref<Market[]>([])

// 加载当前信任签名者
async function loadTrustedSigner() {
  try {
    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)
    
    const signer = await contract.trustedSigner()
    currentTrustedSigner.value = signer
    
    // 如果是零地址，显示未设置
    if (signer === ethers.ZeroAddress) {
      currentTrustedSigner.value = '未设置'
    }
  } catch (err) {
    console.error('加载信任签名者失败:', err)
  }
}

// 设置信任签名者
async function setTrustedSigner() {
  if (!props.account) {
    alert('请先连接钱包')
    return
  }

  if (!ethers.isAddress(trustedSignerAddress.value)) {
    alert('请输入有效的以太坊地址')
    return
  }

  try {
    isSettingSigner.value = true

    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    const tx = await contract.setTrustedSigner(trustedSignerAddress.value)
    console.log('设置信任签名者交易已提交:', tx.hash)
    
    await tx.wait()
    console.log('交易已确认！')
    
    alert('信任签名者设置成功！')
    
    // 重新加载
    await loadTrustedSigner()
    trustedSignerAddress.value = ''
  } catch (err: any) {
    console.error('设置信任签名者失败:', err)
    
    let errorMessage = '设置失败'
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Ownable: caller is not the owner')) {
      errorMessage = '只有 owner 可以设置信任签名者'
    } else {
      errorMessage = `设置失败: ${err.message || '未知错误'}`
    }
    
    alert(errorMessage)
  } finally {
    isSettingSigner.value = false
  }
}

// 加载待结算的市场
async function loadResolvingMarkets() {
  try {
    isLoading.value = true
    resolvingMarkets.value = []

    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)

    // 获取市场总数
    const totalMarkets = await contract.OrderMarketid()
    console.log('市场总数:', totalMarkets.toString())

    // 遍历所有市场，找出状态为 Resolving (2) 的市场
    // 注意：市场 ID 可能从 1 开始，所以从 1 遍历到 totalMarkets
    for (let i = 1; i <= Number(totalMarkets); i++) {
      try {
        const market = await contract.markets(i)
        
        console.log(`市场 ${i}:`, {
          description: market.description,
          state: market.state,
          stateText: ['Open', 'Closed', 'Resolving', 'Resolved'][market.state]
        })
        
        // state: 0=Open, 1=Closed, 2=Resolving, 3=Resolved
        if (market.state === 2) {
          resolvingMarkets.value.push({
            id: i,
            description: market.description,
            state: market.state
          })
        }
      } catch (err) {
        console.error(`加载市场 ${i} 失败:`, err)
      }
    }

    console.log('待结算市场:', resolvingMarkets.value)
    
    if (resolvingMarkets.value.length === 0) {
      console.log('提示：没有找到状态为 Resolving 的市场')
      console.log('请确保：')
      console.log('1. 市场已关闭（当前时间 >= closeTime）')
      console.log('2. 已点击"请求结算"按钮')
      console.log('3. 市场状态已变为 Resolving')
    }
  } catch (err) {
    console.error('加载市场失败:', err)
    alert('加载市场失败')
  } finally {
    isLoading.value = false
  }
}

// 提交结算结果
async function submitResult(marketId: number) {
  if (!props.account) {
    alert('请先连接钱包')
    return
  }

  try {
    isSubmitting.value = true
    submitTxHash.value = null

    const outcome = 1 // 固定为 1

    console.log('开始提交结果...')
    console.log('市场ID:', marketId)
    console.log('结果:', outcome)

    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    // 1. 构造消息哈希
    const messageHash = ethers.solidityPackedKeccak256(
      ['uint256', 'uint8'],
      [marketId, outcome]
    )
    console.log('消息哈希:', messageHash)

    // 2. 签名（ethers 会自动加 "\x19Ethereum Signed Message:\n32" 前缀）
    const signature = await signer.signMessage(ethers.getBytes(messageHash))
    console.log('签名:', signature)

    // 3. 调用合约
    const tx = await contract.submitResult(marketId, outcome, signature)
    console.log('提交结果交易已提交:', tx.hash)
    submitTxHash.value = tx.hash
    
    await tx.wait()
    console.log('交易已确认！')
    
    alert('结算结果提交成功！')
    
    // 刷新市场列表
    await loadResolvingMarkets()
  } catch (err: any) {
    console.error('提交结果失败:', err)
    
    let errorMessage = '提交失败'
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Invalid signature')) {
      errorMessage = '签名无效：请确保当前钱包地址已被设置为信任签名者'
    } else if (err.message?.includes('Not resolving')) {
      errorMessage = '市场状态不是 Resolving'
    } else {
      errorMessage = `提交失败: ${err.message || '未知错误'}`
    }
    
    alert(errorMessage)
  } finally {
    isSubmitting.value = false
  }
}

// 手动提交结果
async function submitManualResult() {
  if (!manualMarketId.value) {
    alert('请输入市场 ID')
    return
  }
  await submitResult(manualMarketId.value)
}

onMounted(() => {
  loadTrustedSigner()
  loadResolvingMarkets()
})
</script>

<style scoped>
.admin-panel {
  max-width: 900px;
  margin: 0 auto;
  padding: 24px;
}

.admin-panel h2 {
  font-size: 28px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 32px;
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
  margin-bottom: 12px;
}

.info {
  font-size: 14px;
  color: #6b7280;
  margin-bottom: 16px;
}

.input-group {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.address-input {
  flex: 1;
  padding: 12px 16px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  font-size: 14px;
  font-family: monospace;
}

.address-input:focus {
  outline: none;
  border-color: #667eea;
}

.current-signer {
  font-size: 13px;
  color: #059669;
  font-weight: 600;
  padding: 8px 12px;
  background: #d1fae5;
  border-radius: 8px;
  font-family: monospace;
}

.btn-primary,
.btn-secondary,
.btn-submit {
  padding: 12px 24px;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(102, 126, 234, 0.4);
}

.btn-secondary {
  background: #f3f4f6;
  color: #374151;
  margin-bottom: 16px;
}

.btn-secondary:hover:not(:disabled) {
  background: #e5e7eb;
}

.btn-submit {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  font-size: 13px;
  padding: 10px 20px;
}

.btn-submit:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(16, 185, 129, 0.4);
}

.btn-primary:disabled,
.btn-secondary:disabled,
.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.empty-message {
  text-align: center;
  padding: 40px 20px;
  color: #6b7280;
  font-size: 14px;
}

.markets-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.market-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #f9fafb;
  border-radius: 12px;
  border: 2px solid #e5e7eb;
}

.market-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.market-id {
  font-size: 13px;
  font-weight: 600;
  color: #667eea;
}

.market-desc {
  font-size: 14px;
  color: #1f2937;
}

.tx-link {
  margin-top: 16px;
  font-size: 13px;
  color: #6b7280;
  text-align: center;
}

.tx-link a {
  color: #667eea;
  text-decoration: none;
  font-weight: 600;
}

@media (max-width: 768px) {
  .input-group {
    flex-direction: column;
  }
  
  .market-item {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }
}
</style>


.market-id-input {
  flex: 1;
  padding: 12px 16px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  font-size: 14px;
  font-family: monospace;
}

.market-id-input:focus {
  outline: none;
  border-color: #667eea;
}

.btn-submit-small {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  font-size: 13px;
  padding: 10px 20px;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-submit-small:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(16, 185, 129, 0.4);
}

.btn-submit-small:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.advanced-section {
  margin-top: 24px;
  padding: 16px;
  background: #f9fafb;
  border-radius: 12px;
}

.advanced-section summary {
  cursor: pointer;
  font-weight: 600;
  color: #374151;
  margin-bottom: 16px;
}

.advanced-section summary:hover {
  color: #667eea;
}
