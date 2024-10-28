CREATE TABLE file_chunks (
                             id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自定义ID字段',
                             created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             deleted_at DATETIME NULL COMMENT '删除时间',
                             file_hash VARCHAR(255) COMMENT '文件MD5',
                             is_uploaded boolean DEFAULT false COMMENT '未上传: 0, 已上传: 1',
                             upload_id VARCHAR(255) UNIQUE COMMENT 'MinIO 上传 ID',
                             total_chunks INT COMMENT '总分片数',
                             file_size BIGINT COMMENT '文件大小',
                             file_name VARCHAR(255) COMMENT '文件名称',
                             object_name VARCHAR(255) COMMENT 'minio 路径',
                             user_id BIGINT NOT NULL COMMENT '用户ID，上传文件的用户',
                             file_type INT NOT NULL COMMENT '文件类型，1: 图片, 2: 视频, 3: 音频, 4: 文档, 5: 其他',
                             UNIQUE INDEX idx_file_hash (user_id,file_hash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件分片表';
