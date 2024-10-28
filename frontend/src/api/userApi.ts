// @ts-ignore
/* eslint-disable */
import request from '@/utility/request'

/** 关注或取消关注用户 POST /v1/auth/user/follow */
export async function followUser(body: API.FollowUserReq, options?: { [key: string]: any }) {
    return request<API.FollowUserResp>('/v1/auth/user/follow', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 更新用户信息 POST /v1/auth/user/info */
export async function updateUserInfo(body: API.UpdateUserInfoReq, options?: { [key: string]: any }) {
    return request<API.UpdateUserInfoResp>('/v1/auth/user/info', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 用户登出 POST /v1/auth/user/logout */
export async function logout(body: API.LogoutReq, options?: { [key: string]: any }) {
    return request<API.LogoutResp>('/v1/auth/user/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: body,
        ...(options || {})
    })
}

/** 刷新令牌 POST /v1/auth/user/refresh */
export async function refreshToken(options?: { [key: string]: any }) {
    return request<API.RefreshTokenResp>('/v1/auth/user/refresh', {
        method: 'POST',
        ...(options || {})
    })
}

/** 获取粉丝列表 GET /v1/user/followers */
export async function followerList(
    // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
    params: API.FollowerListParams,
    options?: { [key: string]: any }
) {
    return request<API.FollowerListResp>('/v1/user/followers', {
        method: 'GET',
        params: {
            ...params
        },
        ...(options || {})
    })
}

/** 获取关注列表 GET /v1/user/following */
export async function followingList(
    // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
    params: API.FollowingListParams,
    options?: { [key: string]: any }
) {
    return request<API.FollowingListResp>('/v1/user/following', {
        method: 'GET',
        params: {
            ...params
        },
        ...(options || {})
    })
}

/** 获取好友列表 GET /v1/user/friends */
export async function friendList(
    // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
    params: API.FriendListParams,
    options?: { [key: string]: any }
) {
    return request<API.FriendListResp>('/v1/user/friends', {
        method: 'GET',
        params: {
            ...params
        },
        ...(options || {})
    })
}

/** 邮箱登录 POST /v1/user/login/email */
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

/** 手机号登录 POST /v1/user/login/phone */
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

/** 用户名登录 POST /v1/user/login/username */
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

/** 用户注册 POST /v1/user/register */
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
