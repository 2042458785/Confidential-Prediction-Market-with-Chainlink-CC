// SPDX-License-Identifier:MIT
pragma solidity ^0.8.34;

import {FunctionsClient} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/FunctionsClient.sol";

import {FunctionsRequest} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/libraries/FunctionsRequest.sol";

import {AutomationCompatibleInterface} from "@chainlink/contracts/src/v0.8/automation/AutomationCompatible.sol";

import {ConfirmedOwner} from "@chainlink/contracts/src/v0.8/shared/access/ConfirmedOwner.sol";

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";// 重入锁保护，用于资金提取等敏感操作
import "@openzeppelin/contracts/utils/Strings.sol"; // 将 uint256 转换为字符串（用于构建 Functions 参数）

import { MarketData } from "./MarketData.sol";

contract ConfidentialMarket is 
FunctionsClient,AutomationCompatibleInterface,ConfirmedOwner,ReentrancyGuard,MarketData{
    //初始化外部参数
    using FunctionsRequest for FunctionsRequest.Request;
    using Strings for uint256;

    //合约创建
    //父合约已完成
    //address public owner; //定义合约部署者 也就是创建人 作为超级管理员 可设置/删除管理员 

    uint48 internal createtimestamp; //定义创建时的时间戳

    constructor(address _functionsRouter,uint64 _subscription_id)
    FunctionsClient(_functionsRouter)
    ConfirmedOwner(msg.sender) //调用ConfirmedOwner构造函数，初始化所有者为部署者
    {
        //父合约已完成
        //owner=msg.sender; //初始化合约部署者 也就是创建人
        createtimestamp=uint48(block.timestamp); //获取合约创建时 记录此时区块的时间
        //初始化 Chainlink Functions 路由器的地址 和 Chainlink Functions 的订阅 ID
        functionsRouter=_functionsRouter; //Chainlink Functions 路由器的地址 每个网络如Sepolia 有固定的路由器地址 在构造函数中设置用于发送 Functions 请求
        functionsSubscriptionid=_subscription_id; //Chainlink Functions 的订阅 ID 订阅需要提前创建并充值 LINK 用于支付 Functions 请求的费用 在构造函数中设置
    }

    //管理员设置/删除

    mapping (address id =>bool isadmin) internal admin_list;

    //超级管理员设置管理员
    function setAdmin(address _admin)
    public 
    ifSuperAdmin(msg.sender){
        //设置管理员
        admin_list[_admin]=true;
    }

    //超级管理员删除管理员
    function deleteAdmin(address _admin)
    public 
    ifSuperAdmin(msg.sender)
    ifAdmin(_admin){
       //删除管理员
       admin_list[_admin]=false;
    }

    modifier ifSuperAdmin(address sender){
        require(sender==owner(),"You not owner!!!"); //非超级管理员没有权限设置/删除管理员
        _;
    }

    modifier ifAdmin(address admin_addr){
        require(admin_list[admin_addr]==true,"You not admin!!!");//检验此用户是否已经是管理员 只有已经是管理员才能删除此管理员
        _;
    }

    //管理员创建预测事件
    function createEvent(
        string memory description, //事件描述
        uint48 deadlinetimestamp, //下注截止时间戳
        uint48 analysiswindowtime, //截止后多久内必须完成解析(在这个时间内一直不断尝试解析也就是发请求去获得一个结果)
        string memory _functionsSource, //js源码
        bytes memory _encryptedSecrets //js源码中的API密钥(在本地加密后传入合约)
    )
    external
    ifAdmin(msg.sender)
    returns(uint256 market_id)
    {
        require(block.timestamp<deadlinetimestamp,"Deadline timestamp must be later!!!"); //截止时间必须在当前时间之后
        require(analysiswindowtime>0,"Resolve time can't equal zero!!!"); //时间窗口不能为0

        market_id=OrderMarketid++;//生成market_id 从1开始 序号id即为第几个Market
        Market storage market = markets[market_id]; //创建market 并且以此作为引用

        market.id=market_id; 
        market.description=description;
        market.state=MarketStatus.Open;
        market.createTimeStamp=uint48(block.timestamp);
        market.closeTime=deadlinetimestamp;
        market.resolveWindow=analysiswindowtime;
        market.resolveTimeStamp=market.closeTime+market.resolveWindow;
        market.functionsSource=_functionsSource;
        market.encryptedSecrets=_encryptedSecrets;
        //完成状态默认是false 初始化的时候已经写入
    }

    //用户下注
    function placeBet(
    uint256 _market_id,
    bytes memory _encryptedBet) 
    external 
    payable  {
        //先找出对应的市场
        Market storage market=markets[_market_id];
        require(market.state==MarketStatus.Open,"Market not open!!!"); //市场不是开启状态 
        require(block.timestamp<market.closeTime,"Betting closed!!!"); //检测在下注截止之前进行提交
        require(msg.value>0,"Must bigger than 0"); //转入金额必须大于0
        //用户下注
        market.bets.push(Bet({
        user: msg.sender,  //下注用户
        encryptedBet: _encryptedBet,//加密后的下注数据 amount和option 也就是下注金额和选项
        amount: msg.value, //用户下注金额 明文存储 但是需要与加密后的下注数据中的金额一致
        option: 0, //用户的选项(解密后才会有) ---结算后才会填充
        claimed:false  //用户是否已经领取了获胜奖励 目前还没有领取奖励
        }));

        //用户下注成功
    }

    //触发Functions
    function resolveMarket(uint256 market_id) 
    external 
    returns 
    (bytes32 request_id){
        //先找出对应的市场
        Market storage market = markets[market_id];
        require(market.state==MarketStatus.Open || market.state==MarketStatus.Closed,"Can't be resolving"); //验证状态 必须为没有解析过的 触发这个解析函数才能进行解析
        //时间状态在投注截止之后和解析时间之前可以触发这个Chainlink的Functions函数
        require(block.timestamp>=market.closeTime,"Betting not closed"); //下注截止之后
        require(block.timestamp<=market.resolveTimeStamp,"Timeout for resolving"); //解析时间戳之前

        //更新市场状态
        market.state=MarketStatus.Resolving;

        //设置请求参数
        string[] memory args=new string[](2); //定义一个存储2个元素的数组
        args[0]=Strings.toString(market_id); //第一个元素
        args[1]=market.description; //第二个元素

        //构建Functions请求
        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(market.functionsSource); //初始化请求
        req.setArgs(args); //添加参数
        req.addSecretsReference(market.encryptedSecrets); //添加加密的密钥引用

        //发送请求
        bytes32 _requestId=_sendRequest(
            req.encodeCBOR(),
            functionsSubscriptionid,
            300000,
        ""); //发送请求之后会返回一个id方便回调 也就是我们的_requestId

        market.functionsRequest_id=_requestId;//我们把_requestid存储起来 方便在后续回调fulfillRequest函数 中能通过该id找到对应的市场

        //返回request_id
        return _requestId;
    }

    function fulfillRequest(
    bytes32 _requestId,
    bytes memory _response,
    bytes memory _err) 
    internal 
    override { //这个函数是留给Functions来回调的
        //先通过这个_requestId来获取对应的market的id
        uint256 market_id=requestIdGetMarketId(_requestId);
        Market storage market=markets[market_id];

        require(market.state==MarketStatus.Resolving,"Only resolving market!!!"); //只能是正在解析的market

        if (_err.length>0){ //需要满足没有错误
            //失败处理:错误回退 回退到closed
            market.state=MarketStatus.Closed;
            return;
        }

        //没有错误 解析response里面的返回结果
        uint8 outcome=abi.decode(_response,(uint8));

        //获得结果进入内部结算
        _settleMarket(market_id, outcome);
    }

    function requestIdGetMarketId(bytes32 inId)internal view returns(uint256 market_id) { //通过Chainlink Functions的回调传入id来获得市场的id
        for (uint256 i=1;i<=OrderMarketid;i++){
            if (markets[i].functionsRequest_id==inId){
                return i; //返回对应的market的id
            }
        }
        revert("Market not found"); //走到这一步说明没有找到 直接revert弹出提示信息,回退交易
    }

    function _settleMarket(uint256 _marketId,uint8 market_outcome)internal{ //解析完成 结果写入 更新状态
        Market storage market=markets[_marketId]; //获取对应的market
        market.outcome=market_outcome; //写入合约获得的结果
        market.resolved=true; //已经解决这个market的resolve
        market.state=MarketStatus.Resolved; //更新状态为Resolved
    }

    //管理员批量解密后传入对应的数据
    function revealBets(
        uint256 market_id,
        uint256[] calldata _betIndices, //对应的所有需要解密的下注索引 管理员可以自由选择 但是传入的这个索引必须是1开始 因为我上面也是这样规定的
        uint8[] calldata _options,
        uint256[] calldata _amounts
    )
    external 
    ifAdmin(msg.sender){
        Market storage market =markets[market_id];
        require(market.state==MarketStatus.Resolved,"Only resolved market!");//确保market已经解析完成
        
        //验证完成 开始统计
        for (uint256 i=0;i<_betIndices.length;i++){
            Bet storage bet=market.bets[_betIndices[i]]; //从第一个bet开始 也就是索引为0
            require(bet.option==0,"Already revealed!!!");//这个bet需要未解密
            bet.option=_options[i]; //获取这个索引的选项
            bet.amount=_amounts[i]; //获取这个索引的金额
            market.optionTotals[_options[i]]+=_amounts[i];
        }

        //计算下注总金额
        for (uint256 i=0;i<market.optionTotals.length;i++){
            marketTotalPools[market_id]+=market.optionTotals[i];
        }
    }

    //用户领取奖励
    function claimWinnings(uint256 _marketId,uint256 _betIndex)external nonReentrant{
        Market storage market =markets[_marketId]; //获取market
        Bet storage bet =market.bets[_betIndex];//获取bet
        
        require(market.state==MarketStatus.Resolved,"Must be resolved!");//必须是已经结算的
        require(bet.user==msg.sender,"Not your bet");//这个bet必须是本人的
        require(!bet.claimed,"Already claimed!");//必须没有被领取
        require(bet.option==market.outcome,"Not winning option!");//必须和结算结果相同

        //计算奖励 下注正确的人 根据比例分配整个下注的池子
        uint256 totalPool= marketTotalPools[_marketId]; //这个market的总池
        uint256 winningPool =market.optionTotals[market.outcome];//获胜的总池
        uint256 reward=(bet.amount/winningPool) * totalPool;

        //先发送奖励
        bool sendRewardSuccess;
        (sendRewardSuccess,)=payable(bet.user).call{value:reward}("");

        require(sendRewardSuccess,"Transfer failed!"); //确保转账成功再改变用户提取状态

        //发送成功 改变用户领取状态
        bet.claimed=true;
    }

    //使用Chainlink的定时任务 使合约到时间自动执行
    function checkUpkeep(bytes calldata) //这个函数相当于眼睛 找到需要执行的任务
    external 
    view 
    override 
    returns(bool upkeepNeeded,bytes memory performData){
        for (uint256 i=1; i<=OrderMarketid;i++){ //遍历所有的market
            Market storage market=markets[i];
            if (market.state==MarketStatus.Open && 
            block.timestamp>=market.closeTime && 
            block.timestamp<=market.resolveTimeStamp){ //到了需要结算的时间(大于等于截止时间 小于等于解析时间戳)开始自动结算
                upkeepNeeded=true; //此市场需要执行任务
                performData =abi.encode(i);
                return (true,performData);
            }
        }
        //说明遍历了market 没有找到需要执行自动任务的market
        return (false,bytes(""));
    }

    function performUpkeep(bytes calldata performData) //这个函数相当于手 执行眼睛反馈的需要执行的market
    external
    override{
        uint256 market_id=abi.decode(performData,(uint256)); //先把bytes解码转为uint256
        //通过这个market的id 调用解析函数
        this.resolveMarket(market_id); //因为这个解析函数是external 需要作为外部函数调用 不能直接调用
    }
}