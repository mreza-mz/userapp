-- +migrate Up
-- Table structure for table users
CREATE TABLE users (
       `id`  bigint unsigned not null primary key auto_increment,
       `created_at`   timestamp default current_timestamp not null,
       `updated_at`   timestamp default current_timestamp on update current_timestamp not null,
       `deleted_at`   timestamp default null,
       `last_login`   timestamp default null,
       `phone_number` varchar(191) unique null,
       `email` varchar(191) unique null,
       `password` varchar(255) default null,
       `fullname` varchar(191) default null,
       `avatar` varchar(255) default null,
       `role` tinyint(1) not null,
       `is_change_password` tinyint(1) default 0,
       `is_active` tinyint(1) default 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;

-- +migrate Down
DROP TABLE users;