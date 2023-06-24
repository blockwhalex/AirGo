<template>
  <div>
    <!-- 支付弹窗 -->
    <el-dialog v-model="state.isShowPurchaseDialog" title="等待支付" width="30%">

      <div class="home-card-item">
        <el-card>
          <div class="card-text">
            <el-button type="primary" plain>套餐：</el-button>
            <el-text class="card-text-right">{{ shopData.currentOrder.subject }}</el-text>
          </div>
          <div class="card-text">
            <el-button type="primary" plain>金额：</el-button>
            <el-text class="card-text-right">{{ shopData.currentOrder.total_amount }}元</el-text>
          </div>
        </el-card>
      </div>

      <div class="home-card-item">
        <el-card>
          <div v-if="shopData.currentOrder.total_amount!=='0'">
            <el-button type="primary" plain>支付方式</el-button>
            <div>
              <el-radio-group style="height: 60px;" v-model="shopData.currentOrder.pay_type" class="ml-4">
                <el-radio :label="'alipay'">
                  <div style="display: flex;align-items: center">
                    <el-image :src="alipayIcon" style="height: 15px;"></el-image>
                    支付宝
                  </div>
                </el-radio>
                <el-radio :label="'wechat'" disabled>
                  <div style="display: flex;align-items: center">
                    <el-image :src="wechatpayIcon" style="height: 15px;"></el-image>
                    微信(暂未接入)
                  </div>
                </el-radio>
              </el-radio-group>
            </div>
          </div>
        </el-card>
      </div>
      <template #footer>
            <span class="dialog-footer">
                <el-button @click="state.isShowPurchaseDialog = false">取消</el-button>
                <el-button color="#FC3D08"
                           :disabled="shopData.currentOrder.pay_type==='' && shopData.currentOrder.total_amount!=='0'"
                           @click="onPurchase({id:shopData.currentGoods.id})">
                    确认支付
                </el-button>
            </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {ElMessage} from 'element-plus';
import {reactive} from "vue";
//导入store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";
import {isMobile} from "/@/utils/other";

const shopStore = useShopStore()
const {shopData} = storeToRefs(shopStore)

//api
import {useShopApi} from '/@/api/shop/index'

const shopApi = useShopApi()
//图标
import alipayIcon from "/@/assets/icon/alipay.jpeg"
import wechatpayIcon from "/@/assets/icon/wechatpay.jpeg"

//定义变量
const state = reactive({
  isShowPurchaseDialog: false,
})
//打开弹窗
const openDialog = () => {
  state.isShowPurchaseDialog = true
}
//调用父组件
const emits = defineEmits(['openQRDialog'])

//购买按钮
const onPurchase = (params?: object) => {

  //api，传out_trade_no，pay_type
  shopApi.purchaseApi({
    out_trade_no: shopData.value.currentOrder.out_trade_no,
    pay_type: shopData.value.currentOrder.pay_type
  }).then((res) => {
    if (res.code === 0 && res.msg === "购买成功") { //0元购
      ElMessage.success("购买成功")
      state.isShowPurchaseDialog = false
      return
    } else if (res.code === 0 && res.data != "") { //后端返回qrcode
      //保存支付二维码
      shopData.value.currentOrder.qr_code = res.data
      const ok = isMobile()            //手机端跳转支付页面
      if (ok) {
        window.location.href = shopData.value.currentOrder.qr_code;
        return
      } else {
        emits('openQRDialog')   //电脑端打开支付二维码弹窗
      }
    }
  }).catch(() => {
  })
  //关闭弹窗
  state.isShowPurchaseDialog = false
}

defineExpose({
  openDialog,// 打开弹窗
})
</script>

<style scoped>
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