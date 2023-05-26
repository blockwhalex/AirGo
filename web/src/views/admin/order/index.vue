<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input v-model="orderManageData.queryParams.search" size="default" placeholder="请输入订单号" style="max-width: 180px"></el-input>
        <el-date-picker
            size="default"
            v-model="orderManageData.queryParams.date"
            type="datetimerange"
            :shortcuts="shortcuts"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
        <el-button size="default" type="primary" class="ml10" @click="onSearch">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
      </div>
      <el-table :data="orderManageData.allOrders.order_list" fit style="width: 100%;flex: 1;">
        <el-table-column type="index" label="序号"/>
        <el-table-column prop="id" label="订单ID"/>
        <el-table-column prop="out_trade_no" label="订单号"/>
        <el-table-column prop="created_at" label="下单日期"/>
        <el-table-column prop="user_name" label="用户"/>
        <el-table-column prop="goods_id" label="商品ID" show-overflow-tooltip/>
        <el-table-column prop="subject" label="商品标题" show-overflow-tooltip/>
        <el-table-column prop="total_amount" label="订单金额" show-overflow-tooltip/>
        <el-table-column prop="receipt_amount" label="实收金额" show-overflow-tooltip/>
        <el-table-column prop="trade_status" label="交易状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.trade_status==='TRADE_SUCCESS'">支付成功</el-tag>
            <el-tag type="warning" v-else>未支付</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button v-if="scope.row.trade_status != 'TRADE_SUCCESS'" size="small" text type="primary"
                       @click="onCompleteOrder(scope.row)">完成
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background
                     class="mt15"
                     layout="total, sizes, prev, pager, next, jumper"
                     :page-sizes="[10, 20, 30]"
                     v-model:current-page="orderManageData.queryParams.page_num"
                     v-model:page-size="orderManageData.queryParams.page_size"
                      :total="orderManageData.allOrders.total"
                     @size-change="onHandleSizeChange"
                     @current-change="onHandleCurrentChange">
      </el-pagination>
    </el-card>
  </div>
</template>

<script setup lang="ts">
//store
import {onMounted} from "vue";
import {useOrderStore} from "/@/stores/orderStore";
import {storeToRefs} from "pinia";

const orderStore = useOrderStore()
const {orderManageData} = storeToRefs(orderStore)

const onSearch = () => {
  orderStore.getAllOrder()

}
onMounted(() => {
  orderStore.getAllOrder()
})
//完成未支付订单
const onCompleteOrder = (row: Order) => {

}
// 分页改变
const onHandleSizeChange = (val: number) => {
  orderManageData.value.queryParams.page_size = val;
  orderStore.getAllOrder()
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  orderManageData.value.queryParams.page_num = val;
  orderStore.getAllOrder()
};
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