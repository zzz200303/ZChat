CREATE DATABASE IF NOT EXISTS zchat CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE zchat;

CREATE TABLE `recodes` (
                           `id` BIGINT NOT NULL AUTO_INCREMENT,
                           `content` TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
                           `from` BIGINT NOT NULL,
                           `to` BIGINT NOT NULL,
                           `send_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,    -- 主键，自增
                         `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `wuid` (
                        `h` int(10) NOT NULL AUTO_INCREMENT,
                        `x` tinyint(4) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`x`),
                        UNIQUE KEY `h` (`h`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;