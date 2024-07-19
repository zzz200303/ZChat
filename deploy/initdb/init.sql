CREATE DATABASE IF NOT EXISTS zchat;
USE zchat;


CREATE TABLE `users` (
                         `id` varchar(24) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                         `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;