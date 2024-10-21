/*========================>database cart <===================================*/
CREATE DATABASE cart;
USE cart;

CREATE TABLE `cart` (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '购物车id',
                        `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
                        `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
                        `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '数量',
                        `checked` int(11) NOT NULL DEFAULT 0 COMMENT '是否选择,1=已勾选,0=未勾选',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        KEY `ix_userid` (`userid`),
                        KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';
