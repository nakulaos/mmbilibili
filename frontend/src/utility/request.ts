import axios from 'axios';
import { store } from '@/store/store';
import { clearAccessToken } from '@/store/userInfo'
import { getIntl } from '@/locales';
import {  ServerErrorKey } from '@/locales/locale';
import { addRequest, RefreshTokenAtUtility } from '@/utility/refresh'

// 创建 Axios 实例
const DEV_BASE_URL = "http://192.168.56.1:18081";
const PROD_BASE_URL = "http://xx.xx.xx.xx";
export const uri = DEV_BASE_URL
const myAxios = axios.create({
    baseURL: uri,
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
        const {config, data } = response;
        const intl = getIntl();
        console.info(data);
        return new Promise((resolve, reject) => {
            if (data.code !== 0) {
                if (data.code === 100001) {
                    // token 失效
                    store.dispatch(clearAccessToken());
                    addRequest(() => resolve(myAxios(config)))
                    RefreshTokenAtUtility();
                } else {
                    throw new Error(data.msg || intl.formatMessage({ id: ServerErrorKey }));
                }
            }

            return resolve(data);
        })


    },
    async function (error) {
        const intl = getIntl();
        if (error.response) {
            const { status, data } = error.response;
            if (status === 401) {
                addRequest(() => myAxios(error.config))
                RefreshTokenAtUtility();
            } else {
                throw new Error(data.msg || intl.formatMessage({ id: ServerErrorKey }));
            }
        } else {
            throw new Error(intl.formatMessage({ id: ServerErrorKey }));
        }
    }
);

export default myAxios;
