import {createRouter, createWebHistory} from 'vue-router'
import type {RouteMeta} from 'vue-router'
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";
import NProgress from "nprogress";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'web',
            component: () => import('../views/web/web.vue'),
            children: [
                {
                    path: "",
                    name: "index",
                    component: () => import('../views/web/index.vue'),
                },
                {
                    path: "news",
                    name: "news",
                    component: () => import('../views/web/news.vue'),
                },
                {
                    path: "search",
                    name: "search",
                    component: () => import('../views/web/search.vue'),
                },
                {
                    path: "chat",
                    name: "chat",
                    component: () => import('../views/web/chat.vue'),
                },
                {
                    path: "bigModel",
                    name: "bigModel",
                    component: () => import('../views/web/big_model/index.vue'),
                    redirect: "/bigModel/square",
                    children:[
                        {
                            path: "square",
                            name: "role_square",
                            component: () => import('../views/web/big_model/role_square.vue'),
                        },
                        {
                            path: "session",
                            name: "role_session",
                            component: () => import('../views/web/big_model/session.vue'),
                            meta: {
                                isLogin: true,
                            },
                        }
                    ]
                },
                {
                    path: "article/404",
                    name: "article_notfound",
                    component: () => import('../views/web/article_notfound.vue'),
                },
                {
                    path: "article/:id",
                    name: "article",
                    component: () => import('../views/web/article.vue'),
                },
            ]
        },
        {
            path: "/login",
            name: "login",
            component: () => import('../views/login/index.vue'),
        },
        {
            path: "/admin",
            name: "admin",
            component: () => import('../views/admin/index.vue'),
            meta: {
                title: "首页",
                isLogin: true,
            },
            children: [
                {
                    path: "",
                    name: "home",
                    meta: {
                        title: "Home"
                    },
                    component: () => import('../views/admin/home/index.vue'),
                },
                {
                    path: "user_center",
                    name: "user_center",
                    meta: {
                        title: "个人中心"
                    },
                    children: [
                        {
                            path: "user_info",
                            name: "user_info",
                            meta: {
                                title: "我的信息"
                            },
                            component: () => import('../views/admin/user_center/user_info.vue'),
                        },
                        {
                            path: "article",
                            name: "user_add_article",
                            meta: {
                                title: "我的发布"
                            },
                            component: () => import('../views/admin/user_center/user_add_article_list.vue'),
                        },
                        {
                            path: "collects",
                            name: "collects",
                            meta: {
                                title: "我的收藏"
                            },
                            component: () => import('../views/admin/user_center/user_collects_list.vue'),
                        },
                        {
                            path: "messages",
                            name: "messages",
                            meta: {
                                title: "我的消息"
                            },
                            component: () => import('../views/admin/user_center/user_message_list.vue'),
                        }
                    ]
                },
                {
                    path: "article",
                    name: "article_mgr",
                    meta: {
                        title: "文章管理",
                        isAdmin: true,
                        isTourist: true,
                    },
                    children: [
                        {
                            path: "article_list",
                            name: "article_list",
                            meta: {
                                title: "文章列表"
                            },
                            component: () => import('../views/admin/article/article_list.vue'),
                        },
                        {
                            path: "image_list",
                            name: "image_list",
                            meta: {
                                title: "图片列表"
                            },
                            component: () => import('../views/admin/article/image_list.vue'),
                        },
                        {
                            path: "comment_list",
                            name: "comment_list",
                            meta: {
                                title: "评论列表"
                            },
                            component: () => import('../views/admin/article/comment_list.vue'),
                        }
                    ]
                },
                {
                    path: "users",
                    name: "users",
                    meta: {
                        title: "用户管理",
                        isAdmin: true,
                        isTourist: true,
                    },
                    children: [
                        {
                            path: "user_list",
                            name: "user_list",
                            meta: {
                                title: "用户列表"
                            },
                            component: () => import('../views/admin/users/user_list.vue'),
                        },
                        {
                            path: "message_list",
                            name: "message_list",
                            meta: {
                                title: "消息列表"
                            },
                            component: () => import('../views/admin/users/message_list.vue'),
                        }
                    ]
                },
                {
                    path: "chat_group",
                    name: "chat_group",
                    meta: {
                        title: "群聊管理",
                        isAdmin: true,
                        isTourist: true,
                    },
                    children: [
                        {
                            path: "chat_list",
                            name: "chat_list",
                            meta: {
                                title: "聊天记录"
                            },
                            component: () => import('../views/admin/chat_group/chat_list.vue'),
                        }
                    ]
                },
                {
                    path: "big_model",
                    name: "big_model",
                    meta: {
                        title: "大模型管理",
                        isAdmin: true,
                        isTourist: false,
                    },
                    component: () => import('../views/admin/big_model/index.vue'),
                    children: [
                        {
                            path: "options",
                            name: "options",
                            meta: {
                                title: "参数配置"
                            },
                            component: () => import('../views/admin/big_model/options/index.vue'),
                            redirect: "/admin/big_model/options/base",
                            children:[
                                {
                                    path: "base",
                                    name: "base_option",
                                    meta: {
                                        title: "基本配置"
                                    },
                                    component: () => import('../views/admin/big_model/options/base_option.vue'),
                                },
                                {
                                    path: "session",
                                    name: "session_option",
                                    meta: {
                                        title: "会话配置"
                                    },
                                    component: () => import('../views/admin/big_model/options/session_option.vue'),
                                },
                                {
                                    path: "reply",
                                    name: "reply_option",
                                    meta: {
                                        title: "自动回复"
                                    },
                                    component: () => import('../views/admin/big_model/options/reply_option.vue'),
                                }
                            ]
                        },
                        {
                            path: "chat_role",
                            name: "chat_role",
                            meta: {
                                title: "角色配置"
                            },
                            component: () => import('../views/admin/big_model/chat_role/index.vue'),
                            redirect: "/admin/big_model/chat_role/role_list",
                            children:[
                                {
                                    path: "role_tag",
                                    name: "role_tag",
                                    meta: {
                                        title: "角色标签"
                                    },
                                    component: () => import('../views/admin/big_model/chat_role/role_tag.vue'),
                                },
                                {
                                    path: "role_list",
                                    name: "role_list",
                                    meta: {
                                        title: "角色列表"
                                    },
                                    component: () => import('../views/admin/big_model/chat_role/role_list.vue'),
                                }
                            ]
                        },
                        {
                            path: "session",
                            name: "session",
                            meta: {
                                title: "会话管理"
                            },
                            component: () => import('../views/admin/big_model/session/index.vue'),
                        }
                    ]
                },
                {
                    path: "system",
                    name: "system",
                    meta: {
                        title: "系统管理",
                        isAdmin: true,
                        isTourist: false,
                    },
                    children: [
                        {
                            path: "menu_list",
                            name: "menu_list",
                            meta: {
                                title: "菜单列表"
                            },
                            component: () => import('../views/admin/system/menu_list.vue'),
                        },
                        {
                            path: "feedback_list",
                            name: "feedback_list",
                            meta: {
                                title: "用户反馈"
                            },
                            component: () => import('../views/admin/system/feedback_list.vue'),
                        },
                        {
                            path: "promotion_list",
                            name: "promotion_list",
                            meta: {
                                title: "广告列表"
                            },
                            component: () => import('../views/admin/system/promotion_list.vue'),
                        },
                        {
                            path: "log_list",
                            name: "log_list",
                            meta: {
                                title: "系统日志"
                            },
                            component: () => import('../views/admin/system/log_list.vue'),
                        },
                        {
                            path: "system",
                            name: "system_system",
                            meta: {
                                title: "系统配置"
                            },
                            redirect: "/admin/system/system/site",
                            component: () => import('../views/admin/system/system.vue'),
                            children: [
                                {
                                    path: "site",
                                    name: "site_config",
                                    meta: {
                                        title: "网站设置"
                                    },
                                    component: () => import('../views/admin/system/config/site_config.vue'),
                                },
                                {
                                    path: "email",
                                    name: "email_config",
                                    meta: {
                                        title: "邮箱设置"
                                    },
                                    component: () => import('../views/admin/system/config/email_config.vue'),
                                },
                                {
                                    path: "qiniu",
                                    name: "qiniu_config",
                                    meta: {
                                        title: "七牛云设置"
                                    },
                                    component: () => import('../views/admin/system/config/qiniu_config.vue'),
                                },
                                {
                                    path: "qq",
                                    name: "qq_config",
                                    meta: {
                                        title: "QQ设置"
                                    },
                                    component: () => import('../views/admin/system/config/qq_config.vue'),
                                },
                                {
                                    path: "jwt",
                                    name: "jwt_config",
                                    meta: {
                                        title: "jwt设置"
                                    },
                                    component: () => import('../views/admin/system/config/jwt_config.vue'),
                                },
                                {
                                    path: "gaode",
                                    name: "gaode_config",
                                    meta: {
                                        title: "高德设置"
                                    },
                                    component: () => import('../views/admin/system/config/gaode_config.vue'),
                                }
                            ]
                        }
                    ]
                },
            ]
        },
        {
            path: '/:pathMatch(.*)*', // 页面不存在的情况下会跳到404页面
            name: 'notFound',
            component: () => import("@/views/web/notfound.vue")
        }
    ]
})

export default router

router.beforeEach((to, from, next) => {
    const store = useStore()
    const meta = to.meta
    if (meta.isLogin && !store.isLogin) {
        // 需要登录，但是我没有登录
        Message.warning("需要登录")
        router.push({name: from.name as string})
        return
    }
    /*
    isLogin 登录了就能看
    isAdmin  管理员能看
    isTourist  游客能看，管理员也能看
     */

    // 如果我是普通用户，那我就不能访问admin或者游客权限的页面
    if (store.userInfo.role === 2 && (meta.isAdmin || meta.isTourist)) {
        Message.warning("权限不足1")
        router.push({name: from.name as string})
        return
    }

    // 如果我是游客，那我就不能访问游客权限为false的
    if (store.isTourist && meta.isTourist === false) {
        Message.warning("权限不足2")
        router.push({name: from.name as string})
        return
    }


    // if (meta.isAdmin){
    //     // 需要管理员权限
    //     // 管理员也能看，游客也能看
    //     if (meta.isTourist && !store.isTourist && !store.isAdmin) {
    //         // 身份不是游客
    //         Message.warning("权限不足2")
    //         router.push({name: from.name as string})
    //         return
    //     }
    //     if (!meta.isTourist && !store.isAdmin){
    //         // 游客不能看，校验他不是管理员
    //         Message.warning("权限不足1")
    //         router.push({name: from.name as string})
    //         return
    //     }
    // }
    NProgress.start()
    next()
})


router.afterEach(() => {
    NProgress.done()//完成进度条
})