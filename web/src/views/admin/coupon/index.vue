<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-button size="default" type="primary" class="ml10" @click="openDialog('add')">
          <el-icon>
            <ele-FolderAdd/>
          </el-icon>
          新建折扣码
        </el-button>
      </div>
      <div>
        <el-table :data="state.couponList" stripe style="width: 100%;flex: 1;">
          <el-table-column type="index" label="序号" width="60"/>
          <el-table-column prop="name" label="名称" width="180"/>
          <el-table-column prop="id" label="ID" width="60"/>
          <el-table-column prop="discount_rate" label="折扣率" width="60"/>
          <el-table-column prop="limit" label="限制次数" width="100"/>
          <el-table-column prop="expired_at" label="到期时间" width="180">
          <template #default="{row}">
            {{DateStrtoTime(row.expired_at)}}
          </template>
          </el-table-column>
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button @click="openDialog('edit',scope.row)">编辑</el-button>
              <el-button @click="opDeleteCoupon(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <CouponDialog ref="couponDialogRef" @refresh="getCoupon()"></CouponDialog>
  </div>
</template>

<script lang="ts" setup>

import {defineAsyncComponent, onMounted, reactive, ref} from "vue";
//引入组件
const CouponDialog = defineAsyncComponent(()=>import('/@/views/admin/coupon/dialog.vue'))
const couponDialogRef=ref()
//api
import {useCouponApi} from "/@/api/coupon";
import {DateStrtoTime} from "/@/utils/formatTime";
import {ElMessage, ElMessageBox} from "element-plus";
const couponApi =useCouponApi()
//时间


const state =reactive({
  isShowDialog:false,
  couponList:[] as Coupon[],
  coupon:{} as Coupon,
})
//
const getCoupon=()=>{
  couponApi.getCouponApi().then((res)=>{
    state.couponList=res.data
  })
}

//
const openDialog=(type:string,data:Coupon)=>{
  couponDialogRef.value.openDialog(type,data)
}
//
const opDeleteCoupon=(row:Coupon)=>{
  ElMessageBox.confirm(`此操作将永久删除折扣：${row.name}, 是否继续?`, '提示', {
    confirmButtonText: '删除',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        //逻辑
        couponApi.deleteCouponApi(row).then((res)=>{
          ElMessage.success(res.msg);
        })
        setTimeout(() => {
          getCoupon()
          //逻辑
        }, 1000);
      })
      .catch(() => {
      });

}
//
onMounted(()=>{
  getCoupon()

});

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