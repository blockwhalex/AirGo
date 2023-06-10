<template>
  <div>
    <!-- 支付弹窗 -->
    <el-dialog v-model="tableData.isShowPurchaseDialog" title="支付详情" width="30%">
      <div>
        <el-button type="success" plain>套餐：</el-button>
        {{ tableData.currentOrder.subject }}
      </div>
      <div>
        <el-button type="success" plain>金额：</el-button>
        {{ tableData.currentOrder.total_amount }}元
      </div>
      <div v-if="tableData.currentOrder.total_amount!=='0'">
        <el-button type="primary" plain>支付方式</el-button>
        <el-radio-group v-model="tableData.currentOrder.pay_type" class="ml-4">
          <el-radio :label="'alipay'">支付宝</el-radio>
          <el-radio :label="'wechat'" disabled>微信(暂未接入)</el-radio>
        </el-radio-group>

      </div>
      <template #footer>
            <span class="dialog-footer">
                <el-button @click="tableData.isShowPurchaseDialog = false">取消</el-button>
                <el-button type="warning" @click="onPurchase({id:tableData.currentGoods.id})">
                    确认支付
                </el-button>
            </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {ElMessage} from 'element-plus';
//导入store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";
import {isMobile} from "/@/utils/other";

const shopStore = useShopStore()
const {tableData} = storeToRefs(shopStore)
// 定义子组件向父组件传值/事件
//const emit = defineEmits(['purchase']);
//api
import {useShopApi} from '/@/api/shop/index'

const shopApi = useShopApi()
//打开弹窗
const openDialog = () => {
  tableData.value.isShowPurchaseDialog = true
}
//调用父组件
const emits = defineEmits(['openQRDialog'])

//购买按钮
const onPurchase = (params?: object) => {
  //api，传out_trade_no，pay_type
  shopApi.purchaseApi({
    out_trade_no: tableData.value.currentOrder.out_trade_no,
    pay_type: tableData.value.currentOrder.pay_type
  }).then((res) => {
    if (res.code === 0 && res.msg === "购买成功") {
      ElMessage.success("购买成功")
      tableData.value.isShowPurchaseDialog = false
      return

    } else if (res.code === 0 && res.data != "") {
      //保存支付二维码
      tableData.value.QRtext = res.data
      //手机端跳转支付页面
      const ok = isMobile()
      if (ok) {
        window.location.href = tableData.value.QRtext;
        return
      } else {
        //电脑端打开支付二维码弹窗
        console.log("电脑端打开支付二维码弹窗")
        emits('openQRDialog')
      }
    }
  }).catch(() => {
  })
  //关闭弹窗
  tableData.value.isShowPurchaseDialog = false
}

defineExpose({
  openDialog,// 打开弹窗
})
</script>

<style scoped>

</style>