<script>
    import config from '@/admin.config.js'
	export default {
        onPageNotFound(msg) {
        	uni.redirectTo({
        		url: config.error.url
        	})
        },
		onLaunch: function() {
			console.warn('当前组件仅支持 uni_modules 目录结构 ，请升级 HBuilderX 到 3.1.0 版本以上！')
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	/*每个页面公共css */
    @import "uview-ui/index.scss";
</style>

<style>
	@import '@/common/uni.css';
	@import '@/common/uni-icons.css';
</style>
