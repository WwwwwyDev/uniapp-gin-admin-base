<template>
    <view class="noData">
        <view class="image">
            <image src="../../static/notFound.png"></image>
        </view>
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
    uni-page-body {
        width: 100%;
        height: 100%;
    }
</style>

<style lang="less" scoped>
    .noData {
        width: 100%;
        height: 100%;
        display: -webkit-flex;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: row;

        .image {
            width: 187upx;
            height: 144upx;
            image {
                width: 100%;
                height: 100%;
            }
        }
    }
</style>
