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