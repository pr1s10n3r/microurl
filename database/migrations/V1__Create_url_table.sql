CREATE TABLE `url` (
    `id`         INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `code`       CHAR(6)          NOT NULL UNIQUE,
    `value`      VARCHAR(1024)    NOT NULL,
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);