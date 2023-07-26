import {defineStore} from "pinia";

export const useCouponStore = defineStore("couponStore", {
    state: () => ({
        couponList:[] as Coupon[],
    }),
    actions: {
    }

})