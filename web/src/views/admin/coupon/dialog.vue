<template>
  <div>
    <el-dialog v-model="state.isShowDialog" title="折扣码" width="50%">
      <el-form :model="state.coupon" label-width="120px">
        <el-form-item label="名称">
          <el-input v-model="state.coupon.name"></el-input>
        </el-form-item>
        <el-form-item label="折扣率">
          <el-col :span="4">
            <el-input v-model.number="state.coupon.discount_rate" type="number"></el-input>
          </el-col>
          <el-col :span="20">(5折填0.5)</el-col>

        </el-form-item>
        <el-form-item label="限制次数">
          <el-col :span="4">
            <el-input v-model.number="state.coupon.limit" type="number"></el-input>
          </el-col>
          <el-col :span="20"></el-col>
        </el-form-item>
        <el-form-item label="到期时间">
          <el-date-picker
              v-model="state.coupon.expired_at"
              type="datetime"
              placeholder="选择到期时间"
              size="default"
          />
        </el-form-item>
      </el-form>
      <template #footer>
            <span class="dialog-footer">
                <el-button @click="closeDialog">取消</el-button>
                <el-button type="primary" @click="onSubmit(state.coupon)" color="#FC3D08">
                    提交
                </el-button>
            </span>
      </template>
    </el-dialog>

  </div>
</template>

<script lang="ts" setup>
import {reactive} from "vue";

import {ElMessage} from "element-plus";
//api
import {useCouponApi} from "/@/api/coupon";

const couponApi = useCouponApi()

const state = reactive({
  type: '',
  isShowDialog: false,
  coupon: {
    name: '',
    discount_rate: 0,
    limit: 0,
    expired_at: '',
  } as Coupon,
})

//打开弹窗
const openDialog = (type: string, data: Coupon) => {
  state.isShowDialog = true
  state.type = type
  switch (type) {
    case "add":
      state.coupon = {} as Coupon
      break
    case "edit":
      state.coupon = JSON.parse(JSON.stringify(data))
      break
  }

}
//关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false
}
//提交按钮
const onSubmit = (params: object) => {
  switch (state.type) {
    case "add":
      couponApi.newCouponApi(params).then((res) => {
        if (res.code === 0) {
          ElMessage.success(res.msg)
          emits('refresh')
        }
      })
      break
    case "edit":
      couponApi.updateCouponApi(params).then((res) => {
        if (res.code === 0) {
          ElMessage.success(res.msg)
          emits('refresh')
        }
      })
      break

  }


  closeDialog()
}
//子组件调用父组件
const emits = defineEmits(['refresh'])
//暴露变量
defineExpose({
  openDialog,
})
</script>

<style scoped>

</style>