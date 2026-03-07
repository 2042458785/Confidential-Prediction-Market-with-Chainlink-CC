// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

contract MarketData {
    enum MarketStatus { Created, Open, Closed, Resolving, Resolved, Canceled }

    struct Bet {
        address user;
        bytes encryptedBet;      // 加密后的下注数据（金额+选项）
        uint256 amount;          // 实际下注金额（与加密数据一致）
        uint8 option;            // 解密后的选项（结算后填充）
        bool claimed;            // 是否已领取奖励
    }

    struct Market {
        uint256 id;
        string description;
        MarketStatus state;
        uint48 createTimeStamp;
        uint48 closeTime;               // 下注截止时间
        uint48 resolveWindow;            // 解析窗口（秒）
        uint48 resolveTimeStamp;         // 最晚解析时间 = closeTime + resolveWindow
        uint256[] optionTotals;          // 各选项总金额（解密后统计）
        Bet[] bets;
        uint8 outcome;                   // 最终结果
        bool resolved;                   // 是否已结算
    }

    mapping(uint256 => Market) public markets;
    uint256 public OrderMarketid=0;
    mapping(uint256 => uint256) public marketTotalPools; // 市场总下注金额

    // 存储 CRE DON 的公钥地址（可以是多个，但通常是一个聚合签名者）
    address public trustedSigner;
}