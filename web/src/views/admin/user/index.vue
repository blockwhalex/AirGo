<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input v-model="userManageData.params.search" size="default" placeholder="请输入用户名称"
                  style="max-width: 180px"></el-input>
        <el-button @click="onSearch" size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="success" class="ml10" @click="onOpenAddUser('add')">
          <el-icon>
            <ele-FolderAdd/>
          </el-icon>
          新增用户
        </el-button>
      </div>
      <el-table :data="userManageData.users.user_list" fit style="width: 100%;flex: 1;">
        <!--				<el-table-column type="index" label="序号" width="60" />-->
        <el-table-column prop="id" label="账户ID" show-overflow-tooltip fixed width="60"></el-table-column>
        <el-table-column prop="user_name" label="账户名称" show-overflow-tooltip fixed width="150"></el-table-column>
        <el-table-column prop="enable" label="用户状态" show-overflow-tooltip width="80">
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.enable">启用</el-tag>
            <el-tag type="danger" v-else>禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="enable" label="订阅状态" show-overflow-tooltip width="80">
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.enable">启用</el-tag>
            <el-tag type="danger" v-else>禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="subscribe_info.expired_at" label="订阅到期时间" show-overflow-tooltip width="150">
          <template #default="scope">
            <el-tag type="info">
              {{DateStrtoTime(scope.row.subscribe_info.expired_at)}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="subscribe_info.goods_id" label="商品ID" show-overflow-tooltip width="60"></el-table-column>
        <el-table-column prop="subscribe_info.t" label="总流量(GB)" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="info">{{ scope.row.subscribe_info.t / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="已用流量(GB)" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="info">
              {{ (scope.row.subscribe_info.d + scope.row.subscribe_info.u) / 1024 / 1024 / 1024 }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button size="small" text type="primary" @click="onOpenEditUser('edit', scope.row)"
            >修改
            </el-button
            >
            <el-button size="small" text type="primary" @click="onRowDel(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
          background
          class="mt15"
          layout="total, sizes, prev, pager, next, jumper"
          :page-sizes="[10, 20, 30]"
          v-model:current-page="userManageData.params.page_num"
          v-model:page-size="userManageData.params.page_size"
          :total="userManageData.users.total"
          @size-change="onHandleSizeChange"
          @current-change="onHandleCurrentChange"
      >
      </el-pagination>
    </el-card>
    <UserDialog ref="userDialogRef" @refresh=""/>
  </div>
</template>

<script setup lang="ts" name="systemUser">
import {defineAsyncComponent, onMounted, ref} from 'vue';
import {ElMessageBox, ElMessage} from 'element-plus';

// 引入组件
const UserDialog = defineAsyncComponent(() => import('/@/views/admin/user/dialog.vue'));
//store
import {storeToRefs} from 'pinia';
import {useUserStore} from '/@/stores/userStore'
import {DateStrtoTime} from "../../../utils/formatTime";

const userStore = useUserStore()
const {userManageData} = storeToRefs(userStore)


// 定义变量内容
const userDialogRef = ref();

// 打开新增用户弹窗
const onOpenAddUser = (type: string) => {
  userDialogRef.value.openDialog(type);
};
// 打开修改用户弹窗
const onOpenEditUser = (type: string, row: SysUser) => {
  userDialogRef.value.openDialog(type, row);
};
//查询用户
const onSearch = () => {
  userStore.getUserList()
}
// 删除用户
const onRowDel = (row: SysUser) => {
  ElMessageBox.confirm(`此操作将永久删除账户名称：“${row.user_name}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        userStore.deleteUser(row)
        setTimeout(() => {
          userStore.getUserList()
        }, 1000)

        ElMessage.success('删除成功');
      })
      .catch(() => {
      });
};
// 分页改变
const onHandleSizeChange = (val: number) => {
  userManageData.value.params.page_size = val;
  userStore.getUserList()
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  userManageData.value.params.page_num = val;
  userStore.getUserList()
};
// 页面加载时
onMounted(() => {
  userStore.getUserList() //获取用户列表

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


