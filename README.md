# Confidential-Prediction-Market-with-Chainlink-CC

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Chainlink Hackathon](https://img.shields.io/badge/Chainlink-Hackathon-375BD2)
![Go Report Card](https://goreportcard.com/badge/github.com/2042458785/Confidential-Prediction-Market-with-Chainlink-CC)

Confidential Prediction Market with Chainlink CRE

A private, automated prediction market built with Chainlink Runtime Environment (CRE). Users place encrypted bets on real-world events (e.g., weather, sports, finance). A Go‑based CRE workflow listens for on‑chain requests, fetches external data from APIs, determines the outcome, cryptographically signs the result, and submits it back to the blockchain – all without manual intervention.

🏆 Built for the Chainlink Hackathon

Youtube: https://www.youtube.com/watch?v=YuDRxk5sFs8

✨ Key Features

🔐 Privacy by Design – Bets are encrypted on‑chain using a temporary AES key; actual choices stay hidden until market resolution.
🤖 Fully Automated Settlement – Chainlink CRE workflow triggered by SettlementRequested event, executes external API calls, signs results, and writes to chain.
🔗 Multi‑Capability Orchestration – Integrates CRE’s Log Trigger, HTTP Capability, Sign Capability, and EVM Write Capability in a single Go workflow.
📊 Real‑World Data – Fetches temperature from public weather APIs (Open‑Meteo) – easily replaceable with any REST API.
⚡ Scalable & Extensible – Designed to support multiple outcome options, ZK‑proof decryption (future), and cross‑chain with CCIP.

## 🧠 Why This Matters

Existing prediction markets like Polymarket suffer from:

- Public bets → front‑running, copycat strategies.
- Manual settlement → slow, prone to manipulation.
- Opaque outcomes → low user trust.

Our solution eliminates these issues by:

- Encrypting bets to protect user strategy.
- Automating settlement via Chainlink CRE – fast, tamper‑proof, verifiable.
- Trusting cryptographic signatures from a decentralized oracle network instead of human judgment.

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Smart Contract (Sepolia)                  │
│  - Markets, encrypted bets, result storage, payouts          │
│  - Emits SettlementRequested event                           │
└───────────────────────────────┬─────────────────────────────┘
                                │
                                │ listens
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Go CRE Workflow (DON)                     │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐  │
│  │ Log Trigger  │ -> │ HTTP Cap     │ -> │ Sign Cap     │  │
│  └──────────────┘    └──────────────┘    └──────────────┘  │
│                           │                                  │
│                           ▼                                  │
│                    ┌──────────────┐                          │
│                    │ EVM Write Cap│ -> submitResult()       │
│                    └──────────────┘                          │
└───────────────────────────────┬─────────────────────────────┘
                                │
                                │ writes signed result
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Smart Contract (on‑chain)                 │
│                    Market resolved, payouts ready            │
└─────────────────────────────────────────────────────────────┘
```

## 🧩 Chainlink Components Used

| Capability | Description | Code Location |
|-----------|-------------|---------------|
| Log Trigger | Listens for SettlementRequested events | `workflow/main.go#L48-L52` |
| HTTP Capability | Calls geocoding & weather APIs | `workflow/main.go#L124-L153`<br>`workflow/main.go#L164-L192` |
| Sign Capability | Signs the outcome report with CRE DON key | `workflow/main.go#L215-L223` |
| EVM Write Capability | Submits signed result to submitResult() | `workflow/main.go#L235` |
| Consensus Aggregation | Combines multiple DON responses for API calls | `workflow/main.go#L153`<br>`workflow/main.go#L192` |

## 🚀 Quick Start

### Prerequisites

- Node.js v18+ & npm
- Go v1.21+
- Hardhat
- CRE CLI (for simulation/deployment)
- Sepolia testnet ETH & LINK (from faucets.chain.link)

### 1. Clone & Install

```bash
git clone https://github.com/2042458785/Confidential-Prediction-Market-with-Chainlink-CC
cd confidential-prediction-market
npm install
go mod download
```

### 2. Set Environment Variables

```bash
PRIVATE_KEY=your_wallet_private_key
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/your_project_id
CONTRACT_ADDRESS=0x...          # will be set after deployment
CHAIN_SELECTOR=16015286601757825753  # Sepolia
WEATHER_API_URL=https://api.open-meteo.com/v1/forecast
```

### 3. Deploy Smart Contract

```bash
npx hardhat run scripts/deploy.js --network sepolia
```

### 4. Run the CRE Workflow (Simulation)

```bash
cd workflow
cre workflow simulate --file main.go --config config.json
```

### 5. Example config.json:

```json
{
  "contractAddress": "0x...",
  "chainSelector": 16015286601757825753,
  "apiConfig": {
    "url": "https://api.open-meteo.com/v1/forecast",
    "method": "GET",
    "headers": {}
  },
  "gasLimit": 500000
}
```

### 6. (Optional) Deploy CRE Workflow

```bash
cre workflow deploy --file main.go --name "prediction-market-settlement" --config config.json
```
## 📝 How to Use

### Create a Market (Admin)

Call `createMarket(description, closeTime, resolveWindow)` on the contract.

Example: `"2025-03-10:London"` – date and city name separated by colon.

### Place Encrypted Bets (Users)

Use a temporary AES key to encrypt `{option, amount}`. Call `placeBet(marketId, encryptedBet)` with the encrypted blob and send the bet amount in ETH.

### Request Settlement (Anyone after deadline)

Call `requestSettlement(marketId)`. This emits the event that triggers the CRE workflow.

### CRE Workflow in Action

1. Log Trigger captures the event.
2. HTTP Capability geocodes the city and fetches current temperature.
3. Outcome determined (temperature >5°C → 1, else 0).
4. Sign Capability signs the outcome.
5. EVM Write Capability submits `submitResult(marketId, outcome, signature)`.

### Claim Winnings (Winners)

Winners call `claimWinnings(marketId, betIndex)`. Contract verifies their bet matched the outcome and sends proportional reward.
## 🧪 Testing

Run Solidity tests:

```bash
npx hardhat test
```

Run Go workflow simulation (as shown above).

## 🔮 Future Roadmap

- **Zero‑Knowledge Decryption** – Replace admin‑reveal with user‑submitted ZK proofs using DECO, eliminating trusted parties.
- **Multi‑Outcome Markets** – Support for more than binary outcomes (e.g., exact score, price range).
- **Cross‑Chain Liquidity** – Use CCIP to allow betting from multiple chains and unify settlement.
- **AI‑Powered Events** – Integrate with Gemini or other LLMs for complex event resolution (e.g., "Did Team A play better?").

🤝 Contributing

Contributions are welcome! Please open an issue or pull request.

📄 License

This project is licensed under the MIT License – see the LICENSE file for details.

## 🙏 Acknowledgements

- Chainlink for the CRE SDK and hackathon opportunity.
- Open‑Meteo for free weather & geocoding APIs.
- OpenZeppelin for secure contract libraries.

---

**Built with ❤️ for the Chainlink Hackathon Spring 2026**

---

# Prediction Market Frontend

A prediction market frontend application in Polymarket style.

## Features

- 🎨 Modern UI design, similar to Polymarket
- 📊 Market list display
- 💰 Market details and betting functionality
- 🔗 MetaMask wallet connection
- 📱 Responsive design, mobile friendly

## Tech Stack

- Vue 3 + TypeScript
- Vite
- Ethers.js v6
- CSS3

## Installation and Running

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build
```

## Configuration

Before using, modify the following configurations:

### 1. Contract Address

Modify in `src/App.vue`:

```typescript
const CONTRACT_ADDRESS = '0xYourContractAddress'
```

### 2. RPC Provider

Modify in the `loadMarkets` function in `src/App.vue`:

```typescript
const provider = new ethers.JsonRpcProvider('https://sepolia.infura.io/v3/YOUR_INFURA_KEY')
```

You can use:
- Infura: https://infura.io/
- Alchemy: https://www.alchemy.com/
- or other RPC providers

## Project Structure

```
src/
├── App.vue                    # Main application component with wallet connection
├── main.ts                    # Application entry
├── types.ts                   # TypeScript type definitions
├── style.css                  # Global styles
└── components/
    ├── MarketList.vue         # Market list with refresh functionality
    ├── MarketCard.vue         # Individual market card display
    ├── MarketDetail.vue       # Market details, betting, and settlement requests
    ├── ClaimWinnings.vue      # Winnings claim interface for resolved markets
    ├── DecodeTools.vue        # Utility to decode encrypted bet data
    ├── AdminRevealBets.vue    # Admin tool to batch decrypt and reveal bets
    ├── AdminSubmitResult.vue  # Admin panel for trusted signer setup and result submission
    └── TestSubmitResult.vue   # Testing tool for submitResult functionality
```

### Component Details

**App.vue** - Main application container that manages wallet connection (MetaMask), view routing between market list and detail views, and market data loading from the smart contract.

**MarketList.vue** - Displays all available markets in a grid layout with refresh functionality. Includes embedded admin tools (DecodeTools, AdminRevealBets, AdminSubmitResult) for market management.

**MarketCard.vue** - Individual market card showing market ID, status badge (Open/Closed/Resolving/Resolved), description, total pool, bet count, and close time. Clickable to view details.

**MarketDetail.vue** - Comprehensive market view with betting interface (option selection, amount input), settlement request button for closed markets, and integrated ClaimWinnings component for resolved markets.

**ClaimWinnings.vue** - Displays user's bets for a resolved market with win/loss status, estimated rewards, and claim buttons. Automatically loads user bets and handles the claimWinnings transaction.

**DecodeTools.vue** - Developer utility to manually decode encrypted bet data (encryptedBet bytes) into readable option and amount values using ABI decoding.

**AdminRevealBets.vue** - Admin interface to load all bets for a market, decode them locally, and batch submit revealed data to the contract via revealBets() function. Required before users can claim winnings.

**AdminSubmitResult.vue** - Admin panel with two functions: (1) Set trusted signer address for result submission authorization, (2) Manually submit market results with signature verification. Includes auto-discovery of markets in "Resolving" state.

**TestSubmitResult.vue** - Testing interface for submitResult functionality. Shows current configuration (contract address, trusted signer, wallet match status), allows manual result submission with signature generation, and displays detailed signing process steps.

## Usage Instructions

1. Click the "Connect Wallet" button in the top right corner to connect MetaMask
2. Browse the market list to view all available prediction markets
3. Click on a market card to enter the detail page
4. Select betting option (Yes/No)
5. Enter bet amount
6. Click "Confirm Bet" to submit the transaction

## Notes

- Ensure MetaMask browser extension is installed
- Ensure your wallet is connected to the correct network (Sepolia testnet)
- Ensure your wallet has sufficient ETH for betting and gas fees
- The current version's encrypted betting data is a simplified implementation; for production environments, real encryption logic should be implemented

Live URL: https://market-vue-nine.vercel.app/