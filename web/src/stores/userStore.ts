import {defineStore} from 'pinia';
//import Cookies from 'js-cookie';
import {Session} from '/@/utils/storage';
//导入api
import {useUserApi} from "../api/user/index";

const userApi = useUserApi()

//bcrypt
import bcrypt from 'bcryptjs'

export const useUserStore = defineStore('userInfo', {
    state: () => ({
        //登录页面数据
        loginData: {
            user_name: '',
            password: '',
            email_code: '',
        },
        //注册页面数据
        registerData: {
            user_name: '',
            password: '',
            re_password: '',
            email_code: '',
        },
        //全局用户信息
        userInfos: {
            created_at:'',
            updated_at:'',
            id: 0,
            uuid: 0,
            user_name: '',
            nick_name: '',
            password:'',
            avatar: '',
            phone: '',
            email: '',
            enable: true,
            user_role: [],	//角色组
            orders: [],      //订单
            subscribe_info: { //附加订阅信息
                expired_at: '',
                t: 0,
                u: 0,
                d: 0,
            }
        } as SysUser,
        //首页数据
        homeTableData: {
            host: '',//用户混淆
        },
        //个人信息页面数据
        changePasswordDialog: {
            isShowChangePasswordDialog: false,
        },
        //用户管理页面数据
        userManageData: {
            loading: false,
            users:{
                total: 0,
                user_list:[] as SysUser[],
            },
            params: {
                search: '',
                page_num: 1,
                page_size: 10,
            },
            dialog: {
                isShowDialog: false,
                type: '',
                title: '',
                submitTxt: '',
                userForm: {} as SysUser,
                check_list: [], //选中的角色
            },
        },
    }),
    getters: {
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
            if (state.userInfos.subscribe_info.goods_id == 0 || state.userInfos.subscribe_info.subscribe_url == '') {
                return ''
            }
            return import.meta.env.VITE_API_URL + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=1"
        },
        //ShadowRocket订阅
        subShadowRocket: (state) => {
            if (state.userInfos.subscribe_info.goods_id == 0 || state.userInfos.subscribe_info.subscribe_url == '') {
                return ''
            }
            return import.meta.env.VITE_API_URL + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=3"

        },
        //Clash订阅
        subClash: (state) => {
            if (state.userInfos.subscribe_info.goods_id == 0 || state.userInfos.subscribe_info.subscribe_url == '') {
                return ''
            }
            return import.meta.env.VITE_API_URL + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=2"

        },
        //Qx订阅
        subQx: (state) => {
            if (state.userInfos.subscribe_info.goods_id == 0 || state.userInfos.subscribe_info.subscribe_url == '') {
                return ''
            }
            return import.meta.env.VITE_API_URL + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url + "&type=4"
        },
        //

    },
    actions: {
        //注册
        async register(form?: object) {
            const res = await userApi.registerApi(this.registerData)
            if (res.code === 0) {
                return true
            }
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
            Session.set("token", res.data.token)
        },
        //修改混淆
        async changeHost() {
            const res = await userApi.changeHostApi(this.homeTableData)
            this.getUserInfo()
            this.homeTableData.host = ''
        },
        //获取自身信息
        async getUserInfo() {
            const res = await userApi.getUserInfoApi()
            this.userInfos = res.data
            Session.set("userInfos", res.data)
        },
        //获取用户列表
        async getUserList() {
            const res = await userApi.getUserListApi(this.userManageData.params)
            this.userManageData.users = res.data
        },
        //新建用户
        async newUser() {
            //密码加密
            this.userManageData.dialog.userForm.password = bcrypt.hashSync(this.userManageData.dialog.userForm.password, 10)
            const res = await userApi.newUserApi({
                user: this.userManageData.dialog.userForm,
                check_list: this.userManageData.dialog.check_list
            })
        },
        //修改用户
        async updateUser() {
            if (!this.userManageData.dialog.userForm.password.startsWith('$2a$10$')) {
                this.userManageData.dialog.userForm.password = bcrypt.hashSync(this.userManageData.dialog.userForm.password, 10)
            }
            const res = await userApi.updateUserApi({
                user: this.userManageData.dialog.userForm,
                check_list: this.userManageData.dialog.check_list
            })
        },
        //删除用户
        async deleteUser(params?: object) {
            const res = await userApi.deleteUserApi(params)
        },
        //修改密码
        async changePassword() {
            return await userApi.changePasswordApi(this.registerData)
        },
        //确认重置密码
        async submitResetPassword() {
            return await userApi.resetPasswordApi(this.loginData)
        }
    },
});
