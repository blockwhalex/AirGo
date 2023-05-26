declare interface Goods {
    created_at:string;
    updated_at:string;
    id: number;
    subject: string; // 订单标题
    total_amount: number;  // 订单总金额，
    product_code: string;
    total_bandwidth: number;//总流量
    expiration_date: number;//有效期
    checked_nodes:number[]; //套餐编辑时选中的节点
    nodes:[];
}