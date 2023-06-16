import service from "/@/utils/request";

export function useOrderApi() {
    return {
        //获取订单详情(下单时）
        getOrderInfoApi: (data?: object) => {
            return service({
                url: '/order/getOrderInfo',
                method: 'POST',
                data
            })
        },
        //获取全部订单
        getAllOrderApi: (data?: object) => {
            return service({
                url: '/order/getAllOrder',
                method: 'POST',
                data
            })
        },
        //获取用户订单
        getOrderApi: (data?: object) => {
            return service({
                url: '/order/getOrderByUserID',
                method: 'POST',
                // data
            })
        },
        //完成未支付订单
        completedOrderApi: (data?: object) => {
            return service({
                url: '/order/completedOrder',
                method: 'POST',
                data
            })
        },
        //获取订单统计
        getMonthOrderStatisticsApi: (data?: object) => {
            return service({
                url: '/order/getMonthOrderStatistics',
                method: 'POST',
                data
            })
        },
    }

}