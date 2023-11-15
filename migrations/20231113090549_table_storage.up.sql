CREATE TABLE IF NOT EXISTS
  `storage` (
    `locat` char(8) NOT NULL,
    `product_id` varchar(255) NOT NULL,
    `inventory_qty` int NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(255) NOT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(255) DEFAULT NULL,
    UNIQUE KEY `uidx_locat_product_id` (`locat`, `product_id`),
    KEY `storage_ibfk_1` (`product_id`),
    CONSTRAINT `storage_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci