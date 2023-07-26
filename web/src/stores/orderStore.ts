import {defineStore} from "pinia";
//api
import {useOrderApi} from "/@/api/order/index";
import {ElMessage} from "element-plus";

const orderApi = useOrderApi()
export const useOrderStore = defineStore("orderStore", {
    state: () => ({
        //订单管理页面数据
        orderManageData: {
            allOrders: {
                order_list: [] as Order[],
                total: 0,
            },
        },
        //个人订单数据
        orderPersonal: {
            allOrders: {
                order_list: [] as Order[],
                total: 0,
            },
        },
    }),
    actions: {
        //获取订单详情(下单时）
        async getOrderInfo(params: object) {
            const res = await orderApi.getOrderInfoApi(params)
            return res
        },
        //获取全部订单
        async getAllOrder(params?: object) {
            const res = await orderApi.getAllOrderApi(params)
            if (res.code === 0) {
                this.orderManageData.allOrders = res.data
                ElMessage.success(res.msg)
            }
        },
        //获取用户最近10次订单
        async getOrder(params?: object) {
            const res = await orderApi.getOrderApi()
            if (res.code === 0) {
                this.orderPersonal.allOrders.order_list = res.data
                ElMessage.success(res.msg)
            }
        },
        //完成未支付订单
        async completedOrder(params?: object) {
            const res = await orderApi.completedOrderApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
    }
})