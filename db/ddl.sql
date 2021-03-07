-- NOTE: Make sure you have to create db with name "crowfunding" on your machine

-- First
CREATE TABLE `crowfunding`.`users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `occupation` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `password_hash` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `avatar_file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `role` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

-- Second
CREATE TABLE `crowfunding`.`campaigns` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned DEFAULT NULL,
    `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `short_description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `perks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `backer_count` bigint unsigned DEFAULT NULL,
    `goal_amount` bigint unsigned DEFAULT NULL,
    `current_amount` bigint unsigned DEFAULT NULL,
    `slug` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `campaigns_user_id` (`user_id`),
    CONSTRAINT `campaigns_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

-- Third
CREATE TABLE `crowfunding`.`campaign_images` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `campaign_id` bigint unsigned DEFAULT NULL,
    `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `is_primary` tinyint DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `campaign_images_campaign_id` (`campaign_id`),
    CONSTRAINT `campaign_images_campaign_id` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

-- Fourth
CREATE TABLE `transactions` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `campaign_id` bigint unsigned DEFAULT NULL,
    `user_id` bigint unsigned DEFAULT NULL,
    `amount` decimal(15,2) DEFAULT NULL,
    `status` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `code` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `payment_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    KEY `transactions_campaign_id` (`campaign_id`),
    KEY `transactions_user_id` (`user_id`),
    CONSTRAINT `transactions_campaign_id` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE,
    CONSTRAINT `transactions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci