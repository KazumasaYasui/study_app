
-- +migrate Up
CREATE TABLE IF NOT EXISTS `studies` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `memo` varchar(255) DEFAULT NULL,
    `status` varchar(255) NOT NULL,
    `execution_date` date DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_studies_on_user_id` (`user_id`),
    CONSTRAINT `fk_studies_0001` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `studies`;
