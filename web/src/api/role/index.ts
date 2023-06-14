import request from '/@/utils/request';

export function useRoleApi() {
    return {
        //获取全部role
        getRoleListApi: (data?: object) => {
            return request({
                url: "/role/getRoleList",
                method: "post",
                data,
            })
        },
        //修改role 信息
        modifyRoleInfoApi: (data?: object) => {
            return request({
                url: "/role/modifyRoleInfo",
                method: "post",
                data,
            })
        },
        //新建角色
        addRoleApi: (data?: object) => {
            return request({
                url: "/role/addRole",
                method: "post",
                data,
            })

        },
        //查询角色 by role name
        findRoleApi: (data?: object) => {
            return request({
                url: "/role/findRole",
                method: "post",
                data,
            })
        },
        //删除角色
        delRoleApi: (data?: object) => {
            return request({
                url: "/role/delRole",
                method: "delete",
                data,
            })
        },
        //获取当前角色的权限
        getPolicyByRoleIdsApi: (data?: object) => {
            return request({
                url: "/casbin/getPolicyByRoleIds",
                method: "post",
                data,
            })
        },
        //获取全部权限list
        getAllPolicyApi: (params?: object) => {
            return request({
                url: "/casbin/getAllPolicy",
                method: "get",
                params,
            })
        },
        //更新角色权限
        updateCasbinPolicyApi: (data?: object) => {
            return request({
                url: "/casbin/updateCasbinPolicyNew",
                method: "post",
                data,
            })
        },

    }
}

