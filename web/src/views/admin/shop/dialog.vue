<template>
  <el-dialog v-model="state.isShowDialog" :title="state.title" width="769px" destroy-on-close>
    <el-form :model="goodsManageData.currentGoods" label-width="120px">
      <el-form-item label="商品标题">
        <el-input v-model="goodsManageData.currentGoods.subject"/>
      </el-form-item>

      <el-form-item label="描述(支持HTML编辑)">
        <el-input v-model="goodsManageData.currentGoods.des" type="textarea" autosize/>
      </el-form-item>
      <el-form-item label="价格">
        <el-col :span="4">
          <el-input v-model="goodsManageData.currentGoods.total_amount"/>
        </el-col>
        <el-col :span="2" style="text-align: center">
          <span>-</span>
        </el-col>
        <el-col :span="18">
          <span class="text-gray-500">RMB</span>
        </el-col>
      </el-form-item>


      <el-form-item label="总流量">
        <el-col :span="4">
          <!-- input会自动将数字类型转换为了字符串类型，导致form表单提交后端报错，解决方案是Vue的修饰符 -->
          <el-input v-model.number="goodsManageData.currentGoods.total_bandwidth" type="number"/>
        </el-col>
        <el-col :span="2" style="text-align: center">
          <span>-</span>
        </el-col>
        <el-col :span="18">
          <span class="text-gray-500">GB</span>
        </el-col>
      </el-form-item>

      <el-form-item label="有效期">
        <el-col :span="4">
          <el-input v-model.number="goodsManageData.currentGoods.expiration_date" type="number"/>
        </el-col>
        <el-col :span="2" style="text-align: center">
          <span>-</span>
        </el-col>
        <el-col :span="18">
          <span class="text-gray-500">天</span>
        </el-col>

      </el-form-item>

      <el-form-item label="是否显示">
        <el-col :span="4">
          <el-switch v-model="goodsManageData.currentGoods.status" inline-prompt active-text="开启" inactive-text="关闭"
                     style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"></el-switch>
        </el-col>
      </el-form-item>

      <el-form-item label="关联节点">
        <el-transfer
            :data="nodeManageData.nodes.node_list"
            v-model="goodsManageData.currentGoods.checked_nodes"
            :right-default-checked="goodsManageData.currentGoods.checked_nodes"
            :props="{
                  key: 'id',
                  label: 'name',
                  }"
            :titles="['全部节点', '选中节点']"

        />
      </el-form-item>
    </el-form>
    <template #footer>
            <span class="dialog-footer">
                <el-button @click="state.isShowDialog = false">取消</el-button>
                <el-button type="primary" @click="onSubmit">
                    确认
                </el-button>
            </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
//导入store
import {storeToRefs} from 'pinia';
import {useShopStore} from "/@/stores/shopStore";

const shopStore = useShopStore()
const {goodsManageData} = storeToRefs(shopStore)
//store
import {useNodeStore} from "/@/stores/node";
import {reactive} from "vue";

const nodeStore = useNodeStore();
const {nodeManageData} = storeToRefs(nodeStore);
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

//定义参数
const state = reactive({
  isShowDialog: false,
  type: '',
  title: '',
})

// 打开弹窗
const openDialog = (type: string, row?: any) => {
  if (type == 'add') {
    state.type = type
    state.title = "新建商品"
    state.isShowDialog = true
    goodsManageData.value.currentGoods.id = 0 //清空上次编辑的id，否则无法新建
  } else {
    state.type = type
    state.title = "修改商品"
    goodsManageData.value.currentGoods = row //将当前row写入pinia
    if (goodsManageData.value.currentGoods.checked_nodes === null || goodsManageData.value.currentGoods.checked_nodes === undefined) {
      goodsManageData.value.currentGoods.checked_nodes = [] //剔除null,否则ts报错
    }
    goodsManageData.value.currentGoods.nodes = [] //清空nodes
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
    shopStore.newGoods(goodsManageData.value.currentGoods)
    setTimeout(() => {
      emit('refresh');
    }, 1000);
  } else {
    shopStore.updateGoods(goodsManageData.value.currentGoods)
    setTimeout(() => {
      emit('refresh');
    }, 1000);
  }
  closeDialog();
}

// 暴露变量
defineExpose({
  openDialog,
});

</script>


<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
  