import service from "/@/utils/request";

export function usePublicApi() {
    return {
        //获取邮箱验证码
        getEmailCodeApi: (data?: object) => {
            return service({
                url: '/public/getEmailCode',
                method: 'post',
                data
            })
        },
        //获取base64Captcha
        getBase64CaptchaApi: (params?: object) => {
            return service({
                url: '/public/getBase64Captcha',
                method: 'get',
                params
            })
        },
    }
}