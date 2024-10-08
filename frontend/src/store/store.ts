// src/store/store.js
import { configureStore, combineReducers } from '@reduxjs/toolkit';
import storage from 'redux-persist/lib/storage'; // 默认使用 localStorage
import { persistStore, persistReducer } from 'redux-persist';
// import thunk from 'redux-thunk';

// 导入你的 slice
import globalReducer from './global';
import userInfoReducer from './userInfo';

// 配置持久化
const persistConfig = {
  key: 'root',
  storage,
  whitelist: ['global', 'userInfo'], // 需要持久化的 reducer
};

// 合并所有的 reducers
const rootReducer = combineReducers({
  global: globalReducer,
  userInfo: userInfoReducer,
});

// 创建持久化的 reducer
const persistedReducer = persistReducer(persistConfig, rootReducer);

// 配置 store
export const store = configureStore({
  reducer: persistedReducer,
//   middleware: (getDefaultMiddleware) =>
//     getDefaultMiddleware({
//       serializableCheck: {
//         // Redux Persist 会使用一些非序列化的值，需要忽略警告
//         ignoredActions: ['persist/PERSIST', 'persist/REHYDRATE'],
//       },
//     }).concat(thunk),
});

// 创建 persistor
export const persistor = persistStore(store);
