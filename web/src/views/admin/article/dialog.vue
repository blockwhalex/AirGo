<template>
  <el-dialog v-model="state.isShowDialog" :title="state.title" width="100%" destroy-on-close align-center>
    <el-form label-position="top">
      <el-form-item label="文章主题">
        <el-input v-model="state.article.title"/>
      </el-form-item>
      <el-form-item label="文章简介">
        <el-input v-model="state.article.introduction"/>
      </el-form-item>
      <el-form-item label="文章内容---使用markdown语法">
        <!--        <el-input v-model="state.article.content" type="textarea" :rows="12" />-->
        <v-md-editor v-model="state.article.content" height="400px"></v-md-editor>
      </el-form-item>
    </el-form>
    <template #footer>
            <span class="dialog-footer">
                <el-button @click="closeDialog">取消</el-button>
                <el-button type="primary" @click="onSubmit">
                    确认
                </el-button>
            </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
//markdown

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

import {reactive} from "vue";
//api
import {useArticleApi} from "/@/api/article/index"

const articleApi = useArticleApi()
//定义参数
const state = reactive({
  type: "",
  title: "",
  isShowDialog: false,
  article: {} as Article,
})

// 打开弹窗
const openDialog = (type: string, row?: any) => {
  console.log("打开弹窗:", type)
  if (type == 'add') {
    state.type = type
    state.title = "新建文章"
    state.isShowDialog = true
  } else {
    state.article = row
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
    articleApi.newArticleApi(state.article)
    setTimeout(() => {
      emit('refresh');
    }, 1000);       //延时。防止没新建完成就重新请求
  } else {
    articleApi.updaterticleApi(state.article)
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

<style scoped>

</style>