<template>
	<div class="system-role-container layout-padding">
		<div class="system-role-padding layout-padding-auto layout-padding-view">
			<div class="system-user-search mb15">
				<el-input v-model="roleStoreRef.tableData.value.param.search" size="default" placeholder="请输入角色名称"
					style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10" @click="onSearch">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAddRole('add')">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增角色
				</el-button>
			</div>
			<el-table :data="roleStoreRef.tableData.value.data" v-loading="roleStoreRef.tableData.value.loading"
				style="width: 100%">
				<el-table-column type="index" label="序号" width="60" />
				<el-table-column prop="role_name" label="角色名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="id" label="角色ID" show-overflow-tooltip></el-table-column>
				<el-table-column prop="status" label="角色状态" show-overflow-tooltip>
					<template #default="scope">
						<el-tag type="success" v-if="scope.row.status">启用</el-tag>
						<el-tag type="info" v-else>禁用</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="操作" width="300">
					<template #default="scope">
						<el-button :disabled="scope.row.roleName === '超级管理员'" size="small" text type="primary"
							@click="onOpenEditRole('edit', scope.row)">修改路由权限</el-button>
						<el-button :disabled="scope.row.roleName === '超级管理员'" size="small" text type="primary"
							@click="onOpenEditApi(scope.row)">修改api权限</el-button>
						<el-button :disabled="scope.row.roleName === '超级管理员'" size="small" text type="primary"
							@click="onRowDel(scope.row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<el-pagination
        background
				:pager-count="5" :page-sizes="[10, 20, 30]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange" class="mt15"
				v-model:current-page="roleStoreRef.tableData.value.param.page_num"
				v-model:page-size="roleStoreRef.tableData.value.param.page_size"
				 :total="roleStoreRef.tableData.value.total">
			</el-pagination>
		</div>
		<RoleDialog ref="roleDialogRef" @refresh="getTableData()" />
		<RoleDialogEditApi ref="roleDialogEditApiRef" />
	</div>
</template>

<script setup lang="ts" name="systemRole">
import { defineAsyncComponent, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
//导入 role store
import { storeToRefs } from 'pinia';
import { useRoleStore } from "/@/stores/roleStore";

//role api
import { useRoleApi } from "/@/api/role/index";
const roleApi = useRoleApi()

// 引入组件
const RoleDialog = defineAsyncComponent(() => import('/@/views/admin/role/dialog.vue'));
const RoleDialogEditApi = defineAsyncComponent(()=>import('/@/views/admin/role/dialog_editApi.vue'))
const roleDialogRef = ref();
const roleDialogEditApiRef = ref();

// 定义变量内容
const roleStore = useRoleStore()
const roleStoreRef = storeToRefs(roleStore)



const getTableData = function () {
	//请求一次role list
	roleStore.setRoleList()
}

// 打开新增角色弹窗
const onOpenAddRole = (type: string) => {
	roleDialogRef.value.openDialog(type);
};
// 打开修改角色弹窗
const onOpenEditRole = (type: string, row: Object) => {
	roleDialogRef.value.openDialog(type, row);
};
// 打开修改api弹窗
const onOpenEditApi=(row: Object)=>{
	roleDialogEditApiRef.value.openDialog(row);
}
// 删除角色
const onRowDel = (row: RowRoleType) => {
	ElMessageBox.confirm(`此操作将永久删除角色名称：“${row.role_name}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	}).then(() => {
		//请求
		//console.log("删除角色:", row.roleName)
		roleApi.delRoleApi({ id: row.id }).then((res) => {
			if (res.code != 0) {
				ElMessage.success('删除失败');
			} else {
				ElMessage.success('删除成功');
				getTableData();
			}
		})

	})
		.catch(() => { });
};
//查询
const onSearch = () => {
	roleStore.setRoleList()
}
// 分页改变
const onHandleSizeChange = (val: number) => {
	roleStoreRef.tableData.value.param.page_size = val;
	//getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
	roleStoreRef.tableData.value.param.page_num = val;
	//getTableData();
};
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>

<style scoped lang="scss">
.system-role-container {
	.system-role-padding {
		padding: 15px;

		.el-table {
			flex: 1;
		}
	}
}
</style>
