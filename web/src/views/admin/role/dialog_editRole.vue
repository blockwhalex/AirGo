<template>
  <div class="system-role-dialog-container">
    <el-dialog :title="state.title" v-model="state.isShowDialog" width="769px" destroy-on-close>
      <el-form ref="roleDialogFormRef" :model="dialog.ruleForm" size="default" label-width="90px">
        <el-row :gutter="35">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="角色名称">
              <el-input v-model="dialog.ruleForm.role_name" placeholder="请输入角色名称" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="角色状态">
              <el-switch v-model="dialog.ruleForm.status" inline-prompt active-text="启"
                         inactive-text="禁"></el-switch>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="角色描述">
              <el-input v-model="dialog.ruleForm.description" type="textarea" placeholder="请输入角色描述"
                        maxlength="150"></el-input>
            </el-form-item>
          </el-col>
          <!-- 权限开始 -->
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
            <el-form-item label="菜单权限">
              <el-tree ref="tree_ref" node-key="route_id" :data="routesTree"
                       :props="{children:'children',label:'title'}"
                       :default-checked-keys="dialog.ruleForm.nodes" show-checkbox class="menu-data-tree"/>
            </el-form-item>
          </el-col>
          <!-- 权限结束 -->
        </el-row>
      </el-form>
      <template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.submitTxt }}</el-button>
				</span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {reactive, ref} from 'vue';
import {useRoutesStore} from "/@/stores/routesStore";
import {useRoleStore} from "/@/stores/roleStore";
import {storeToRefs} from 'pinia';
import {ElMessage} from 'element-plus';
//
import {arrayExtractionNodes} from "/@/utils/arrayOperation";
//role api
import {useRoleApi} from "/@/api/role/index";

const roleApi = useRoleApi()
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
//定义store
const routesStore = useRoutesStore()
const {routesTree} = storeToRefs(routesStore)
const roleStore = useRoleStore()
const {dialog} = storeToRefs(roleStore)

// 定义变量内容
const roleDialogFormRef = ref();
const state=reactive({
  isShowDialog: false,
  type: '',
  title: '',
  submitTxt: '',
})
//拿到选中的数据
const tree_ref = ref()


// 打开弹窗
const openDialog = (type: string, row: RowRoleType) => {
  state.isShowDialog = true;
  state.type = type
  if (type === 'edit') {
    dialog.value.ruleForm = row;
    //获取当前role的菜单节点
    dialog.value.ruleForm.nodes = arrayExtractionNodes(row.menus)
    state.title = '修改角色';
    state.submitTxt = '修 改';
  } else {
    state.title = '新增角色';
    state.submitTxt = '新 增';
  }
  //请求全部菜单tree
  routesStore.setRoutesTree() //全部菜单
};
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = () => {

  if (state.type === 'edit') {
    //获取全选节点
    //const checkMenu = tree_ref.value.getCheckedNodes();
    //获取半选节点
    // const halfCheckMenu = tree_ref.value.getHalfCheckedNodes()

    dialog.value.ruleForm.menus = [] //清空menu，没必要传输
    dialog.value.ruleForm.nodes = tree_ref.value.getCheckedKeys()
    roleApi.modifyRoleInfoApi(dialog.value.ruleForm).then((res) => {
      if (res.code === 0) {
        ElMessage.success('修改成功');
        //父组件重新加载
        emit('refresh');
      }
    })
    //关闭编辑弹窗
    closeDialog();
  } else {
    dialog.value.ruleForm.id = 0 //清空上次编辑的id
    dialog.value.ruleForm.nodes = tree_ref.value.getCheckedKeys()
    roleApi.addRoleApi(dialog.value.ruleForm).then((res) => {
      if (res.code === 0) {
        ElMessage.success('新建角色成功');
        //父组件重新加载
        emit('refresh');
      }
    })
    //关闭编辑弹窗
    closeDialog();
  }
};


// 暴露变量
defineExpose({
  openDialog,
});
</script>

<style scoped lang="scss">
.system-role-dialog-container {
  .menu-data-tree {
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
    padding: 5px;
  }
}
</style>
