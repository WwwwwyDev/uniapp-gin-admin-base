export default {
    login: {
        url: '/pages/login/login' // 登录页面路径
    },
    index: {
        url: '/pages/index/index' // 登录后跳转的第一个页面
    },
    error: {
        url: '/pages/error/404' // 404 Not Found 错误页面路径
    },
    showIfm: {
        url: '/windows/pages/showIfm' // 展示用户信息界面
    },
    navBar: { // 顶部导航
        logo: '/static/logo.png', // 左侧 Logo
    },
    sideBar: { // 左侧菜单
        // 配置静态菜单列表
        defaultUrl: '/pages/index/index',
        staticMenu: [{
            menu_id: "index",
            text: '首页',
            icon: '',
            value: '/pages/index/index',
        }, {
            menu_id: "demo",
            text: '静态功能演示',
            icon: 'uni-icons-list',
            value: "",
            children: [{
                menu_id: "test",
                text: '测试',
                icon: '',
                value: '/pages/test/test',
            }]
        }]
    }
}
