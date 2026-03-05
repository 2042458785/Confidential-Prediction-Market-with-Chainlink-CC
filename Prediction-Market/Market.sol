// SPDX-License-Identifier:MIT
pragma solidity ^0.8.34;

import { MarketData } from "Prediction-Market/MarketData.sol";

contract ConfidentialMarket is MarketData{
    //合约创建
    address public owner; //定义合约部署者 也就是创建人 作为超级管理员 可设置/删除管理员 

    uint48 internal createtimestamp; //定义创建时的时间戳

    constructor(){
        owner=msg.sender; //初始化合约部署者 也就是创建人
        createtimestamp=uint48(block.timestamp); //获取合约创建时 记录此时区块的时间
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
        require(sender==owner,"You not owner!!!"); //非超级管理员没有权限设置/删除管理员
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


}
