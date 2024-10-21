/*========================>database video<===================================*/
create database video;
use video;
-- Create Video table
CREATE TABLE IF NOT EXISTS `videos` (
                                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                        `created_at` DATETIME(3) DEFAULT NULL,
                                        `updated_at` DATETIME(3) DEFAULT NULL,
                                        `deleted_at` DATETIME(3) DEFAULT NULL,
                                        `author_id` BIGINT(20) unsigned NOT NULL,
                                        `title` VARCHAR(255) NOT NULL,
                                        `play_url` VARCHAR(255) NOT NULL,
                                        `cover_url` VARCHAR(255) NOT NULL,
                                        `favorite_count` BIGINT DEFAULT 0,
                                        `comment_count` BIGINT DEFAULT 0,
                                        PRIMARY KEY (`id`),
                                        INDEX `idx_author_id` (`author_id`),
                                        INDEX `idx_title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create Comment table
CREATE TABLE IF NOT EXISTS `comments` (
                                          `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                          `created_at` DATETIME(3) DEFAULT NULL,
                                          `updated_at` DATETIME(3) DEFAULT NULL,
                                          `deleted_at` DATETIME(3) DEFAULT NULL,
                                          `user_id` BIGINT(20) UNSIGNED NOT NULL,
                                          `video_id` BIGINT UNSIGNED NOT NULL,
                                          `content` TEXT NOT NULL,
                                          PRIMARY KEY (`id`),
                                          INDEX `idx_user_id` (`user_id`),
                                          INDEX `idx_video_id` (`video_id`),
                                          CONSTRAINT `fk_video_comment` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create Favorite table
CREATE TABLE IF NOT EXISTS `favorites` (
                                           `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                           `created_at` DATETIME(3) DEFAULT NULL,
                                           `updated_at` DATETIME(3) DEFAULT NULL,
                                           `deleted_at` DATETIME(3) DEFAULT NULL,
                                           `user_id` BIGINT(20) UNSIGNED NOT NULL,
                                           `video_id` BIGINT UNSIGNED NOT NULL,
                                           PRIMARY KEY (`id`),
                                           INDEX `idx_user_id` (`user_id`),
                                           INDEX `idx_video_id` (`video_id`),
                                           CONSTRAINT `fk_video_favorite` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;