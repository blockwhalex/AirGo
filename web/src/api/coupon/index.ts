import service from "/@/utils/request";

export function useCouponApi() {
    return {
        //新建
        newCouponApi: (data?: object) => {
            return service({
                url: '/coupon/newCoupon',
                method: 'POST',
                data
            })
        },
        // 删除折扣
        deleteCouponApi: (data?: object) => {
            return service({
                url: '/coupon/deleteCoupon',
                method: 'POST',
                data
            })
        },
        // 更新折扣
        updateCouponApi: (data?: object) => {
            return service({
                url: '/coupon/updateCoupon',
                method: 'POST',
                data
            })
        },
        // 获取折扣列表
        getCouponApi: (data?: object) => {
            return service({
                url: '/coupon/getCoupon',
                method: 'POST',
                data
            })
        },
    }

}