
-- +migrate Up
CREATE TABLE IF NOT EXISTS `books` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `genre_id` bigint(20) NOT NULL,
    `publisher_id` bigint(20) NOT NULL,
    `author_id` bigint(20) NOT NULL,
    `title` varchar(255) NOT NULL,
    `image_url` varchar(255) DEFAULT NULL,
    `publish_date` date DEFAULT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_books_on_genre_id` (`genre_id`),
    KEY `index_books_on_publisher_id` (`publisher_id`),
    KEY `index_books_on_author_id` (`author_id`),
    CONSTRAINT `fk_books_0000` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_books_0001` FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_books_0002` FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `books`;
