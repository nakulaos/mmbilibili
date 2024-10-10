import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { clear } from 'console';

interface UserInfo {
  id: number;
  username: string;
  nickname: string;
  avatar: string;
  gender: number;
  role: number;
  followerCount: number;
  followingCount: number;
  likeCount: number;
  starCount: number;
  selfStarCount: number;
  selfLikeCount: number;
  liveCount: number;
  workCount: number;
  friendCount: number;
  phone: string;
  email: string;
  status: number;
  token?: string;
}

const initialState: UserInfo = {
  id: 0,
  username: '',
  nickname: '',
  avatar: '',
  gender: 0,
  role: 0,
  followerCount: 0,
  followingCount: 0,
  likeCount: 0,
  starCount: 0,
  selfStarCount: 0,
  selfLikeCount: 0,
  liveCount: 0,
  workCount: 0,
  friendCount: 0,
  phone: '',
  email: '',
  status: 0,
  token: '',
};

const userInfoSlice = createSlice({
  name: 'userInfo',
  initialState,
  reducers: {
    setUserInfo(state, action: PayloadAction<Partial<UserInfo>>) {
      return {
        ...state,
        ...action.payload, // 合并当前状态和传入的 payload
      };
    },

    clearUserInfo(state) {
      return initialState;
    },
    setToken(state, action: PayloadAction<string>) {
      state.token = action.payload;
    },
    clearToken(state) {
      state.token = '';
    },
  },
});

export const { setUserInfo, clearUserInfo } = userInfoSlice.actions;
export default userInfoSlice.reducer;
