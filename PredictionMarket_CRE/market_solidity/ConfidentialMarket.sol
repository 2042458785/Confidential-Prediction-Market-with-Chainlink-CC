// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "./MarketData.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ConfidentialMarket is MarketData, ReentrancyGuard, Ownable {

    // 管理员映射
    mapping(address => bool) public admins;

    event MarketCreated(uint256 indexed marketId, string description, uint48 closeTime);
    event BetPlaced(uint256 indexed marketId, address indexed user, uint256 amount, bytes encryptedBet);
    event SettlementRequested(uint256 indexed marketId);  // 触发 CRE 工作流的事件
    event MarketResolved(uint256 indexed marketId, uint8 outcome);
    event WinningsClaimed(uint256 indexed marketId, address indexed user, uint256 amount);

    modifier onlyAdmin() {
        require(admins[msg.sender] || msg.sender == owner(), "Not admin");
        _;
    }

    constructor() Ownable(msg.sender) {}

    // 设置管理员
    function setAdmin(address admin, bool status) external onlyOwner {
        admins[admin] = status;
    }

    // 创建市场
    function createMarket(
        string memory description,
        uint48 closeTime,
        uint48 resolveWindow
    ) external onlyAdmin returns (uint256 marketId) {
        require(closeTime > block.timestamp, "Close time must be future");
        require(resolveWindow > 0, "Resolve window required");

        marketId = OrderMarketid++;
        Market storage market = markets[marketId];
        market.id = marketId;
        market.description = description;
        market.state = MarketStatus.Open;
        market.createTimeStamp = uint48(block.timestamp);
        market.closeTime = closeTime;
        market.resolveWindow = resolveWindow;
        market.resolveTimeStamp = closeTime + resolveWindow;
        // 初始化选项总额数组（假设二元期权，可根据需要动态调整）
        market.optionTotals = new uint256[](2);
        market.resolved = false;

        emit MarketCreated(marketId, description, closeTime);
    }

    // 用户下注
    function placeBet(uint256 marketId, bytes calldata encryptedBet) external payable {
        Market storage market = markets[marketId];
        require(market.state == MarketStatus.Open, "Market not open");
        require(block.timestamp < market.closeTime, "Betting closed");
        require(msg.value > 0, "Amount must be >0");

        market.bets.push(Bet({
            user: msg.sender,
            encryptedBet: encryptedBet,
            amount: msg.value,
            option: 0,
            claimed: false
        }));

        // 累加总池（明文记录，但选项总额需解密后统计）
        marketTotalPools[marketId] += msg.value;

        emit BetPlaced(marketId, msg.sender, msg.value, encryptedBet);
    }

    // 任何人（或前端）调用以请求结算，触发事件
    function requestSettlement(uint256 marketId) external {
        Market storage market = markets[marketId];
        require(market.state == MarketStatus.Open || market.state == MarketStatus.Closed, "Invalid state");
        require(block.timestamp >= market.closeTime, "Betting not closed");
        require(block.timestamp <= market.resolveTimeStamp, "Resolve window expired");

        market.state = MarketStatus.Resolving;  // 可选，标记正在解析
        emit SettlementRequested(marketId);
    }

    // 由 CRE 工作流调用的结果提交函数（需验证签名）
    function submitResult(
        uint256 marketId,
        uint8 outcome,
        bytes memory signature
    ) external {
        // 验证签名来自可信的 CRE DON（此处简化，仅验证调用者是预先设置的 CRE 地址）
        // 实际可使用 EIP-712 或简单验证调用者白名单
        require(msg.sender == trustedCREAddress, "Invalid caller"); // 需预先设置 trustedCREAddress

        Market storage market = markets[marketId];
        require(market.state == MarketStatus.Resolving, "Not resolving");

        market.outcome = outcome;
        market.resolved = true;
        market.state = MarketStatus.Resolved;

        emit MarketResolved(marketId, outcome);
    }

    // 设置可信的 CRE 地址（由 owner 设置）
    address public trustedCREAddress;
    function setTrustedCRE(address _addr) external onlyOwner {
        trustedCREAddress = _addr;
    }

    // 管理员批量解密下注数据（或在领取时由用户自行解密并验证，本合约保留此函数）
    function revealBets(
        uint256 marketId,
        uint256[] calldata betIndices,
        uint8[] calldata options,
        uint256[] calldata amounts
    ) external onlyAdmin {
        Market storage market = markets[marketId];
        require(market.resolved, "Market not resolved");

        for (uint i = 0; i < betIndices.length; i++) {
            Bet storage bet = market.bets[betIndices[i]];
            require(bet.option == 0, "Already revealed");
            bet.option = options[i];
            bet.amount = amounts[i];
            market.optionTotals[options[i]] += amounts[i];
        }
    }

    // 用户领取奖励
    function claimWinnings(uint256 marketId, uint256 betIndex) external nonReentrant {
        Market storage market = markets[marketId];
        Bet storage bet = market.bets[betIndex];

        require(market.resolved, "Market not resolved");
        require(bet.user == msg.sender, "Not your bet");
        require(!bet.claimed, "Already claimed");
        require(bet.option == market.outcome, "Not winning option");

        uint256 totalPool = marketTotalPools[marketId];
        uint256 winningPool = market.optionTotals[market.outcome];
        require(winningPool > 0, "Winning pool zero");

        uint256 reward = (bet.amount * totalPool) / winningPool;

        bet.claimed = true;

        (bool success, ) = payable(msg.sender).call{value: reward}("");
        require(success, "Transfer failed");

        emit WinningsClaimed(marketId, msg.sender, reward);
    }
}