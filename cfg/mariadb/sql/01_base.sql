create table if not exists `club_pages` (
    `club_uuid` char(36) not null primary key,
    `club_slug` char(15) not null,
    `name` varchar(63) not null,
    `description` text not null,
    `short_description` varchar(50) not null,
    `campus` tinyint unsigned not null,
    `club_type` tinyint unsigned not null,
    `club_remark` text,
    `schedule_remark` text,
    `visible` tinyint(1) not null,
    `updated_at` datetime not null
);

create table if not exists `users` (
    `user_uuid` char(36) not null primary key,
    `role` varchar(16) not null
);
