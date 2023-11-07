CREATE TABLE IF NOT EXISTS
  `account` (
    `id` varchar(255) NOT NULL,
    `username` varchar(50) NOT NULL,
    `email` varchar(255) NOT NULL,
    `hash_password` longblob NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci