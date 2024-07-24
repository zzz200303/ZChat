CREATE TABLE `gmember` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `gid` int NOT NULL,
                         `uid` int NOT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `gmember`
ADD INDEX `gid_uid` (`gid`, `uid`);