// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 完成分块上传 POST /v1/auth/file/complete_multipart */
export async function completeMultipart(body: API.CompleteMultipartReq, options?: { [key: string]: any }) {
    return request<API.CompleteMultipartResp>('/v1/auth/file/complete_multipart', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 上传文件 POST /v1/auth/file/multi_upload */
export async function newMultiUpload(body: API.NewMultiUploadReq, options?: { [key: string]: any }) {
    return request<API.NewMultiUploadResp>('/v1/auth/file/multi_upload', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 获取大文件上传请求的 URI POST /v1/auth/file/multi_upload_uri */
export async function getMultiUploadUri(body: API.GetMultiUploadUriReq, options?: { [key: string]: any }) {
    return request<API.GetMultiUploadUriResp>('/v1/auth/file/multi_upload_uri', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 获取已经上传的分块 POST /v1/auth/file/success_chunks */
export async function getSuccessChunks(body: API.GetSuccessChunksReq, options?: { [key: string]: any }) {
    return request<API.GetSuccessChunksResp>('/v1/auth/file/success_chunks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}
