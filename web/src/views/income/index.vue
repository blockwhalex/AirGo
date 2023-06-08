<template>
  <div>
    <div >
      <el-row :gutter="15" >
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" >
          <div class="home-card-item">
            <el-card class="box-card">
              <el-row :gutter="10" justify="space-around" align="middle" >
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                  今日订单：{{state.statisticsData.dayOrderDataCurrent.total}}
                </el-col>
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                  今日收入：{{state.statisticsData.dayOrderDataCurrent.total_amount}} ¥
                </el-col>
              </el-row>
            </el-card>
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" >
          <div class="home-card-item">
            <el-card class="box-card">
              <el-row :gutter="10" justify="space-around" align="middle" >
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                 当月订单：{{state.statisticsData.monthOrderDataCurrent.total}}
                </el-col>
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                  当月收入：{{state.statisticsData.monthOrderDataCurrent.total_amount}} ¥
                </el-col>
              </el-row>
            </el-card>
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" >
          <div class="home-card-item">
            <el-card class="box-card">
              <el-row :gutter="10" justify="space-around" align="middle" >
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                  上月订单：{{state.statisticsData.monthOrderDataLast.total}}
                </el-col>
                <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                  上月收入：{{state.statisticsData.monthOrderDataLast.total_amount}} ¥
                </el-col>
              </el-row>
            </el-card>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="home-card-item">
      <el-divider>
        当月数据
      </el-divider>
      <el-table :data="state.statisticsData.monthNodeCurrent.node_list" height="100%" style="width: 100%;flex: 1;" stripe fit show-summary  :summary-method="getSummaries">
        <el-table-column fixed type="index" label="序号" width="60"/>
        <el-table-column prop="name" label="节点名称" show-overflow-tooltip width="120"></el-table-column>
        <el-table-column prop="address" label="节点地址" show-overflow-tooltip width="120"></el-table-column>
        <el-table-column prop="total_up" label="上行流量(GB)" show-overflow-tooltip width="200">
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.total_up / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="total_down" label="下行流量(GB)" show-overflow-tooltip >
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.total_down / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="home-card-item">
      <el-divider>
        上月数据
      </el-divider>
      <el-table :data="state.statisticsData.monthNodeLast.node_list" height="100%" style="width: 100%;flex: 1;" stripe fit show-summary :summary-method="getSummaries">
        <el-table-column fixed type="index" label="序号" width="60"/>
        <el-table-column prop="name" label="节点名称" show-overflow-tooltip width="120"></el-table-column>
        <el-table-column prop="address" label="节点地址" show-overflow-tooltip width="120"></el-table-column>
        <el-table-column prop="total_up" label="上行流量(GB)" show-overflow-tooltip width="200">
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.total_up / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="total_down" label="下行流量(GB)" show-overflow-tooltip >
          <template #default="scope">
            <el-tag type="warning">{{ scope.row.total_down / 1024 / 1024 / 1024 }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script lang="ts" setup>

import {ref, reactive, onMounted, computed} from 'vue'
import {ElMessage,TableColumnCtx} from "element-plus";
import {NextLoading} from "/@/utils/loading";
import {formatDate} from "/@/utils/formatTime";

//node store
import { storeToRefs } from 'pinia';
import {useNodeStore} from "/@/stores/node";
const nodeStore = useNodeStore()
const {nodeManageData} = storeToRefs(nodeStore)
//api
import {useNodeApi} from "/@/api/node/index";
import {useOrderApi} from "/@/api/order/index"
const nodeApi = useNodeApi()
const orderApi = useOrderApi()

const state = reactive({
  statisticsData: {
    dayOrderDataCurrent:{
      total_amount:0,
      total:0,
    } as OrdersWithTotal,
    monthOrderDataCurrent:{
      total_amount:0,
      total:0,
    } as OrdersWithTotal,
    monthOrderDataLast:{
      total_amount:0,
      total:0,
    } as OrdersWithTotal,
    monthNodeCurrent:{
      total: 0,
      node_list: [] as NodeInfo[],
    },
    monthNodeLast:{
      total: 0,
      node_list: [] as NodeInfo[],
    },
  },
  monthOptions: {
    dateDay:{
      page_num:1,
      page_size:200,
      date: ['', ''],
    },
    dateCurrent:{
      page_num:1,
      page_size:200,
      date: ['', ''],
    },
    dateLast:{
      page_num:1,
      page_size:200,
      date: ['', ''],
    },
  },
})
function getMonthOrder(params?:object){
  orderApi.getMonthOrderStatisticsApi(state.monthOptions.dateDay).then((res)=>{
    if (res.code === 0) {
      state.statisticsData.dayOrderDataCurrent=res.data
      // ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
  orderApi.getMonthOrderStatisticsApi(state.monthOptions.dateCurrent).then((res)=>{
    if (res.code === 0) {
      state.statisticsData.monthOrderDataCurrent=res.data
     // ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
  orderApi.getMonthOrderStatisticsApi(state.monthOptions.dateLast).then((res)=>{
    if (res.code === 0) {
      state.statisticsData.monthOrderDataLast=res.data
     // ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
}
function getMonthNodeInfo(params?:object){
  nodeApi.getNodeWithTrafficApi(state.monthOptions.dateCurrent).then((res)=>{
    if (res.code === 0) {
      state.statisticsData.monthNodeCurrent=res.data
      //ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
  nodeApi.getNodeWithTrafficApi(state.monthOptions.dateLast).then((res)=>{
    if (res.code === 0) {
      state.statisticsData.monthNodeLast=res.data
     // ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
}

function initDate() {
  // 时间范围格式 "2023-05-09 11:56:02"
  var currentDate = new Date();
  var currentY = currentDate.getFullYear();
  var currentM = currentDate.getMonth() + 1;
 // var MonthDayNum = new Date(currentY,currentM,0).getDate();  //计算当月的天数

  //当月
  var startDate = new Date(currentY, currentM - 1, 1);
  var endDate = new Date(currentY, currentM, 0, 23, 59, 59); // new Date(2020,11,0);//表示2020/11/30这天
  var start = formatDate(startDate, "YYYY-mm-dd HH:MM:SS")
  var end = formatDate(endDate, "YYYY-mm-dd HH:MM:SS")
  //上月
  var lastStartDate = new Date(currentY, currentM - 2, 1);
  var lastEndDate = new Date(currentY, currentM - 1, 0, 23, 59, 59); // new Date(2020,11,0);//表示2020/11/30这天
  var lastStart = formatDate(lastStartDate, "YYYY-mm-dd HH:MM:SS")
  var lastEnd = formatDate(lastEndDate, "YYYY-mm-dd HH:MM:SS")
  //当天范围
  var currentDayStartDate =new Date();
  currentDayStartDate.setHours(0);
  currentDayStartDate.setMinutes(0);
  currentDayStartDate.setSeconds(0);
  currentDayStartDate.setMilliseconds(0);
  var currentDayEndDate = new Date(currentDayStartDate.getTime() + 3600*1000*24 -1000)
  var currentDayStart=formatDate(currentDayStartDate, "YYYY-mm-dd HH:MM:SS")
  var currentDayEnd=formatDate(currentDayEndDate, "YYYY-mm-dd HH:MM:SS")
  state.monthOptions.dateCurrent.date = [start, end]
  state.monthOptions.dateLast.date = [lastStart, lastEnd]
  state.monthOptions.dateDay.date=[currentDayStart,currentDayEnd]
}


// 页面加载前
onMounted(() => {
  initDate()
  getMonthOrder()
  getMonthNodeInfo()

});
interface SummaryMethodProps<T = Product> {
  columns: TableColumnCtx<T>[]
  data: T[]
}

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 0) {
      sums[index] = '合计'
      return
    }
    const values = data.map((item) => Number(item[column.property]))
    if (!values.every((value) => Number.isNaN(value))) {
      sums[index] = `${values.reduce((prev, curr) => {
        const value = Number(curr)
        if (!Number.isNaN(value)) {
          return prev + curr/(1024*1024*1024)
        } else {
          return prev/(1024*1024*1024)
        }
      }, 0)} GB`
    } else {
      sums[index] = 'N/A'
    }
  })
  return sums
}

</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.home-card-item {
  width: 100%;
  /*height: 100%;*/
  border-radius: 4px;
  transition: all ease 0.3s;
  padding: 20px;
  overflow: hidden;
  background: var(--el-color-white);
  color: var(--el-text-color-primary);
  border: 1px solid var(--next-border-color-light);
}
</style>