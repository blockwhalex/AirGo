import {defineStore} from "pinia";
//api
import {useOrderApi} from "/@/api/order/index";

const orderApi = useOrderApi()
export const useOrderStore = defineStore("orderStore", {
    state: () => ({
        //订单管理页面数据
        orderManageData: {
            queryParams: {
                page_num: 1,
                page_size: 10,
                search: '',
                date: [],
            } as QueryParams,
            allOrders: {
                order_list: [] as Order[],
                total: 0,
            },
        },
        //个人订单数据
        orderPersonal: {
            queryParams: {
                page_num: 1,
                page_size: 10,
                search: '',
                date: [],
            } as QueryParams,
            allOrders: {
                order_list: [] as Order[],
                total: 0,
            },
        },
    }),
    actions: {
        //获取订单详情(下单时）
        async getOrderInfo(id?: number) {
            const res = await orderApi.getOrderInfoApi({goods_id: id})
            return res

        },
        //获取全部订单
        async getAllOrder(params?: object) {
            const res = await orderApi.getAllOrderApi(this.orderManageData.queryParams)
            this.orderManageData.allOrders = res.data
        },
        //获取用户全部订单
        async getOrder(params?: object) {
            const res = await orderApi.getOrderApi(this.orderPersonal.queryParams)
            this.orderPersonal.allOrders.order_list = res.data
        },
        //完成未支付订单
        async completedOrder(params?: object) {
            const res = await orderApi.completedOrderApi(params)
            return res
        },
    }
})