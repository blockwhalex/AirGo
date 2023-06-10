//角色-store
import {defineStore} from "pinia";
import {Session} from "/@/utils/storage";

//role api
import {useRoleApi} from "/@/api/role/index";
import {ElMessage} from "element-plus";

const roleApi = useRoleApi()

//暴露role store
export const useRoleStore = defineStore("roleStore", {
    state: () => ({
        //编辑角色弹窗参数
        dialog: {
            isShowDialog: false,
            type: '',
            title: '',
            submitTxt: '',
            ruleForm: {           //当前编辑角色的参数
                id: 0,        //角色ID
                role_name: '',     // 角色名称
                description: '',  //角色描述
                status: true,     //角色状态
                menus: [],        //角色菜单
                nodes: [],
            } as RowRoleType,
        },
        //编辑角色Api弹窗参数
        dialogEditApi: {
            isShowDialog: false,
            type: '',
            title: '',
            submitTxt: '',
            casbinInfo: {    //当前角色Api
                roleID: 0,
                casbinItems: [''],
            },
            allCasbinInfo: {    //全部api
                roleID: 0,
                casbinItems: [],
            } as CasbinInfo,
        },

        //角色管理参数
        roleManageData: {
            roles: {
                total: 0,
                role_list: [] as RowRoleType[],
            },
            loading: false,
            params: {
                search: '',
                page_num: 1,
                page_size: 10,
            },
        }
    }),
    actions: {
        //params===undefined,角色列表,分页;否则查询全部
        async getRoleList(params?: object) {
            if (params != undefined) {
                const res: any = await roleApi.getRoleListApi(params)
                this.roleManageData.roles = res.data
            } else {
                const res: any = await roleApi.getRoleListApi(this.roleManageData.params)
                this.roleManageData.roles = res.data
            }
        },

        //获取当前角色的权限
        async getPolicyByRoleIds() {
            const res = await roleApi.getPolicyByRoleIdsApi()
            // this.dialogEditApi.casbinInfo = res.data
            var casbinRes: CasbinInfo = res.data
            if (casbinRes.casbinItems !== null) {
                var oldArr: string[] = []
                casbinRes.casbinItems.forEach((item: CasbinItem) => {
                    oldArr.push(item.path)
                });
                this.dialogEditApi.casbinInfo.casbinItems = oldArr
                // console.log("oldArr:",oldArr)
            }
        },
        //获取全部权限
        async getAllPolicy() {
            const res = await roleApi.getAllPolicyApi()
            this.dialogEditApi.allCasbinInfo = res.data
        },
        //更新角色权限
        async updateCasbinPolicy(): Promise<boolean> {
            const res = await roleApi.updateCasbinPolicyApi({
                roleID: this.dialogEditApi.casbinInfo.roleID,
                data: this.dialogEditApi.casbinInfo.casbinItems
            })
            if (res.code === 0) {
                ElMessage.success(res.msg)
            } else {

            }
            return true
        }
    }
})