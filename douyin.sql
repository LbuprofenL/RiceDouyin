-- Active: 1690553765068@@127.0.0.1@3306@douyin

DROP TABLE IF EXISTS `user`;

CREATE TABLE
    `user` (
        `id` BIGINT UNSIGNED not NULL COMMENT '主键id',
        `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
        `password` varchar(16) NOT NULL DEFAULT '' COMMENT '用户密码',
        `create_time` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
        `work_count` integer NOT NULL DEFAULT 0 COMMENT '作品数量',
        `favorite_count` integer NOT NULL DEFAULT 0 COMMENT '获赞数量',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

CREATE TABLE
    `video` (
        `id` BIGINT UNSIGNED not NULL COMMENT '主键id',
        `user_id` BIGINT UNSIGNED not NULL COMMENT '作者id',
        `title` varchar(32) NOT NULL DEFAULT '' COMMENT '视频标题',
        `favorite_count` integer NOT NULL DEFAULT 0 COMMENT '视频获赞数量',
        `comment_count` integer NOT NULL DEFAULT 0 COMMENT '评论数量',
        `cover_url` varchar(255) NOT NULL DEFAULT '' COMMENT '封面路径',
        `video_url` varchar(255) NOT NULL DEFAULT '' COMMENT '视频路径',
        `create_time` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
        PRIMARY KEY (`id`),
        INDEX idk_video (`create_time` DESC)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '视频表';

ALTER TABLE `video`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);