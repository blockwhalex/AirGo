import {defineStore} from "pinia";
//api
import {useNodeApi} from "/@/api/node/index";
import {ElMessage} from "element-plus";


const nodeApi = useNodeApi()

export const useNodeStore = defineStore("nodeStore", {
    state: () => ({
        //节点管理页数据
        nodeManageData: {
            loading: true,
            nodes: {
                total: 0,
                node_list: [] as NodeInfo[],
            },
            params: {
                search: '',
                page_num: 1,
                page_size: 10,
                date: [],
            },

        },
        //弹窗页数据
        dialogData: {
            type: "",
            title: "",
            isShowDialog: false,
            nodeInfo: {
                id: 0,
                name: '',
                address: '',
                port: '',
                sort: 11,          //类型sort==11  V2Ray vmess
                nodespeed_limit: 0, //限速
                traffic_rate: 1,     //倍率
                status: true,          //节点状态，true启用，flase禁用
                //vmess参数
                v: '',
                aid: '',
                scy: '', //加密方式
                net: '', //传输协议 ws tcp kcp quic grpc等
                disguise_type: '',  //伪装类型 none http
                host: '',
                path: '',
                tls: "",
                sni: "",
            } as NodeInfo,
        },
        //节点状态页面数据
        serverStatusData: {
            type: 0,
            data: [] as ServerStatusInfo[],
        },

    }),
    actions: {
        //获取全部节点
        async getAllNode() {
            const res = await nodeApi.getAllNodeApi()
            this.nodeManageData.nodes.node_list = res.data
            this.nodeManageData.loading = false

        },
        //获取全部节点 with Traffic,分页
        async getNodeWithTraffic(params?: object) {
            const res = await nodeApi.getNodeWithTrafficApi(this.nodeManageData.params)
            this.nodeManageData.nodes = res.data
            this.nodeManageData.loading = false

        },
        //获取全部节点 with Traffic,订单统计
        async getNodeStatistics(params?: object) {
            const res = await nodeApi.getNodeWithTrafficApi(params)
            return res
        },
        //更新节点
        async updateNode() {
            const res = await nodeApi.updateNodeApi(this.dialogData.nodeInfo)
            if (res.code === 0) {
                ElMessage.success("更新节点成功")
            }
        },
        //删除节点
        async deleteNode(params: object) {
            const res = await nodeApi.deleteNodeApi(params)
            if (res.code === 0) {
                ElMessage.success("删除节点成功")
            }
        },
        //新建节点
        async newNode() {
            const res = await nodeApi.newNodeApi(this.dialogData.nodeInfo)
            if (res.code === 0) {
                ElMessage.success("新建节点成功")
            }
        },
    }
})