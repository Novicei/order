# 项目需求

## 1.模拟数据源

提供一个假的数据源，可以源源不断的推送用户订单数据，用户订单数据结构如下：

```json
{
    "orderId": int // 订单id
    "orderTime": int // 下单时间
    "skuId": int // 商品id
    "userId": int // 用户id
    "status": int // 订单状态 (0待付款,1已付款,2取消)
    "price": double  // 价格
    "pay": double // 支付金额
}
```

要求：

1. 数据源可以针对同一订单 (orderId 相同的订单) 提供一条或多条数据，来表征同一订单的订单状态变更

## 2.订单处理服务

设计一个服务来消费用户订单数据，我们对这个服务的要求是：

\* 支持通过用户id获取用户订单历史

\* 支持按skuId查询从下单到实际付款概率

\* 支持按时段获得成交量最高的sku top10

# 项目架构

分为client端和server端，其中client端主要负责模拟发送订单请求;server端主要负责对订单请求进行落库并进行相应的功能实现，二者之间通过Kafka进行交互。

# 项目启动

1.在测试机器上部署mysql和kafka环境

2.安装go开发环境，sdk 1.15.14；设置GOPROXY=https://goproxy.cn；go mod tidy

3.修改client/conf文件夹和server/conf文件夹下的config.ini文件，将其更换成测试机的环境配置

4.运行server文件夹下main.go文件;运行client文件夹下main.go文件

此时，client端不断的mock数据到数据库当中

5.运行server/mysql文件夹下code_test.go文件中的测试函数，即可测试相应需求

# 项目回顾

因为开发时间只用了几个小时，所以只是一个demo项目，对于其中的并发请求或性能以及日志输出等都不是很好，只提供借鉴思路。

并且也可以考虑使用redis或者memcache等组件，应该会取得更好的效果。

2021年11月4日10点-2021年11月4日16点