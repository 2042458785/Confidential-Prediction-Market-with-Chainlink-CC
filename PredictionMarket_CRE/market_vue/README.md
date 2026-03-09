# 预测市场前端

类似 Polymarket 风格的预测市场前端应用。

## 功能特性

- 🎨 现代化的 UI 设计，类似 Polymarket
- 📊 市场列表展示
- 💰 市场详情和下注功能
- 🔗 MetaMask 钱包连接
- 📱 响应式设计，支持移动端

## 技术栈

- Vue 3 + TypeScript
- Vite
- Ethers.js v6
- CSS3

## 安装和运行

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build
```

## 配置

在使用前，请修改以下配置：

### 1. 合约地址

在 `src/App.vue` 中修改：

```typescript
const CONTRACT_ADDRESS = '0x你的合约地址'
```

### 2. RPC 提供商

在 `src/App.vue` 的 `loadMarkets` 函数中修改：

```typescript
const provider = new ethers.JsonRpcProvider('https://sepolia.infura.io/v3/YOUR_INFURA_KEY')
```

你可以使用：
- Infura: https://infura.io/
- Alchemy: https://www.alchemy.com/
- 或其他 RPC 提供商

## 项目结构

```
src/
├── App.vue                    # 主应用组件
├── main.ts                    # 应用入口
├── types.ts                   # TypeScript 类型定义
├── style.css                  # 全局样式
└── components/
    ├── Header.vue             # 头部组件（钱包连接）
    ├── MarketList.vue         # 市场列表组件
    ├── MarketCard.vue         # 市场卡片组件
    └── MarketDetail.vue       # 市场详情组件
```

## 使用说明

1. 点击右上角"连接钱包"按钮连接 MetaMask
2. 浏览市场列表，查看所有可用的预测市场
3. 点击市场卡片进入详情页
4. 选择下注选项（是/否）
5. 输入下注金额
6. 点击"确认下注"提交交易

## 注意事项

- 确保已安装 MetaMask 浏览器扩展
- 确保钱包连接到正确的网络（Sepolia 测试网）
- 确保钱包有足够的 ETH 用于下注和 gas 费用
- 当前版本的加密下注数据是简化实现，生产环境需要实现真正的加密逻辑

## 开发计划

- [ ] 添加市场搜索和筛选功能
- [ ] 添加用户下注历史
- [ ] 添加实时价格更新
- [ ] 添加图表展示
- [ ] 优化移动端体验
- [ ] 添加多语言支持
