import request from '/@/utils/request';

export function useMenuApi() {
    return {
        //获取全部动态路由list
        getAllRouteListApi: (params?: object) => {
            return request({
                url: '/menu/getAllRouteList',
                method: 'get',
                params,
            })
        },
        //获取当前角色的动态路由list
        getRouteListApi: (params?: object) => {
            return request({
                url: '/menu/getRouteList',
                method: 'get',
                params,
            });
        },
        //获取全部动态路由tree
        getAllRouteTreeApi: (params?: object) => {
            return request({
                url: '/menu/getAllRouteTree',
                method: 'get',
                params
            });
        },
        //获取当前角色动态路由tree
        getRouteTreeApi: (params?: object) => {
            return request({
                url: "/menu/getRouteTree",
                method: "get",
                params
            })
        },
        //查询动态路由 by meta.title
        findDynamicRouteApi: (data?: object) => {
            return request({
                url: "/menu/findDynamicRoute",
                method: "post",
                data
            })
        },
        //新建动态路由
        newDynamicRouteApi: (data?: object) => {
            return request({
                url: "/menu/newDynamicRoute",
                method: "post",
                data
            })

        },
        //删除动态路由
        delDynamicRouteApi: (data?: object) => {
            return request({
                url: "/menu/delDynamicRoute",
                method: "post",
                data
            })

        },
        //更新动态路由
        updateDynamicRouteApi: (data?: object) => {
            return request({
                url: "/menu/updateDynamicRoute",
                method: "post",
                data
            })
        }
    }
}