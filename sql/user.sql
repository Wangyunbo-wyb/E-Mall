/*========================>database user <===================================*/
CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                        `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                        `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
                        `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
                        `question` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码问题',
                        `answer` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码答案',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uniq_phone` (`phone`),
                        UNIQUE KEY `uniq_username` (`username`),
                        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `user_receive_address` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                        `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
                                        `name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人名称',
                                        `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
                                        `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否为默认地址',
                                        `post_code` varchar(100) NOT NULL DEFAULT '' COMMENT '邮政编码',
                                        `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份/直辖市',
                                        `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
                                        `region` varchar(100) NOT NULL DEFAULT '' COMMENT '区',
                                        `detail_address` varchar(128) NOT NULL DEFAULT '' COMMENT '详细地址(街道)',
                                        `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
                                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
                                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
                                        PRIMARY KEY (`id`),
                                        KEY `idx_uid` (`uid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户收货地址表';

CREATE TABLE `user_collection` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '收藏Id',
                                   `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
                                   `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
                                   `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
                                   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
                                   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `UN_collection_uid_product_id`(uid,product_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏表';
