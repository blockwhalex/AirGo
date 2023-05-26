<template>
	<div>
		<el-row :gutter="10" class="home-card-two mb15">
			<el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
				<div class="home-card-item">
					<el-card class="box-card">
						<template #header>
							<div class="card-header">
								<div class="block">
								</div>
								<el-tag class="ml-2" type="success">套餐详情</el-tag>
							</div>
						</template>
						<div class="text item">
							<el-tag class="ml-2" type="info">剩余流量:</el-tag><span>{{ userInfos.used }}GB</span>
						</div>
						<div class="text item">
							<el-tag class="ml-2" type="info">到期时间:</el-tag>{{ userInfos.expired }}
						</div>
					</el-card>
				</div>
			</el-col>
			<el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
				<div class="home-card-item">
					<el-card class="box-card" style="width: 100%;flex: 1;">
						<template #header>
							<div class="card-header">
								<el-tag class="ml-2" type="success">当前混淆:</el-tag>
								<el-tag class="ml-2" type="info">{{ userInfos.userInfos.value.subscribe_info.host }}</el-tag>
							</div>
						</template>
						<div class="mt-4">
							<el-input v-model="userInfos.homeTableData.value.host" placeholder="输入混淆"
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

//store
import { storeToRefs } from "pinia";
import { useUserStore } from "/@/stores/userStore";
import {  onMounted } from 'vue';
const userInfo = useUserStore()
const userInfos = storeToRefs(userInfo)

//icon 
import { Select } from '@element-plus/icons-vue'
//复制剪切板
import commonFunction from '/@/utils/commonFunction';
const { copyText } = commonFunction();


//修改混淆
const onChangeHost = () => {
	userInfo.changeHost()
}
const v2rayNGSub = (type: number) => {
	switch (type) {
		case 1:
			copyText(userInfos.subV2rayNG.value)
			break
		case 2:
			copyText(userInfos.subClash.value)
			break
		case 3:
			copyText(userInfos.subShadowRocket.value)
			break
		case 4:
			copyText(userInfos.subQx.value)
			break
		default:
			copyText(userInfos.subV2rayNG.value)
			break;
	}
}
// 页面加载时
onMounted(() => {
		//设置用户信息
		userInfo.getUserInfo()
})



</script>

<style scoped lang="scss">
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