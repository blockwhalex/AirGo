import service from "/@/utils/request";

export function usePublicApi() {
    return {
        //从服务器获取主题配置
        getThemeConfigApi: (params?: object) => {
            return service({
                url: "/public/getThemeConfig",
                method: "get",
                params
            })
        },
        //获取邮箱验证码
        getEmailCodeApi: (data?: object) => {
            return service({
                url: '/public/getEmailCode',
                method: 'POST',
                data
            })
        },
    }

}