<template>
  <div class="container report_layout-padding">
    <div>
      <el-button type="primary" @click="addCondition">新增条件</el-button>
    </div>
    <el-table :data="state.reportTable.field_params_list" height="100%" style="width: 100%;flex: 1;" stripe>
      <el-table-column align="left" type="index" label="序列" width="60"/>
      <el-table-column align="left" prop="field" label="字段" width="160">
        <template #default="{row}">
          <el-select v-if="dbInfo.db_type==='sqlite'" v-model="row.field" class="m-2" placeholder="选择字段">
            <el-option
                v-for="(v,k) in sqliteColumn"
                :key="k"
                :label="v.name"
                :value="v.name"
            />
          </el-select>
          <el-select v-if="dbInfo.db_type==='mysql'" v-model="row.field" class="m-2" placeholder="选择字段">
            <el-option
                v-for="(v,k) in mysqlColumn"
                :key="k"
                :label="v.column_name"
                :value="v.column_name"
            />
          </el-select>
        </template>

      </el-table-column>
      <el-table-column align="left" prop="field_chinese_name " label="字段中文名" width="160">
        <template #default="{row}">
          <el-text>{{ row.field_chinese_name }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" prop="field_type" label="字段类型" width="160">
        <template #default="{row}">
          <el-text>{{ row.field_type }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" prop="condition" label="搜索条件" width="160">
        <template #default="{row}">
          <el-select v-model="row.condition" class="m-2" placeholder="搜索条件">
            <el-option
                v-for="(v,k) in conditionList"
                :key="k"
                :label="v"
                :value="v"
            />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column align="left" prop="conditionValue" label="条件值">
        <template #default="{row}">
          <el-date-picker
              v-if="row.field_type==='datetime'"
              v-model="row.condition_value"
              type="datetime"
              value-format="YYYY-MM-DD HH:mm:ss"
          />
          <el-input v-else v-model="row.condition_value"></el-input>
        </template>
      </el-table-column>
      <el-table-column align="left" label="操作">
        <template #default="{row}">
          <el-button type="primary" @click="deleteCurrrentCondition(row)">删除</el-button>
        </template>
      </el-table-column>

    </el-table>
    <div style="margin-top: 20px">
      <el-button @click="onSubmit()" type="primary">确定</el-button>
    </div>

  </div>
</template>

<script lang="ts" setup>


import {onMounted, reactive, watch} from "vue";
//report store
import {useReportStore} from "/@/stores/reportStore"
import {storeToRefs} from "pinia";

const reportStore = useReportStore()
const {
  dbInfo,
  sqliteTable,
  mysqlColumn,
  mysqlColumnTypeMap,
  mysqlColumnChineseNameMap,
  sqliteColumn,
  sqliteColumnTypeMap
} = storeToRefs(reportStore)
//report api
import {useReportApi} from "/@/api/report";
import {ElMessage} from "element-plus";

const reportApi = useReportApi()
//定义的参数
const state = reactive({
  //选中的数据库，库表，用来请求获取数据库的数据表的所有字段名,类型值 请求参数
  checkedDbInfo: {
    database: '',
    table_name: '',
  },
  reportTable: {
    table_name: '',
    field_params_list: [] as any[]    //搜索条件列表
  },
})
//搜索条件
const conditionList = ['<', '>', "=", "<>", "like"]
//子传父
const emits = defineEmits(['getReportData'])

//删除当前条件
const deleteCurrrentCondition = (row: any) => {
  // console.log("当前所有条件:",state.reportTable.field_params_list)
  state.reportTable.field_params_list = state.reportTable.field_params_list.filter(item => item !== row);
}
//增加新条件
const addCondition = () => {
  state.reportTable.field_params_list.push({
    field: '',
    field_chinese_name: '',
    field_type: '',
    condition: '=',
    condition_value: '',
  })
};

//提交
const onSubmit = (params?: object) => {
  if (state.checkedDbInfo.table_name === '') {
    return
  }
  //将查询参数返回父组件
  emits('getReportData', state.reportTable)
}

//监听
watch(
    () => state.reportTable,
    () => {
      console.log("state.reportTable.field_params_list:",state.reportTable.field_params_list)
      state.reportTable.field_params_list.forEach((value,index,array) => {
        if (dbInfo.value.db_type === 'mysql') {
          value.field_type = mysqlColumnTypeMap.value.get(value.field)
          value.field_chinese_name = mysqlColumnChineseNameMap.value.get(value.field)
        } else {
          value.field_type = sqliteColumnTypeMap.value.get(value.field)
          value.field_chinese_name = sqliteColumnTypeMap.value.get(value.field)
        }
      })
    },
    {
      deep: true,
    }
);
// onMounted(() => {
//
// });

//打开时加载数据
const openReportComponent = (params: string) => {
  state.checkedDbInfo.table_name = params
  state.reportTable.table_name = params
  //加载数据库类型
  reportStore.getDB()
  //加载字段
  reportStore.getColumn(state.checkedDbInfo)
}

//暴露变量
defineExpose({
  openReportComponent,
});


</script>

<style scoped lang="scss">
.container {
  :deep(.el-card__body) {
    display: flex;
    flex-direction: column;
    flex: 1;
    overflow: auto;
  }

}

.home-card-item {
  font-size: 16px;
  width: 100%;
  height: 100%;
  border-radius: 4px;
  transition: all ease 0.3s;
  padding: 10px;
  overflow: hidden;
  background: var(--el-color-white);
  color: var(--el-text-color-primary);
  border: 1px solid var(--next-border-color-light);
}

.report_layout-padding {
  padding: 15px;
}

//.el-card {
//  background-image: url("../../assets/bgc/3.png");
//  background-repeat: no-repeat;
//}
</style>