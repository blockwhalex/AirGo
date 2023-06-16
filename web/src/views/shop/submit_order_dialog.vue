<template>
  <div>
    <el-dialog v-model="state.isShowSubmitOrderDialog" title="订单详情" width="30%">
      <div>
        <el-button type="success" plain>套餐：</el-button>
        {{ shopData.currentOrder.subject }}
      </div>
      <div>
        <el-button type="success" plain>金额：</el-button>
        {{ shopData.currentOrder.total_amount }}元
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
const {shopData} = storeToRefs(shopStore)
//order store
import {useOrderStore} from "/@/stores/orderStore"

const orderStore = useOrderStore()

//api
import {useShopApi} from '/@/api/shop/index'

const shopApi = useShopApi()

//定义变量
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
      shopData.value.currentOrder = res.data
    }
  }).catch()
}

//提交订单按钮
const onSubmitOrder = () => {
  //传goods_id
  shopApi.preCreatePayApi({goods_id: shopData.value.currentOrder.goods_id}).then((res) => {
    if (res.code === 0) {
      //保存订单信息到pinia
      shopData.value.currentOrder = res.data
      //关闭弹窗
      state.isShowSubmitOrderDialog = false
      //调用父组件 支付弹窗
      emits('openPurchaseDialog')
    } else {
      state.isShowSubmitOrderDialog = false
    }
  })
}
//子组件调用父组件
const emits = defineEmits(['openPurchaseDialog'])
//暴露变量
defineExpose({
  openDialog,
})
</script>

<style scoped>

</style>