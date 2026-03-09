<template>
  <div class="app">
    <Header :account="account" @connect="connectWallet" @disconnect="disconnectWallet" />
    
    <main class="container">
      <MarketList 
        v-if="currentView === 'list'"
        :markets="markets"
        :loading="loadingMarkets"
        :account="account"
        @select-market="selectMarket"
        @refresh="loadMarkets"
      />
      
      <MarketDetail
        v-else-if="currentView === 'detail' && selectedMarket"
        :market="selectedMarket"
        :account="account"
        @back="currentView = 'list'"
        @bet-placed="handleBetPlaced"
      />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ethers } from 'ethers'
import Header from './components/Header.vue'
import MarketList from './components/MarketList.vue'
import MarketDetail from './components/MarketDetail.vue'
import type { Market } from './types'

const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS
const RPC_URL = import.meta.env.VITE_RPC_URL || 'https://sepolia.infura.io/v3/YOUR_INFURA_KEY'

const CONTRACT_ABI = [
  'function markets(uint256) view returns (uint256 id, string description, uint8 state, uint48 createTimeStamp, uint48 closeTime, uint48 resolveWindow, uint48 resolveTimeStamp, uint8 outcome, bool resolved)',
  'function OrderMarketid() view returns (uint256)',
  'function marketTotalPools(uint256) view returns (uint256)',
  'function placeBet(uint256 marketId, bytes encryptedBet) payable',
  'event BetPlaced(uint256 indexed marketId, address indexed user, uint256 amount, bytes encryptedBet)',
]

const account = ref<string | null>(null)
const markets = ref<Market[]>([])
const loadingMarkets = ref(false)
const currentView = ref<'list' | 'detail'>('list')
const selectedMarket = ref<Market | null>(null)

async function connectWallet() {
  if (!(window as any).ethereum) {
    alert('请先安装 MetaMask')
    return
  }

  try {
    const [addr] = await (window as any).ethereum.request({
      method: 'eth_requestAccounts',
    })
    account.value = addr
  } catch (err) {
    console.error(err)
    alert('连接钱包失败')
  }
}

function disconnectWallet() {
  account.value = null
  alert('已断开钱包连接')
}

async function loadMarkets() {
  loadingMarkets.value = true
  try {
    const provider = new ethers.JsonRpcProvider(RPC_URL)
    const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider)
    
    console.log('开始加载市场...')
    const marketCount = await contract.OrderMarketid()
    console.log('市场总数:', Number(marketCount))
    
    const marketPromises = []
    
    for (let i = 1; i <= Number(marketCount); i++) {
      marketPromises.push(loadMarket(contract, i))
    }
    
    markets.value = (await Promise.all(marketPromises)).filter(m => m !== null) as Market[]
    console.log('加载完成，市场数量:', markets.value.length)
  } catch (err) {
    console.error('加载市场失败:', err)
    alert('加载市场失败，请检查网络连接或刷新页面重试')
  } finally {
    loadingMarkets.value = false
  }
}

async function loadMarket(contract: ethers.Contract, marketId: number): Promise<Market | null> {
  try {
    console.log(`开始加载市场 ${marketId}...`)
    
    const [id, description, state, createTimeStamp, closeTime, resolveWindow, resolveTimeStamp, outcome, resolved] = await contract.markets(marketId)
    const totalPool = await contract.marketTotalPools(marketId)
    
    console.log(`市场 ${marketId} 基本信息加载完成`)
    
    // 通过事件统计下注数量（添加超时保护）
    let betsCount = 0
    try {
      const provider = contract.runner?.provider
      if (provider) {
        const currentBlock = await provider.getBlockNumber()
        // 只查询最近 5000 个区块（减少查询范围，加快速度）
        const fromBlock = Math.max(0, currentBlock - 5000)
        
        console.log(`查询市场 ${marketId} 事件，区块范围: ${fromBlock} - ${currentBlock}`)
        
        // 添加超时保护
        const filter = contract.filters.BetPlaced(marketId)
        const eventsPromise = contract.queryFilter(filter, fromBlock, currentBlock)
        const timeoutPromise = new Promise((_, reject) => 
          setTimeout(() => reject(new Error('查询超时')), 10000) // 10秒超时
        )
        
        const events = await Promise.race([eventsPromise, timeoutPromise]) as any[]
        betsCount = events.length
        console.log(`市场 ${marketId} 找到 ${betsCount} 个下注`)
      }
    } catch (err) {
      console.warn(`统计市场 ${marketId} 下注数量失败（将显示为0）:`, err)
      // 失败时设置为 0，不影响其他数据加载
      betsCount = 0
    }
    
    return {
      id: Number(id),
      description,
      state: Number(state),
      createTimeStamp: Number(createTimeStamp),
      closeTime: Number(closeTime),
      resolveWindow: Number(resolveWindow),
      resolveTimeStamp: Number(resolveTimeStamp),
      outcome: Number(outcome),
      resolved,
      totalPool: ethers.formatEther(totalPool),
      betsCount,
    }
  } catch (err) {
    console.error(`加载市场 ${marketId} 失败:`, err)
    return null
  }
}

function selectMarket(market: Market) {
  selectedMarket.value = market
  currentView.value = 'detail'
}

function handleBetPlaced() {
  loadMarkets()
}

onMounted(() => {
  loadMarkets()
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.app {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 24px 16px;
}

@media (max-width: 768px) {
  .container {
    padding: 16px 12px;
  }
}
</style>
