# Confidential-Prediction-Market-with-Chainlink-CC
Confidential Prediction Market with Chainlink CRE

https://img.shields.io/badge/License-MIT-yellow.svg
https://img.shields.io/badge/Chainlink-Hackathon-375BD2
https://goreportcard.com/badge/github.com/yourname/confidential-prediction-market

A private, automated prediction market built with Chainlink Runtime Environment (CRE). Users place encrypted bets on real-world events (e.g., weather, sports, finance). A Go‑based CRE workflow listens for on‑chain requests, fetches external data from APIs, determines the outcome, cryptographically signs the result, and submits it back to the blockchain – all without manual intervention.

🏆 Built for the Chainlink Hackathon

✨ Key Features

🔐 Privacy by Design – Bets are encrypted on‑chain using a temporary AES key; actual choices stay hidden until market resolution.
🤖 Fully Automated Settlement – Chainlink CRE workflow triggered by SettlementRequested event, executes external API calls, signs results, and writes to chain.
🔗 Multi‑Capability Orchestration – Integrates CRE’s Log Trigger, HTTP Capability, Sign Capability, and EVM Write Capability in a single Go workflow.
📊 Real‑World Data – Fetches temperature from public weather APIs (Open‑Meteo) – easily replaceable with any REST API.
⚡ Scalable & Extensible – Designed to support multiple outcome options, ZK‑proof decryption (future), and cross‑chain with CCIP.

🧠 Why This Matters

Existing prediction markets like Polymarket suffer from:

Public bets → front‑running, copycat strategies.
Manual settlement → slow, prone to manipulation.
Opaque outcomes → low user trust.
Our solution eliminates these issues by:

Encrypting bets to protect user strategy.
Automating settlement via Chainlink CRE – fast, tamper‑proof, verifiable.
Trusting cryptographic signatures from a decentralized oracle network instead of human judgment.

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

🧩 Chainlink Components Used

Capability	Description	Code Location
Log Trigger	Listens for SettlementRequested events	workflow/main.go#L48-L52
HTTP Capability	Calls geocoding & weather APIs	workflow/main.go#L124-L153
workflow/main.go#L164-L192
Sign Capability	Signs the outcome report with CRE DON key	workflow/main.go#L215-L223
EVM Write Capability	Submits signed result to submitResult()	workflow/main.go#L235
Consensus Aggregation	Combines multiple DON responses for API calls	workflow/main.go#L153
workflow/main.go#L192

🚀 Quick Start

Prerequisites

Node.js v18+ & npm
Go v1.21+
Hardhat
CRE CLI (for simulation/deployment)
Sepolia testnet ETH & LINK (from faucets.chain.link)

1. Clone & Install
   git clone https://github.com/2042458785/Confidential-Prediction-Market-with-Chainlink-CC
   cd confidential-prediction-market
   npm install
   go mod download
   2. Set Environment Variables
   PRIVATE_KEY=your_wallet_private_key
   SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/your_project_id
   CONTRACT_ADDRESS=0x...          # will be set after deployment
   CHAIN_SELECTOR=16015286601757825753  # Sepolia
   WEATHER_API_URL=https://api.open-meteo.com/v1/forecast
3.  Deploy Smart Contract
    npx hardhat run scripts/deploy.js --network sepolia
4. Run the CRE Workflow (Simulation)
   cd workflow
   cre workflow simulate --file main.go --config config.json
5. Example config.json:
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
6. (Optional) Deploy CRE Workflow
   cre workflow deploy --file main.go --name "prediction-market-settlement" --config config.json
   7. 📝 How to Use
      Create a Market (Admin)
      Call createMarket(description, closeTime, resolveWindow) on the contract.
      Example: "2025-03-10:London" – date and city name separated by colon.
      Place Encrypted Bets (Users)
      Use a temporary AES key to encrypt {option, amount}. Call placeBet(marketId, encryptedBet) with the encrypted blob and send the bet amount in ETH.
      Request Settlement (Anyone after deadline)
      Call requestSettlement(marketId). This emits the event that triggers the CRE workflow.
      CRE Workflow in Action

       Log Trigger captures the event.
   HTTP Capability geocodes the city and fetches current temperature.
   Outcome determined (temperature >5°C → 1, else 0).
   Sign Capability signs the outcome.
   EVM Write Capability submits submitResult(marketId, outcome, signature).
   Claim Winnings (Winners)
   Winners call claimWinnings(marketId, betIndex). Contract verifies their bet matched the outcome and sends proportional reward.
8. 🧪 Testing
   Run Solidity tests:
   npx hardhat test
   Run Go workflow simulation (as shown above).

🔮 Future Roadmap

Zero‑Knowledge Decryption – Replace admin‑reveal with user‑submitted ZK proofs using DECO, eliminating trusted parties.
Multi‑Outcome Markets – Support for more than binary outcomes (e.g., exact score, price range).
Cross‑Chain Liquidity – Use CCIP to allow betting from multiple chains and unify settlement.
AI‑Powered Events – Integrate with Gemini or other LLMs for complex event resolution (e.g., "Did Team A play better?").

🤝 Contributing

Contributions are welcome! Please open an issue or pull request.

📄 License

This project is licensed under the MIT License – see the LICENSE file for details.

🙏 Acknowledgements

Chainlink for the CRE SDK and hackathon opportunity.
Open‑Meteo for free weather & geocoding APIs.
OpenZeppelin for secure contract libraries.

Built with ❤️ for the Chainlink Hackathon Spring 2026