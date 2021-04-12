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
    `achievement` text not null,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_contents` (
    `content_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `content` text not null,
    unique (`content`) using hash,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_images` (
    `image_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `path` text not null,
    unique (`path`) using hash,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_videos` (
     `video_id` int unsigned not null primary key auto_increment,
     `club_uuid` char(36) not null,
     `path` text not null,
     unique (`path`) using hash,
     foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_schedules` (
    `schedule_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `month` tinyint unsigned not null,
    `schedule` text not null,
    `remarks` text,
    unique (`month`, `schedule`) using hash,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_links` (
    `link_id` int unsigned not null primary key auto_increment,
    `club_uuid` char(36) not null,
    `label` varchar(255) not null,
    `url` varchar(2047) not null,
    unique (`label`, `url`) using hash,
    foreign key (`club_uuid`) references `club_pages` (`club_uuid`) on delete cascade on update restrict
);

create table if not exists `club_times` (
    `time_id` int unsigned not null primary key auto_increment,
    `date` varchar(3) not null,
    `time` varchar(255) not null,
    unique (`date`, `time`)
);

create table if not exists `club_places` (
    `place_id` int unsigned not null primary key auto_increment,
    `place` text not null unique
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

create table if not exists `club_remarks` (
    `remark_id` int unsigned not null primary key auto_increment,
    `time_id` int unsigned not null,
    `place_id` int unsigned not null,
    `club_uuid` char(36) not null,
    `place_remarks` text,
    `time_remarks` text,
    foreign key (`club_uuid`) references `activity_details` (`club_uuid`) on delete cascade on update restrict,
    foreign key (`time_id`) references `activity_details` (`time_id`) on delete cascade on update restrict,
    foreign key (`place_id`) references `activity_details` (`place_id`) on delete cascade on update restrict
)
