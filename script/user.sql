CREATE DATABASE `cms_account` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `account` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` varchar(64) DEFAULT '' COMMENT '用户名',
    `password` varchar(64) DEFAULT '' COMMENT '密码',
    `nickname` varchar(64) DEFAULT '' COMMENT '昵称',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';