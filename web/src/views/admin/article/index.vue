<template>
  <div class="container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="mb15">
        <el-input v-model="state.params.search" size="default" placeholder="请输入名称"
                  style="max-width: 180px"></el-input>
        <el-button size="default" type="primary" class="ml10" @click="getArticleList(state.params)">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="success" class="ml10" @click="onOpenDialog('add')">
          <el-icon>
            <ele-FolderAdd/>
          </el-icon>
          新建文章
        </el-button>
      </div>
      <el-tag type="info">*右上角通知只会显示最近3条数据</el-tag>
      <el-table :data="state.articleDate.article_list" stripe height="100%" style="width: 100%;flex: 1;">
        <el-table-column fixed type="index" label="序号" width="60"/>
        <el-table-column prop="id" label="ID" show-overflow-tooltip width="60"></el-table-column>
        <el-table-column prop="title" label="标题" show-overflow-tooltip width="200"></el-table-column>
        <el-table-column prop="introduction" label="简介" show-overflow-tooltip></el-table-column>
        <!--        <el-table-column prop="content" label="内容" show-overflow-tooltip></el-table-column>-->
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button size="small" text type="primary"
                       @click="onOpenDialog('edit', scope.row)">修改
            </el-button>
            <el-button size="small" text type="primary"
                       @click="deleteArticle(scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
          background
          :page-sizes="[10, 20, 30, 40]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange"
          v-model:current-page="state.params.page_num"
          v-model:page-size="state.params.page_size"
          :total="state.articleDate.total"
      />
    </el-card>
    <ArticleDialog ref="articleDialogRef" @refresh="getArticleList(state.params)"></ArticleDialog>
  </div>
</template>

<script lang="ts" setup>

//api
import {useArticleApi} from "/@/api/article/index"

const articleApi = useArticleApi()
import {ElMessage, ElMessageBox} from "element-plus";
import {defineAsyncComponent, onMounted, reactive, ref} from "vue";
//导入组件
const ArticleDialog = defineAsyncComponent(() => import('/@/views/admin/article/dialog.vue'))
const articleDialogRef = ref()

//

//定义变量
const state = reactive({
  articleDate: {
    total: 0,
    article_list: [] as Article[],
  },

  params: {
    search: '',
    page_num: 1,
    page_size: 10,
  },
})

function deleteArticle(row: any) {
  ElMessageBox.confirm(`此操作将永久删除文章：${row.title}, 是否继续?`, '提示', {
    confirmButtonText: '删除',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        //逻辑
        articleApi.deleteArticleApi(row)
        setTimeout(() => {
          //逻辑
          getArticleList(state.params)
          ElMessage.success('成功');
        }, 1000);
      })
      .catch(() => {
      });
}

//打开弹窗
const onOpenDialog = (type: string, row?: any) => {
  // console.log("打开弹窗:",type)
  articleDialogRef.value.openDialog(type, row)
};
//获取article列表
const getArticleList = (params: object) => {
  articleApi.getArticleApi(params).then((res) => {
    if (res.code === 0) {
      state.articleDate = res.data
    }
  })
}

// 分页改变
const onHandleSizeChange = (val: number) => {
  state.params.page_size = val;
  getArticleList(state.params)
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  state.params.page_num = val;
  getArticleList(state.params)
};

//加载时
onMounted(() => {
  getArticleList(state.params)
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