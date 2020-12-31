
-- +migrate Up
CREATE TABLE IF NOT EXISTS `study_result_items` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `study_id` bigint(20) NOT NULL,
    `book_id` bigint(20) NOT NULL,
    `memo` varchar(255) DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_study_result_items_on_user_id` (`study_id`),
    KEY `index_study_result_items_on_book_id` (`book_id`),
    CONSTRAINT `fk_study_result_items_0000` FOREIGN KEY (`study_id`) REFERENCES `studies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_study_result_items_0001` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `study_result_items`;
