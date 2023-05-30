declare interface NodeInfo {
    created_at:string;
    updated_at:string;
    id: number;
    //基础参数
    name: string;
    address: string;
    port: string;
    sort: number;             //类型sort==11  V2Ray vmess
    nodespeed_limit: number;  //限速
    traffic_rate: number;     //倍率
    status:boolean;          //节点状态，true启用，flase禁用
    total_up:number;
    total_down:number;
    //vmess参数
    v: string;
    aid: string;
    scy:string;  //加密方式
    net: string; //传输协议 ws tcp kcp quic grpc等
    disguise_type: string;  //伪装类型 none http
    host: string;
    path: string;
    tls: string;
    sni:string;
}

