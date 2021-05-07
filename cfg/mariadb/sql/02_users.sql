create table if not exists `domain_users` (
    `user_uuid` char(36) not null primary key,
    `email` varchar(255) not null,
    `name` varchar(32) not null,
    unique (`email`, `name`),
    foreign key (`user_uuid`) references `users` (`user_uuid`) on delete cascade on update restrict
);

create table if not exists `general_users` (
    `user_uuid` char(36) not null primary key,
    `email` varchar(255) not null,
    `name` varchar(32) not null,
    `club_uuid` char(36),
    unique (`email`, `name`),
    foreign key (`user_uuid`) references `users` (`user_uuid`) on delete cascade on update restrict,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete set null on update restrict
    );

create table if not exists `admin_users` (
    `user_uuid` char(36) not null primary key,
    `email` varchar(255) not null,
    `name` varchar(32) not null,
    unique (`email`, `name`),
    foreign key (`user_uuid`) references `users` (`user_uuid`) on delete cascade on update restrict
);

create table if not exists `favorite_clubs` (
    `user_uuid` char(36) not null,
    `club_uuid` char(36) not null,
    primary key (`user_uuid`, `club_uuid`),
    foreign key (`user_uuid`) references `users` (`user_uuid`) on delete cascade on update restrict,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `uploaded_images` (
    `image_id` int unsigned not null primary key auto_increment,
    `path` varchar(255) not null,
    `owner` char(36) not null,
    `created_at` datetime not null,
    unique (`path`) using hash,
    foreign key (`owner`) references `users` (`user_uuid`) on delete cascade on update restrict
);
