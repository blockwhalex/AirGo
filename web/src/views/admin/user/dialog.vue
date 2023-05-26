<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="userManageData.dialog.title" v-model="userManageData.dialog.isShowDialog" width="769px">
			<el-form ref="userDialogFormRef" :model="userManageData.userInfos" size="default" label-width="90px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户邮箱">
							<el-input v-model="userManageData.userInfos.user_name" placeholder="请输入账户名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="关联角色">
              <el-checkbox-group v-model="userManageData.dialog.checkList"  >
                <el-checkbox :label="v.roleName" v-for="(v,index) in tableData.data" :key="index"></el-checkbox>
              </el-checkbox-group>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户密码">
							<el-input v-model="userManageData.userInfos.password" placeholder="请输入" type="password" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="用户状态">
<!--							<el-switch v-model="userManageData.userInfos.enable" inline-prompt :active-value="1" :inactive-value="0" ></el-switch>-->
              <el-switch v-model="userManageData.userInfos.enable" inline-prompt  ></el-switch>
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

<script setup lang="ts" name="systemUserDialog">
import { reactive,ref } from 'vue';
//user store
import {useUserStore} from '/@/stores/userStore'
const userStore=useUserStore()
import { storeToRefs } from 'pinia';
const { userManageData } = storeToRefs(userStore)
//role store
import {useRoleStore} from '/@/stores/roleStore'
const roleStore=useRoleStore()
const {tableData} =storeToRefs(roleStore)

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const userDialogFormRef = ref();
const state = reactive({
	ruleForm: {
		userName: '', // 账户名称
		userNickname: '', // 用户昵称
		roleSign: '', // 关联角色
		department: [] as string[], // 部门
		phone: '', // 手机号
		email: '', // 邮箱
		sex: '', // 性别
		password: '', // 账户密码
		overdueTime: '', // 账户过期
		status: true, // 用户状态
		describe: '', // 用户描述
	},
	deptData: [] as DeptTreeType[], // 部门数据
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: SysUser) => {
	
	if (type === 'edit') {
		userManageData.value.userInfos = row;
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
//打开时加载角色信息
  roleStore.setRoleList()
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
  if (userManageData.value.dialog.title === '新增用户'){
    userStore.newUser()
  } else {
    userStore.updateUser()
  }
  // emit('refresh');
  closeDialog();
  setTimeout(()=>{
    userStore.getUserList()
  },1000)

};

// 暴露变量
defineExpose({
	openDialog,
});
</script>
