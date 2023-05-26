<template>
	<el-form size="large" class="login-content-form">

		<el-form-item class="login-animation1">
			<el-input text placeholder="请输入邮箱" v-model="registerData.user_name" clearable autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-User /></el-icon>
				</template>
			</el-input>
		</el-form-item>

		<el-form-item class="login-animation1">
			<el-input text placeholder="请输入密码" v-model="registerData.password" clearable autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-Unlock /></el-icon>
				</template>
			</el-input>
		</el-form-item>

		<el-form-item class="login-animation1">
			<el-input text placeholder="重新输入密码" v-model="registerData.re_password" clearable autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-Unlock /></el-icon>
				</template>
			</el-input>
		</el-form-item>

    <el-form-item class="login-animation3" v-if="themeConfig.enable_email_code">
      <el-col :span="15">
        <el-input text maxlength="4" placeholder="请输入验证码" v-model="state.ruleForm.email_code" clearable autocomplete="off">
          <template #prefix>
            <el-icon class="el-input__icon"><ele-Position /></el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="8">
        <el-button class="login-content-code" type="primary" @click="onGetEmailCode">{{ state.isCountDown ? `${state.countDowdTime}s后重新获取` : "获取验证码" }}</el-button>
      </el-col>
    </el-form-item>

		<el-form-item class="login-animation3">
			<el-button @click="onRegister"  round type="primary" v-waves class="login-content-submit">
				<span>注 册</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts" name="loginMobile">
import { reactive } from 'vue';
//user store
import { useUserStore } from "/@/stores/userStore";
import { storeToRefs } from 'pinia';
import { ElMessage } from 'element-plus';
const userInfo = useUserStore()
const { registerData } = storeToRefs(userInfo)
//theme store
import { useThemeConfig } from '/@/stores/themeConfig';
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
//router
// import { useRouter } from 'vue-router';
// const router=useRouter()
//api
import {usePublicApi} from '/@/api/public/index'
const publicApi =usePublicApi()
//定义参数
const state = reactive({
  isShowPassword: false,
  isShowEmailCode:false,
  isCountDown:false,
  countDowdTime:60,
  ruleForm: {
    user_name: 'admin@oicq.com',
    password: 'admin',
    email_code: '',
  },
  loading: {
    signIn: false,
  },
});

//注册
const onRegister=()=>{
	userInfo.register().then((res)=>{
		if(res){
			ElMessage.success('注册成功，即将跳转...')
			setTimeout(()=>{
				window.location.href = '/'; // 去登录页
				//router.push('/'); // 去登录页
			},1000)	
		}

	})
}
const onGetEmailCode=()=>{
  publicApi.getEmailCodeApi(state.ruleForm).then((res)=>{
    if (res.code===0){
      state.ruleForm.email_code=res.data
      handleTimeChange()
    }
  })
};
//
const handleTimeChange = () => {
  if (state.countDowdTime <= 0) {
    state.isCountDown = false;
    state.countDowdTime  = 60;
  } else {
    setTimeout(() => {
      state.countDowdTime--;
      handleTimeChange();
    }, 1000);
  }
};




</script>




<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;

	@for $i from 1 through 4 {
		.login-animation#{$i} {
			opacity: 0;
			animation-name: error-num;
			animation-duration: 0.5s;
			animation-fill-mode: forwards;
			animation-delay: calc($i/10) + s;
		}
	}

	.login-content-code {
		width: 100%;
		padding: 0;
	}

	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}

}</style>
