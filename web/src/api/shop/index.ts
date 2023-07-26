import request from '/@/utils/request';


export function useShopApi() {
    return {
        preCreatePayApi: (data?: object) => {
            return request({
                url: "/shop/preCreatePay",
                method: "post",
                data
            })
        },
        alipayTradePreCreatePayApi: (data?: object) => {
            return request({
                url: "/shop/alipayTradePreCreatePay",
                method: "post",
                data
            })
        },
        //发起支付
        purchaseApi: (data?: object) => {
            return request({
                url: "/shop/purchase",
                method: "post",
                data
            })
        },
        //获取全部商品
        getAllEnabledGoodsApi: () => {
            return request({
                url: "/shop/getAllEnabledGoods",
                method: "get",
            })
        },
        //获取全部商品
        getAllGoodsApi: () => {
            return request({
                url: "/shop/getAllGoods",
                method: "get",
            })
        },
        //添加商品
        newGoodsApi: (data?: object) => {
            return request({
                url: "/shop/newGoods",
                method: "post",
                data
            })

        },
        //修改商品
        updateGoodsApi: (data?: object) => {
            return request({
                url: "/shop/updateGoods",
                method: "post",
                data
            })
        },
        //删除商品
        deleteGoodsApi: (data?: object) => {
            return request({
                url: "/shop/deleteGoods",
                method: "post",
                data
            })
        },
        //商品排序
        goodsSortApi: (data?: object) => {
            return request({
                url: "/shop/goodsSort",
                method: "post",
                data
            })
        },
    }
}
