<template>
    <view>
        <uni-list>
            <uni-list-chat :avatar-circle="true" :avatar="userInfo.headerImg"></uni-list-chat>
            <uni-list-item title="用户名" :rightText="userInfo.username"></uni-list-item>
            <uni-list-item title="昵称" :rightText="userInfo.nickName"></uni-list-item>
            <uni-list-item title="联系电话" :rightText="userInfo.phone"></uni-list-item>
            <uni-list-item title="邮箱" :rightText="userInfo.email"></uni-list-item>
        </uni-list>
        <view style="padding: 20upx 100upx 20upx 100upx;">
            <button type="primary">修改</button>
        </view>
    </view>
</template>

<script>
    import {
        mapActions
    } from 'vuex'
    export default {
        data() {
            return {
                userInfo: {
                    username: '',
                    nickName: '',
                    email: '',
                    phone: '',
                    headerImg: '',
                },
            }
        },
        async mounted() {
            let res = await this.getUserInfo();
            if (res === null) {
                this.isValidSession()
                return;
            }
            this.userInfo = res;
        },
        methods: {
            ...mapActions({
                isValidSession: 'user/isValidSession',
                getUserInfo: 'user/getUserInfo'
            })
        }
    }
</script>

<style>
</style>
