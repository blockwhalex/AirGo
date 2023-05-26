
declare interface SysUser {
	created_at:string;
	updated_at:string;
	id: number;
	uuid: number;
	user_name: string;
	nick_name: string;
	password:string;
	avatar: string;
	phone: string;
	email: string;
	enable: boolean;
	user_role: [];	//角色组
	orders: [];      //订单
	subscribe_info: { //附加订阅信息
		host: string;
		client_ip: string;
		sub_status: boolean;
		subscribe_url: string;
		goods_id: int;
		expired_at: string;
		t: number;
		u: number;
		d: number;
		reset_day: number;
		node_speedlimit: number;
		node_connector: number;
	}
}


