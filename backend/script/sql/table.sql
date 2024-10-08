create table articles
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    title      varchar(255)    not null comment '文章标题',
    content    text            not null comment '文章内容',
    author_id  bigint unsigned not null comment '作者ID'
);

create index idx_articles_deleted_at
    on articles (deleted_at);

create table categories
(
    name        varchar(255) not null,
    type        varchar(255) not null,
    description text         null,
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3)  null,
    updated_at  datetime(3)  null,
    deleted_at  datetime(3)  null,
    constraint idx_name_type
        unique (name, type)
);

create table article_categories
(
    article_id  bigint unsigned not null,
    category_id bigint unsigned not null,
    primary key (article_id, category_id),
    constraint fk_article_categories_article
        foreign key (article_id) references articles (id),
    constraint fk_article_categories_category
        foreign key (category_id) references categories (id)
);

create index idx_categories_deleted_at
    on categories (deleted_at);

create table histories
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    uid        bigint unsigned not null,
    work_id    bigint unsigned not null,
    type       bigint          not null comment '1: video, 2: live 3. article'
);

create index idx_histories_deleted_at
    on histories (deleted_at);

create table tags
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3)             null,
    updated_at  datetime(3)             null,
    deleted_at  datetime(3)             null,
    name        varchar(255) default '' not null comment '标签名',
    description text                    null
);

create table article_tags
(
    article_id bigint unsigned not null,
    tag_id     bigint unsigned not null,
    primary key (article_id, tag_id),
    constraint fk_article_tags_article
        foreign key (article_id) references articles (id),
    constraint fk_article_tags_tag
        foreign key (tag_id) references tags (id)
);

create table category_tags
(
    category_id bigint unsigned not null,
    tag_id      bigint unsigned not null,
    primary key (category_id, tag_id),
    constraint fk_category_tags_category
        foreign key (category_id) references categories (id),
    constraint fk_category_tags_tag
        foreign key (tag_id) references tags (id)
);

create index idx_tags_deleted_at
    on tags (deleted_at);

create table users
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      datetime(3)              null,
    updated_at      datetime(3)              null,
    deleted_at      datetime(3)              null,
    username        varchar(48)              not null comment '账号名',
    nickname        varchar(48)              not null,
    description     varchar(256)             null comment '简介',
    status          int unsigned default '1' not null,
    phone           varchar(48)              null,
    email           varchar(48)              null,
    password        varchar(256)             not null,
    avatar          varchar(255)             not null,
    role            varchar(36)  default '1' not null,
    gender          int unsigned default '1' not null,
    salt            varchar(36)              not null,
    follower_count  bigint       default 0   not null comment '粉丝数',
    following_count bigint       default 0   not null comment '关注数',
    like_count      bigint       default 0   not null comment '被点赞数',
    star_count      bigint       default 0   not null comment '被收藏数',
    self_star_count bigint                   null comment '自己收藏作品数',
    self_like_count bigint       default 0   not null comment '自己点赞作品数',
    live_count      bigint       default 0   not null comment '直播次数',
    work_count      bigint       default 0   not null comment '作品数',
    friend_count    bigint       default 0   not null comment '朋友数',
    room_id         bigint       default 0   not null comment '房间ID',
    constraint idx_users_username
        unique (username)
);

create table comments
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3)      null,
    updated_at    datetime(3)      null,
    deleted_at    datetime(3)      null,
    uid           bigint unsigned  not null,
    owner_id      bigint unsigned  not null,
    owner_type    longtext         not null,
    content       text             not null,
    parent_id     bigint unsigned  null,
    like_count    bigint default 0 null,
    un_like_count bigint default 0 null,
    constraint fk_comments_replies
        foreign key (parent_id) references comments (id),
    constraint fk_users_comments
        foreign key (uid) references users (id)
);

create index idx_comments_deleted_at
    on comments (deleted_at);

create index idx_comments_parent_id
    on comments (parent_id);

create table danmus
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3)      null,
    updated_at    datetime(3)      null,
    deleted_at    datetime(3)      null,
    uid           bigint unsigned  not null,
    owner_id      bigint unsigned  not null,
    owner_type    varchar(191)     not null,
    content       text             not null,
    send_time     double           not null comment '在视频的哪个点发送的弹幕',
    type          bigint           not null comment '弹幕类型，1 为视频，2 为直播',
    like_count    bigint default 0 null,
    un_like_count bigint default 0 null,
    constraint fk_users_danmus
        foreign key (uid) references users (id)
);

create index idx_danmus_deleted_at
    on danmus (deleted_at);

create index idx_owner
    on danmus (owner_id, owner_type);

create table lives
(
    uid           bigint unsigned              not null comment '用户ID',
    title         varchar(255)                 not null comment '直播标题',
    description   varchar(255)                 null comment '直播描述',
    start_time    datetime(3)                  not null comment '开始时间',
    end_time      datetime(3)                  null comment '结束时间',
    status        tinyint unsigned default '0' not null comment '直播状态',
    view_count    bigint unsigned  default '0' not null comment '观看人数',
    like_count    bigint unsigned  default '0' not null comment '点赞数',
    comment_count bigint unsigned  default '0' not null comment '评论数',
    play_url      varchar(255)                 not null comment '直播地址',
    cover_url     varchar(255)                 not null comment '封面地址',
    stream_id     bigint unsigned  default '0' not null comment '流ID',
    is_over       bigint           default 0   not null comment '是否结束',
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3)                  null,
    updated_at    datetime(3)                  null,
    deleted_at    datetime(3)                  null,
    constraint fk_users_lives
        foreign key (uid) references users (id)
);

create table live_category
(
    live_id     bigint unsigned not null,
    category_id bigint unsigned not null,
    primary key (live_id, category_id),
    constraint fk_live_category_category
        foreign key (category_id) references categories (id),
    constraint fk_live_category_live
        foreign key (live_id) references lives (id)
);

create table live_tag
(
    live_id bigint unsigned not null,
    tag_id  bigint unsigned not null,
    primary key (live_id, tag_id),
    constraint fk_live_tag_live
        foreign key (live_id) references lives (id),
    constraint fk_live_tag_tag
        foreign key (tag_id) references tags (id)
);

create index idx_lives_deleted_at
    on lives (deleted_at);

create table user_favorite_articles
(
    user_id    bigint unsigned not null,
    article_id bigint unsigned not null,
    primary key (user_id, article_id),
    constraint fk_user_favorite_articles_article
        foreign key (article_id) references articles (id),
    constraint fk_user_favorite_articles_user
        foreign key (user_id) references users (id)
);

create table user_favorite_lives
(
    user_id bigint unsigned not null,
    live_id bigint unsigned not null,
    primary key (user_id, live_id),
    constraint fk_user_favorite_lives_live
        foreign key (live_id) references lives (id),
    constraint fk_user_favorite_lives_user
        foreign key (user_id) references users (id)
);

create table user_follows
(
    followed_id bigint unsigned not null,
    follower_id bigint unsigned not null,
    primary key (followed_id, follower_id),
    constraint fk_user_follows_followers
        foreign key (follower_id) references users (id),
    constraint fk_user_follows_user
        foreign key (followed_id) references users (id)
);

create table user_star_articles
(
    user_id    bigint unsigned not null,
    article_id bigint unsigned not null,
    primary key (user_id, article_id),
    constraint fk_user_star_articles_article
        foreign key (article_id) references articles (id),
    constraint fk_user_star_articles_user
        foreign key (user_id) references users (id)
);

create table user_tag
(
    user_id bigint unsigned not null,
    tag_id  bigint unsigned not null,
    primary key (user_id, tag_id),
    constraint fk_user_tag_tag
        foreign key (tag_id) references tags (id),
    constraint fk_user_tag_user
        foreign key (user_id) references users (id)
);

create index idx_users_deleted_at
    on users (deleted_at);

create index idx_users_email
    on users (email);

create index idx_users_phone
    on users (phone);

create table videos
(
    id             bigint unsigned auto_increment
        primary key,
    created_at     datetime(3)                 null,
    updated_at     datetime(3)                 null,
    deleted_at     datetime(3)                 null,
    author_id      bigint unsigned             not null comment '上传用户Id',
    title          varchar(255)                not null comment '视频标题',
    cover_url      varchar(255)                not null comment '封面url',
    play_url       varchar(255)                not null comment '视频播放url',
    favorite_count bigint unsigned default '0' not null comment '点赞数',
    star_count     bigint unsigned             not null comment '收藏数',
    comment_count  bigint unsigned default '0' not null comment '评论数目',
    category       bigint unsigned             not null comment '视频分类',
    duration       varchar(255)                null comment '视频时长',
    constraint fk_users_videos
        foreign key (author_id) references users (id)
);

create table user_favorite_videos
(
    user_id  bigint unsigned not null,
    video_id bigint unsigned not null,
    primary key (user_id, video_id),
    constraint fk_user_favorite_videos_user
        foreign key (user_id) references users (id),
    constraint fk_user_favorite_videos_video
        foreign key (video_id) references videos (id)
);

create table user_star_videos
(
    user_id  bigint unsigned not null,
    video_id bigint unsigned not null,
    primary key (user_id, video_id),
    constraint fk_user_star_videos_user
        foreign key (user_id) references users (id),
    constraint fk_user_star_videos_video
        foreign key (video_id) references videos (id)
);

create table video_category
(
    video_id    bigint unsigned not null,
    category_id bigint unsigned not null,
    primary key (video_id, category_id),
    constraint fk_video_category_category
        foreign key (category_id) references categories (id),
    constraint fk_video_category_video
        foreign key (video_id) references videos (id)
);

create table video_tag
(
    video_id bigint unsigned not null,
    tag_id   bigint unsigned not null,
    primary key (video_id, tag_id),
    constraint fk_video_tag_tag
        foreign key (tag_id) references tags (id),
    constraint fk_video_tag_video
        foreign key (video_id) references videos (id)
);

create index idx_videos_deleted_at
    on videos (deleted_at);

