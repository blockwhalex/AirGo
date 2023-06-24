import service from "/@/utils/request";

export function useArticleApi() {
    return {
        newArticleApi: (data?: object) => {
            return service({
                url: '/article/newArticle',
                method: 'POST',
                data
            })
        },
        deleteArticleApi: (data?: object) => {
            return service({
                url: '/article/deleteArticle',
                method: 'POST',
                data
            })
        },
        updaterticleApi: (data?: object) => {
            return service({
                url: '/article/updaterticle',
                method: 'POST',
                data
            })
        },
        getArticleApi: (data?: object) => {
            return service({
                url: '/article/getArticle',
                method: 'POST',
                data
            })
        },
    }

}