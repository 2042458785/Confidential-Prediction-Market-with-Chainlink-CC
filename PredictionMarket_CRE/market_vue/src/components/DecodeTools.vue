<template>
  <div class="decode-tools">
    <div class="tools-card">
      <h3>🔓 解码工具</h3>
      <p class="description">解码下注数据（encryptedBet）</p>
      
      <div class="input-group">
        <label>编码后的数据 (Hex)</label>
        <textarea
          v-model="encodedData"
          placeholder="0x0000000000000000000000000000000000000000000000000000000000000000..."
          rows="3"
        ></textarea>
      </div>
      
      <button class="decode-btn" @click="decodeData" :disabled="!encodedData">
        解码
      </button>
      
      <div v-if="decodedResult" class="result">
        <h4>解码结果：</h4>
        <div class="result-item">
          <span class="label">选项 (option):</span>
          <span class="value">{{ decodedResult.option }}</span>
        </div>
        <div class="result-item">
          <span class="label">金额 (amount):</span>
          <span class="value">{{ decodedResult.amount }} ETH</span>
        </div>
      </div>
      
      <div v-if="decodeError" class="error">
        ❌ {{ decodeError }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ethers } from 'ethers'

const encodedData = ref('')
const decodedResult = ref<{ option: number; amount: string } | null>(null)
const decodeError = ref('')

function decodeData() {
  decodedResult.value = null
  decodeError.value = ''
  
  try {
    if (!encodedData.value.trim()) {
      decodeError.value = '请输入编码数据'
      return
    }
    
    // 确保数据以 0x 开头
    let data = encodedData.value.trim()
    if (!data.startsWith('0x')) {
      data = '0x' + data
    }
    
    // 使用 ABI 解码
    const abiCoder = ethers.AbiCoder.defaultAbiCoder()
    const decoded = abiCoder.decode(['uint8', 'uint256'], data)
    
    decodedResult.value = {
      option: Number(decoded[0]),
      amount: ethers.formatEther(decoded[1])
    }
  } catch (err: any) {
    console.error('解码失败:', err)
    decodeError.value = `解码失败: ${err.message || '数据格式不正确'}`
  }
}
</script>

<style scoped>
.decode-tools {
  margin-top: 24px;
}

.tools-card {
  background: white;
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.tools-card h3 {
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

.input-group textarea {
  width: 100%;
  padding: 12px;
  border: 2px solid #e5e7eb;
  border-radius: 8px;
  font-size: 14px;
  font-family: 'Courier New', monospace;
  resize: vertical;
  transition: all 0.2s;
}

.input-group textarea:focus {
  outline: none;
  border-color: #667eea;
}

.decode-btn {
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

.decode-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.decode-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.result {
  margin-top: 20px;
  padding: 16px;
  background: #f0fdf4;
  border: 2px solid #86efac;
  border-radius: 8px;
}

.result h4 {
  font-size: 14px;
  font-weight: 600;
  color: #166534;
  margin-bottom: 12px;
}

.result-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #bbf7d0;
}

.result-item:last-child {
  border-bottom: none;
}

.result-item .label {
  font-size: 13px;
  color: #166534;
  font-weight: 500;
}

.result-item .value {
  font-size: 14px;
  font-weight: 700;
  color: #15803d;
  font-family: 'Courier New', monospace;
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
</style>
