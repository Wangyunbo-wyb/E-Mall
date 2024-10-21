/*========================>database reply <===================================*/
CREATE DATABASE reply;
USE reply;

CREATE TABLE `reply`(
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论表id',
                        `business` varchar(64) NOT NULL DEFAULT '' COMMENT '评论业务类型',
                        `targetid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论目标id',
                        `reply_userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复用户id',
                        `be_reply_userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '被回复用户id',
                        `parentid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论id',
                        `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
                        `image` varchar(255) NOT NULL DEFAULT '' COMMENT '评论图片',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        KEY `ix_targetid` (`targetid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论列表';