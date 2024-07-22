CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,    -- 主键，自增
                         `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
