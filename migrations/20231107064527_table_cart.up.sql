CREATE TABLE IF NOT EXISTS
  `cart` (
    `account_id` varchar(255) NOT NULL,
    `product_id` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(255) NOT NULL,
    PRIMARY KEY (`account_id`),
    KEY `cart_ibfk_1` (`product_id`),
    KEY `cart_ibfk_2` (`created_by`),
    CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
    CONSTRAINT `cart_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `account` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci