<template>
    <!-- 二维码弹窗 -->
    <el-dialog v-model="tableData.isShowQRDialog" title="扫描二维码支付" width="30%">
        <div>
        <el-card shadow="hover">
            <div class="qrcode-img-warp">
                <div class="mb30 mt30 qrcode-img">
                    <div id="qrcode" class="qrcode" ref="qrcodeRef"></div>
                </div>
            </div>
        </el-card>
    </div>
    <template #footer>
            <span class="dialog-footer">
                <el-button @click="tableData.isShowQRDialog = false">取消</el-button>
                <el-button type="primary" @click="tableData.isShowQRDialog = false">
                    已支付
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref ,nextTick} from "vue";
//导入二维码 js
import QRCode from 'qrcodejs2-fixes';
// 定义二维码变量内容
const qrcodeRef = ref();
//导入store
import { storeToRefs } from 'pinia';
import { useShopStore } from "/@/stores/shopStore";
const shopStore = useShopStore()
const { tableData} = storeToRefs(shopStore)

//二维码
const  onInitQrcode = () => {
    //清除上一次二维码
    const codeHtml = document.getElementById("qrcode");
      codeHtml.innerHTML = "";

    tableData.value.QRcode= new QRCode(qrcodeRef.value, {
        text: tableData.value.QRtext,
        width: 125,
        height: 125,
        colorDark: '#000000',
        colorLight: '#ffffff',
    });
}
//打开弹窗
const openDialog=(goods?:object)=>{
    tableData.value.isShowQRDialog=true
    nextTick(() => {
        onInitQrcode()
    })
}

defineExpose({
    openDialog,
})
</script>

<style scoped></style>