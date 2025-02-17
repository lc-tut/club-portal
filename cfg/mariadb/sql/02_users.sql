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

# general_usersとclub_pagesのnameを同期させるトリガー
create trigger club_name_sync_by_club_pages
    after update on club_pages
    for each row
    update general_users set name = new.name where club_uuid = new.club_uuid;

create trigger club_name_sync_by_general_users
    after update on general_users
    for each row
    update club_pages set name = new.name where club_uuid = new.club_uuid;

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

create table if not exists `uploaded_thumbnails` (
    `thumbnail_id` int unsigned not null primary key auto_increment,
    `path` varchar(255) not null,
    unique (`path`) using hash
)
