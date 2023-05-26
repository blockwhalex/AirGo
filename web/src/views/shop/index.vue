<template>
    <div>
        <el-row :gutter="15" class="home-card-one mb15">
            <el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="6" v-for="(v, k) in goodsList" :key="k">
                <div class="home-card-item">
                    <el-card class="box-card">
                        <template #header>
                            <div class="card-header">
                                <div class="block">
                                </div>
                                <el-tag class="ml-2" type="success">{{ v.subject }}</el-tag>
                            </div>
                        </template>
                        <div class="text item">
                            <el-tag class="ml-2" type="warning">套餐流量:</el-tag><span>{{ v.total_bandwidth }}GB</span>
                        </div>
                        <div class="text item">
                            <el-tag class="ml-2" type="warning">有效期:</el-tag><span>{{ v.expiration_date }}天</span>
                        </div>
                        <div class="text item">
                            <el-tag class="ml-2" type="warning">价格:</el-tag><span>{{ v.total_amount }}元</span>
                        </div>
                        <el-tag class="ml-2" type="info">有效期内不清零、不重置，长期有效</el-tag>
                        <div>
                            <el-button @click="openSubmitOrderDialog(v)" type="primary">立即购买</el-button>
                        </div>

                    </el-card>
                </div>
                   
            </el-col>
        </el-row>
        <!--引入提交订单弹窗-->
      <SubmitOrderDialog ref="SubmitOrderDialogRef" @openPurchaseDialog="openPurchaseDialog"></SubmitOrderDialog>
        <!-- 引入确认支付弹窗组件 -->
        <PurchaseDialog ref="PurchaseDialogRef" @openQRDialog="openQRDialog" ></PurchaseDialog>
        <!-- 引入二维码弹窗 -->
        <QRDialog ref="QRDialogRef" ></QRDialog>

    </div>
    
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue';
//导入api
import { useShopApi } from "../../api/shop/index";
const shopApi = useShopApi()
//导入二维码 js
import QRCode from 'qrcodejs2-fixes';

//判断移动端
import { isMobile } from "/@/utils/other";
// 定义二维码变量内容
const qrcodeRef = ref();
//导入store
import { storeToRefs } from 'pinia';
import { useShopStore } from "/@/stores/shopStore";
import Submit_order_dialog from "/@/views/shop/submit_order_dialog.vue";
const shopStore = useShopStore()
const { goodsList ,tableData} = storeToRefs(shopStore)
//引入弹窗组件
const SubmitOrderDialog =defineAsyncComponent(()=>import('/@/views/shop/submit_order_dialog.vue'))
const PurchaseDialog = defineAsyncComponent(()=>import('/@/views/shop/purchase_dialog.vue'))
const QRDialog = defineAsyncComponent(()=>import('/@/views/shop/QR_dialog.vue'))
const PurchaseDialogRef=ref()
const QRDialogRef=ref()
const SubmitOrderDialogRef=ref()

//加载时获取全部商品
onMounted(() => {
    shopStore.getAllGoods()
})
//打开提交订单弹窗
const openSubmitOrderDialog=(v:Goods)=>{
  tableData.value.currentGoods=v
  SubmitOrderDialogRef.value.openDialog()
}
//打开确认支付弹窗
const openPurchaseDialog=(goods:Goods)=>{
    //调用子组件打开弹窗
    PurchaseDialogRef.value.openDialog()
}
//打开二维码弹窗
const openQRDialog=()=>{
    //调用子组件打开弹窗
    QRDialogRef.value.openDialog()
}

</script>

<style scoped>
.home-card-item {
    width: 100%;
    height: 100%;
    border-radius: 4px;
    transition: all ease 0.3s;
    padding: 20px;
    overflow: hidden;
    background: var(--el-color-white);
    color: var(--el-text-color-primary);
    border: 1px solid var(--next-border-color-light);
}
</style>