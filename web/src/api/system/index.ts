import service from "/@/utils/request";

export function useSystemApi() {
    return {
        //获取邮箱验证码
        getEmailCodeApi: (data?: object) => {
            return service({
                url: '/email/getEmailCode',
                method: 'POST',
                data
            })
        },
        //从服务器获取主题配置
        getThemeConfigApi: (params?: object) => {
            return service({
                url: "system/getThemeConfig",
                method: "get",
                params
            })
        },
        //设置主题
        updateThemeConfigApi: (data?: object) => {
            return service({
                url: "system/updateThemeConfig",
                method: "post",
                data
            })
        },
        //获取系统设置
        getServerConfig: () => {
            return service({
                url: "system/getSetting",
                method: "get",
            })
        },
        //获取公共系统设置
        getPublicServerConfig: () => {
            return service({
                url: "system/getPublicSetting",
                method: "get",
            })
        },
        //修改系统设置
        updateServerConfig: (data?: object) => {
            return service({
                url: "system/updateSetting",
                method: "post",
                data
            })
        },
    }
}