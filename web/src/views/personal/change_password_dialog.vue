<template>
  <div >
    <el-dialog v-model="changePasswordDialog.isShowChangePasswordDialog" title="修改密码" width="500px">
      <el-form ref="userDialogFormRef" size="default" label-width="90px">
        <el-form-item label="新密码">
          <el-input v-model="registerData.password" placeholder="请输入密码" clearable></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input v-model="registerData.re_password" placeholder="请输入密码" clearable></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
				<span class="dialog-footer">
					<el-button @click="closeDialog" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">确认</el-button>
				</span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {ElMessage} from "element-plus";
//user store
import { useUserStore } from "/@/stores/userStore";
import { storeToRefs } from 'pinia';
const userStore=useUserStore()
const {registerData,changePasswordDialog} =storeToRefs(userStore)

// 打开弹窗
const openDialog =() => {
  changePasswordDialog.value.isShowChangePasswordDialog=true
};

// 关闭弹窗
const closeDialog = () => {
  changePasswordDialog.value.isShowChangePasswordDialog=false
};

// 提交
const onSubmit = () => {
  userStore.changePassword().then((res)=>{
    if (res.code===0){
      ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
  closeDialog()

};
defineExpose({
  openDialog,
})
</script>
