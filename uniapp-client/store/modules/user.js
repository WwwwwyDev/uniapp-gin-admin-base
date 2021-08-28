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
            let res = await this.$http.post('/api/v1/c/jwt');
            if(res.code == -1){
                return null
            }
            if(res.code == 20000){
                return res
            }
        }
    }
}
