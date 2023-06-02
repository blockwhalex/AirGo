<template>
	<div>
		<el-row :gutter="10" class="home-card-two mb15">
			<el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
				<div class="home-card-item">
					<el-card class="box-card">
						<template #header>
							<div class="card-header">
								<el-tag class="ml-2" plain>套餐详情</el-tag>
							</div>
						</template>
						<div class="text item">
							<el-tag class="ml-2" type="info">剩余流量:</el-tag><span>{{ (userInfos.subscribe_info.t-userInfos.subscribe_info.d-userInfos.subscribe_info.u)/1024/1024/1024 }}GB</span>
						</div>
						<div class="text item">
							<el-tag class="ml-2" type="info">到期时间:</el-tag>{{ DateStrtoTime(userInfos.subscribe_info.expired_at) }}
						</div>
					</el-card>
				</div>
			</el-col>
			<el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
				<div class="home-card-item">
					<el-card class="box-card" style="width: 100%;flex: 1;">
						<template #header>
							<div class="card-header">
								<el-tag   plain>当前混淆:</el-tag>
								<el-tag  type="info" >{{ userInfos.subscribe_info.host }}</el-tag>
							</div>
						</template>
						<div class="mt-4">
							<el-input v-model="homeTableData.host" placeholder="输入混淆"
								class="input-with-select">
								<template #append>
									<el-button @click="onChangeHost" :icon="Select">确认修改</el-button>
								</template>
							</el-input>
						</div>

					</el-card>
				</div>
			</el-col>
		</el-row>
		<el-row :gutter="10">
			<el-col :xs="24" :sm="14" :md="14" :lg="16" :xl="16">
				<div class="home-card-item"> 
					<el-card class="box-card">
						<template #header>
							<div class="card-header">
								<span>订阅地址</span>
                <el-button type="primary" plain class="button" @click="onResetSub" >重置订阅链接</el-button>
							</div>
						</template>
						<div>
							<el-button @click="v2rayNGSub(1)" type="primary" plain>复制v2rayNG订阅</el-button>
							<el-button @click="v2rayNGSub(2)" type="success" plain>复制Clash Meta订阅</el-button>
							<el-button @click="v2rayNGSub(3)" type="info" plain>复制shadowrocket订阅</el-button>
							<el-button @click="v2rayNGSub(4)" type="warning" plain>复制Quantumult X订阅</el-button>
						</div>
					</el-card>
				</div>
			</el-col>
		</el-row>
	</div>
</template>

<script setup lang="ts">
//时间转换
import {DateStrtoTime} from "/@/utils/formatTime";
//api
import {useUserApi} from '/@/api/user/index'
const userApi=useUserApi()

//store
import { storeToRefs } from "pinia";
import { useUserStore } from "/@/stores/userStore";
import {  onMounted } from 'vue';
const userStore = useUserStore()
const {userInfos,homeTableData} = storeToRefs(userStore)
//ElMessage
import { ElMessage } from 'element-plus';
//icon 
import { Select } from '@element-plus/icons-vue'
//复制剪切板
import commonFunction from '/@/utils/commonFunction';

const { copyText } = commonFunction();


//修改混淆
const onChangeHost = () => {
  userStore.changeHost()
}
//重置订阅
const onResetSub=()=>{
  userApi.resetSubApi().then((res)=>{
    if (res.code === 0) {
      ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
}
const v2rayNGSub = (type: number) => {
	switch (type) {
		case 1:
			copyText(userStore.subV2rayNG)
			break
		case 2:
			copyText(userStore.subClash)
			break
		case 3:
			copyText(userStore.subShadowRocket)
			break
		case 4:
			copyText(userStore.subQx)
			break
		default:
			copyText(userStore.subV2rayNG)
			break;
	}
}
// 页面加载时
onMounted(() => {
		//设置用户信息
  userStore.getUserInfo()
});



</script>

<style scoped lang="scss">
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.home-card-item {
	width: 100%;
	height: 100%;
	border-radius: 4px;
	transition: all ease 0.3s;
	padding: 20px;
	overflow: hidden;
	background: var(--el-color-white);
	color: var(--el-text-color-primary);
	border: 1px solid var(--next-border-color-light);
}
</style>