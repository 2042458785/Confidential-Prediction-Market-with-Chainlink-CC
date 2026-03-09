<template>
  <div class="market-list">
    <div class="list-header">
      <h2>所有市场</h2>
      <button class="refresh-btn" @click="$emit('refresh')" :disabled="loading">
        <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
          <path d="M4 2v6h6M16 18v-6h-6M17.65 6.35A8 8 0 1 0 6.35 17.65"/>
        </svg>
        刷新
      </button>
    </div>

    <div v-if="loading" class="loading">
      <div class="spinner"></div>
      <p>加载市场中...</p>
    </div>

    <div v-else-if="markets.length === 0" class="empty">
      <p>暂无市场</p>
    </div>

    <div v-else class="market-grid">
      <MarketCard
        v-for="market in markets"
        :key="market.id"
        :market="market"
        @click="$emit('select-market', market)"
      />
    </div>
    
    <!-- 解码工具 -->
    <DecodeTools />
    
    <!-- 管理员批量解密工具 -->
    <AdminRevealBets />
    
    <!-- 管理员提交结算结果 -->
    <AdminSubmitResult :account="account" />
  </div>
</template>

<script setup lang="ts">
import MarketCard from './MarketCard.vue'
import DecodeTools from './DecodeTools.vue'
import AdminRevealBets from './AdminRevealBets.vue'
import AdminSubmitResult from './AdminSubmitResult.vue'
import type { Market } from '../types'

defineProps<{
  markets: Market[]
  loading: boolean
  account: string | null
}>()

defineEmits<{
  'select-market': [market: Market]
  refresh: []
}>()
</script>

<style scoped>
.market-list {
  padding: 24px 0;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.list-header h2 {
  font-size: 32px;
  font-weight: 700;
  color: white;
}

.refresh-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.refresh-btn:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.3);
}

.refresh-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  color: white;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.empty {
  text-align: center;
  padding: 60px 20px;
  color: white;
  font-size: 18px;
}

.market-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
}

@media (max-width: 768px) {
  .list-header h2 {
    font-size: 24px;
  }
  
  .market-grid {
    grid-template-columns: 1fr;
  }
}
</style>
