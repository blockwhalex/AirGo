<template>
  <div>
    <div>标题</div>
    <div>简介</div>
    <div>内容</div>
  </div>

</template>

<script lang="ts" setup>
import {useRoute} from 'vue-router'

const route = useRoute()

let Id = route.query.id;
console.log("传过来的id：", Id)
</script>

<style scoped>

</style>