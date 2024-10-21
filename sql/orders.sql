/*========================>database orders <===================================*/
CREATE DATABASE orders;
USE orders;

CREATE TABLE `orders` (
                          `id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                          `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                          `shoppingid` bigint(20) NOT NULL DEFAULT 0 COMMENT '收货信息表id',
                          `payment` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT '实际付款金额,单位是元,保留两位小数',
                          `paymenttype` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付类型,1-在线支付',
                          `postage` int(10)  NOT NULL DEFAULT 0 COMMENT '运费,单位是元',
                          `status` smallint(6) NOT NULL DEFAULT 10 COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
                          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`),
                          KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE `orderitem` (
                             `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单子表id',
                             `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                             `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                             `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
                             `product_name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
                             `product_image` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片地址',
                             `current_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '生成订单时的商品单价，单位是元,保留两位小数',
                             `quantity` int(10) NOT NULL DEFAULT 0 COMMENT '商品数量',
                             `total_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品总价,单位是元,保留两位小数',
                             `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             KEY `ix_orderid` (`order_id`),
                             KEY `ix_userid` (`user_id`),
                             KEY `ix_proid` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';

CREATE TABLE `shipping` (
                            `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '收货信息表id',
                            `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
                            `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                            `receiver_name` varchar(20) NOT NULL DEFAULT '' COMMENT '收货姓名',
                            `receiver_phone` varchar(20) NOT NULL DEFAULT '' COMMENT '收货固定电话',
                            `receiver_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '收货移动电话',
                            `receiver_province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
                            `receiver_city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
                            `receiver_district` varchar(20) NOT NULL DEFAULT '' COMMENT '区/县',
                            `receiver_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`),
                            KEY `ix_orderid` (`orderid`),
                            KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货信息表';
