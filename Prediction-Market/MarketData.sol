// SPDX-License-Identifier:MIT
pragma solidity ^0.8.34;

contract MarketData{

    enum MarketStatus { Created,Open,Closed,Resolving,Resolved,Canceled }

    struct Bet{
        address user;  //下注用户
        bytes encryptedBet; //加密后的下注数据 amount和option 也就是下注金额和选项
        uint256 amount; //用户下注金额 明文存储 但是需要与加密后的下注数据中的金额一致
        uint8 option; //用户的选项(解密后才会有) ---结算后才会填充 默认初始为0
        bool claimed; //用户是否已经领取了获胜奖励 默认为false
    }

    struct Market {
        uint256 id; //此Market的id
        string description; //描述
        MarketStatus state; //Market状态
        uint48 createTimeStamp; //创建Market时间戳
        uint48 closeTime; //下注截止时间戳
        uint48 resolveWindow; //解析时间窗口
        uint48 resolveTimeStamp; //解析时间戳 在这个时间戳之前才可以进行解析
        uint256[] optionTotals; //各选项总金额(解密之后才可以进行统计)
        Bet[] bets; //所有的下注
        bytes32 functionsRequest_id;  //最近一次发送的 Chainlink Functions 请求的 ID 用于在回调时关联市场
        string functionsSource; //IPFS的CID或者直接的JS代码
        bytes encryptedSecrets; //通过公钥加密的API密钥 通过后续的chainlink的Don的私钥来进行解析
        uint8 outcome; //调用API之后返回的最终结果
        bool resolved; //是否已经成功结算
    }

    mapping (uint256 id => uint256 totalPool) public marketTotalPools; //该市场的所有下注金额 每次下注时累加

    mapping (uint256 market_id => Market) public markets; //所有的Market
    uint256 public OrderMarketid; //总Market数量(也作为生成Market的序号)
    address internal  functionsRouter; //Chainlink Functions 路由器的地址 每个网络（如 Sepolia）有固定的路由器地址，在构造函数中设置 用于向此地址发送 Functions 请求
    uint64 internal  functionsSubscriptionid; //Chainlink Functions 的订阅 ID 订阅需要提前创建并充值 LINK，用于支付 Functions 请求的费用 在构造函数中设置

}