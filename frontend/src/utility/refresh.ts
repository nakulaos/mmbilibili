import { refreshToken } from '@/api/userApi';
import myAxios, { uri } from '@/utility/request'
import { store } from '@/store/store';
import { clearToken, clearUserInfo, setAccessToken, setRefreshToken } from '@/store/userInfo'
import axios from 'axios'
import { message } from 'antd'
import { AuthorizationErrorKey, ServerErrorKey } from '@/locales/locale'
import { getIntl } from '@/locales'

type RequestFunction = () => void;

let subscribes: RequestFunction[] = [];
let flag = false;

export const addRequest = (request: RequestFunction): void => {
    subscribes.push(request);
};

export const retryRequest = (): void => {
    subscribes.forEach((request) => request());
    subscribes = [];
};

export const RefreshTokenAtUtility = (): void => {
    const intl = getIntl()
    if (!flag) {
        flag = true;
        const state = store.getState();
        const r_tk = state.userInfo.refreshToken;
        if (r_tk !== "") {
            axios.create({
                baseURL: uri,
                timeout: 10000,
            }).post('/v1/auth/user/refresh', {},{
                headers: {
                    'Kind': 'refresh',
                    'Authorization': r_tk,
                },
            }).then((res) => {
                    console.log("res.data",res.data);
                    console.log("res.data.code",res.data.code);
                    console.log("res.data.data",res.data.data);

                    if (res.data && res.data.code === 0) {
                        console.log("res.data.data.access_token",res.data.data.access_token);
                        store.dispatch(setAccessToken(res.data.data.access_token));
                        store.dispatch(setRefreshToken(res.data.data.refresh_token));
                        flag = false;
                        retryRequest();

                    } else { // @ts-ignore
                        if (res.data && res.data.code === 100001) {
                            flag = false;
                            store.dispatch(clearToken());
                            store.dispatch(clearUserInfo());
                            message.error(intl.formatMessage({ id: AuthorizationErrorKey }));
                        }
                    }
                })
                .catch((err) => {
                    if(err.response.status === 401){
                        flag = false;
                        store.dispatch(clearToken());
                        store.dispatch(clearUserInfo());
                        message.error(intl.formatMessage({ id: AuthorizationErrorKey }));
                    }else{
                        flag = false;
                        store.dispatch(clearToken());
                        store.dispatch(clearUserInfo());
                        message.error(intl.formatMessage({ id: ServerErrorKey }));
                    }
                });
        }
    }
};
