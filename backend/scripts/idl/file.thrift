
struct NewMultiUploadReq {
    1: string FileHash (api.body = "file_hash,required" api.vd="len($)==64") // 文件内容的唯一哈希值;
    2: i64 ChunkTotalNumber (api.body = "chunk_total_number,required" api.vd="$>=1") // 分块总数;
    3: string FileName (api.body = "file_name,required" ) // 文件名;
    4: i64 FileSize (api.body = "file_size,required") // 文件大小;
    5: i32 FileType (api.body = "file_type,required" api.vd="$>=1") // 文件类型，文件属性比如图片，视频等;
}

struct NewMultiUploadResp {

}

struct GetSuccessChunksReq {
    1: string FileHash (api.body = "file_hash,required") // ;
}

struct GetSuccessChunksResp {
    1: bool IsUpload (api.body = "is_upload") // ;
    2: string Chunks (api.body = "chunks") // ;
}

struct GetMultiUploadUriReq {
    1: string FileHash (api.body = "file_hash,required") // ;
    2: i64 ChunkID (api.body = "chunk_id,required") // 分块ID;
    3: i64 ChunkSize (api.body = "chunk_size,required") // 分块大小;
}

struct GetMultiUploadUriResp {
    1: string Uri (api.body = "uri") // ;
}

struct CompleteMultipartReq {
    1: string FileHash (api.body = "file_hash,required") // ;
}

struct CompleteMultipartResp {
}


service FileAPI {
   NewMultiUploadResp NewMultiUpload(1: NewMultiUploadReq req) (api.post = "/v1/auth/file/multi_upload");
   GetMultiUploadUriResp GetMultiUploadUri(1: GetMultiUploadUriReq req) (api.post = "/v1/auth/file/multi_upload_uri");
   CompleteMultipartResp CompleteMultipart(1: CompleteMultipartReq req) (api.post = "/v1/auth/file/complete_multipart");
   GetSuccessChunksResp GetSuccessChunks(1: GetSuccessChunksReq req) (api.post = "/v1/auth/file/success_chunks");
}