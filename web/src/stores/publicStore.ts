import {defineStore} from 'pinia';

import {ElMessage} from "element-plus";
import {usePublicApi} from "/@/api/public";
const publicApi = usePublicApi()

export const usePublicStore = defineStore('publicStore', {
    state: () => ({
        base64CaptchaData:{
            id: '',
            b64s:'',

        } as Base64CaptchaInfo,

    }),
    actions: {
        //获取base64Captcha
        async getBase64Captcha(){
            const res = await publicApi.getBase64CaptchaApi()
            this.base64CaptchaData=res.data
        }

    }

})