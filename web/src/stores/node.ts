import {defineStore} from "pinia";
//api
import {useNodeApi} from "/@/api/node/index";
import {ElMessage} from "element-plus";


const nodeApi = useNodeApi()

export const useNodeStore = defineStore("nodeStore", {
    state: () => ({
        nodeList: [] as Node[],
        //默认页数据
        tableData: {
            total: 0,
            loading: true,
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
            } as Node,
        },

    }),
    actions: {
        //获取全部节点
        async getAllnode() {
            const res = await nodeApi.getAllNodeApi(this.tableData.params)
            this.nodeList = res.data
            this.tableData.loading = false
            this.tableData.total = this.nodeList.length

        },
        //更新节点
        async updateNode() {
            const res = await nodeApi.updateNodeApi(this.dialogData.nodeInfo)
            //console.log("更新节点:", res)
            if (res.code === 0) {
                ElMessage.success("更新节点成功")
            }
        },
        //删除节点
        async deleteNode(params: object) {
            const res = await nodeApi.deleteNodeApi(params)
            //console.log("更新节点:", res)
            if (res.code === 0) {
                ElMessage.success("删除节点成功")
            }
        },
        //新建节点
        async newNode() {
            const res = await nodeApi.newNodeApi(this.dialogData.nodeInfo)
            //console.log("新建节点:", res)
            if (res.code === 0) {
                ElMessage.success("新建节点成功")
            }
        },
        async getNodeByName() {
            const res = await nodeApi.getNodeByNameApi(this.tableData.params)
            this.nodeList = res.data
            this.tableData.loading = false
            this.tableData.total = this.nodeList.length
        },
        async getNodeTraffic() {
            const res = await nodeApi.getNodeTrafficApi(this.tableData.params)
        }

    }
})