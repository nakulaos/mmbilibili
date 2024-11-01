CREATE TABLE users (
                       id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       deleted_at DATETIME NULL COMMENT '删除时间',
                       role TINYINT UNSIGNED NOT NULL DEFAULT 1,
                       gender TINYINT UNSIGNED NOT NULL DEFAULT 1,
                       status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1:正常 2:封禁',
                       username VARCHAR(48) NOT NULL UNIQUE COMMENT '账号名',
                       nickname VARCHAR(48) NOT NULL,
                       description VARCHAR(256) DEFAULT NULL COMMENT '简介',
                       salt VARCHAR(36) NOT NULL,
                       phone VARCHAR(48) UNIQUE DEFAULT NULL,
                       email VARCHAR(48) UNIQUE DEFAULT NULL,
                       password VARCHAR(256) NOT NULL,
                       avatar VARCHAR(255) NOT NULL,
                       INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE user_relationships
(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,  -- 主键ID
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',   -- 用户ID
    related_user_id BIGINT UNSIGNED NOT NULL COMMENT '相关用户ID',   -- 相关用户ID
    relationship_attr BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '关系属性',  -- 属性
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',  -- 创建时间
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',  -- 更新时间
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',  -- 删除时间，用于软删除

    UNIQUE KEY idx_user_related (user_id, related_user_id),
    CONSTRAINT fk_user_relationships_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_relationships_related
        FOREIGN KEY (related_user_id) REFERENCES users (id) ON DELETE CASCADE,
    INDEX idx_deleted_at (deleted_at)
);


CREATE TABLE user_relevant_counts (
                                      id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                      user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
                                      follower_count BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝数',
                                      following_count BIGINT NOT NULL DEFAULT 0 COMMENT '关注数',
                                      like_count BIGINT NOT NULL DEFAULT 0 COMMENT '被点赞数',
                                      star_count BIGINT NOT NULL DEFAULT 0 COMMENT '被收藏数',
                                      self_star_count BIGINT DEFAULT NULL COMMENT '自己收藏作品数',
                                      self_like_count BIGINT NOT NULL DEFAULT 0 COMMENT '自己点赞作品数',
                                      live_count BIGINT NOT NULL DEFAULT 0 COMMENT '直播次数',
                                      work_count BIGINT NOT NULL DEFAULT 0 COMMENT '作品数',
                                      friend_count BIGINT NOT NULL DEFAULT 0 COMMENT '朋友数',
                                      whisper_count BIGINT NOT NULL DEFAULT 0 COMMENT '私信数',
                                      black_count BIGINT NOT NULL DEFAULT 0 COMMENT '黑名单数',
                                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
                                      INDEX idx_deleted_at (deleted_at),
                                      UNIQUE INDEX idx_user_id (user_id),
                                      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

