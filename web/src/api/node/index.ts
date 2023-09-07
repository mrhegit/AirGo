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
        //共享节点api
        newNodeSharedApi: (data?: object) => {
            return service({
                url: '/node/newNodeShared',
                method: 'post',
                data
            })
        },
        getNodeSharedListApi: (params?: object) => {
            return service({
                url: '/node/getNodeSharedList',
                method: 'get',
                params
            })
        },
        deleteNodeSharedApi: (data?: object) => {
            return service({
                url: '/node/deleteNodeShared',
                method: 'post',
                data
            })
        },
    }
}