<template>
  <div>
    <el-dialog v-model="state.isShowSubmitOrderDialog" title="订单详情" width="30%">
      <div>
        <el-button type="success" plain>套餐：</el-button>
        {{ tableData.currentOrder.subject }}
      </div>
      <div>
        <el-button type="success" plain>金额：</el-button>
        {{ tableData.currentOrder.total_amount }}元
      </div>
      <template #footer>
            <span class="dialog-footer">
                <el-button @click="state.isShowSubmitOrderDialog = false">取消</el-button>
                <el-button type="primary" @click="onSubmitOrder">
                    提交订单
                </el-button>
            </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {reactive} from 'vue';
import {ElMessage} from "element-plus";
//shop store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";

const shopStore = useShopStore()
const {tableData} = storeToRefs(shopStore)
//order store
import {useOrderStore} from "/@/stores/orderStore"

const orderStore = useOrderStore()

//api
import {useShopApi} from '/@/api/shop/index'

const shopApi = useShopApi()
import {useOrderApi} from '/@/api/order/index'

const orderApi = useOrderApi()
//变量
const state = reactive({
  isShowSubmitOrderDialog: false,

})
//打开弹窗
const openDialog = (id: number) => {
  state.isShowSubmitOrderDialog = true
  //获取订单详情（计算价格等）
  orderStore.getOrderInfo(id).then((res) => {
    if (res.code === 0) {
      ElMessage.success(res.msg)
      tableData.value.currentOrder = res.data

    } else {
      ElMessage.error(res.msg)
    }
  }).catch()

}
//子组件调用父组件
const emits = defineEmits(['openPurchaseDialog'])
//提交订单按钮
const onSubmitOrder = () => {
  //传goods_id
  shopApi.preCreatePayApi({goods_id: tableData.value.currentOrder.goods_id}).then((res) => {
    if (res.code === 0) {
      //保存订单信息到pinia
      tableData.value.currentOrder = res.data
      //关闭弹窗
      state.isShowSubmitOrderDialog = false
      //调用父组件 支付弹窗
      emits('openPurchaseDialog')

    } else {
      ElMessage.error("提交订单错误" + res.msg)
      //关闭弹窗
      state.isShowSubmitOrderDialog = false
    }
  })

}
defineExpose({
  openDialog,// 打开弹窗
})
</script>

<style scoped>

</style>