<template>
  <div class="market-detail">
    <button class="back-btn" @click="$emit('back')">
      <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
        <path d="M12 4l-8 8 8 8" stroke="currentColor" stroke-width="2" fill="none"/>
      </svg>
      返回
    </button>

    <div class="detail-card">
      <div class="detail-header">
        <div>
          <span class="market-id">#{{ market.id }}</span>
          <h2 class="market-title">{{ market.description }}</h2>
        </div>
        <span class="status-badge" :class="statusClass">
          {{ statusText }}
        </span>
      </div>

      <div class="market-info">
        <div class="info-item">
          <span class="info-label">总池</span>
          <span class="info-value">{{ formatPool(market.totalPool) }} ETH</span>
        </div>
        <div class="info-item">
          <span class="info-label">创建时间</span>
          <span class="info-value">{{ formatDate(market.createTimeStamp) }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">截止时间</span>
          <span class="info-value">{{ formatDate(market.closeTime) }}</span>
        </div>
        <div class="info-item" v-if="market.resolved">
          <span class="info-label">结算结果</span>
          <span class="info-value outcome">{{ market.outcome === 1 ? '是' : '否' }}</span>
        </div>
      </div>

      <div v-if="canBet && account" class="betting-section">
        <h3>下注</h3>
        
        <div class="bet-options">
          <button
            class="option-btn"
            :class="{ active: selectedOption === 0 }"
            @click="selectedOption = 0"
          >
            <span class="option-label">选项 0</span>
            <span class="option-odds">否</span>
          </button>
          
          <button
            class="option-btn"
            :class="{ active: selectedOption === 1 }"
            @click="selectedOption = 1"
          >
            <span class="option-label">选项 1</span>
            <span class="option-odds">是</span>
          </button>
        </div>

        <div class="bet-input">
          <div class="input-header">
            <label>下注金额 (ETH)</label>
            <span class="balance">余额: {{ parseFloat(userBalance).toFixed(4) }} ETH</span>
          </div>
          <input
            v-model="betAmount"
            type="number"
            step="0.001"
            min="0"
            placeholder="0.01"
          />
        </div>

        <button
          class="place-bet-btn"
          @click="placeBet"
          :disabled="isPlacingBet || !betAmount"
        >
          {{ isPlacingBet ? '提交中...' : '确认下注' }}
        </button>

        <p v-if="txHash" class="tx-link">
          交易哈希: 
          <a :href="`${EXPLORER_URL}/tx/${txHash}`" target="_blank">
            {{ txHash.slice(0, 10) }}...{{ txHash.slice(-8) }}
          </a>
        </p>
      </div>

      <div v-else-if="!account" class="connect-prompt">
        <p>请先连接钱包以参与下注</p>
      </div>

      <div v-else class="closed-message">
        <p>该市场已关闭，无法继续下注</p>
      </div>

      <div v-if="canRequestSettlement && account" class="settlement-section">
        <h3>请求结算</h3>
        <p class="settlement-info">市场已关闭，可以请求 CRE 工作流进行结算</p>
        
        <button
          class="settlement-btn"
          @click="requestSettlement"
          :disabled="isRequestingSettlement"
        >
          {{ isRequestingSettlement ? '提交中...' : '请求结算' }}
        </button>

        <p v-if="settlementTxHash" class="tx-link">
          交易哈希: 
          <a :href="`${EXPLORER_URL}/tx/${settlementTxHash}`" target="_blank">
            {{ settlementTxHash.slice(0, 10) }}...{{ settlementTxHash.slice(-8) }}
          </a>
        </p>
      </div>

      <!-- 领取奖励 -->
      <ClaimWinnings 
        v-if="market.resolved"
        :market="market" 
        :account="account"
        @claimed="$emit('bet-placed')"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ethers } from 'ethers'
import type { Market } from '../types'
import { MarketStatus } from '../types'
import ClaimWinnings from './ClaimWinnings.vue'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS || '0x35E3e0a0C9BDB4f4FFc2e09f7500dbf5249464dF'
const EXPLORER_URL = import.meta.env.VITE_EXPLORER_URL || 'https://sepolia.etherscan.io'

const CONTRACT_ABI = [
  'function placeBet(uint256 marketId, bytes encryptedBet) payable',
  'function requestSettlement(uint256 marketId)',
]

const props = defineProps<{
  market: Market
  account: string | null
}>()

const emit = defineEmits<{
  back: []
  'bet-placed': []
}>()

const selectedOption = ref(0)
const betAmount = ref('')
const isPlacingBet = ref(false)
const txHash = ref<string | null>(null)
const userBalance = ref<string>('0')
const isRequestingSettlement = ref(false)
const settlementTxHash = ref<string | null>(null)

// 加载用户余额
async function loadBalance() {
  if (!props.account) return
  
  try {
    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const balance = await provider.getBalance(props.account)
    userBalance.value = ethers.formatEther(balance)
  } catch (err) {
    console.error('获取余额失败:', err)
  }
}

// 当账户变化时重新加载余额
import { watch } from 'vue'
watch(() => props.account, () => {
  if (props.account) {
    loadBalance()
  }
}, { immediate: true })

const statusText = computed(() => {
  // 如果市场已解决，显示已结算
  if (props.market.resolved) {
    return '已结算'
  }
  
  // 根据 state 判断状态
  // state: 0=Open, 1=Closed, 2=Resolving, 3=Resolved
  if (props.market.state === 2) {
    return '结算中'
  }
  
  // 根据当前时间判断状态
  const now = Math.floor(Date.now() / 1000)
  
  if (now >= props.market.closeTime) {
    return '已关闭'
  }
  
  return '进行中'
})

const statusClass = computed(() => {
  // 如果市场已解决，显示已结算样式
  if (props.market.resolved) {
    return 'status-resolved'
  }
  
  // 根据 state 判断状态
  if (props.market.state === 2) {
    return 'status-resolving'
  }
  
  // 根据当前时间判断状态
  const now = Math.floor(Date.now() / 1000)
  
  if (now >= props.market.closeTime) {
    return 'status-closed'
  }
  
  return 'status-open'
})

// 判断市场是否可以下注
const canBet = computed(() => {
  if (props.market.resolved) return false
  const now = Math.floor(Date.now() / 1000)
  return now < props.market.closeTime
})

// 判断是否可以请求结算
const canRequestSettlement = computed(() => {
  if (props.market.resolved) return false
  const now = Math.floor(Date.now() / 1000)
  return now >= props.market.closeTime
})

function formatPool(pool: string): string {
  // 直接显示完整数值，不限制小数位数
  return pool
}

function formatDate(timestamp: number): string {
  const date = new Date(timestamp * 1000)
  return date.toLocaleString('zh-CN')
}

async function placeBet() {
  if (!props.account || !betAmount.value) {
    alert('请输入下注金额')
    return
  }

  try {
    isPlacingBet.value = true
    txHash.value = null

    console.log('开始下注...')
    console.log('市场ID:', props.market.id)
    console.log('选项:', selectedOption.value)
    console.log('金额:', betAmount.value, 'ETH')

    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    // 将金额转换为字符串，然后解析为 Wei
    const value = ethers.parseEther(String(betAmount.value))
    
    // 编码下注数据：将选项和金额编码为 bytes
    // 格式：ABI 编码 (uint8 option, uint256 amount)
    const abiCoder = ethers.AbiCoder.defaultAbiCoder()
    const encryptedBet = abiCoder.encode(
      ['uint8', 'uint256'],
      [selectedOption.value, value]
    )

    console.log('编码后的数据:', encryptedBet)
    console.log('发送的 ETH 金额:', ethers.formatEther(value))

    const tx = await contract.placeBet(props.market.id, encryptedBet, { value })
    console.log('交易已提交:', tx.hash)
    txHash.value = tx.hash
    
    console.log('等待交易确认...')
    await tx.wait()
    console.log('交易已确认！')
    
    alert('下注成功！')
    
    // 重置表单
    betAmount.value = ''
    selectedOption.value = 0
    
    emit('bet-placed')
  } catch (err: any) {
    console.error('下注失败 - 完整错误:', err)
    console.error('错误代码:', err.code)
    console.error('错误消息:', err.message)
    console.error('错误数据:', err.data)
    
    // 提供更友好的错误提示
    let errorMessage = '下注失败'
    
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Market not open')) {
      errorMessage = '市场未开放'
    } else if (err.message?.includes('Betting closed')) {
      errorMessage = '下注已截止'
    } else if (err.message?.includes('insufficient funds')) {
      errorMessage = '余额不足'
    } else if (err.data?.message) {
      errorMessage = `合约错误: ${err.data.message}`
    } else if (err.reason) {
      errorMessage = `错误: ${err.reason}`
    } else {
      errorMessage = `下注失败: ${err.message || '未知错误'}`
    }
    
    alert(errorMessage)
  } finally {
    isPlacingBet.value = false
  }
}

async function requestSettlement() {
  if (!props.account) {
    alert('请先连接钱包')
    return
  }

  try {
    isRequestingSettlement.value = true
    settlementTxHash.value = null

    console.log('请求结算市场:', props.market.id)

    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    const tx = await contract.requestSettlement(props.market.id)
    console.log('结算请求已提交:', tx.hash)
    settlementTxHash.value = tx.hash
    
    console.log('等待交易确认...')
    await tx.wait()
    console.log('结算请求已确认！')
    
    alert('结算请求成功！CRE 工作流将处理结算。')
    
    emit('bet-placed') // 触发刷新
  } catch (err: any) {
    console.error('请求结算失败:', err)
    
    let errorMessage = '请求结算失败'
    
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Invalid state')) {
      errorMessage = '市场状态无效'
    } else if (err.message?.includes('Betting not closed')) {
      errorMessage = '市场尚未关闭'
    } else if (err.message?.includes('Resolve window expired')) {
      errorMessage = '结算窗口已过期'
    } else if (err.data?.message) {
      errorMessage = `合约错误: ${err.data.message}`
    } else if (err.reason) {
      errorMessage = `错误: ${err.reason}`
    } else {
      errorMessage = `请求结算失败: ${err.message || '未知错误'}`
    }
    
    alert(errorMessage)
  } finally {
    isRequestingSettlement.value = false
  }
}
</script>

<style scoped>
.market-detail {
  padding: 24px 0;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  margin-bottom: 20px;
  transition: all 0.2s;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

.detail-card {
  background: white;
  border-radius: 20px;
  padding: 32px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
  gap: 16px;
}

.market-id {
  font-size: 14px;
  font-weight: 600;
  color: #667eea;
}

.market-title {
  font-size: 28px;
  font-weight: 700;
  color: #1f2937;
  margin-top: 8px;
  line-height: 1.3;
}

.status-badge {
  padding: 6px 16px;
  border-radius: 16px;
  font-size: 14px;
  font-weight: 600;
  white-space: nowrap;
}

.status-open {
  background: #d1fae5;
  color: #065f46;
}

.status-closed {
  background: #fef3c7;
  color: #92400e;
}

.status-resolved {
  background: #dbeafe;
  color: #1e40af;
}

.status-resolving {
  background: #fef3c7;
  color: #d97706;
}

.market-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  padding: 24px;
  background: #f9fafb;
  border-radius: 12px;
  margin-bottom: 32px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-label {
  font-size: 13px;
  color: #6b7280;
  font-weight: 500;
}

.info-value {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.info-value.outcome {
  color: #667eea;
}

.betting-section {
  border-top: 2px solid #e5e7eb;
  padding-top: 32px;
}

.betting-section h3 {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 20px;
}

.bet-options {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 24px;
}

.option-btn {
  padding: 20px;
  background: #f9fafb;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.option-btn:hover {
  border-color: #667eea;
  background: #f3f4f6;
}

.option-btn.active {
  border-color: #667eea;
  background: #eef2ff;
}

.option-label {
  font-size: 14px;
  font-weight: 600;
  color: #6b7280;
}

.option-odds {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
}

.bet-input {
  margin-bottom: 20px;
}

.bet-input .input-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.bet-input label {
  font-size: 14px;
  font-weight: 600;
  color: #374151;
}

.bet-input .balance {
  font-size: 12px;
  color: #6b7280;
}

.bet-input input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  font-size: 16px;
  transition: all 0.2s;
}

.bet-input input:focus {
  outline: none;
  border-color: #667eea;
}

.place-bet-btn {
  width: 100%;
  padding: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
}

.place-bet-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(102, 126, 234, 0.4);
}

.place-bet-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
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

.connect-prompt,
.closed-message {
  text-align: center;
  padding: 40px 20px;
  color: #6b7280;
  font-size: 16px;
}

.settlement-section {
  border-top: 2px solid #e5e7eb;
  padding-top: 32px;
  margin-top: 32px;
}

.settlement-section h3 {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 12px;
}

.settlement-info {
  font-size: 14px;
  color: #6b7280;
  margin-bottom: 20px;
}

.settlement-btn {
  width: 100%;
  padding: 16px;
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
}

.settlement-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(245, 158, 11, 0.4);
}

.settlement-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

@media (max-width: 768px) {
  .detail-card {
    padding: 20px;
  }
  
  .market-title {
    font-size: 22px;
  }
  
  .market-info {
    grid-template-columns: 1fr;
  }
  
  .bet-options {
    grid-template-columns: 1fr;
  }
}
</style>
