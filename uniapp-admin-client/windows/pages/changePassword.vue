<template>
    <view>
        <view class="uni-header">
            <view class="uni-group">
                <view class="uni-title">修改密码</view>
            </view>
        </view>
        <view class="uni-container">
            <uni-forms ref="resetPwdForm" v-model="resetPwdFormData" :rules="resetPwdrules" @submit="submitPwd">
                <uni-forms-item label="用户名" name="oldPassword" labelWidth="85">
                	<input class="uni-input-border" disabled="true" v-model="userInfo.username" />
                </uni-forms-item>
                <uni-forms-item label="旧密码" name="oldPassword" labelWidth="85">
                    <input class="uni-input-border" :password="showOldPassword" placeholder="旧密码"
                        v-model="resetPwdFormData.oldPassword" />
                    <text class="uni-icon-password-eye pointer" :class="[!showOldPassword ? 'uni-eye-active' : '']"
                        @click="showOldPassword=!showOldPassword">&#xe568;</text>
                </uni-forms-item>
                <uni-forms-item label="新密码" name="newPassword" labelWidth="85">
                    <input class="uni-input-border" :password="showNewPassword" placeholder="新密码"
                        v-model="resetPwdFormData.newPassword" />
                    <text class="uni-icon-password-eye pointer" :class="[!showNewPassword ? 'uni-eye-active' : '']"
                        @click="showNewPassword=!showNewPassword">&#xe568;</text>
                </uni-forms-item>
                <uni-forms-item label="确认新密码" name="newPasswordAgain" labelWidth="85">
                    <input @confirm="updatePwd" class="uni-input-border" :password="showNewPasswordAgain"
                        placeholder="确认新密码" v-model="resetPwdFormData.newPasswordAgain" />
                    <text class="uni-icon-password-eye pointer" :class="[!showNewPasswordAgain ? 'uni-eye-active' : '']"
                        @click="showNewPasswordAgain=!showNewPasswordAgain">&#xe568;</text>
                </uni-forms-item>
                <view class="uni-button-group pointer">
                    <button class="uni-button uni-button-full" type="primary" @click="updatePwd">修改</button>
                </view>
            </uni-forms>
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
                showNewPassword: true,
                showNewPasswordAgain: true,
                showOldPassword: true,
                userInfo: {
                    username: '',
                    nickName: '',
                    email: '',
                    phone: '',
                    headerImg: '',
                },
                resetPwdFormData:{
                    username: '',
                    oldPassword: '',
                    newPassword: '',
                    newPasswordAgain: ''
                },
                resetPwdrules: {
                    // 对password字段进行必填验证
                    oldPassword: {
                        rules: [{
                                required: true,
                                errorMessage: '请输入正确的密码',
                            },
                            {
                                pattern: '^.{6,18}$', //正则校验(与后端匹配)
                                minLength: 6,
                                maxLength: 16,
                                errorMessage: '密码长度在{minLength}到{maxLength}个字符',
                            }
                        ]
                    },
                    newPassword: {
                        rules: [{
                                required: true,
                                errorMessage: '请输入正确的密码',
                            },
                            {
                                pattern: '^.{6,18}$', //正则校验(与后端匹配)
                                minLength: 6,
                                maxLength: 16,
                                errorMessage: '密码长度在{minLength}到{maxLength}个字符',
                            }
                        ]
                    },
                    newPasswordAgain: {
                        rules: [{
                                required: true,
                                errorMessage: '请输入正确的密码',
                            },
                            {
                                pattern: '^.{6,18}$', //正则校验(与后端匹配)
                                minLength: 6,
                                maxLength: 16,
                                errorMessage: '密码长度在{minLength}到{maxLength}个字符',
                            }
                        ]
                    },
                }
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
            async submitPwd() {
                if (this.resetPwdFormData.newPassword != this.resetPwdFormData.newPasswordAgain) {
                    uni.showToast({
                        title: " 两次输入密码不一致",
                        icon: "none"
                    })
                    return;
                }
                if (this.resetPwdFormData.oldPassword == this.resetPwdFormData.newPasswordAgain) {
                    uni.showToast({
                        title: " 新旧密码不能相同",
                        icon: "none"
                    })
                    return;
                }
                this.resetPwdFormData.username = this.userInfo.username;
                let res = await this.$http.post('/api/v1/sysUser/changePassword', this.resetPwdFormData);
                if (res.code == -1) {
                    uni.showToast({
                        title: "后台出错",
                        icon: "none"
                    })
                }
                if (res.code == 50002) {
                    uni.showToast({
                        title: res.msg,
                        icon: "none"
                    })
                }
                if (res.code == 20000) {
                    uni.showToast({
                        title: res.msg,
                        icon: "none"
                    })
                    await setTimeout(function() {
                        uni.navigateBack();
                    }, 600);
                }
            },
            updatePwd() {
                this.$refs.resetPwdForm.submit();
            },
            ...mapActions({
                isValidSession: 'user/isValidSession',
                getUserInfo: 'user/getUserInfo'
            })
        }
    }
</script>

<style>

</style>
