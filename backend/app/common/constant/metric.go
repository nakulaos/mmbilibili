package constant

const (
	PromCacheUserDetailUsername = "cache.user_detail_username"
	PromCacheUserDetailUID      = "cache.user_detail_uid"
	PromRedisUserTokenBlackList = "redis.user_token_black_list"
	PromDBUser                  = "db.http"
	PromDBRelevantCount         = "db.relevant_count"
	PromDBFileChunk             = "db.file_chunk"

	PromMinIOUploadID = "minio.upload_id"
)

const (
	PromUserExistError     = "user_exist_error"
	PromUserNotExistError  = "user_not_exist_error"
	PromGenerateTokenError = "generate_token_error"
)
