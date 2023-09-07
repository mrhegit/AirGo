import service from "/@/utils/request";

export function useISPApi() {
    return {

        sendCodeApi: (data?: object) => {
            return service({
                url: '/isp/sendCode',
                method: 'POST',
                data
            })
        },
        ispLoginApi: (data?: object) => {
            return service({
                url: '/isp/ispLogin',
                method: 'POST',
                data
            })
        },
        queryPackageApi: (data?: object) => {
            return service({
                url: '/isp/queryPackage',
                method: 'POST',
                data
            })
        },
        getMonitorByUserIDApi: (data?: object) => {
            return service({
                url: '/isp/getMonitorByUserID',
                method: 'POST',
                data
            })
        },

    }

}