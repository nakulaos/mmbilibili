// src/store/globalSlice.js
import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  theme: 'light',
  language: 'zh-CN',
  mode: 'normal',
};

const globalSlice = createSlice({
  name: 'global',
  initialState,
  reducers: {
    setTheme(state, action) {
      state.theme = action.payload;
    },
    setLanguage(state, action) {
      state.language = action.payload;
    },
    setMode(state, action) {
        state.mode = action.payload;
    },
    
  },
});

export const { setTheme, setLanguage } = globalSlice.actions;
export default globalSlice.reducer;
