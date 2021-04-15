CREATE TABLE `full_table_scan_test`
(
    `id`            int(11)   NOT NULL AUTO_INCREMENT,
    `account`       varchar(50)    DEFAULT NULL,
    `client_type`   tinyint(4)     DEFAULT NULL,
    `security_code` varchar(50)    DEFAULT NULL,
    `create_at`     timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_three` (`account`, `security_code`, `create_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;