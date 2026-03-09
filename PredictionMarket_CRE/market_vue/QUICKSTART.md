# 快速启动指南

## 第一步：安装依赖

```bash
cd PredictionMarket_CRE/market_vue
npm install
```

## 第二步：配置环境变量（可选）

复制 `.env.example` 为 `.env` 并修改配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件：

```env
VITE_CONTRACT_ADDRESS=你的合约地址
VITE_RPC_URL=https://sepolia.infura.io/v3/你的INFURA_KEY
VITE_EXPLORER_URL=https://sepolia.etherscan.io
```

如果不配置环境变量，应用会使用默认值。

## 第三步：启动开发服务器

```bash
npm run dev
```

应用将在 http://localhost:5173 启动

## 第四步：连接钱包

1. 确保已安装 MetaMask 浏览器扩展
2. 切换到 Sepolia 测试网络
3. 点击页面右上角"连接钱包"按钮
4. 在 MetaMask 中确认连接

## 第五步：使用应用

### 浏览市场
- 主页显示所有可用的预测市场
- 每个市场卡片显示：市场 ID、描述、状态、总池、截止时间
- 点击"刷新"按钮更新市场数据

### 参与下注
1. 点击任意市场卡片进入详情页
2. 查看市场详细信息
3. 选择下注选项（选项 0 或选项 1）
4. 输入下注金额（ETH）
5. 点击"确认下注"
6. 在 MetaMask 中确认交易
7. 等待交易确认

### 查看交易
- 下注成功后，页面会显示交易哈希
- 点击交易哈希可在区块链浏览器中查看详情

## 常见问题

### Q: 无法连接钱包？
A: 确保已安装 MetaMask 并刷新页面

### Q: 交易失败？
A: 检查以下几点：
- 钱包是否有足够的 ETH
- 是否连接到正确的网络（Sepolia）
- 市场是否仍在开放状态
- 查看浏览器控制台的错误信息

### Q: 看不到市场数据？
A: 检查：
- RPC URL 是否正确配置
- 合约地址是否正确
- 网络连接是否正常
- 打开浏览器控制台查看错误信息

### Q: 如何获取测试 ETH？
A: 访问 Sepolia 水龙头：
- https://sepoliafaucet.com/
- https://www.alchemy.com/faucets/ethereum-sepolia

## 生产部署

```bash
# 构建生产版本
npm run build

# 构建产物在 dist/ 目录
# 可以部署到任何静态网站托管服务
# 如：Vercel, Netlify, GitHub Pages 等
```

## 技术支持

如有问题，请查看：
- 浏览器控制台的错误信息
- MetaMask 的交易历史
- 区块链浏览器的交易详情
