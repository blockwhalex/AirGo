<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input size="default" placeholder="请输入名称" style="max-width: 180px"></el-input>
        <el-button size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>


import {ElMessage, ElMessageBox} from "element-plus";

function test(row: any) {
  ElMessageBox.confirm(`此操作将永久删除路由：${row.path}, 是否继续?`, '提示', {
    confirmButtonText: '删除',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        //逻辑
        setTimeout(() => {
          //逻辑
          ElMessage.success('成功');
        }, 1000);
      })
      .catch(() => {
      });
}

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
//定义参数
import {reactive} from "vue";

const state = reactive({
  type: "",
  title: "",
  isShowDialog: false,
  article: {} as Article,
})

// 打开弹窗
const openDialog = (type: string, row?: any) => {
  if (type == 'add') {
    state.type = type
    state.title = "新建文章"
    state.isShowDialog = true
  } else {
    state.type = type
    state.title = "修改文章"
    state.isShowDialog = true
  }
}
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false
};

//确认提交
function onSubmit() {
  if (state.type === 'add') {
    setTimeout(() => {
      emit('refresh');
    }, 1000);       //延时。防止没新建完成就重新请求
  } else {
    setTimeout(() => {
      emit('refresh');
    }, 1000);
  }
  closeDialog()
}

// 暴露变量
defineExpose({
  openDialog,   // 打开弹窗
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