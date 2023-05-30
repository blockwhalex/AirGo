<template>
  <div>
    <el-dialog v-model="tableData.isShowSubmitOrderDialog" title="订单详情" width="30%">
      <div>
        <el-button type="success" plain>套餐：</el-button>
        {{ tableData.currentGoods.subject }}
      </div>
      <div>
        <el-button type="success" plain>金额：</el-button>
        {{ tableData.currentGoods.total_amount }}元
      </div>
      <template #footer>
            <span class="dialog-footer">
                <el-button @click="tableData.isShowSubmitOrderDialog = false">取消</el-button>
                <el-button type="primary" @click="onSubmitOrder">
                    提交订单
                </el-button>
            </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>

import {ElMessage} from "element-plus";
//导入store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";
const shopStore = useShopStore()
const {tableData} = storeToRefs(shopStore)
//api
import {useShopApi} from '/@/api/shop/index'
const shopApi = useShopApi()
//打开弹窗
const openDialog = () => {
  tableData.value.isShowSubmitOrderDialog = true
}
//子组件调用父组件
const emits = defineEmits(['openPurchaseDialog'])
//提交订单按钮
const onSubmitOrder = () => {
  //传goods_id
  shopApi.preCreatePayApi({goods_id:tableData.value.currentGoods.id}).then((res) => {
    if (res.code === 0) {
      //保存订单信息到pinia
      tableData.value.currentOrder = res.data
      //关闭弹窗
      tableData.value.isShowSubmitOrderDialog = false
      //调用父组件 支付弹窗
      emits('openPurchaseDialog')

    } else {
      ElMessage.error("提交订单错误" + res.msg)
      //关闭弹窗
      tableData.value.isShowSubmitOrderDialog = false
    }
  })

}
defineExpose({
  openDialog,// 打开弹窗
})
</script>

<style scoped>

</style>