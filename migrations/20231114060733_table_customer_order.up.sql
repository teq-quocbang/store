CREATE TABLE IF NOT EXISTS
  `customer_order` (
    `account_id` varchar(255) NOT NULL,
    `product_id` varchar(255) NOT NULL,
    `sold_qty` int NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `customer_order_ibfk_1` (`account_id`),
    KEY `customer_order_ibfk_2` (`product_id`),
    CONSTRAINT `customer_order_ibfk_1` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`),
    CONSTRAINT `customer_order_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci