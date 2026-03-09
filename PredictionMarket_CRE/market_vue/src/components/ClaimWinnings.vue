<template>
  <div class="claim-section">
    <h3>领取奖励</h3>
    
    <div v-if="!account" class="connect-prompt">
      <p>请先连接钱包查看您的下注</p>
    </div>

    <div v-else-if="loading" class="loading">
      <div class="spinner"></div>
      <p>加载中...</p>
    </div>

    <div v-else-if="userBets.length === 0" class="empty-message">
      <p>您在此市场没有下注</p>
    </div>

    <div v-else class="bets-list">
      <div v-for="bet in userBets" :key="bet.index" class="bet-item" :class="getBetClass(bet)">
        <div class="bet-info">
          <div class="bet-header">
            <span class="bet-label">下注 #{{ bet.index }}</span>
            <span class="bet-status" :class="getStatusClass(bet)">
              {{ getStatusText(bet) }}
            </span>
          </div>
          
          <div class="bet-details">
            <div class="detail-item">
              <span class="label">选项:</span>
              <span class="value">{{ bet.option === 0 ? '否' : '是' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">金额:</span>
              <span class="value">{{ formatAmount(bet.amount) }} ETH</span>
            </div>
            <div v-if="market.resolved && bet.option === market.outcome" class="detail-item">
              <span class="label">预计奖励:</span>
              <span class="value reward">{{ calculateReward(bet) }} ETH</span>
            </div>
          </div>
        </div>

        <button
          v-if="canClaim(bet)"
          @click="claimWinnings(bet.index)"
          :disabled="isClaiming"
          class="claim-btn"
        >
          {{ isClaiming ? '领取中...' : '领取奖励' }}
        </button>
      </div>

      <p v-if="claimTxHash" class="tx-link">
        交易哈希: 
        <a :href="`${EXPLORER_URL}/tx/${claimTxHash}`" target="_blank">
          {{ claimTxHash.slice(0, 10) }}...{{ claimTxHash.slice(-8) }}
        </a>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { ethers } from 'ethers'
import type { Market } from '../types'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS || '0x3f2E7c68Ec5fC4E6419C5922a444C0e20Db13b60'
const RPC_URL = import.meta.env.VITE_RPC_URL || 'https://ethereum-sepolia-rpc.publicnode.com'
const EXPLORER_URL = import.meta.env.VITE_EXPLORER_URL || 'https://sepolia.etherscan.io'

const CONTRACT_ABI = [
  'function getBet(uint256 marketId, uint256 index) view returns (address user, bytes encryptedBet, uint256 amount, uint8 option, bool claimed)',
  'function marketTotalPools(uint256) view returns (uint256)',
  'function markets(uint256) view returns (uint256 id, string description, uint8 state, uint48 createTimeStamp, uint48 closeTime, uint48 resolveWindow, uint48 resolveTimeStamp, uint8 outcome, bool resolved)',
  'function claimWinnings(uint256 marketId, uint256 betIndex)',
  'event BetPlaced(uint256 indexed marketId, address indexed user, uint256 amount, bytes encryptedBet)',
]

const props = defineProps<{
  market: Market
  account: string | null
}>()

const emit = defineEmits<{
  'claimed': []
}>()

interface UserBet {
  index: number
  user: string
  amount: string
  option: number
  claimed: boolean
}

const userBets = ref<UserBet[]>([])
const loading = ref(false)
const isClaiming = ref(false)
const claimTxHash = ref<string | null>(null)

async function loadUserBets() {
  if (!props.account) return

  try {
    loading.value = true
    userBets.value = []

    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)

    // 使用 Map 来去重，key 是 encryptedBet 的哈希
    const betMap = new Map<string, UserBet>()
    
    // 尝试读取下注，直到遇到错误为止
    let index = 0
    let consecutiveErrors = 0
    const maxConsecutiveErrors = 3
    
    while (consecutiveErrors < maxConsecutiveErrors) {
      try {
        const bet = await contract.getBet(props.market.id, index)
        consecutiveErrors = 0 // 重置错误计数
        
        // 检查是否是当前用户的下注
        if (bet.user.toLowerCase() === props.account.toLowerCase()) {
          // 使用 encryptedBet 作为唯一标识
          const betKey = bet.encryptedBet
          
          if (!betMap.has(betKey)) {
            // 解码 encryptedBet 获取真实的 option
            const abiCoder = ethers.AbiCoder.defaultAbiCoder()
            const decoded = abiCoder.decode(['uint8', 'uint256'], bet.encryptedBet)
            const decodedOption = Number(decoded[0])
            
            betMap.set(betKey, {
              index: index,
              user: bet.user,
              amount: ethers.formatEther(bet.amount),
              option: decodedOption,
              claimed: bet.claimed
            })
            
            console.log(`下注 #${index}: 选项=${decodedOption}, 金额=${ethers.formatEther(bet.amount)}, 已领取=${bet.claimed}`)
          } else {
            console.log(`跳过重复下注 #${index}`)
          }
        }
        
        index++
      } catch (err: any) {
        // 如果是 "Index out of bounds" 错误，说明已经读完了
        if (err.message?.includes('Index out of bounds')) {
          console.log(`已读取所有下注，总数: ${index}`)
          break
        }
        consecutiveErrors++
        index++
      }
    }

    // 将 Map 转换为数组
    userBets.value = Array.from(betMap.values())
    console.log(`找到 ${userBets.value.length} 个用户下注`)
  } catch (err) {
    console.error('加载用户下注失败:', err)
  } finally {
    loading.value = false
  }
}

async function getTotalBets(contract: ethers.Contract, marketId: number): Promise<number> {
  try {
    // 尝试通过事件计数
    const provider = contract.runner?.provider
    if (provider) {
      const currentBlock = await provider.getBlockNumber()
      const fromBlock = Math.max(0, currentBlock - 10000)
      const filter = contract.filters.BetPlaced(marketId)
      const events = await contract.queryFilter(filter, fromBlock, currentBlock)
      return events.length
    }
  } catch (err) {
    console.error('获取下注总数失败:', err)
  }
  return 0
}

function canClaim(bet: UserBet): boolean {
  return props.market.resolved && 
         bet.option === props.market.outcome && 
         !bet.claimed
}

function getBetClass(bet: UserBet): string {
  if (!props.market.resolved) return ''
  if (bet.option === props.market.outcome) {
    return bet.claimed ? 'bet-claimed' : 'bet-winning'
  }
  return 'bet-losing'
}

function getStatusClass(bet: UserBet): string {
  if (!props.market.resolved) return 'status-pending'
  if (bet.option === props.market.outcome) {
    return bet.claimed ? 'status-claimed' : 'status-winning'
  }
  return 'status-losing'
}

function getStatusText(bet: UserBet): string {
  if (!props.market.resolved) return '待结算'
  if (bet.option === props.market.outcome) {
    return bet.claimed ? '已领取' : '可领取'
  }
  return '未中奖'
}

function calculateReward(bet: UserBet): string {
  if (!props.market.resolved || bet.option !== props.market.outcome) {
    return '0'
  }

  // 简化计算：实际奖励需要从合约获取 optionTotals
  // 这里只是估算
  const totalPool = parseFloat(props.market.totalPool)
  const betAmount = parseFloat(bet.amount)
  
  // 假设获胜池是总池的一半（实际需要从合约读取）
  const estimatedReward = (betAmount * totalPool) / (totalPool / 2)
  
  return estimatedReward.toFixed(6)
}

function formatAmount(amount: string): string {
  return parseFloat(amount).toFixed(6)
}

async function claimWinnings(betIndex: number) {
  if (!props.account) {
    alert('请先连接钱包')
    return
  }

  try {
    isClaiming.value = true
    claimTxHash.value = null

    console.log('开始领取奖励...')
    console.log('市场ID:', props.market.id)
    console.log('下注索引:', betIndex)

    // 直接使用钱包的 provider，不使用 JsonRpcProvider
    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)

    // 先验证下注信息
    const bet = await contract.getBet(props.market.id, betIndex)
    console.log('下注信息:', {
      user: bet.user,
      amount: ethers.formatEther(bet.amount),
      option: bet.option,
      claimed: bet.claimed
    })

    // 解码获取真实选项
    const abiCoder = ethers.AbiCoder.defaultAbiCoder()
    const decoded = abiCoder.decode(['uint8', 'uint256'], bet.encryptedBet)
    const decodedOption = Number(decoded[0])
    
    console.log('解码后的选项:', decodedOption)
    console.log('市场结果:', props.market.outcome)

    // 检查是否需要先 reveal
    if (bet.option === 0 && decodedOption !== 0) {
      alert('⚠️ 管理员还未揭示下注数据\n\n请联系管理员使用 AdminRevealBets 工具调用 revealBets() 函数后再领取奖励。\n\n技术说明：合约中 bet.option 当前为 0（未揭示状态），需要管理员解密并更新后才能领取。')
      return
    }
    
    if (bet.option === 0 && decodedOption === 0) {
      // 两者都是 0，可能已经 reveal 过了，或者真的下注的是 0
      console.log('bet.option 和 decodedOption 都是 0，继续尝试领取')
    }

    const tx = await contract.claimWinnings(props.market.id, betIndex)
    console.log('领取交易已提交:', tx.hash)
    claimTxHash.value = tx.hash
    
    // 等待交易确认，增加超时保护
    const receipt = await Promise.race([
      tx.wait(),
      new Promise((_, reject) => 
        setTimeout(() => reject(new Error('交易确认超时')), 60000)
      )
    ])
    
    console.log('交易已确认！', receipt)
    
    alert('奖励领取成功！')
    
    // 重新加载下注列表
    await loadUserBets()
    emit('claimed')
  } catch (err: any) {
    console.error('领取奖励失败:', err)
    
    let errorMessage = '领取失败'
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Market not resolved')) {
      errorMessage = '市场尚未结算'
    } else if (err.message?.includes('Not your bet')) {
      errorMessage = '这不是您的下注'
    } else if (err.message?.includes('Already claimed')) {
      errorMessage = '已经领取过了'
    } else if (err.message?.includes('Not winning option')) {
      errorMessage = '您的下注未中奖'
    } else if (err.message?.includes('Failed to fetch')) {
      errorMessage = 'RPC 连接失败，请检查网络或稍后重试'
    } else if (err.message?.includes('交易确认超时')) {
      errorMessage = '交易已提交但确认超时，请在区块浏览器查看交易状态'
    } else {
      errorMessage = `领取失败: ${err.message || '未知错误'}`
    }
    
    alert(errorMessage)
  } finally {
    isClaiming.value = false
  }
}

watch(() => props.account, () => {
  if (props.account) {
    loadUserBets()
  } else {
    userBets.value = []
  }
}, { immediate: true })

watch(() => props.market.resolved, () => {
  if (props.market.resolved && props.account) {
    loadUserBets()
  }
})

onMounted(() => {
  if (props.account) {
    loadUserBets()
  }
})
</script>

<style scoped>
.claim-section {
  border-top: 2px solid #e5e7eb;
  padding-top: 32px;
  margin-top: 32px;
}

.claim-section h3 {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 20px;
}

.connect-prompt,
.empty-message {
  text-align: center;
  padding: 40px 20px;
  color: #6b7280;
  font-size: 14px;
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  padding: 40px 20px;
  color: #6b7280;
}

.spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #e5e7eb;
  border-top-color: #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.bets-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.bet-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background: #f9fafb;
  border-radius: 12px;
  border: 2px solid #e5e7eb;
  transition: all 0.2s;
}

.bet-item.bet-winning {
  background: #d1fae5;
  border-color: #10b981;
}

.bet-item.bet-claimed {
  background: #e5e7eb;
  border-color: #9ca3af;
}

.bet-item.bet-losing {
  background: #fee2e2;
  border-color: #ef4444;
}

.bet-info {
  flex: 1;
}

.bet-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.bet-label {
  font-size: 14px;
  font-weight: 600;
  color: #374151;
}

.bet-status {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.status-pending {
  background: #fef3c7;
  color: #92400e;
}

.status-winning {
  background: #10b981;
  color: white;
}

.status-claimed {
  background: #6b7280;
  color: white;
}

.status-losing {
  background: #ef4444;
  color: white;
}

.bet-details {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 12px;
}

.detail-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.detail-item .label {
  font-size: 12px;
  color: #6b7280;
}

.detail-item .value {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.detail-item .value.reward {
  color: #10b981;
  font-size: 16px;
}

.claim-btn {
  padding: 12px 24px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.claim-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(16, 185, 129, 0.4);
}

.claim-btn:disabled {
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

@media (max-width: 768px) {
  .bet-item {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }
  
  .claim-btn {
    width: 100%;
  }
  
  .bet-details {
    grid-template-columns: 1fr;
  }
}
</style>
