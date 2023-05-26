import service from "/@/utils/request";

export function useNodeApi(){
    return {
        getAllNodeApi:(data?:object)=>{
            return service({
                url: '/node/getTraffic',
				method: 'post',
                data
            })

        },
        newNodeApi:(data?:object)=>{
            return service({
                url: '/node/newNode',
				method: 'post',
                data
            })

        },
        updateNodeApi:(data?:object)=>{
            return service({
                url: '/node/updateNode',
				method: 'post',
                data
            })
        },
        deleteNodeApi:(data?:object)=>{
            return service({
                url: '/node/deleteNode',
				method: 'post',
                data
            })
        },
        getNodeByNameApi:(data?:object)=>{
            return service({
                url: '/node/getNodeByName',
				method: 'post',
                data
            })
        },
        getNodeTrafficApi:(data?:object)=>{
            return service({
                url: '/node/getTraffic',
                method: 'post',
                data
            })
        }
    }
}