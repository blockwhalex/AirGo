import { defineStore } from "pinia";
//api
import { useOrderApi } from "/@/api/order/index";
const orderApi = useOrderApi()
export const useOrderStore = defineStore("orderStore", {
    state: () => ({
        //订单管理页面数据
        orderManageData: {
            queryParams:{
                page_num: 1,
                page_size: 10,
                search: '',
                date: [],
            } as QueryParams,
            allOrders: {
                order_list:[],
                total:0,
            } as OrdersWithTotal,
        },
        //个人订单数据
        orderPersonal: {
            queryParams:{
                page_num: 1,
                page_size: 10,
                search: '',
                date: [],
            } as QueryParams,
            allOrders: {
                order_list:[],
                total:0,
            } as OrdersWithTotal,
        }
    }),
    actions: {
        //获取全部订单
        async getAllOrder(params?: object) {
            const res = await orderApi.getAllOrderApi(this.orderManageData.queryParams)
            this.orderManageData.allOrders = res.data
           // this.orderManageData.queryParams.total=this.orderManageData.allOrders.length
        },
        //获取用户全部订单
        async getOrder(params?: object) {
            const res = await orderApi.getOrderApi(this.orderPersonal.queryParams)
            this.orderPersonal.allOrders.order_list = res.data
           // this.orderPersonal.queryParams.total=this.orderPersonal.allOrders.length
        },
        //完成未支付订单
        async completedOrder(params?: object) {
            const res = await orderApi.completedOrderApi(params)
            return res

        },
    }
})