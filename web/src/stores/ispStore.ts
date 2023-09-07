import {defineStore} from "pinia";
import {useISPApi} from "/@/api/isp";
import {ElMessage} from "element-plus";

const ispApi = useISPApi()
export const useISPStore = defineStore("ispStore", {
    state: () => ({
        isp:{
            user_id: 0,
            isp_type: '',
            status: false,
            mobile: '',
            unicom_config: {
                version: '',
                app_id: '',
                cookie: '',
                unicomMobile: '',
                password: '',
            },
            telecom_config:{
                phoneNum: '',
                telecomPassword: '',
                timestamp: '',
                loginAuthCipherAsymmertric: '',
                deviceUid: '',
                telecomToken: '',
            },

        } as Isp,
        isCountDown: false,
        countDownTime: 60,
    }),
    actions: {
        async sendCode(params?: object) {
            const res = await ispApi.sendCodeApi(params)
            ElMessage.success(res.msg)

        },
        async ispLogin(params?: object) {
            const res = await ispApi.ispLoginApi(params)
            if (res.msg==='获取成功'){
                this.isp=res.data
                return
            }
            this.getMonitorByUserID()
        },
        async queryPackage(params?: object) {
            const res = await ispApi.queryPackageApi(params)
            ElMessage.success(res.msg)

        },
        async getMonitorByUserID(params?: object) {
            const res = await ispApi.getMonitorByUserIDApi(params)
            ElMessage.success(res.msg)
            this.isp=res.data

        },
    }
})