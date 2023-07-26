<template>
  <el-dialog v-model="state.isShowDialog" :title="state.title" width="769px" destroy-on-close align-center>
    <el-divider content-position="left">节点参数</el-divider>
    <el-form :model="dialogData.nodeInfo" label-width="120px">
      <el-form-item label="Protocol">
        <el-radio-group v-model="dialogData.nodeInfo.sort">
          <el-radio :label="11">vmess</el-radio>
          <el-radio :label="15">vless</el-radio>
          <el-radio :label="14">trojan</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Name">
        <el-input v-model="dialogData.nodeInfo.remarks" placeholder="这是一个节点"/>
      </el-form-item>
      <el-form-item label="Address">
        <el-input v-model="dialogData.nodeInfo.address" placeholder="110.110.110.110"/>
      </el-form-item>
      <el-form-item label="Port">
        <el-input v-model.number="dialogData.nodeInfo.port" placeholder="80"/>
      </el-form-item>
      <el-form-item label="Host">
        <el-input v-model="dialogData.nodeInfo.host" placeholder="www.189.cn"/>
      </el-form-item>
      <el-form-item label="Path">
        <el-input v-model="dialogData.nodeInfo.path" placeholder="/"/>
      </el-form-item>
      <el-form-item label="Security" v-if="dialogData.nodeInfo.sort !== 14">
        <el-radio-group v-model="dialogData.nodeInfo.scy">
          <el-radio :label="'auto'">auto</el-radio>
          <el-radio :label="'none'">none</el-radio>
          <el-radio :label="'chacha20-poly1305'">chacha20-poly1305</el-radio>
          <el-radio :label="'aes-128-gcm'">aes-128-gcm</el-radio>
          <el-radio :label="'zero'">zero</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Network">
        <!--              <el-input v-model="dialogData.nodeInfo.net"  placeholder="ws"/>-->
        <el-radio-group v-model="dialogData.nodeInfo.network">
          <el-radio :label="'ws'">ws</el-radio>
          <el-radio :label="'tcp'">tcp</el-radio>
          <el-radio :label="'grpc'">grpc</el-radio>
          <el-radio :label="'quic'">quic</el-radio>
          <el-radio :label="'kcp'">kcp</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Tls">
        <el-radio-group v-model="dialogData.nodeInfo.security">
          <el-radio :label="''">none</el-radio>
          <el-radio :label="'tls'">tls</el-radio>
          <el-radio :label="'reality'">reality</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Sni" v-if="dialogData.nodeInfo.security !== ''">
        <el-input v-model="dialogData.nodeInfo.sni" placeholder=""/>
      </el-form-item>
    </el-form>
    <el-divider content-position="left">其他参数</el-divider>
    <el-form :model="dialogData.nodeInfo" label-width="120px">
      <el-form-item label="是否启用">
        <el-switch
            size="small"
            v-model="dialogData.nodeInfo.enabled"
            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
        />
      </el-form-item>
      <el-form-item label="节点限速">
        <el-input type="number" v-model="dialogData.nodeInfo.node_speedlimit" placeholder="0"/>
      </el-form-item>
      <el-form-item label="节点倍率">
        <el-input type="number" v-model="dialogData.nodeInfo.traffic_rate" placeholder="1"/>
      </el-form-item>
      <el-form-item label="启用中转">
        <el-switch
            size="small"
            v-model="dialogData.nodeInfo.enable_transfer"
            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
        />
      </el-form-item>
      <el-form-item label="中转ip" v-if="dialogData.nodeInfo.enable_transfer">
        <el-input v-model="dialogData.nodeInfo.transfer_address" placeholder=""/>
      </el-form-item>
      <el-form-item label="中转端口" v-if="dialogData.nodeInfo.enable_transfer">
        <el-input v-model="dialogData.nodeInfo.transfer_port" placeholder=""/>
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

import {storeToRefs} from "pinia";
//store
import {useNodeStore} from "/@/stores/node";
import {reactive} from "vue";

const nodeStore = useNodeStore()
const {dialogData} = storeToRefs(nodeStore)
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
//定义参数
const state = reactive({
  type: "",
  title: "",
  isShowDialog: false,
})

// 打开弹窗
const openDialog = (type: string, row?: any) => {
  if (type == 'add') {
    dialogData.value.nodeInfo.id = 0 //编辑和添加公用一个store，清空id,否则服务器无法插入
    state.type = type
    state.title = "新建节点"
    state.isShowDialog = true
  } else {
    state.type = type
    state.title = "修改节点"
    dialogData.value.nodeInfo = row  //将当前row写入pinia
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
    //新建节点
    nodeStore.newNode(dialogData.value.nodeInfo)
    setTimeout(() => {
      emit('refresh');
    }, 1000);       //延时。防止没新建完成就重新请求
  } else {
    //更新节点
    nodeStore.updateNode(dialogData.value.nodeInfo)
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
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
  