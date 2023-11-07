CREATE TABLE IF NOT EXISTS
  `product` (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `product_type` varchar(255) NOT NULL,
    `producer_id` varchar(255) NOT NULL,
    `amount` int NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL,
    `created_by` varchar(255) NOT NULL,
    `updated_by` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `product_ibfk_1` (`producer_id`),
    KEY `product_ibfk_2` (`created_by`),
    KEY `product_ibfk_3` (`updated_by`),
    CONSTRAINT `product_ibfk_1` FOREIGN KEY (`producer_id`) REFERENCES `producer` (`id`),
    CONSTRAINT `product_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `account` (`id`),
    CONSTRAINT `product_ibfk_3` FOREIGN KEY (`updated_by`) REFERENCES `account` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci