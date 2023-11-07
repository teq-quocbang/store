CREATE TABLE IF NOT EXISTS
  `producer` (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `country` varchar(255) NOT NULL,
    `created_by` varchar(255) NOT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(255) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `producer_ibfk_1` (`created_by`),
    KEY `producer_ibfk_2` (`updated_by`),
    CONSTRAINT `producer_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `account` (`id`),
    CONSTRAINT `producer_ibfk_2` FOREIGN KEY (`updated_by`) REFERENCES `account` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci