import {defineStore} from "pinia";
//api
import {useSystemApi} from "/@/api/system/index";

const systemApi = useSystemApi()
export const useServerStore = defineStore("serverStore", {
    state: () => ({
        serverConfig: {
            created_at:'',
            updated_at:'',
            id: 0,
            jwt: {
                signing_key: '',
                expires_time: '',
                buffer_time: '',
                issuer: '',
            },
            system: {
                enable_register: true,
                enable_email_code: false,
                is_multipoint:true,
                sub_name:'',
            },
            captcha: {
                key_long: 0,
                img_width: 0,
                img_height: 0,
                open_captcha: 0,
                open_captcha_time_out: 0,
            },
            pay: {
                return_url: '',
                app_id: '',
                private_key: '',
                ali_public_key: '',
                encrypt_key: '',
            },
            email: {
                email_from: '',
                email_secret: '',
                email_host: '',
                email_port: 0,
                email_is_ssl: true,
                email_nickname: '',
                email_subject:'',
                email_content:'',
            },
        } as Server,
        other: 0,

    }),
    actions: {
        //获取系统设置
        async getServerConfig() {
            const res = await systemApi.getServerConfig()
            this.serverConfig = res.data
        },
        //修改系统设置
        async updateServerConfig() {
            const res = await systemApi.updateServerConfig(this.serverConfig)
            //this.serverConfig=res.data
        }

    }
})