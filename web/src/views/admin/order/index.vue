<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input v-model="state.params.search" placeholder="请输入订单号"
                  style="max-width: 180px"></el-input>
        <el-date-picker
            size="small"
            v-model="state.params.date"
            type="datetimerange"
            :shortcuts="shortcuts"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
        <el-button size="default" type="primary" class="ml10" @click="onSearch(state.params)">

          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
      </div>
      <el-table :data="orderManageData.allOrders.order_list" fit style="width: 100%;flex: 1;">
        <el-table-column type="index" label="序号" fixed/>
        <!--        <el-table-column prop="id" label="订单ID" fixed/>-->
        <el-table-column prop="out_trade_no" label="订单号" fixed width="200"/>
        <el-table-column prop="created_at" label="下单日期" width="150">
          <template #default="scope">
            <el-tag type="success">{{ DateStrtoTime(scope.row.created_at) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="user_name" label="用户" width="120"/>
        <el-table-column prop="goods_id" label="商品ID" show-overflow-tooltip width="60"/>
        <el-table-column prop="subject" label="商品标题" show-overflow-tooltip width="200"/>
        <el-table-column prop="total_amount" label="订单金额" show-overflow-tooltip width="80"/>
        <el-table-column prop="receipt_amount" label="实收金额" show-overflow-tooltip width="80"/>
        <el-table-column prop="trade_status" label="交易状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.trade_status==='TRADE_SUCCESS'">支付成功</el-tag>
            <el-tag type="warning" v-else-if="scope.row.trade_status==='WAIT_BUYER_PAY'">等待买家付款</el-tag>
            <el-tag type="danger" v-else-if="scope.row.trade_status==='TRADE_CLOSED'">交易超时关闭</el-tag>
            <el-tag type="success" v-else-if="scope.row.trade_status==='TRADE_FINISHED'">交易结束</el-tag>
            <el-tag type="info" v-else-if="scope.row.trade_status==='created'">订单已创建</el-tag>
            <el-tag type="success" v-else-if="scope.row.trade_status==='completed'">订单已完成</el-tag>
            <el-tag type="danger" v-else>未知状态</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button v-if="scope.row.trade_status === 'WAIT_BUYER_PAY' || scope.row.trade_status ==='created'"
                       size="small" text type="primary"
                       @click="onCompleteOrder(scope.row)">完成
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background
                     class="mt15"
                     layout="total, sizes, prev, pager, next, jumper"
                     :page-sizes="[10, 20, 30]"
                     v-model:current-page="state.params.page_num"
                     v-model:page-size="state.params.page_size"
                     :total="orderManageData.allOrders.total"
                     @size-change="onHandleSizeChange"
                     @current-change="onHandleCurrentChange">
      </el-pagination>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import {ElMessage} from 'element-plus';
//store
import {onMounted, reactive} from "vue";
import {useOrderStore} from "/@/stores/orderStore";
import {storeToRefs} from "pinia";
//格式化时间
import {DateStrtoTime} from "/@/utils/formatTime"
// console.log(DateStrtoTime("2023-05-29T17:28:47.50276+08:00"))

const orderStore = useOrderStore()
const {orderManageData} = storeToRefs(orderStore)

//定义参数
const state = reactive({
  params: {
    page_num: 1,
    page_size: 10,
    search: '',
    date: [],
  } as QueryParams,
})
//
const onSearch = (params?:object) => {
  orderStore.getAllOrder(params)
}
onMounted(() => {
  onSearch(state.params)
})
//完成未支付订单
const onCompleteOrder = (row: Order) => {
  orderStore.completedOrder(row)
}
// 分页改变
const onHandleSizeChange = (val: number) => {
  state.params.page_size = val;
  orderStore.getAllOrder(state.params)
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  state.params.page_num = val;
  orderStore.getAllOrder(state.params)
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