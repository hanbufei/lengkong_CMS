DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `header`        varchar(100) NOT NULL COMMENT '大标题',
    `class_id`      int(10) DEFAULT NULL COMMENT '类别ID',
    `label_id_list` varchar(100) NOT NULL COMMENT '标签id列表',
    `create_at`     datetime DEFAULT NULL COMMENT '创建时间',
    `path`          text DEFAULT NULL COMMENT '文章路径',
    `last_img`      text DEFAULT NULL COMMENT '最新图片',
    `detail`        varchar(300) NOT NULL COMMENT '文章描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB3;