CREATE TABLE IF NOT EXISTS
  `cart` (
    `account_id` varchar(255) NOT NULL,
    `product_id` varchar(255) NOT NULL,
    `qty` int NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `cart_ibfk_1` (`account_id`),
    KEY `cart_ibfk_2` (`product_id`),
    CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`),
    CONSTRAINT `cart_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
    UNIQUE KEY `uidx_account_id_product_id` (`account_id`, `product_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci