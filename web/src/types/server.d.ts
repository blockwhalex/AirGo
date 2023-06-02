declare interface Server{
    created_at:string;
    updated_at:string;
    id: number;
    jwt:Jwt;
    system:System;
    captcha:Captcha;
    pay:Pay;
    email:Email;
}
declare interface Jwt{
    signing_key:string;
    expires_time:string;
    buffer_time:string;
    issuer:string;
}
declare interface System{
    enable_register:Boolean;
    enable_email_code:Boolean;
    is_multipoint:boolean;
    sub_name:string;
    muKey:string;
    default_goods:string;
}
declare interface Captcha{
    key_long:number;
    img_width:number;
    img_height:number;
    open_captcha:number;
    open_captcha_time_out:number;
}
declare interface Pay{
    return_url:string;
    app_id:string;
    private_key:string;
    ali_public_key:string;
    encrypt_key:string;
}
declare interface Email{
    email_from:string;
    email_secret:string;
    email_host:string;
    email_port:number;
    email_is_ssl:Boolean;
    email_nickname:string;
    email_subject:string;
    email_content:string;
}