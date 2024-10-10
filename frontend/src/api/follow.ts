// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 关注用户 POST /v1/user/follow */
export async function followUser(body: API.FollowUserReq, options?: { [key: string]: any }) {
    return request<string>('/v1/user/follow', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 粉丝列表 POST /v1/user/follower/list */
export async function followerList(body: API.FollowerListReq, options?: { [key: string]: any }) {
    return request<API.FollowerListResp>('/v1/user/follower/list', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 关注列表 POST /v1/user/following/list */
export async function followingList(body: API.FollowingListReq, options?: { [key: string]: any }) {
    return request<API.FollowingListResp>('/v1/user/following/list', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 好友列表 POST /v1/user/friend/list */
export async function friendList(body: API.FriendListReq, options?: { [key: string]: any }) {
    return request<API.FriendListResp>('/v1/user/friend/list', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/logout */
export async function logout(body: API.LogoutReq, options?: { [key: string]: any }) {
    return request<string>('/v1/user/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 此处后端没有提供注释 POST /v1/user/userinfo */
export async function updateUserInfo(body: API.UpdateUserInfoReq, options?: { [key: string]: any }) {
    return request<API.UpdateUserInfoResp>('/v1/user/userinfo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}
