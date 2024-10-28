import axios from 'axios';
import { store } from '@/store/store';
import { clearToken, clearUserInfo, setAccessToken, setRefreshToken } from '@/store/userInfo'
import { message } from 'antd';
import { getIntl } from '@/locales';
import { AuthorizationErrorKey, ServerErrorKey } from '@/locales/locale';
import { refreshToken as refreshTokenApi  } from '@/api/userApi';

// 创建 Axios 实例
const DEV_BASE_URL = "http://192.168.56.1:18081";
const PROD_BASE_URL = "http://xx.xx.xx.xx";
const myAxios = axios.create({
    baseURL: DEV_BASE_URL,
    timeout: 10000,
});

// 请求拦截器：添加访问令牌
myAxios.interceptors.request.use(
    function (config) {
        const state = store.getState();
        const accessToken = state.userInfo.accessToken;
        if (accessToken) {
            config.headers.Authorization = accessToken;
        }
        return config;
    },
    function (error) {
        return Promise.reject(error);
    },
);

// 响应拦截器：处理错误和自动刷新令牌
myAxios.interceptors.response.use(
    function (response) {
        const { data } = response;
        const intl = getIntl();

        if (data.code !== 0) {
            if (data.code === 100001) {
                store.dispatch(clearUserInfo());
                store.dispatch(clearToken());
                message.warning(intl.formatMessage({ id: AuthorizationErrorKey }));

            } else {
                message.error(data.msg || intl.formatMessage({ id: ServerErrorKey }));
            }
            return Promise.reject(data?.data || {});
        }

        message.success(data.msg);
        return data;
    },
    async function (error) {
        const originalRequest = error.config;
        const intl = getIntl();

        if ((error.response?.status === 401) && !originalRequest._retry) {
            originalRequest._retry = true;
            const state = store.getState();

            try {
                const state = store.getState();
                const refreshToken = state.userInfo.refreshToken;
                const refreshResponse = await refreshTokenApi({

                },{
                    headers: {
                        "Authorization": refreshToken,
                        "Kind": "refresh",
                    }
                });
                const { access_token,refresh_token } = refreshResponse.data;

                // 更新store中的访问令牌
                store.dispatch(setAccessToken(access_token));
                store.dispatch(setRefreshToken(refresh_token));


                originalRequest.headers.Authorization = access_token;
                return myAxios(originalRequest);
            } catch (refreshError) {
                store.dispatch(clearUserInfo());
                store.dispatch(clearToken());
                message.warning(intl.formatMessage({ id: AuthorizationErrorKey }));
                return Promise.reject(refreshError);
            }
        }

        message.error("network error");
        return Promise.reject(error);
    }
);

export default myAxios;
