//路由-store
import {defineStore} from "pinia";
import {useMenuApi} from "/@/api/menu/index";

const menuApi = useMenuApi()

export const useRoutesStore = defineStore("routesStore", {
    state: () => ({
        //所有的动态路由list
        allRoutesList: [] as Route[],
        //当前角色的动态路由list
        routesListSate: {
            routesList: [],
            isColumnsMenuHover: false,
            isColumnsNavHover: false,
        } as RoutesListState,
        //所有动态路由tree（路由tree）
        routesTree: [] as RoutesTree[],
        //当前角色动态路由tree（路由tree）
        currentRoleRoutesTree: [] as RoutesTree[],
    }),
    actions: {
        // 设置动态路由到 pinia routesList 中
        async setRoutesList(data: Array<Route>) {
            this.routesListSate.routesList = data;
        },
        async setColumnsMenuHover(bool: Boolean) {
            this.routesListSate.isColumnsMenuHover = bool;
        },
        async setColumnsNavHover(bool: Boolean) {
            this.routesListSate.isColumnsNavHover = bool;
        },
        //所有动态路由tree（路由tree）
        async setRoutesTree() {
            const res: any = await menuApi.getAllRouteTreeApi()
            this.routesTree = res.data

        },
        //当前角色动态路由tree（路由tree）
        async setCurrentRoleRoutesTree(params: any) {
            const res: any = await menuApi.getRouteTreeApi(params)
            this.currentRoleRoutesTree = res.data

        },
        //所有的动态路由list
        async setAllRoutesList() {
            const res = await menuApi.getAllRouteListApi()
            this.allRoutesList = res.data
        },
        //查询动态路由 by meta.title
        async findRoutesListByTitle(params: object) {
            const res = await menuApi.findDynamicRouteApi(params)
            this.allRoutesList = res.data
        },
        //新建动态路由
        async newDynamicRoute(params: object) {
            const res = menuApi.newDynamicRouteApi(params)
        },
        //删除动态路由
        async delDynamicRoute(params: object) {
            const res = menuApi.delDynamicRouteApi(params)
        },
        //更新动态路由
        async updateDynamicRoute(params: object) {
            const res = menuApi.updateDynamicRouteApi(params)
        },

    }
})