// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 直播信息 POST /v1/live/info */
export async function liveInfo(body: API.LiveDetailReq, options?: { [key: string]: any }) {
    return request<API.LiveDetailResp>('/v1/live/info', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 直播评论列表 POST /v1/live/list/comment */
export async function liveCommentList(body: API.LiveCommentListReq, options?: { [key: string]: any }) {
    return request<API.LiveCommentListResp>('/v1/live/list/comment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 直播列表 POST /v1/live/list/liveinfo */
export async function liveList(body: API.LiveListReq, options?: { [key: string]: any }) {
    return request<API.LiveListResp>('/v1/live/list/liveinfo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/login/email */
export async function loginWithEmail(body: API.LoginWithEmailReq, options?: { [key: string]: any }) {
    return request<API.LoginResp>('/v1/user/login/email', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/login/phone */
export async function loginWithPhone(body: API.LoginWithPhoneReq, options?: { [key: string]: any }) {
    return request<API.LoginResp>('/v1/user/login/phone', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/login/username */
export async function loginWithUsername(body: API.LoginWithUsernameReq, options?: { [key: string]: any }) {
    return request<API.LoginResp>('/v1/user/login/username', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/register */
export async function register(body: API.RegisterReq, options?: { [key: string]: any }) {
    return request<API.LoginResp>('/v1/user/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}
