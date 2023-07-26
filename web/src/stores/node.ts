import {defineStore} from "pinia";
import {ElMessage} from "element-plus";
//api
import {useNodeApi} from "/@/api/node/index";

const nodeApi = useNodeApi()

export const useNodeStore = defineStore("nodeStore", {
    state: () => ({
        //节点管理页数据
        nodeManageData: {
            nodes: {
                total: 0,
                node_list: [] as NodeInfo[],
            },
        },
        //弹窗页数据
        dialogData: {
            nodeInfo: {
                // created_at: string;
                // updated_at: string;
                id: 0,
                node_speedlimit: 0, //节点限速/Mbps
                traffic_rate: 0,    //倍率
                sort: 15,            //类型 vless(15) vmess(11) trojan(14)
                server: '',
                // type: string;//显示与隐藏
                //基础参数
                remarks: '',//别名
                address: '',
                port: 0,
                node_order: 0,//节点排序
                enabled: true,  //是否为激活节点
                //中转参数
                enable_transfer: false,//是否启用中转
                transfer_address: '',//中转ip
                transfer_port: 0,   //中转port
                //
                total_up: 0,
                total_down: 0,

                goods: [],//多对多,关联商品
                //vmess参数
                v: '',
                scy: '',//加密方式 auto,none,chacha20-poly1305,aes-128-gcm,zero
                aid: 0,//额外ID
                //vless参数
                flow: '',//流控 none,xtls-rprx-vision,xtls-rprx-vision-udp443
                encryption: '',//加密方式 none

                network: '',//传输协议 tcp,kcp,ws,h2,quic,grpc
                type: '',   //伪装类型 ws,h2：无    tcp,kcp：none，http    quic：none，srtp，utp，wechat-video，dtls，wireguard
                host: '',   //伪装域名
                path: '',   //path
                mode: '',   //grpc传输模式 gun，multi

                security: '',//传输层安全类型 none,tls,reality
                sni: '',
                fp: '',
                alpn: '',
                allowInsecure: true,//tls 跳过证书验证
                pbk: '',
                sid: '',
                spx: '',
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
        async getAllNode(params?: object) {
            const res = await nodeApi.getAllNodeApi()
            if (res.code === 0) {
                ElMessage.success(res.msg)
                this.nodeManageData.nodes.node_list = res.data
            }
        },
        //获取全部节点 with Traffic,分页
        async getNodeWithTraffic(params?: object) {
            const res = await nodeApi.getNodeWithTrafficApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
                this.nodeManageData.nodes = res.data
            }
        },
        //获取节点 with Traffic(营收概览)
        async getNodeStatistics(params?: object) {
            const res = await nodeApi.getNodeWithTrafficApi(params)
            return res
        },
        //更新节点
        async updateNode(params?: object) {
            const res = await nodeApi.updateNodeApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //删除节点
        async deleteNode(params: object) {
            const res = await nodeApi.deleteNodeApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //新建节点
        async newNode(params?: object) {
            const res = await nodeApi.newNodeApi(params)
            if (res.code === 0) {
                ElMessage.success("新建节点成功")
            }
        },
    }
})