CREATE TABLE `record`
(
    `id`        BIGINT                          NOT NULL AUTO_INCREMENT,
    `content`   TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
    `from`      BIGINT                          NOT NULL,
    `to`        BIGINT                          NOT NULL,
    `type`      BIGINT                          NOT NULL,
    `send_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;