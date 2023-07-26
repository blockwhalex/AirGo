declare interface Goods {
    id: number;
    created_at: string;
    updated_at: string;

    good_order: number;    //排序
    status: boolean; //是否启用
    des: string;     //商品描述
    subject: string;       // 标题
    total_amount: string;  // 总金额
    product_code: string;
    total_bandwidth: number; //总流量
    expiration_date: number; //有效期
    checked_nodes: number[]; //套餐编辑时选中的节点
    nodes: Node[];

}