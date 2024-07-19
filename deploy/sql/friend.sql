CREATE TABLE `friend` (
                           `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                           `user_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                           `friend_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                           `created_at` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;