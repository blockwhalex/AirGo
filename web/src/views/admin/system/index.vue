<template>
  <div style="padding: 15px;">
    <el-card>
      <el-tabs stretch style="height: 100%" class="demo-tabs">
        <el-tab-pane label="登录/注册">
          <el-form :model="serverConfig" label-width="120px">
            <el-form-item label="是否开启注册">
              <el-switch v-model="serverConfig.system.enable_register" inline-prompt active-text="开启"
                         inactive-text="关闭"
                         style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"></el-switch>
            </el-form-item>
            <el-form-item label="注册邮箱验证码">
              <el-switch v-model="serverConfig.system.enable_email_code" inline-prompt active-text="开启"
                         inactive-text="关闭"
                         style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"></el-switch>
            </el-form-item>


            <el-divider></el-divider>
            <el-form-item label="是否多点登录">
              <el-switch v-model="serverConfig.system.is_multipoint" inline-prompt active-text="开启"
                         inactive-text="关闭"
                         style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"></el-switch>
            </el-form-item>

            <el-form-item label="IP限流">
              <el-col :span="4">
                <el-input v-model.number="serverConfig.rate_limit_params.ip_role_param" type="number"/>
              </el-col>
              <el-col :span="2" style="text-align: center">
                <span>-</span>
              </el-col>
              <el-col :span="18">
                <span class="text-gray-500">请求/分钟</span>
              </el-col>
            </el-form-item>

            <el-form-item label="用户限流">
              <el-col :span="4">
                <el-input v-model.number="serverConfig.rate_limit_params.visit_param" type="number"/>
              </el-col>
              <el-col :span="2" style="text-align: center">
                <span>-</span>
              </el-col>
              <el-col :span="18">
                <span class="text-gray-500">请求/分钟</span>
              </el-col>
            </el-form-item>


            <el-divider></el-divider>
            <el-form-item label="通信密钥">
              <el-col :span="12">
                <el-input v-model="serverConfig.system.muKey" placeholder="务必前后端保持一致！"/>
              </el-col>

            </el-form-item>
            <el-form-item label="订阅名称">
              <el-col :span="12">
                <el-input v-model="serverConfig.system.sub_name"/>
              </el-col>

            </el-form-item>
            <el-form-item label="新注册分配套餐">
              <el-select v-model="serverConfig.system.default_goods" class="m-2" placeholder="选择套餐">
                <el-option
                    v-for="item in goodsList"
                    :key="item.id"
                    :label="item.subject"
                    :value="item.subject"
                />
              </el-select>
            </el-form-item>
            <el-divider></el-divider>
            <el-form-item>
              <el-button @click="onSubmit" type="primary">保存</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>


        <el-tab-pane label="支付">
          <el-divider content-position="left">支付宝 支付</el-divider>
          <el-form :model="serverConfig" label-width="120px">
            <el-form-item label="支付宝appID">
              <el-input v-model="serverConfig.pay.app_id"/>
            </el-form-item>
            <el-form-item label="支付宝应用私钥">
              <el-input v-model="serverConfig.pay.private_key" type="textarea" autosize/>
            </el-form-item>
            <el-form-item label="支付宝公钥">
              <el-input v-model="serverConfig.pay.ali_public_key" type="textarea" autosize/>
            </el-form-item>
            <el-form-item label="支付宝加密密钥">
              <el-input v-model="serverConfig.pay.encrypt_key" type="password"/>
            </el-form-item>
            <el-divider content-position="left">微信 支付</el-divider>
            <el-form-item>
              <el-button @click="onSubmit" type="primary">保存</el-button>
            </el-form-item>
          </el-form>

        </el-tab-pane>
        <el-tab-pane label="邮件">
          <el-form :model="serverConfig" label-width="100px">
            <el-form-item label="服务器地址">
              <el-input v-model="serverConfig.email.email_host" placeholder="mail.example.com"/>
            </el-form-item>
            <el-form-item label="端口">
              <el-input v-model.number="serverConfig.email.email_port" type="number"/>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model="serverConfig.email.email_from" placeholder="admin@oicq.com"/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="serverConfig.email.email_secret" type="password"/>
            </el-form-item>
            <el-form-item label="邮件主题">
              <el-input v-model="serverConfig.email.email_subject"/>
            </el-form-item>
            <el-form-item label="邮件内容格式">
              <el-input v-model="serverConfig.email.email_content" type="textarea"/>
              <el-tag class="ml-2" type="warning">*emailcode 字段不可删除！</el-tag>
            </el-form-item>
            <el-divider></el-divider>
            <el-form-item>
              <el-button @click="onSubmit" type="primary">保存</el-button>
            </el-form-item>

          </el-form>
        </el-tab-pane>


        <el-tab-pane label="json web token">
          <el-form :model="serverConfig" label-width="100px">
            <el-form-item label="jwt签名">
              <el-input v-model="serverConfig.jwt.signing_key"/>
            </el-form-item>
            <el-form-item label="签发者">
              <el-input v-model="serverConfig.jwt.issuer"/>
            </el-form-item>
            <el-form-item label="过期时间">
              <el-input v-model="serverConfig.jwt.expires_time"/>
            </el-form-item>
            <el-divider></el-divider>
            <el-form-item>
              <el-button @click="onSubmit" type="primary">保存</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <!--        <el-tab-pane label="其他">-->
        <!--          <el-form :model="serverConfig" label-width="100px">-->
        <!--            <el-form-item>-->
        <!--              <el-button @click="onSubmit" type="primary">保存</el-button>-->
        <!--            </el-form-item>-->
        <!--          </el-form>-->
        <!--        </el-tab-pane>-->
      </el-tabs>
    </el-card>

  </div>
</template>

<script lang="ts" setup>
import {ElMessage} from 'element-plus';
import {onMounted} from "vue";
//server store
import {useServerStore} from "/@/stores/serverStore";
import {storeToRefs} from "pinia";

const serverStore = useServerStore()
const {serverConfig} = storeToRefs(serverStore)
//shop store
import {useShopStore} from "/@/stores/shopStore";

const shopStore = useShopStore()
const {goodsList} = storeToRefs(shopStore)
//保存提交
const onSubmit = () => {
  serverStore.updateServerConfig().then((res) => {
    if (res.code === 0) {
      ElMessage.success(res.msg)
    } else {
      ElMessage.error(res.msg)
    }
  })
  setTimeout(() => {
    serverStore.getServerConfig()
  }, 2000)
}
//加载时
onMounted(() => {
  serverStore.getServerConfig()
  shopStore.getAllGoods()

});

</script>

<style lang="scss" scoped>

</style>