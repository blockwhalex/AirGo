import service from "/@/utils/request";

export function useNodeApi() {
    return {
        getAllNodeApi: () => {
            return service({
                url: '/node/getAllNode',
                method: 'get',
            })

        },
        getNodeWithTrafficApi: (data?: object) => {
            return service({
                url: '/node/getTraffic',
                method: 'post',
                data
            })

        },
        newNodeApi: (data?: object) => {
            return service({
                url: '/node/newNode',
                method: 'post',
                data
            })

        },
        updateNodeApi: (data?: object) => {
            return service({
                url: '/node/updateNode',
                method: 'post',
                data
            })
        },
        deleteNodeApi: (data?: object) => {
            return service({
                url: '/node/deleteNode',
                method: 'post',
                data
            })
        },
        //
        nodeSortApi: (data?: object) => {
            return service({
                url: '/node/nodeSort',
                method: 'post',
                data
            })
        },
    }
}