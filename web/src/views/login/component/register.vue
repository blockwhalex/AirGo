<template>
  <el-form size="large" class="login-content-form">

    <el-form-item>
      <el-input text placeholder="请输入邮箱" v-model="registerData.user_name" clearable autocomplete="off">
        <template #prefix>
          <el-icon>
            <ele-User/>
          </el-icon>
        </template>
        <template #append>
          <el-select v-model="email_suffix" style="width: 130px;">
            <el-option label="qq.com" value="qq.com"/>
            <el-option label="gmail.com" value="gmail.com"/>
            <el-option label="163.com" value="163.com"/>
            <el-option label="outlook.com" value="163.com"/>
            <el-option label="hotmail.com" value="163.com"/>
            <el-option label="foxmail.com" value="163.com"/>
          </el-select>
        </template>
      </el-input>
    </el-form-item>

    <el-form-item>
      <el-input text placeholder="请输入密码" v-model="registerData.password" clearable autocomplete="off">
        <template #prefix>
          <el-icon>
            <ele-Unlock/>
          </el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item>
      <el-input text placeholder="重新输入密码" v-model="registerData.re_password" clearable autocomplete="off">
        <template #prefix>
          <el-icon>
            <ele-Unlock/>
          </el-icon>
        </template>
      </el-input>
    </el-form-item>

    <el-form-item v-if="themeConfig.enable_email_code">
      <el-col :span="13">
        <el-input text maxlength="4" placeholder="请输入验证码" v-model="registerData.email_code" clearable
                  autocomplete="off">
          <template #prefix>
            <el-icon>
              <ele-Position/>
            </el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="10">
        <el-button class="login-content-code" type="primary" :disabled="state.isCountDown" @click="onGetEmailCode">
          {{ state.isCountDown ? `${state.countDownTime}s后重新获取` : "获取验证码" }}
        </el-button>
      </el-col>
    </el-form-item>

    <el-form-item>
      <el-button @click="onRegister" round type="primary" v-waves class="login-content-submit">
        <span>注 册</span>
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts" name="loginMobile">
import {reactive} from 'vue';
//user store
import {useUserStore} from "/@/stores/userStore";
import {storeToRefs} from 'pinia';
import {ElMessage} from 'element-plus';

const userStore = useUserStore()
const {registerData, email_suffix} = storeToRefs(userStore)
//theme store
import {useThemeConfig} from '/@/stores/themeConfig';

const storesThemeConfig = useThemeConfig();
const {themeConfig} = storeToRefs(storesThemeConfig);
//api
import {usePublicApi} from '/@/api/public/index'

const publicApi = usePublicApi()
//定义参数
const state = reactive({
  isShowPassword: false,
  isCountDown: false,
  countDownTime: 60,
  loading: {
    signIn: false,
  },
});

//注册
const onRegister = () => {
  userStore.register().then((res) => {
    if (res.code === 0) {
      ElMessage.success('注册成功，前往登录...')
      setTimeout(() => {
        window.location.href = '/'; // 去登录页
        //router.push('/'); // 去登录页
      }, 1000)
    }
  })
}
const onGetEmailCode = () => {
  if (registerData.value.user_name === '') {
    return
  }
  state.isCountDown = true
  //console.log("userStore.userFormReq:",userStore.userFormReq)
  publicApi.getEmailCodeApi(userStore.userFormReq).then((res) => {
    if (res.code === 0) {
      ElMessage.success(res.msg)
      handleTimeChange()
    } else {

    }
  })

};
//倒计时
const handleTimeChange = () => {
  if (state.countDownTime <= 0) {
    state.isCountDown = false;
    state.countDownTime = 60;
  } else {
    setTimeout(() => {
      state.countDownTime--;
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
  }

}</style>
