CREATE TABLE file_chunks (
                             id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自定义ID字段',
                             created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             deleted_at DATETIME NULL COMMENT '删除时间',
                             uuid VARCHAR(255) UNIQUE COMMENT '文件唯一标识',
                             file_hash VARCHAR(255) UNIQUE COMMENT '文件MD5',
                             is_uploaded INT DEFAULT 0 COMMENT '未上传: 0, 已上传: 1',
                             upload_id VARCHAR(255) UNIQUE COMMENT 'MinIO 上传 ID',
                             total_chunks INT COMMENT '总分片数',
                             file_size BIGINT COMMENT '文件大小',
                             file_name VARCHAR(255) COMMENT '文件名称',
                             origin_name VARCHAR(255) COMMENT 'minio 路径',
                             completed_parts TEXT COMMENT '已完成的分片信息: chunkNumber+etag eg: ,1-asqwewqe21312312.2-123hjkas',
                             user_id BIGINT NOT NULL COMMENT '用户ID，上传文件的用户',
                             file_type INT NOT NULL COMMENT '文件类型，1: 图片, 2: 视频, 3: 音频, 4: 文档, 5: 其他',
                             INDEX idx_user_id (user_id), -- 为user_id添加索引
                             UNIQUE INDEX idx_uuid (uuid) -- 为uuid添加唯一索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件分片表';
