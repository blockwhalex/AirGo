<template>
  <div class="report_layout-padding">
    <el-card shadow="hover">

      <el-row>
        <el-col :span="8">
          <el-button type="primary">当前数据库</el-button>
          <el-select v-model="state.checkedDbInfo.database" class="m-2" placeholder="选择数据库">
            <el-option
                v-for="(v,k) in dbInfo.database_list"
                :key="k"
                :label="v"
                :value="v"
            />
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-button @click="onGetTables(state.checkedDbInfo)" type="primary">加载库表</el-button>
          <el-select @change="onGetColumn(state.checkedDbInfo)" v-model="state.checkedDbInfo.table_name" class="m-2"
                     placeholder="选择库表">
            <el-option
                v-for="(v,k) in sqliteTable"
                :key="v.name"
                :label="v.name"
                :value="v.name"
            />
          </el-select>
        </el-col>
      </el-row>

    </el-card>
  </div>
  <div class="container report_layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div>
        <el-button type="primary" @click="addCondition">新增条件</el-button>
      </div>
      <el-table :data="state.reportTable.field_params_list" height="100%" style="width: 100%;flex: 1;">
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
                  :value="v.data_type"
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
            <el-input v-model="row.condition_value"></el-input>
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
    </el-card>
  </div>
  <div class="container report_layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <el-table :data="state.tableData.table_data" style="width: 100%">
        <el-table-column
            :prop="index"
            :label="item"
            v-for="(item, index) in state.tableData.table_header"
            :key="index"
        >
        </el-table-column>
      </el-table>
    </el-card>
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
  //选中的数据库，库表;获取数据库的数据表的所有字段名,类型值 请求参数
  checkedDbInfo: {
    database: '',
    table_name: '',
  },
  reportTable: {
    table_name: '',
    field_params_list: [] as any[]    //搜索条件列表
  },
  //展示的表格数据
  tableData: {
    table_header: {},
    table_data: [],
  }
})


//搜索条件
const conditionList = ['<', '>', "=", "<>", "like"]
//加载库表
const onGetTables = (params?: object) => {
  reportStore.getTables(params)
}
//加载字段
const onGetColumn = (params?: object) => {
  reportStore.getColumn(params)
}
//提交
const onSubmit = (params?: object) => {
  if (state.checkedDbInfo.table_name === '') {
    return
  }
  state.reportTable.table_name = state.checkedDbInfo.table_name
  reportApi.submitReportApi(state.reportTable).then((res) => {
    if (res.code === 0) {
      state.tableData = res.data
    }

  })
}

//删除当前条件
const deleteCurrrentCondition = (row: any) => {
  // console.log("当前所有条件:",state.reportTable.field_params_list)
  state.reportTable.field_params_list = state.reportTable.field_params_list.filter(item => item !== row);
}
//增加新条件
const addCondition = () => {
  if (dbInfo.value.db_type === '' || (mysqlColumn.value.length === 0 && sqliteColumn.value.length === 0)) {
    ElMessage.warning("请选择数据库，并加载数据库表！")
    return
  }
  state.reportTable.field_params_list.push({
    field: '',
    field_chinese_name: '',
    field_type: '',
    condition: '=',
    condition_value: '',
  })
};

//监听
watch(
    () => state.reportTable,
    () => {
      state.reportTable.field_params_list.forEach((item) => {
        if (dbInfo.value.db_type === 'mysql') {
          item.field_type = mysqlColumnTypeMap.value.get(item.field)
          item.field_chinese_name = mysqlColumnChineseNameMap.value.get(item.field)
        } else {
          item.field_type = sqliteColumnTypeMap.value.get(item.field)
          item.field_chinese_name = sqliteColumnTypeMap.value.get(item.field)
        }
      })
    },
    {
      deep: true,
    }
);
onMounted(() => {
  reportStore.getDB()
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