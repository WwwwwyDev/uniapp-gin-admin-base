<template>
    <scroll-view class="sidebar" scroll-y="true">
        <uni-data-menu :value="defaultUrl" :localdata="staticMenu" :unique-opened="true" active-text-color="#409eff"
            @select="select"></uni-data-menu>
    </scroll-view>
</template>

<script>
    import config from '@/admin.config.js'
    export default {
        data() {
            return {
                ...config.sideBar,
            }
        },
        mounted(){
            // #ifdef H5
            this.defaultUrl = window.location.hash.slice(1);
            // #endif
        },
        methods: {
            select(e) {
                let url = e.value
                this.clickMenuItem(url)
            },
            clickMenuItem(url) {
                // #ifdef H5
                if (url.indexOf('http') === 0) {
                	return window.open(url)
                }
                // #endif
                
                // url 开头可用有 / ，也可没有
                if (url[0] !== '/' && url.indexOf('http') !== 0) {
                	url = '/' + url
                }
                uni.redirectTo({
                	url: url
                })
            },
        }
    }
</script>

<style lang="scss">
    .sidebar {
        position: fixed;
        top: var(--top-window-height);
        width: 240px;
        height: calc(100vh - (var(--top-window-height)));
        box-sizing: border-box;
        border-right: 1px solid darken($left-window-bg-color, 8%);
        background-color: $left-window-bg-color;
        padding-bottom: 10px;
    }

    .title {
        margin-left: 5px;
    }
</style>
