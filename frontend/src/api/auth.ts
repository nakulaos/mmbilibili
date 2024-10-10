// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 直播评论 POST /v1/live/comment */
export async function liveComment(body: API.LiveCommentReq, options?: { [key: string]: any }) {
    return request<API.LiveCommentResp>('/v1/live/comment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 直播评论删除 POST /v1/live/delete/comment */
export async function liveDeleteComment(body: API.LiveDeleteCommentReq, options?: { [key: string]: any }) {
    return request<API.LiveCommentResp>('/v1/live/delete/comment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 结束直播 POST /v1/live/end */
export async function endLive(body: API.EndLiveReq, options?: { [key: string]: any }) {
    return request<API.LiveDetailResp>('/v1/live/end', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 直播点赞 POST /v1/live/like */
export async function liveLike(body: API.LiveLikeReq, options?: { [key: string]: any }) {
    return request<API.LiveLikeResp>('/v1/live/like', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 直播列表 POST /v1/live/list/gorse */
export async function liveListByGorse(body: API.LiveListReq, options?: { [key: string]: any }) {
    return request<API.LiveListResp>('/v1/live/list/gorse', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 开启直播 POST /v1/live/start */
export async function startLive(body: API.StartLiveReq, options?: { [key: string]: any }) {
    return request<API.LiveDetailResp>('/v1/live/start', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}
