import {defineStore} from 'pinia';
//import Cookies from 'js-cookie';
import {Local, Session} from '/@/utils/storage';
//导入api
import {useUserApi} from "/@/api/user";

import {ElMessage} from "element-plus";
const userApi = useUserApi()
//server store
import {useServerStore} from "/@/stores/serverStore";
//publicStore
import {usePublicStore} from "/@/stores/publicStore";
const publicStore = usePublicStore()

export const useUserStore = defineStore('userInfo', {
    state: () => ({
        //登录页面数据
        loginData: {
            user_name: '',
            password: '',
            email_code: '',
            base64_captcha:{
                id:'',
                b64s:'',
            }
        } as LoginForm,
        //注册数据
        registerData: {
            user_name: '',
            email_suffix: '@qq.com',
            password: '',
            re_password: '',
            email_code: '',
            referrer_code: '',
            base64_captcha:{
                id:'',
                b64s:'',
            }
        } as RegisterForm,
        //全局用户信息
        userInfos: {
            created_at: '',
            updated_at: '',
            id: 0,
            uuid: 0,
            user_name: '',
            nick_name: '',
            password: '',
            avatar: '',
            phone: '',
            email: '',
            enable: true,
            invitation_code:'',
            referrer_code:'',
            remain:0,
            role_group: [] as RowRoleType[],	//角色组
            orders: [],      //订单
            subscribe_info: { //附加订阅信息
                expired_at: '',
                t: 0,
                u: 0,
                d: 0,
            }
        } as SysUser,
        //用户管理页面数据
        userManageData: {
            users: {
                total: 0,
                user_list: [] as SysUser[],
            },
            dialog: {
                user: {
                    user_name: '',
                    nick_name: '',
                    password: '123456',
                    avatar: '',
                    phone: '',
                    email: '',
                    enable: true,
                    role_group: [] as RowRoleType[],
                    subscribe_info: {
                        sub_status: true,
                        expired_at: '',
                        t: 0,
                        u: 0,
                        d: 0,
                        node_speedlimit: 0,
                        node_connector: 3,
                    }
                } as SysUser,
                check_list: ['普通用户'], //选中的角色
            },
        },
    }),
    getters: {
        used: (state): number => {
            return +((state.userInfos.subscribe_info.t - state.userInfos.subscribe_info.u - state.userInfos.subscribe_info.d) / 1024 / 1024 / 1024).toFixed(2)
        },
        expired: (state): string => {
            if (state.userInfos.subscribe_info.expired_at === null) {
                return ""
            } else {
                return state.userInfos.subscribe_info.expired_at.slice(0, 10)
            }
        },
        //订阅
        subUrl: (state): string => {
            const serverStore= useServerStore()
            return serverStore.publicServerConfig.sub_url_pre + "user/getSub?link=" + state.userInfos.subscribe_info.subscribe_url
        },

    },
    actions: {
        // 重置数据
        async resetData() {
            this.userManageData.dialog.user = {
                user_name: '',
                nick_name: '',
                password: '123456',
                avatar: '',
                phone: '',
                email: '',
                enable: true,
                role_group: [] as RowRoleType[],
                subscribe_info: {
                    sub_status: true,
                    expired_at: '',
                    t: 0,
                    u: 0,
                    d: 0,
                    node_speedlimit: 0,
                    node_connector: 3,
                }
            } as SysUser
            this.userManageData.dialog.check_list=['普通用户']
        },

        //注册
        async register(form?: object) {
            const referrerCode:string = Local.get('invitation')
            if (referrerCode.length === 8){
                this.registerData.referrer_code =referrerCode
            }
            this.registerData.base64_captcha.id=publicStore.base64CaptchaData.id
            // console.log("用户注册信息：",this.registerData)
            const res = await userApi.registerApi(this.registerData)
            return res
        },
        //登录
        async userLogin(form?: any) {
            // 存储用户信息到浏览器缓存
            const res: any = await userApi.signIn(form);
            //保存用户信息到pinia
            this.userInfos = res.data.user;
            //保存用户信息到Session
            Session.set("userInfos", res.data.user)
            //保存token
            // Session.set("token", res.data.token)
            Local.set("token", res.data.token)
        },
        //修改混淆
        async changeHost(params?: object) {
            const res = await userApi.changeHostApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
                await this.getUserInfo()
            }
        },
        //获取自身信息
        async getUserInfo() {
            const res = await userApi.getUserInfoApi()
            if (res.code === 0) {
                // ElMessage.success(res.msg)
                this.userInfos = res.data
                Session.set("userInfos", res.data)
            }
        },
        //获取用户列表
        async getUserList(params?: object) {
            const res = await userApi.getUserListApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
                this.userManageData.users = res.data
            }
        },
        //新建用户
        async newUser(params?: object) {
            const res = await userApi.newUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //修改用户
        async updateUser(params?: object) {
            const res = await userApi.updateUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //删除用户
        async deleteUser(params?: object) {
            const res = await userApi.deleteUserApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //修改密码
        async changePassword(params?: object) {
            const res = await userApi.changePasswordApi(params)
            if (res.code === 0) {
                ElMessage.success(res.msg)
            }
        },
        //确认重置密码
        async submitResetPassword() {
            return await userApi.resetPasswordApi(this.loginData)
        }
    },
});
