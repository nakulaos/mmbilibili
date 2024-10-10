// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 上传文件 POST /v1/user/upload/file */
export async function userUploadFile(
    body: {
        kind: string
    },
    file?: File,
    options?: { [key: string]: any }
) {
    const formData = new FormData()

    if (file) {
        formData.append('file', file)
    }

    Object.keys(body).forEach((ele) => {
        const item = (body as any)[ele]

        if (item !== undefined && item !== null) {
            if (typeof item === 'object' && !(item instanceof File)) {
                if (item instanceof Array) {
                    item.forEach((f) => formData.append(ele, f || ''))
                } else {
                    formData.append(ele, JSON.stringify(item))
                }
            } else {
                formData.append(ele, item)
            }
        }
    })

    return request<API.UserUploadFileResp>('/v1/user/upload/file', {
        method: 'POST',
        data: formData,
        requestType: 'form',
        ...(options || {})
    })
}
