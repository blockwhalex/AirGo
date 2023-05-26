<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <el-button type="info" plain>只显示最近10次订单</el-button>
      <el-table :data="orderPersonal.allOrders.order_list" stripe fit  height="100%" style="width: 100%;">
        <el-table-column prop="subject"  label="商品标题" show-overflow-tooltip/>
<!--        <el-table-column prop="total_amount" label="订单金额" show-overflow-tooltip/>-->
        <el-table-column prop="trade_status" label="交易状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="info" v-if="scope.row.trade_status==='created'">订单已创建</el-tag>
            <el-tag type="warning" v-if="scope.row.trade_status==='WAIT_BUYER_PAY'">等待付款</el-tag>
            <el-tag type="success" v-if="scope.row.trade_status==='TRADE_SUCCESS'">支付成功</el-tag>
            <el-tag type="warning" v-if="scope.row.trade_status==='TRADE_CLOSED'">交易关闭</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button v-if="scope.row.trade_status == 'WAIT_BUYER_PAY'" size="small" text type="primary"
                       @click="onCompleteOrder(scope.row)">去支付
            </el-button>
          </template>
        </el-table-column>
      </el-table>
<!--      <el-pagination background v-model:current-page="orderPersonal.queryParams.page_num" v-model:page-size="orderPersonal.queryParams.page_size"-->
<!--                     :page-sizes="[10, 20, 30, 40]" layout="sizes, prev, pager, next" :total="orderPersonal.total"-->
<!--                     @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange"/>-->
    </el-card>
    <!-- 引入确认支付弹窗组件 -->
    <PurchaseDialog ref="PurchaseDialogRef" @openQRDialog="openQRDialog"></PurchaseDialog>
    <!-- 引入二维码弹窗 -->
    <QRDialog ref="QRDialogRef"></QRDialog>
  </div>
</template>

<script setup lang="ts">
import {storeToRefs} from "pinia";
import {defineAsyncComponent, onMounted, ref} from "vue";
//store
import {useOrderStore} from "/@/stores/orderStore";

const orderStore = useOrderStore()
const {orderPersonal} = storeToRefs(orderStore)
//store
import {useShopStore} from "/@/stores/shopStore";
import {isMobile} from "/@/utils/other";

const shopStore = useShopStore()
const {tableData} = storeToRefs(shopStore)
//引入弹窗组件
const PurchaseDialog = defineAsyncComponent(() => import('/@/views/shop/purchase_dialog.vue'))
const QRDialog = defineAsyncComponent(() => import('/@/views/shop/QR_dialog.vue'))
const PurchaseDialogRef = ref()
const QRDialogRef = ref()


onMounted(() => {
  orderStore.getOrder({pageNum: 1, pageSize: 5})
})
//完成未支付订单
const onCompleteOrder = (row: Order) => {
  //当前订单存入pinia
  tableData.value.currentOrder = row
  tableData.value.QRtext = row.qr_code
  if (tableData.value.currentOrder.qr_code === '') {
    //打开支付弹窗
    PurchaseDialogRef.value.openDialog()
  } else {
    //直接付款
    //手机端跳转支付页面
    const ok = isMobile()
    if (ok) {
      window.location.href = tableData.value.QRtext;
      return
    } else {
      //电脑端打开支付二维码弹窗
      openQRDialog()
    }
  }
}
//打开二维码弹窗
const openQRDialog = () => {
  //调用子组件打开弹窗
  QRDialogRef.value.openDialog()
}
// 分页改变
const onHandleSizeChange = (val: number) => {
  orderPersonal.value.queryParams.page_size = val;
  //getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  orderPersonal.value.queryParams.page_num = val;
  //getTableData();
};

</script>

<style scoped lang="scss">
.container {
  :deep(.el-card__body) {
    display: flex;
    flex-direction: column;
    flex: 1;
    overflow: auto;
    .el-table {
      flex: 1;
    }
  }
}
</style>