import service from "/@/utils/request";

export function useSystemApi() {
    return {
        //从服务器获取主题配置
        // getThemeConfigApi: (params?: object) => {
        //     return service({
        //         url: "system/getThemeConfig",
        //         method: "get",
        //         params
        //     })
        // },
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