<template>
    <view>
        首页
        <!-- #ifndef H5 -->
        <fix-window />
        <!-- #endif -->
    </view>
</template>

<script>
    import {
        mapActions
    } from 'vuex'
    import config from '@/admin.config.js'
    export default {
        data() {
            return {

            }
        },
        mounted() {
            this.isValidSession()
        },
        methods: {
            ...mapActions({
                isValidSession: 'user/isValidSession',
            })
        }
    }
</script>

<style>
    /* #ifndef H5 */
    page {
        padding-top: 85px;
    }

    /* #endif */
</style>
