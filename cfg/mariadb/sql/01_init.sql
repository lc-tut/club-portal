create table if not exists `club_pages` (
    `club_uuid` char(36) not null primary key,
    `club_slug` char(15) not null,
    `name` varchar(63) not null,
    `description` text not null,
    `campus` tinyint not null,
    `club_type` tinyint not null,
    `visible` tinyint(1) not null,
    `updated_at` datetime not null
);

create table if not exists `club_achievements` (
    `achievement_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `achievements` text not null,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_activities` (
    `activity_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `activity` text not null unique,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_images` (
    `image_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `path` text not null unique,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_videos` (
     `video_id` int unsigned not null primary key auto_increment,
     `club_uuid` char(36) not null,
     `path` text not null unique,
     foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_schedules` (
    `schedule_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `month` tinyint unsigned not null unique,
    `schedule` text not null unique,
    `remarks` text,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_links` (
    `link_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `label` varchar(255) not null unique,
    `url` varchar(2047) not null unique,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_times` (
    `time_id` int unsigned not null primary key auto_increment,
    `date` varchar(3) not null unique,
    `time` varchar(255) not null unique,
    `remarks` text
);

create table if not exists `club_places` (
    `place_id` int unsigned not null primary key auto_increment,
    `place` text not null unique,
    `remarks` text
);

create table if not exists `activity_details` (
    `time_id` int unsigned not null,
    `place_id` int unsigned not null,
    `club_uuid` char(36) not null,
    primary key (`time_id`, `place_id`),
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict,
    foreign key (`time_id`) references `club_times` (`time_id`) on delete cascade on update restrict,
    foreign key (`place_id`) references `club_places` (`place_id`) on delete cascade on update restrict
);
