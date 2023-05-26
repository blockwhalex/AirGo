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
                casbinItems: [],
            } as CasbinInfo,
            allCasbinInfo: {    //全部api
                roleID: 0,
                casbinItems: [],
            } as CasbinInfo,
        },

        //角色管理默认页参数
        tableData: {
            data: [] as RowRoleType[],//所有角色列表
            total: 0,
            loading: false,
            param: {
                search: '',
                page_num: 1,
                page_size: 10,
            },
        },
    }),
    actions: {
        //所有角色列表,分页
        async setRoleList() {
            const res: any = await roleApi.getAllRoleListApi(this.tableData.param)
            // console.log("所有rolelist：",res.data)
            this.tableData.data = res.data
        },
        //修改role 信息
        //新建角色
        //查询角色 by role name
        //删除角色

        //获取当前角色的权限
        async getPolicyByRoleIds() {
            const res = await roleApi.getPolicyByRoleIdsApi()
            //console.log("获取当前角色的权限:",res.data)
        },
        //获取全部权限
        async getAllPolicy() {
            const res = await roleApi.getAllPolicyApi()
            this.dialogEditApi.allCasbinInfo = res.data
            // console.log("获取全部权限:",this.adminCasbinInfo)

        },
        //更新角色权限
        async updateCasbinPolicy(): Promise<boolean> {
            const res = await roleApi.updateCasbinPolicyApi({
                roleID: this.dialogEditApi.casbinInfo.roleID,
                data: this.dialogEditApi.casbinInfo.casbinItems
            })
            // console.log("更新角色权限:",res.data)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            } else {
                ElMessage.error(res.msg)
            }
            return true
        }
    }
})