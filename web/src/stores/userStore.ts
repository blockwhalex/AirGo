import {defineStore} from 'pinia';
//import Cookies from 'js-cookie';
import {Local, Session} from '/@/utils/storage';
//导入api
import {useUserApi} from "../api/user/index";
//bcrypt
import bcrypt from 'bcryptjs'
import {ElMessage} from "element-plus";
const userApi = useUserApi()
//server store
import {useServerStore} from "/@/stores/serverStore";

export const useUserStore = defineStore('userInfo', {
    state: () => ({
        //登录页面数据
        loginData: {
            user_name: '',
            password: '',
            email_code: '',
        },
        //注册页面数据
        email_suffix: 'qq.com',
        registerReq: {
            user_name: '',
            password: '',
            re_password: '',
            email_code: '',
            referrer_code: '',
        },
        registerData: {
            user_name: '',
            password: '',
            re_password: '',
            email_code: '',
            referrer_code: '',
        },
        //全局用户信息
        userInfos: {
            created_at: '',
            updated_at: '',
            id: 0,
            uuid: 0,
            user_name: '',
            nick_name: '',
            password: '',
            avatar: '',
            phone: '',
            email: '',
            enable: true,
            role_group: [] as RowRoleType[],	//角色组
            orders: [],      //订单
            subscribe_info: { //附加订阅信息
                expired_at: '',
                t: 0,
                u: 0,
                d: 0,
            }
        } as SysUser,
        //用户管理页面数据
        userManageData: {
            users: {
                total: 0,
                user_list: [] as SysUser[],
            },
            dialog: {
                user: {
                    user_name: '',
                    nick_name: '',
                    password: '',
                    avatar: '',
                    phone: '',
                    email: '',
                    enable: true,
                    role_group: [] as RowRoleType[],
                    subscribe_info: {
                        sub_status: true,
                        expired_at: '',
                        t: 0,
                        u: 0,
                        d: 0,
                        node_speedlimit: 0,
                        node_connector: 3,
                    }
                } as SysUser,
                check_list: [''], //选中的角色
            },
        },
    }),
    getters: {
        userFormReq: (state): object => {
            state.registerReq.password = state.registerData.password
            state.registerReq.re_password = state.registerData.re_password
            state.registerReq.email_code = state.registerData.email_code
            state.registerReq.user_name = state.registerData.user_name + '@' + state.email_suffix
            state.registerReq.referrer_code = state.registerData.referrer_code
            return state.registerReq
        },
        used: (state): number => {
            return +((state.userInfos.subscribe_info.t - state.userInfos.subscribe_info.u - state.userInfos.subscribe_info.d) / 1024 / 1024 / 1024).toFixed(2)
        },
        expired: (state): string => {
            if (state.userInfos.subscribe_info.expired_at === null) {
                return ""
            } else {
                return state.userInfos.subscribe_info.expired_at.slice(0, 10)
            }
        },
        //v2rayNG订阅
        subV2rayNG: (state): string => {
            const serverStore= useServerStore()
            return serverStore.publicServerConfig.sub_url_pre + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=1"
        },
        //ShadowRocket订阅
        subShadowRocket: (state) => {
            const serverStore= useServerStore()
            return serverStore.publicServerConfig.sub_url_pre + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=3"

        },
        //Clash订阅
        subClash: (state) => {
            const serverStore= useServerStore()
            return serverStore.publicServerConfig.sub_url_pre + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=2"

        },
        //Qx订阅
        subQx: (state) => {
            const serverStore= useServerStore()
            return serverStore.publicServerConfig.sub_url_pre + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=4"
        },
        //

    },
    actions: {
        //注册
        async register(form?: object) {
            this.registerData.referrer_code = Local.get('invitation')
            const res = await userApi.registerApi(this.userFormReq)
            return res
        },
        //登录
        async userLogin(form?: any) {
            // 存储用户信息到浏览器缓存
            const res: any = await userApi.signIn(form);
            //保存用户信息到pinia
            this.userInfos = res.data.user;
            //保存用户信息到Session
            Session.set("userInfos", res.data.user)
            //保存token
            // Session.set("token", res.data.token)
            Local.set("token", res.data.token)
        },
        //修改混淆
        async changeHost(params?: object) {
            const res = await userApi.changeHostApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
                await this.getUserInfo()
            }
        },
        //获取自身信息
        async getUserInfo() {
            const res = await userApi.getUserInfoApi()
            if (res.code === 0) {
                // ElMessage.success(res.msg)
                this.userInfos = res.data
                Session.set("userInfos", res.data)
            }
        },
        //获取用户列表
        async getUserList(params?: object) {
            const res = await userApi.getUserListApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
                this.userManageData.users = res.data
            }
        },
        //新建用户
        async newUser(params?: object) {
            //密码加密
            this.userManageData.dialog.user.password = bcrypt.hashSync(this.userManageData.dialog.user.password, 10)
            const res = await userApi.newUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //修改用户
        async updateUser(params?: object) {
            if (!this.userManageData.dialog.user.password.startsWith('$2a$10$')) {
                this.userManageData.dialog.user.password = bcrypt.hashSync(this.userManageData.dialog.user.password, 10)
            }
            const res = await userApi.updateUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //删除用户
        async deleteUser(params?: object) {
            const res = await userApi.deleteUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //修改密码
        async changePassword(params?: object) {
            const res = await userApi.changePasswordApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //确认重置密码
        async submitResetPassword() {
            return await userApi.resetPasswordApi(this.loginData)
        }
    },
});
