
-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_book_items` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `book_id` bigint(20) NOT NULL,
    `status` varchar(255) NOT NULL,
    `progress_rate` float NOT NULL,
    `urgency` int(11) NOT NULL,
    `priority` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_user_book_items_on_user_id` (`user_id`),
    KEY `index_user_book_items_on_book_id` (`book_id`),
    CONSTRAINT `fk_user_book_items_0000` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_user_book_items_0001` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `user_book_items`;
