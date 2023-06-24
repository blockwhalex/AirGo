<template>
  <div>
    <el-row :gutter="15" class="home-card-one mb15">
      <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" v-for="(v, k) in goodsList" :key="k">
        <div class="home-card-item">
          <el-card>
            <template #header>
              <div>
                <el-text class="card-header-left">{{ v.subject }}</el-text>
              </div>
            </template>
            <div class="card-text">
              <el-tag class="card-text-left">套餐流量</el-tag>
              <span class="card-text-right">{{ v.total_bandwidth }}GB</span>
            </div>
            <div class="card-text">
              <el-tag class="card-text-left" type="warning">有效期</el-tag>
              <span class="card-text-right">{{ v.expiration_date }}天</span>
            </div>
            <div class="card-text">
              <el-tag class="card-text-left" type="warning">价格</el-tag>
              <span class="card-text-right">¥{{ v.total_amount }}</span>
            </div>
            <div v-html="v.des"></div>
            <div style="margin-top: 10px;margin-bottom: 10px">
              <el-button size="large" @click="openSubmitOrderDialog(v)" color="#FC3D08">立即购买</el-button>
            </div>
          </el-card>
        </div>

      </el-col>
    </el-row>
    <!--引入提交订单弹窗-->
    <SubmitOrderDialog ref="SubmitOrderDialogRef" @openPurchaseDialog="openPurchaseDialog"></SubmitOrderDialog>
    <!-- 引入确认支付弹窗组件 -->
    <PurchaseDialog ref="PurchaseDialogRef" @openQRDialog="openQRDialog"></PurchaseDialog>
    <!-- 引入二维码弹窗 -->
    <QRDialog ref="QRDialogRef"></QRDialog>

  </div>

</template>

<script setup lang="ts">
import {defineAsyncComponent, onMounted, reactive, ref} from 'vue';
//导入store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";

const shopStore = useShopStore()
const {goodsList, shopData} = storeToRefs(shopStore)
//引入弹窗组件
const SubmitOrderDialog = defineAsyncComponent(() => import('/@/views/shop/submit_order_dialog.vue'))
const PurchaseDialog = defineAsyncComponent(() => import('/@/views/shop/purchase_dialog.vue'))
const QRDialog = defineAsyncComponent(() => import('/@/views/shop/QR_dialog.vue'))
const PurchaseDialogRef = ref()
const QRDialogRef = ref()
const SubmitOrderDialogRef = ref()

//加载时获取全部已启用商品
onMounted(() => {
  shopStore.getAllEnabledGoods()
})
//打开提交订单弹窗
const openSubmitOrderDialog = (goood: Goods) => {
  shopData.value.currentGoods = goood
  SubmitOrderDialogRef.value.openDialog()
}
//打开确认购买弹窗
const openPurchaseDialog = (goods: Goods) => {
  PurchaseDialogRef.value.openDialog()
}
//打开二维码弹窗
const openQRDialog = () => {
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

.el-card {
  background-image: url("../../assets/bgc/bg-3.svg");
  background-repeat: no-repeat;
  background-position: 100%, 100%;
}

.card-text {
  display: flex;
  justify-content: space-between;
  height: 35px
}

.card-text-left {
  margin-top: auto;
  margin-bottom: auto;
}

.card-text-right {
  margin-top: auto;
  margin-bottom: auto;
  font-size: 20px;
}

.card-header-left {
  font-size: 30px;
  color: #FC3D08;
}
</style>