<template>
  <div class="system-user-dialog-container">
    <el-dialog :title="userManageData.dialog.title" v-model="userManageData.dialog.isShowDialog" width="769px">
      <el-form ref="userDialogFormRef" :model="userManageData.dialog.userForm" size="default" label-width="90px">
        <el-row :gutter="35">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="账户邮箱">
              <el-input v-model="userManageData.dialog.userForm.user_name" placeholder="请输入账户名称"
                        clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="关联角色">
              <el-checkbox-group v-model="userManageData.dialog.check_list">
                <el-checkbox :label="v.role_name" v-for="(v,index) in roleManageData.roles.role_list"
                             :key="index"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="账户密码">
              <el-input v-model="userManageData.dialog.userForm.password" placeholder="请输入" type="password"
                        clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="用户状态">
              <!--							<el-switch v-model="userManageData.dialog.userForm.enable" inline-prompt :active-value="1" :inactive-value="0" ></el-switch>-->
              <el-switch v-model="userManageData.dialog.userForm.enable" inline-prompt></el-switch>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ userManageData.dialog.submitTxt }}</el-button>
				</span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue';
//user store
import {useUserStore} from '/@/stores/userStore'

const userStore = useUserStore()
import {storeToRefs} from 'pinia';

const {userManageData} = storeToRefs(userStore)
//role store
import {useRoleStore} from '/@/stores/roleStore'

const roleStore = useRoleStore()
const {roleManageData} = storeToRefs(roleStore)

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const userDialogFormRef = ref();

// 打开弹窗
const openDialog = (type: string, row: SysUser) => {

  if (type === 'edit') {
    userManageData.value.dialog.userForm = row;
    userManageData.value.dialog.title = '修改用户';
    userManageData.value.dialog.submitTxt = '修 改';
  } else {
    userManageData.value.dialog.title = '新增用户';
    userManageData.value.dialog.submitTxt = '新 增';
    // 清空表单，此项需加表单验证才能使用
    // nextTick(() => {
    // 	userDialogFormRef.value.resetFields();
    // });
  }
  userManageData.value.dialog.isShowDialog = true;
//打开时加载全部角色，用来设置用户角色
  //配置参数，不分页
  roleManageData.value.params.page_num = 0
  roleManageData.value.params.page_size = 10000
  roleStore.getRoleList()
};
// 关闭弹窗
const closeDialog = () => {
  userManageData.value.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = () => {
  if (userManageData.value.dialog.title === '新增用户') {
    userStore.newUser()
  } else {
    userStore.updateUser()
  }
  // emit('refresh');
  closeDialog();
  setTimeout(() => {
    userStore.getUserList()
  }, 1000)

};

// 暴露变量
defineExpose({
  openDialog,
});
</script>
