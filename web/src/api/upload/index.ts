import request from '/@/utils/request';

export function useUploadApi() {
    return {
        //获取全部role
        newPictureUrlApi: (params?: object) => {
            return request({
                url: "/upload/newPictureUrl",
                method: "get",
                params,
            })
        },
        getPictureListApi: (data?: object) => {
            return request({
                url: "/upload/getPictureList",
                method: "post",
                data,
            })
        },
    }
}