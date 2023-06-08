<template>
  <div >
  <el-row :gutter="15" >
    <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24"  v-for="(v,k) in serverStatusData.data" :key="k">
      <div class="home-card-item">
        <el-card class="box-card">
          <el-row :gutter="10" justify="space-around" align="middle" >
            <el-col :xs="12" :sm="12" :md="5" :lg="5" :xl="5" style="margin-top: 20px;margin-bottom: 10px">
              {{ v.name }}
            </el-col>
            <el-col :xs="12" :sm="12" :md="5" :lg="5" :xl="5" style="margin-top: 20px;margin-bottom: 10px">
              倍率：{{v.traffic_rate}}
            </el-col>
            <el-col :xs="12" :sm="12" :md="5" :lg="5" :xl="5" style="margin-top: 20px;margin-bottom: 10px">
              <el-button v-if="v.status" type="success" plain>在线</el-button>
              <el-button v-else type="info" plain>离线</el-button>
            </el-col>
            <el-col :xs="12" :sm="12" :md="5" :lg="5" :xl="5" style="margin-top: 20px;margin-bottom: 10px">
              {{ v.user_amount }}人在线
            </el-col>
            <el-col :xs="12" :sm="12" :md="5" :lg="5" :xl="5" style="margin-top: 20px;margin-bottom: 10px">
              <el-icon color="#409EFC"><Top /></el-icon><span>{{ v.u }}MB/s</span>
            </el-col>
            <el-col :xs="12" :sm="12" :md="4" :lg="4" :xl="4" style="margin-top: 20px;margin-bottom: 10px">
              <el-icon color="#409EFC"><Bottom /></el-icon><span>{{ v.d }}MB/s</span>
            </el-col>
          </el-row>
        </el-card>

      </div>
    </el-col>
  </el-row>
  </div>

</template>
<!--<script setup lang="ts" name="personal">-->
<script setup lang="ts">
import {reactive, computed, onMounted, onUnmounted} from "vue";
import {storeToRefs} from "pinia";

//node store
import {useNodeStore} from "/@/stores/node";
import {Session} from "/@/utils/storage";
const nodeStore = useNodeStore()
const {serverStatusData} = storeToRefs(nodeStore)

//var ws = new WebSocket("ws://localhost/ws",[token]);
const token =Session.get('token')
var ws = new WebSocket('ws:///192.168.0.8:8899/websocket/msg',token);
var interval = null;//计时器
//监听是否连接成功
function initWS(){
  ws.onopen = function () {
    console.log('ws连接成功,连接状态：' + ws.readyState);
    ws.send('{"type":1,"data":"hi"}');
    interval = setInterval(() => {
      ws.send('{"type":1,"data":"hi"}');
    }, 3000);
  }
//接收服务器发回的信息
  ws.onmessage = function (data) {
    console.log('ws接收服务器发回的信息：' + ws.readyState);
    serverStatusData.value = JSON.parse(data.data)
   // console.log("JSON.parse:", JSON.parse(data.data))
  }
//监听连接关闭事件
  ws.onclose = function () {
   // console.log('ws连接关闭,连接状态：' + ws.readyState);
    clearInterval(interval)
    ws.close();
  }
//监听并处理error事件
  ws.onerror = function (error) {
   // console.log(error);
    ws.close();
    clearInterval(interval)
  }
}

onMounted(() => {
  initWS()
});
onUnmounted(() => {
  //完成通信后关闭WebSocket连接
   ws.close();
  clearInterval(interval)
});
</script>

<style scoped>
.home-card-item {
  font-size: 16px;
  width: 100%;
  height: 100%;
  border-radius: 4px;
  transition: all ease 0.3s;
  padding: 20px;
  overflow: hidden;
  background: var(--el-color-white);
  color: var(--el-text-color-primary);
  border: 1px solid var(--next-border-color-light);
}
</style>