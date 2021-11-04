DROP TABLE IF EXISTS `orderT`;
CREATE TABLE `orderT`
(
    `orderID`   varchar(30) NOT NULL COMMENT '订单ID',
    `orderTime` varchar(30) NOT NULL COMMENT '订单时间',
    `skuId`     varchar(30) NOT NULL COMMENT '商品ID',
    `userId`    varchar(30) NOT NULL COMMENT '用户ID',
    `status`    varchar(30) NOT NULL COMMENT '订单状态',
    `price`     float NOT NULL COMMENT '商品价格',
    `pay`       float NOT NULL COMMENT '支付',
    PRIMARY KEY (`orderID`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';