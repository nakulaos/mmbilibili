// src/store/store.js
import { configureStore, combineReducers } from '@reduxjs/toolkit';
import storage from 'redux-persist/lib/storage';
import { persistStore, persistReducer } from 'redux-persist';
// import thunk from 'redux-thunk';


import globalReducer from './global';
import userInfoReducer from './userInfo';


const persistConfig = {
  key: 'root',
  storage,
  whitelist: ['global', 'userInfo'], // 需要持久化的 reducer
};


const rootReducer = combineReducers({
  global: globalReducer,
  userInfo: userInfoReducer,
});


const persistedReducer = persistReducer(persistConfig, rootReducer);


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
