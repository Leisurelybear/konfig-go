CREATE TABLE `config` (
                          `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '配置集合自增id',
                          `collection_id` int(11) NOT NULL COMMENT '对应配置集合中的自增ID',
                          `name` varchar(255) DEFAULT NULL COMMENT '配置名称，为了用户识别',
                          `key` varchar(255) NOT NULL COMMENT '配置key',
                          `value` text DEFAULT NULL COMMENT '存放配置内容',
                          `version` tinyint(2) NOT NULL DEFAULT 0 COMMENT '0=草稿，1=生产环境,，准备弃用，只对集粒度控制',
                          `updated_by` varchar(40) DEFAULT NULL COMMENT '更新者用户名',
                          `created_by` varchar(40) DEFAULT NULL COMMENT '创建者用户名',
                          `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                          `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                          `deleted_at` datetime DEFAULT NULL COMMENT '0存在，1已经删除',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `collection` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `name` varchar(255) DEFAULT NULL,
                              `version` int(11) NOT NULL DEFAULT 0,
                              `updated_by` varchar(40) DEFAULT NULL,
                              `created_by` varchar(40) DEFAULT NULL,
                              `updated_at` datetime DEFAULT NULL,
                              `created_at` datetime DEFAULT NULL,
                              `deleted_at` datetime DEFAULT '0000-00-00 00:00:00',
                              `desc` varchar(255) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;