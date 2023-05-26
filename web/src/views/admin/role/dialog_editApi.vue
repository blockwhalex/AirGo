<template>
  <el-dialog v-model="dialogEditApi.isShowDialog" title="修改角色api权限" style="width: 80%;flex: 1;">
    <el-transfer v-model="dialogEditApi.casbinInfo.casbinItems"
    :props="{
      key: 'path',
      label: 'path',
    }" 
    :data="dialogEditApi.allCasbinInfo.casbinItems" />

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogEditApi.isShowDialog = false">取消</el-button>
        <el-button type="primary" @click="onSubmit">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
  
<script lang="ts" setup>
import { storeToRefs } from 'pinia';
// role store
import { useRoleStore } from "/@/stores/roleStore";
import { reactive, ref } from 'vue';
const roleStore = useRoleStore()
const { dialogEditApi } = storeToRefs(roleStore)


// 打开弹窗
const openDialog = (row:RowRoleType) => { //RowRoleType 角色类型
  //获取当前roleID
  //console.log("获取当前roleID",row.roleID)
  dialogEditApi.value.casbinInfo.roleID=row.roleID
  //获取全部api list
  roleStore.getAllPolicy()
  dialogEditApi.value.isShowDialog = true;

};
// 关闭弹窗
const closeDialog = () => {
  dialogEditApi.value.isShowDialog = false;
};
//提交
const onSubmit = () => {
  //console.log("选中的api：",casbinInfo.value.casbinItems)
    roleStore.updateCasbinPolicy()
    closeDialog()
}

// 暴露变量
defineExpose({
  openDialog,
});


</script>
<style> 
/* 定义两边的el-transfer-panel大小的方法,直接设置是没有用的,需要去掉scoped即可。才能成功覆盖原生的样式 */
.el-transfer-panel {
  width: 300px;
  height: 300px;
}

.el-transfer-panel__list.is-filterable {
  height: 400px;
}

</style>
  