import axios from 'axios'
import { store } from '@/store/store'
import { clearUserInfo } from '@/store/userInfo';
import { message } from 'antd';
import { useIntl } from '@ant-design/pro-components';
import { getIntl } from '@/locales';
import { AuthorizationErrorKey, ServerErrorKey } from '@/locales/locale';
import { i } from 'vitest/dist/reporters-yx5ZTtEV';


// 创建 Axios 实例
// 区分开发和生产环境
const DEV_BASE_URL = "http://localhost:8082";
const PROD_BASE_URL = "http://xx.xx.xx.xx";
const myAxios = axios.create({
  baseURL: DEV_BASE_URL,
  timeout: 10000,
});

// 创建请求拦截器
myAxios.interceptors.request.use(
  function (config) {
    const state = store.getState();
    const loginToken = state.userInfo.token
    if(loginToken !== ''){
        config.headers.Authorization = loginToken
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  },
);


// 创建响应拦截器
myAxios.interceptors.response.use(
    // 2xx 响应触发
    function (response) {
      // 处理响应数据
      const { data } = response;
      const intl =getIntl()
  
      if (data.code !== 0) {
        // 认证未通过
        if (data.code === 100001) {
          // 重定向到登录页面
          // 调用 Redux 的清除用户信息的方法
          store.dispatch(clearUserInfo());
          store.dispatch(clearToken());
          // 使用 Ant Design 的 message 组件提示用户
          message.warning(intl.formatMessage({id: AuthorizationErrorKey}));
        } else {
          // 其他错误使用 Ant Design 的 message 组件提示用户
          message.error(data.msg || intl.formatMessage({id: ServerErrorKey}));
        }
        Promise.reject(data?.data || {});
      }
      return data;
    },
    function (error) {
      // 处理响应错误
      message.error("网络错误，请检查您的连接！");
      return Promise.reject(error);
    },
  );

export default myAxios;
function clearToken(): any {
    throw new Error('Function not implemented.');
}

