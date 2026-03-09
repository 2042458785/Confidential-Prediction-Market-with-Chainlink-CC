<template>
  <div class="admin-reveal">
    <div class="reveal-card">
      <h3>🔐 管理员批量解密下注</h3>
      <p class="description">查看并解密市场的所有下注数据</p>
      
      <div class="input-group">
        <label>市场 ID</label>
        <input
          v-model.number="marketId"
          type="number"
          placeholder="输入市场 ID"
          min="0"
        />
      </div>
      
      <button class="load-btn" @click="loadBets" :disabled="loading || !marketId">
        {{ loading ? '加载中...' : '加载下注数据' }}
      </button>
      
      <div v-if="error" class="error">
        ❌ {{ error }}
      </div>
      
      <div v-if="bets.length > 0" class="bets-list">
        <h4>下注列表（共 {{ bets.length }} 个）</h4>
        
        <div class="bet-item" v-for="(bet, index) in bets" :key="index">
          <div class="bet-header">
            <span class="bet-index">#{{ index }}</span>
            <span class="bet-user">{{ formatAddress(bet.user) }}</span>
          </div>
          
          <div class="bet-details">
            <div class="detail-row">
              <span class="label">加密数据:</span>
              <span class="value mono">{{ bet.encryptedBet.slice(0, 20) }}...</span>
            </div>
            
            <div class="detail-row">
              <span class="label">金额:</span>
              <span class="value">{{ bet.amount }} ETH</span>
            </div>
            
            <div class="detail-row">
              <span class="label">选项:</span>
              <span class="value" :class="bet.option === 0 ? 'option-no' : 'option-yes'">
                {{ bet.option === 0 ? '否' : bet.option === 1 ? '是' : '未解密' }}
              </span>
            </div>
            
            <div class="detail-row">
              <span class="label">已领取:</span>
              <span class="value">{{ bet.claimed ? '是' : '否' }}</span>
            </div>
            
            <div v-if="bet.decoded" class="decoded-info">
              <div class="detail-row">
                <span class="label">解码选项:</span>
                <span class="value" :class="bet.decoded.option === 0 ? 'option-no' : 'option-yes'">
                  {{ bet.decoded.option === 0 ? '否' : '是' }}
                </span>
              </div>
              <div class="detail-row">
                <span class="label">解码金额:</span>
                <span class="value">{{ bet.decoded.amount }} ETH</span>
              </div>
            </div>
            
            <button 
              v-if="!bet.decoded" 
              class="decode-btn-small" 
              @click="decodeBet(index)"
            >
              解码
            </button>
          </div>
        </div>
        
        <div class="batch-actions">
          <button class="batch-decode-btn" @click="decodeAllBets" :disabled="allDecoded">
            {{ allDecoded ? '全部已解码' : '批量解码全部' }}
          </button>
          
          <button 
            class="reveal-btn" 
            @click="revealBetsOnChain" 
            :disabled="!allDecoded || revealing"
          >
            {{ revealing ? '提交中...' : '📤 提交到链上（revealBets）' }}
          </button>
          
          <p v-if="txHash" class="tx-link">
            交易哈希: 
            <a :href="`${EXPLORER_URL}/tx/${txHash}`" target="_blank">
              {{ txHash.slice(0, 10) }}...{{ txHash.slice(-8) }}
            </a>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ethers } from 'ethers'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS || '0x3f2E7c68Ec5fC4E6419C5922a444C0e20Db13b60'
const RPC_URL = import.meta.env.VITE_RPC_URL || 'https://ethereum-sepolia-rpc.publicnode.com'

const CONTRACT_ABI = [
  'function getBetsCount(uint256 marketId) view returns (uint256)',
  'function getBet(uint256 marketId, uint256 index) view returns (address user, bytes encryptedBet, uint256 amount, uint8 option, bool claimed)',
  'function revealBets(uint256 marketId, uint256[] betIndices, uint8[] options, uint256[] amounts)',
]

interface Bet {
  user: string
  encryptedBet: string
  amount: string
  option: number
  claimed: boolean
  decoded?: {
    option: number
    amount: string
  }
}

const marketId = ref<number>(0)
const bets = ref<Bet[]>([])
const loading = ref(false)
const error = ref('')
const revealing = ref(false)
const txHash = ref<string | null>(null)

const EXPLORER_URL = import.meta.env.VITE_EXPLORER_URL || 'https://sepolia.etherscan.io'

const allDecoded = computed(() => {
  return bets.value.length > 0 && bets.value.every(bet => bet.decoded)
})

async function loadBets() {
  if (!marketId.value && marketId.value !== 0) {
    error.value = '请输入市场 ID'
    return
  }
  
  loading.value = true
  error.value = ''
  bets.value = []
  
  try {
    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)
    
    // 获取下注总数
    const count = await contract.getBetsCount(marketId.value)
    const totalBets = Number(count)
    
    if (totalBets === 0) {
      error.value = '该市场没有下注记录'
      return
    }
    
    // 逐个获取下注数据
    for (let i = 0; i < totalBets; i++) {
      const [user, encryptedBet, amount, option, claimed] = await contract.getBet(marketId.value, i)
      
      bets.value.push({
        user,
        encryptedBet,
        amount: ethers.formatEther(amount),
        option: Number(option),
        claimed
      })
    }
    
  } catch (err: any) {
    console.error('加载下注数据失败:', err)
    error.value = `加载失败: ${err.message || '未知错误'}`
  } finally {
    loading.value = false
  }
}

function decodeBet(index: number) {
  const bet = bets.value[index]
  if (!bet || bet.decoded) return
  
  try {
    const abiCoder = ethers.AbiCoder.defaultAbiCoder()
    const decoded = abiCoder.decode(['uint8', 'uint256'], bet.encryptedBet)
    
    bet.decoded = {
      option: Number(decoded[0]),
      amount: ethers.formatEther(decoded[1])
    }
  } catch (err: any) {
    console.error(`解码下注 ${index} 失败:`, err)
    error.value = `解码失败: ${err.message}`
  }
}

function decodeAllBets() {
  bets.value.forEach((_, index) => {
    if (!bets.value[index].decoded) {
      decodeBet(index)
    }
  })
}

function formatAddress(address: string): string {
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

async function revealBetsOnChain() {
  if (bets.value.length === 0) {
    error.value = '没有下注数据'
    return
  }
  
  // 检查是否所有下注都已解码
  const allDecoded = bets.value.every(bet => bet.decoded)
  if (!allDecoded) {
    error.value = '请先解码所有下注数据'
    return
  }
  
  revealing.value = true
  error.value = ''
  txHash.value = null
  
  try {
    // 连接钱包
    if (!(window as any).ethereum) {
      throw new Error('请先安装 MetaMask')
    }
    
    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const signer = await provider.getSigner()
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, signer)
    
    // 准备数据
    const betIndices: number[] = []
    const options: number[] = []
    const amounts: bigint[] = []
    
    bets.value.forEach((bet, index) => {
      if (bet.decoded && bet.option === 0) { // 只处理未揭示的下注（option === 0）
        betIndices.push(index)
        options.push(bet.decoded.option)
        amounts.push(ethers.parseEther(bet.decoded.amount))
      }
    })
    
    if (betIndices.length === 0) {
      error.value = '所有下注已经揭示，无需再次调用'
      return
    }
    
    console.log('调用 revealBets:', {
      marketId: marketId.value,
      betIndices,
      options,
      amounts: amounts.map(a => ethers.formatEther(a))
    })
    
    // 调用合约
    const tx = await contract.revealBets(marketId.value, betIndices, options, amounts)
    txHash.value = tx.hash
    
    console.log('交易已提交:', tx.hash)
    await tx.wait()
    
    console.log('交易已确认！')
    alert('批量解密成功！')
    
    // 重新加载数据
    await loadBets()
    
  } catch (err: any) {
    console.error('批量解密失败:', err)
    
    let errorMessage = '批量解密失败'
    if (err.code === 'ACTION_REJECTED' || err.code === 4001) {
      errorMessage = '用户取消了交易'
    } else if (err.message?.includes('Not admin')) {
      errorMessage = '只有管理员可以执行此操作'
    } else if (err.message?.includes('Market not resolved')) {
      errorMessage = '市场尚未结算'
    } else if (err.message?.includes('Already revealed')) {
      errorMessage = '部分下注已经揭示'
    } else if (err.reason) {
      errorMessage = `错误: ${err.reason}`
    } else {
      errorMessage = `批量解密失败: ${err.message || '未知错误'}`
    }
    
    error.value = errorMessage
  } finally {
    revealing.value = false
  }
}
</script>

<style scoped>
.admin-reveal {
  margin-top: 24px;
}

.reveal-card {
  background: white;
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.reveal-card h3 {
  font-size: 20px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 8px;
}

.description {
  font-size: 14px;
  color: #6b7280;
  margin-bottom: 20px;
}

.input-group {
  margin-bottom: 16px;
}

.input-group label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: #374151;
  margin-bottom: 8px;
}

.input-group input {
  width: 100%;
  padding: 12px;
  border: 2px solid #e5e7eb;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.2s;
}

.input-group input:focus {
  outline: none;
  border-color: #667eea;
}

.load-btn {
  width: 100%;
  padding: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.load-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.load-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.error {
  margin-top: 16px;
  padding: 12px;
  background: #fef2f2;
  border: 2px solid #fca5a5;
  border-radius: 8px;
  color: #991b1b;
  font-size: 13px;
}

.bets-list {
  margin-top: 24px;
}

.bets-list h4 {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 16px;
}

.bet-item {
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 12px;
}

.bet-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e5e7eb;
}

.bet-index {
  font-size: 14px;
  font-weight: 700;
  color: #667eea;
}

.bet-user {
  font-size: 13px;
  font-family: 'Courier New', monospace;
  color: #6b7280;
}

.bet-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 13px;
}

.detail-row .label {
  color: #6b7280;
  font-weight: 500;
}

.detail-row .value {
  color: #1f2937;
  font-weight: 600;
}

.detail-row .value.mono {
  font-family: 'Courier New', monospace;
  font-size: 12px;
}

.option-yes {
  color: #059669;
}

.option-no {
  color: #dc2626;
}

.decoded-info {
  margin-top: 8px;
  padding-top: 8px;
  border-top: 1px dashed #d1d5db;
}

.decode-btn-small {
  margin-top: 8px;
  padding: 6px 12px;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.decode-btn-small:hover {
  background: #5568d3;
}

.batch-actions {
  margin-top: 16px;
  padding-top: 16px;
  border-top: 2px solid #e5e7eb;
}

.batch-decode-btn {
  width: 100%;
  padding: 12px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.batch-decode-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4);
}

.batch-decode-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.reveal-btn {
  width: 100%;
  padding: 12px;
  margin-top: 12px;
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.reveal-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.4);
}

.reveal-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.tx-link {
  margin-top: 12px;
  font-size: 13px;
  color: #6b7280;
  text-align: center;
}

.tx-link a {
  color: #667eea;
  text-decoration: none;
  font-weight: 600;
}
</style>
