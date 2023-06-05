<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input  v-model="nodeManageData.params.search" size="default" placeholder="请输入名称" style="max-width: 180px"></el-input>
        <el-date-picker
            size="default"
            v-model="nodeManageData.params.date"
            type="datetimerange"
            :shortcuts="shortcuts"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
        <el-button  @click="nodeStore.getNodeWithTraffic()" size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="success" class="ml10" @click="onOpenEditNode('add')">
          <el-icon>
            <ele-FolderAdd />
          </el-icon>
          新增节点
        </el-button>
      </div>
      <el-table :data="nodeManageData.nodes.node_list" height="100%" style="width: 100%;flex: 1;">
        <el-table-column fixed type="index" label="序号" width="60"/>
        <el-table-column prop="name" label="节点名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="address" label="节点地址" show-overflow-tooltip></el-table-column>
        <el-table-column prop="port" label="节点端口" show-overflow-tooltip></el-table-column>
        <el-table-column prop="sort" label="类型" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.sort ===11">vmess</el-tag>
            <el-tag type="success" v-if="scope.row.sort ===15">vless</el-tag>
            <el-tag type="success" v-if="scope.row.sort ===14">trojan</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="totalUp" label="上行流量(GB)" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.totalUp / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="下行流量(GB)" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.totalDown / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="节点状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.status">启用</el-tag>
            <el-tag type="info" v-else>禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button :disabled="userInfos.id !== 1" size="small" text type="primary"
                       @click="onOpenEditNode('edit', scope.row)">修改
            </el-button>
            <el-button :disabled="userInfos.id !== 1" size="small" text type="primary"
                       @click="onRowDel(scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
          background
          :page-sizes="[10, 20, 30, 40]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange"
          v-model:current-page="nodeManageData.params.page_num"
          v-model:page-size="nodeManageData.params.page_size"
          :total="nodeManageData.nodes.total"
      />
    </el-card>
    <nodeDialog ref="nodeDialogRef" @refresh="nodeStore.getNodeWithTraffic()"/>
  </div>
</template>

<script setup lang="ts" name="NodeManage">

import {defineAsyncComponent, onMounted, ref} from "vue";

import {storeToRefs} from "pinia";
//导入弹出层
const nodeDialog = defineAsyncComponent(() => import('/@/views/admin/node/dialog.vue'))
const nodeDialogRef = ref()
//node store
import {useNodeStore} from "/@/stores/node";
const nodeStore = useNodeStore()
const {nodeManageData} = storeToRefs(nodeStore)
//user store
import {useUserStore} from "/@/stores/userStore";
const userStore = useUserStore()
const {userInfos} = storeToRefs(userStore)
//时间范围
const shortcuts = [
  {
    text: '上周',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      return [start, end]
    },
  },
  {
    text: '上月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '最近3个月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
]

//新建节点
function onOpenAddNode(type: string) {
  nodeDialogRef.value.openDialog(type)
}

//修改节点
function onOpenEditNode(type: string, row: Object) {
  nodeDialogRef.value.openDialog(type, row)
}

//删除节点
function onRowDel(row: Object) {
  nodeStore.deleteNode(row)
  setTimeout(() => {
    nodeStore.getNodeWithTraffic()
  }, 2000);


}

// 分页改变
const onHandleSizeChange = (val: number) => {
  nodeManageData.value.params.page_size = val;
  nodeStore.getNodeWithTraffic()
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  nodeManageData.value.params.page_num = val;
  nodeStore.getNodeWithTraffic()
};

//查询节点
function onSearch() {
  nodeStore.getNodeWithTraffic()
}

onMounted(() => {
  nodeStore.getNodeWithTraffic()
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