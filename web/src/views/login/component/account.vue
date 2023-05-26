<!-- eslint-disable no-console -->
<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation1">
			<el-input text placeholder="用户名" v-model="loginData.user_name" clearable
				autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-User /></el-icon>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation2">
			<el-input :type="state.isShowPassword ? 'text' : 'password'" placeholder="密码"
				v-model="loginData.password" autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-Unlock /></el-icon>
				</template>
				<template #suffix>
					<i class="iconfont el-input__icon login-content-password"
						:class="state.isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
						@click="state.isShowPassword = !state.isShowPassword">
					</i>
				</template>
			</el-input>
		</el-form-item>

    <el-form-item class="login-animation3" v-if="state.enableResetPassword">
        <el-col :span="15">
          <el-input text maxlength="4" placeholder="请输入验证码" v-model="loginData.email_code" clearable autocomplete="off">
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

    <el-form-item>
      <el-button v-if="!state.enableResetPassword" type="primary" class="login-content-submit" round @click="onSignIn">
				<span>登 录</span>
			</el-button>
      <el-button v-if="!state.enableResetPassword" class="login-content-resetPassword" round  @click="onResetPassword">
        <span>重置密码</span>
      </el-button>
      <el-button v-if="state.enableResetPassword" @click="onSubmitResetPassword" class="login-content-resetPassword" round type="danger" plain>
        <span>确认重置密码</span>
      </el-button>
      <el-button v-if="state.enableResetPassword" @click="state.enableResetPassword=false" class="login-content-resetPassword" round type="primary" plain>
        <span>返回登录</span>
      </el-button>
		</el-form-item>
		<div class="font12 mt30  login-msg">
			* 建议使用谷歌、Microsoft Edge，360浏览器请使用极速模式
		</div>
	</el-form>
</template>

<script setup lang="ts" name="loginAccount">
import {reactive, computed, onMounted} from 'vue';
import { ElMessage } from 'element-plus';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { NextLoading } from '/@/utils/loading';
//后端路由
import { initBackEndControlRoutes } from '/@/router/backEnd';
//route
import { useRoute, useRouter } from 'vue-router';
const route = useRoute();
const router = useRouter();
//user store
import { storeToRefs } from 'pinia';
import { useUserStore } from "/@/stores/userStore";
const userStore = useUserStore()
const { loginData } = storeToRefs(userStore)
//theme store
import { useThemeConfig } from '/@/stores/themeConfig';
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
//api
import {usePublicApi} from '/@/api/public/index'
const publicApi =usePublicApi()
//定义参数
const state = reactive({
  enableResetPassword:false,
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

// 时间获取
const currentTime = computed(() => {
	return formatAxis(new Date());
});
//重置密码
const onResetPassword=()=>{
  state.enableResetPassword=true
}
//确认重置密码
const onSubmitResetPassword=()=>{
  userStore.submitResetPassword().then((res)=>{
    if (res.code===0){
      ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
}
// 登录
const onSignIn = async () => {
	//state.loading.signIn = true;
  await userStore.userLogin(loginData.value)
	//添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
	const isNoPower = await initBackEndControlRoutes();
	//执行完 initBackEndControlRoutes，再执行 signInSuccess
	signInSuccess(isNoPower);

};
// 登录成功后的跳转
const signInSuccess = (isNoPower: boolean | undefined) => {
	if (isNoPower) {
		ElMessage.warning('抱歉，您没有登录权限');
		Session.clear();
	} else {
		// 初始化登录成功时间问候语
		let currentTimeInfo = currentTime.value;
		// 登录成功，跳到转首页
		// 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
		if (route.query?.redirect) {
			router.push({
				path: <string>route.query?.redirect,
				query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
			});
		} else {
			router.push('/');
		}
		// 登录成功提示
		const signInText = '欢迎回来！';
		ElMessage.success(`${currentTimeInfo}，${signInText}`);
		// 添加 loading，防止第一次进入界面时出现短暂空白
		NextLoading.start();
	}
	//state.loading.signIn = false;
};
//获取邮箱验证码
const onGetEmailCode=()=>{
  publicApi.getEmailCodeApi(loginData.value).then((res)=>{
    if (res.code===0){
      loginData.value.email_code=res.data
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
//
onMounted(()=>{

});



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

	.login-content-password {
		display: inline-block;
		width: 20px;
		cursor: pointer;

		&:hover {
			color: #909399;
		}
	}

	.login-content-code {
		width: 100%;
		padding: 0;
		font-weight: bold;
	//	letter-spacing: 5px;
	}

	.login-content-submit {
		width: 45%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
  .login-content-resetPassword{
    width: 45%;
    letter-spacing: 2px;
    font-weight: 300;
    margin-top: 15px;
  }
	.login-msg {
		color: var(--el-text-color-placeholder);
	}
}</style>
