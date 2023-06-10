<template>
  <div class="layout-navbars-breadcrumb-user-news">
    <div class="head-box">
      <div class="head-box-title">通知</div>
      <div class="head-box-btn" v-if="state.newsList.length > 0" @click="onAllReadClick">全部已读</div>
    </div>
    <div class="content-box">
      <template v-if="state.newsList.length > 0">
        <div class="content-box-item" v-for="(v, k) in state.newsList" :key="k">
          <div>{{ v.label }}</div>
          <div class="content-box-msg">
            {{ v.value }}
          </div>
          <div class="content-box-time">{{ v.time }}</div>
        </div>
      </template>
      <el-empty description="暂无通知" v-else></el-empty>
    </div>
    <div class="foot-box" @click="onGoToGiteeClick" v-if="state.newsList.length > 0">加入我们</div>
  </div>
</template>

<script setup lang="ts" name="layoutBreadcrumbUserNews">
import {reactive} from 'vue';

// 定义变量内容
const state = reactive({
  newsList: [
    {
      label: '欢迎使用',
      value: '禁止一切违法行为，所有违法行为与作者和群友无关。',
      time: '2023-06-06',
    },
  ],
});

// 全部已读点击
const onAllReadClick = () => {
  state.newsList = [];
};
// 前往通知中心点击
const onGoToGiteeClick = () => {
  window.open('http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=qVqTTvfkP2FACYzeKKQ1yS8m9DugxnCA&authKey=CGvwGFUdQvr9fWxmWbCqAyZDpLR0oIyOhAPzAgxqbsBKUigTkgbrpOCyTpULmQrT&noverify=0&group_code=144545032');
};
</script>

<style scoped lang="scss">
.layout-navbars-breadcrumb-user-news {
  .head-box {
    display: flex;
    border-bottom: 1px solid var(--el-border-color-lighter);
    box-sizing: border-box;
    color: var(--el-text-color-primary);
    justify-content: space-between;
    height: 35px;
    align-items: center;

    .head-box-btn {
      color: var(--el-color-primary);
      font-size: 13px;
      cursor: pointer;
      opacity: 0.8;

      &:hover {
        opacity: 1;
      }
    }
  }

  .content-box {
    font-size: 13px;

    .content-box-item {
      padding-top: 12px;

      &:last-of-type {
        padding-bottom: 12px;
      }

      .content-box-msg {
        color: var(--el-text-color-secondary);
        margin-top: 5px;
        margin-bottom: 5px;
      }

      .content-box-time {
        color: var(--el-text-color-secondary);
      }
    }
  }

  .foot-box {
    height: 35px;
    color: var(--el-color-primary);
    font-size: 13px;
    cursor: pointer;
    opacity: 0.8;
    display: flex;
    align-items: center;
    justify-content: center;
    border-top: 1px solid var(--el-border-color-lighter);

    &:hover {
      opacity: 1;
    }
  }

  :deep(.el-empty__description p) {
    font-size: 13px;
  }
}
</style>
