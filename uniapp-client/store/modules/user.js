import $http from '@/uni_modules/zhouWei-request/js_sdk/requestConfig';
import config from '@/admin.config.js'
export default {
    namespaced: true,
    state: {
        token: uni.getStorageSync("j-token"),
    },
    getters: {
        getToken: (state) => {
            return state.token
        },
    },
    mutations: {
        setToken: (state, token) => {
            state.token = token
        },
    },
    actions: {
        async getUserInfo() {
            let res = await $http.post('/api/v1/c/jwt');
            if(res.code != 20000){
                return null
            }
            return res.data
        },
        async isValidToken() {
            if (uni.getStorageSync("j-token")===""){return false}
            let res = await $http.post('/api/v1/c/jwt');
            if(res.code != 20000){
                return false
            }
            return true
        },
        async isValidSession() {
            if (uni.getStorageSync("j-token")===""){
                uni.showToast({
                    title: "未登录,请登录后访问",
                    icon: "none"
                })
                await setTimeout(function() {
                    uni.redirectTo({
                        url: config.login.url
                    })
                }, 600);
                return;
            }
            let f = await this.dispatch('user/isValidToken')
            if (!f) {
                uni.removeStorageSync("j-token")
                uni.showToast({
                    title: "会话过期,请重新登录",
                    icon: "none"
                })
                await setTimeout(function() {
                    uni.redirectTo({
                        url: config.login.url
                    })
                }, 600);
                return;
            }
        },
        async delToken() {
            let res = await $http.delete('/api/v1/c/jwt');
            if(res.code != 20000){
                return false
            }
            uni.removeStorageSync("j-token")
            return true
        }
    }
}
