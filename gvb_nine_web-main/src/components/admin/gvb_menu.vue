<template>
  <div class="gvb_menu">
    <a-menu
        @menu-item-click="clickMenu"
        v-model:selected-keys="selectedKeys"
        v-model:open-keys="openKeys"
        show-collapse-button
        @collapse="collapse"
    >
      <template v-for="item in menuList" :key="item.name">
        <a-menu-item :key="item.name" v-if="item.child?.length === 0">
          {{ item.title }}
          <template #icon>
            <component :is="item.icon"></component>
          </template>
        </a-menu-item>
        <a-sub-menu v-if="item.child?.length!==0 " :key="item.name">
          <template #icon>
            <component :is="item.icon"></component>
          </template>
          <template #title>{{ item.title }}</template>
          <a-menu-item :key="sub.name" v-for="sub in item.child">
            {{ sub.title }}
            <template #icon>
              <component :is="sub.icon"></component>
            </template>
          </a-menu-item>
        </a-sub-menu>

      </template>
    </a-menu>
  </div>
</template>

<script setup lang="ts">
import {defineComponent, h, ref, watch} from "vue";
import type {Component} from "vue";
import {useRoute, useRouter} from "vue-router";
import {useStore} from "@/stores";
import {IconMenu, IconUser, IconSettings, IconMessage, IconUserGroup} from "@arco-design/web-vue/es/icon";
import {IconBook, IconHome, IconStorage, IconFile, IconShareAlt} from "@arco-design/web-vue/es/icon";
import {IconImage} from "@arco-design/web-vue/es/icon";

const store = useStore()
const route = useRoute()
const router = useRouter()

function getFontComponent(name: string): Component {
  return defineComponent({
    render: () => {
      return h("i", {class: name})
    }
  })
}


interface MenuType {
  title: string
  icon?: Component
  name?: string // 路由名字
  child?: MenuType[]
}


let menuList: MenuType[] = [
  {title: "首页", icon: IconHome, name: "home", child: []},
  {
    title: "个人中心", icon: IconUser, name: "user_center", child: [
      {title: "我的信息", icon: getFontComponent("fa fa-vcard"), name: "user_info"},
      {title: "我的发布", icon: getFontComponent("fa fa-book"), name: "user_add_article"},
      {title: "我的收藏", icon: getFontComponent("fa fa-star"), name: "collects"},
      {title: "我的消息", icon: IconMessage, name: "messages"},
    ]
  },
  {
    title: "文章管理", icon: IconBook, name: "article_mgr", child: [
      {title: "文章列表", icon: IconBook, name: "article_list"},
      {title: "图片列表", icon: IconImage, name: "image_list"},
      {title: "评论列表", icon: getFontComponent("fa fa-comments"), name: "comment_list"},
    ]
  },
  {
    title: "用户管理", icon: IconUserGroup, name: "users", child: [
      {title: "用户列表", icon: IconUserGroup, name: "user_list"},
      {title: "消息列表", icon: IconMessage, name: "message_list"},
    ]
  },
  {
    title: "群聊管理", icon: IconMessage, name: "chat_group", child: [
      {title: "聊天记录", icon: IconMessage, name: "chat_list"},
    ]
  },
  {
    title: "大模型管理", icon: IconMessage, name: "big_model", child: [
      {title: "参数配置", icon: IconMessage, name: "options"},
      {title: "角色配置", icon: IconMessage, name: "chat_role"},
      {title: "会话管理", icon: IconMessage, name: "session"},
    ]
  },
  {
    title: "系统管理", icon: IconSettings, name: "system", child: [
      {title: "菜单列表", icon: IconMenu, name: "menu_list"},
      {title: "用户反馈", icon: IconMenu, name: "feedback_list"},
      {title: "广告列表", icon: IconShareAlt, name: "promotion_list"},
      {title: "系统日志", icon: IconFile, name: "log_list"},
      {title: "系统配置", icon: IconStorage, name: "system_system"},
    ]
  },
]

if (store.isGeneral) {
  menuList = [
    {title: "首页", icon: IconHome, name: "home", child: []},
    {
      title: "个人中心", icon: IconUser, name: "user_center", child: [
        {title: "我的信息", icon: getFontComponent("fa fa-vcard"), name: "user_info"},
        {title: "我的收藏", icon: getFontComponent("fa fa-star"), name: "collects"},
        {title: "我的消息", icon: IconMessage, name: "messages"},
      ]
    },
  ]
}


const selectedKeys = ref([route.name])
const openKeys = ref([route.matched[1].name])

function collapse(collapsed: boolean) {
  store.setCollapsed(collapsed)
}


function clickMenu(name: string) {
  router.push({
    name: name,
  })
}

watch(() => route.name, () => {
  selectedKeys.value = [route.name]
  openKeys.value = [route.matched[1].name]
})

</script>


<style lang="scss">
.gvb_menu {
  .arco-menu {
    position: inherit;
  }

  .arco-menu-collapse-button {
    position: absolute;
    right: 0;
    top: 50%;
    transform: translate(50%, -50%);
    opacity: 0;
    transition: all .3s;
  }
}

aside:hover {
  .gvb_menu {
    .arco-menu-collapse-button {
      opacity: 1;
    }
  }
}
</style>