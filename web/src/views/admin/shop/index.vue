<template>
    <div class="container layout-padding">
        <el-card shadow="hover" class="layout-padding-auto">
            <div class="system-user-search mb15">
                <el-button size="default" type="success" class="ml10" @click="onOpenAddGoods">
                    <el-icon>
                        <ele-FolderAdd />
                    </el-icon>
                    新增套餐
                </el-button>
            </div>
            <el-table :data="goodsList" height="100%" style="width: 100%;flex: 1;">
                <el-table-column type="index" label="序号" width="60" />
                <el-table-column prop="subject" label="套餐名称" show-overflow-tooltip  width="300"></el-table-column>
                <el-table-column prop="total_amount" label="套餐价格(元)" show-overflow-tooltip  width="100"></el-table-column>
                <el-table-column prop="total_bandwidth" label="总流量(GB)" show-overflow-tooltip  width="100"></el-table-column>
                <el-table-column prop="expiration_date" label="有效期(天)" show-overflow-tooltip ></el-table-column>
              <el-table-column prop="nodes" label="关联节点" show-overflow-tooltip >
                <template #default="scope">
                  <el-tag class="ml-2" type="success" v-for="(v,k) in scope.row.nodes" :key="k">{{v.name}}</el-tag>
                </template>
              </el-table-column>
                <el-table-column label="操作" width="100">
                    <template #default="scope">
                        <el-button size="small" text type="primary"
                            @click="onOpenEditGoods(scope.row)">修改</el-button>
                        <el-button size="small" text type="primary"
                            @click="onRowDel(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination background v-model:current-page="goodsManageData.pageNum"
                v-model:page-size="goodsManageData.pageSize" :page-sizes="[10, 20, 30, 40]"
                layout="sizes, prev, pager, next" :total="goodsManageData.total" @size-change="onHandleSizeChange"
                @current-change="onHandleCurrentChange" />
            </el-card>
            <!-- 引入弹窗组件 -->
            <shopDialog ref="shopDialogRef" @refresh="shopStore.getAllGoods()"></shopDialog>
    </div>
</template>

<script setup lang="ts">

//导入store
import { defineAsyncComponent, onMounted,ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useShopStore } from "/@/stores/shopStore";
const shopStore = useShopStore()
const { goodsList, goodsManageData } = storeToRefs(shopStore)
//store
import { useNodeStore } from "/@/stores/node";
const nodeStore = useNodeStore()

//引入弹窗组件
const shopDialog=defineAsyncComponent(()=>import('/@/views/admin/shop/dialog.vue'))
const shopDialogRef=ref()


//修改套餐
const onOpenEditGoods = (row: Goods) => {
    //打开tanc
    shopDialogRef.value.openDialog('edit',row)
}
//删除套餐
const onRowDel = (row: Goods) => {
    goodsManageData.value.currentGoods=row
    shopStore.deleteGoods(row)
    //延迟2秒
    setTimeout(()=>{
        shopStore.getAllGoods()
    },2000)
}
//添加套餐
const onOpenAddGoods = () => {
    //打开tanc
    shopDialogRef.value.openDialog('add')
    //延迟2秒重新获取数据
    setTimeout(()=>{
        shopStore.getAllGoods()
    },2000)
}
// 分页改变
const onHandleSizeChange = (val: number) => {
    goodsManageData.value.pageSize = val;
    //getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
    goodsManageData.value.pageNum = val;
    //getTableData();
};
//加载时
onMounted(() => {
    shopStore.getAllGoods() //获取全部商品
    nodeStore.getAllNode()  //获取全部节点
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