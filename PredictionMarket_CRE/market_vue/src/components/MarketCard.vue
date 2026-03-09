<template>
  <div class="market-card" @click="$emit('click')">
    <div class="card-header">
      <span class="market-id">#{{ market.id }}</span>
      <span class="status-badge" :class="statusClass">
        {{ statusText }}
      </span>
    </div>

    <h3 class="market-title">{{ market.description }}</h3>

    <div class="market-stats">
      <div class="stat">
        <span class="stat-label">总池</span>
        <span class="stat-value">{{ formatPool(market.totalPool) }} ETH</span>
      </div>
      
      <div class="stat">
        <span class="stat-label">下注数</span>
        <span class="stat-value">{{ market.betsCount || 0 }}</span>
      </div>
      
      <div class="stat">
        <span class="stat-label">截止时间</span>
        <span class="stat-value">{{ formatDate(market.closeTime) }}</span>
      </div>
    </div>

    <div v-if="market.resolved" class="outcome">
      <span class="outcome-label">结果:</span>
      <span class="outcome-value">{{ market.outcome === 1 ? '是' : '否' }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Market } from '../types'
import { MarketStatus } from '../types'

const props = defineProps<{
  market: Market
}>()

defineEmits<{
  click: []
}>()

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

function formatPool(pool: string): string {
  // 直接显示完整数值，不限制小数位数
  return pool
}

function formatDate(timestamp: number): string {
  const date = new Date(timestamp * 1000)
  return date.toLocaleDateString('zh-CN', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}
</script>

<style scoped>
.market-card {
  background: white;
  border-radius: 16px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.market-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.market-id {
  font-size: 12px;
  font-weight: 600;
  color: #667eea;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
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

.market-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 16px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.market-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 12px;
}

@media (max-width: 768px) {
  .market-stats {
    grid-template-columns: 1fr;
  }
}

.stat {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-label {
  font-size: 12px;
  color: #6b7280;
}

.stat-value {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.outcome {
  padding-top: 12px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  gap: 8px;
}

.outcome-label {
  font-size: 12px;
  color: #6b7280;
}

.outcome-value {
  font-size: 14px;
  font-weight: 600;
  color: #667eea;
}
</style>
