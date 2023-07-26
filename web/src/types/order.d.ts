declare interface Order {
    created_at: string;
    updated_at: string;
    id: number;
    userID: number;
    user_name: string;
    user: any;

    out_trade_no: string;
    goods_id: number;
    subject: string;
    price: string;
    pay_type: string;
    coupon: number;
    coupon_name: string;
    coupon_amount: string;
    deduction_amount: string;
    remain_amount: string;

    qr_code: string;
    trade_no: string;
    buyer_logon_id: string;
    trade_status: string;
    total_amount: string;
    receipt_amount: string;
    buyer_pay_amount: string;
}

declare interface OrdersWithTotal {
    total_amount: number;
    total: number;
}