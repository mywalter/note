import axios from 'axios';
export function request(config) {
    // 1.创建axios的实例
    const instance = axios.create({
        method:"POST",
        baseURL: '/api',
        timeout: 5000,
    })

    // 2.拦截器（请求成功，请求失败，响应成功，响应失败）
    // 2.1.0 请求成功，请求失败
    instance.interceptors.request.use(config => {
        // 2.1拦截
        // 2.1.1 比如config种的一些信息不符合服务器的要求
        // 2.1.2 比如每次发送网络请求时，都希望在界面种显示一一个请求的图标（UI过渡）
        // 2.1.3 某些网络请求（比如登录），必须携带一些特殊的信息
        // console.log(config, "请求成功");
        // 2.2 放行
        return config;
    }, err => {
        console.log(err)
    });
    // 2.1.0 响应成功，响应失败
    instance.interceptors.response.use(res => {
        // console.log(res, "响应成功")
        return res.data
    }, err => {
        console.log(err)
    });

    // 3.发送请求,不能再次处理数据，应该要回调
    return instance(config)
}