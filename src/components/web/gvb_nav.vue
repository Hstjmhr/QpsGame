<template>
  <div :class="{gvb_nav: true, isShow: isShow}">
    <div class="gvb_nav_container">
      <div class="left">
        <a class="logo" href="/">
          <div class="slogan">{{ store.siteInfo.slogan }}</div>
          <div class="en_slogan">{{ store.siteInfo.slogan_en }}</div>
        </a>
        <div class="nav">
          <template v-for="item in navList">
            <a target="_blank" :href="item.path" v-if="item.path.startsWith('http')"
               :title="item.title">{{ item.title }}</a>
            <router-link :to="item.path" v-else>{{ item.title }}</router-link>
          </template>
        </div>
        <gvb_search></gvb_search>
      </div>
      <div class="right">
        <div class="login" v-if="!store.isLogin">
          <router-link :to="{name: 'login'}">登录</router-link>
        </div>
        <div class="user_info" v-else>
          <gvb_user_info_menu></gvb_user_info_menu>
        </div>
        <div class="theme">
          <gvb_theme/>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {IconSearch} from "@arco-design/web-vue/es/icon";
import Gvb_theme from "@/components/common/gvb_theme.vue";
import {useStore} from "@/stores";
import Gvb_user_info_menu from "@/components/common/gvb_user_info_menu.vue";
import {menuNameListApi} from "@/api/menu_api";
import type {menuNameType} from "@/api/menu_api";
import {reactive, ref} from "vue";
import {onUnmounted} from "vue";
import Gvb_search from "@/components/web/gvb_search.vue";
import {bigModelSettingsApi} from "@/api/big_model_api";
import type {bigModelSettingsType} from "@/api/big_model_api";

interface Props {
  noScroll?: boolean // 不需要滚动监听
}

const props = defineProps<Props>()
const {noScroll = false} = props

const store = useStore()

const navList = ref<menuNameType[]>([])

const isShow = ref(true)

if (!noScroll) {
  isShow.value = false
  window.addEventListener("scroll", scroll)
  scroll()
}


function scroll() {
  let top = document.documentElement.scrollTop
  if (top >= 200) {
    isShow.value = true
  } else {
    isShow.value = false
  }
}

onUnmounted(() => {
  if (!noScroll) {
    window.removeEventListener("scroll", scroll)
  }
})


async function getData() {

  const val = sessionStorage.getItem("navList")
  if (val !== null) {
    try {
      navList.value = JSON.parse(val)
      return
    } catch (e) {
    }
  }
  let res = await menuNameListApi()
  navList.value = res.data
  if (store.bigModelInfo.enable){
    // 大模型启用
    // 在特定位置加入大模型菜单
    // navList.value
    navList.value.splice(store.bigModelInfo.order, 0, {
      id: 12345,
      title: "大模型",
      path: "/bigModel/square"
    })
  }
  sessionStorage.setItem("navList", JSON.stringify(navList.value))
}

getData()

</script>
<style lang="scss">


.gvb_nav {
  width: 100%;
  top: 0;
  display: flex;
  justify-content: center;
  position: fixed;
  transition: all .3s;
  color: var(--nav_text_color);
  z-index: 100;

  &.isShow {
    background-color: var(--color-bg-1);
    color: var(--color-text-1);
    box-shadow: 5px 0 5px 0 rgba(0, 0, 0, 0.1);

    a {
      color: var(--color-text-1);
    }
  }

  a {
    text-decoration: none;
    color: var(--nav_text_color);

    &.router-link-exact-active {
      color: var(--active);
    }
  }

  .gvb_nav_container {
    width: var(--container_width);
    height: 60px;
    display: flex;
    align-items: center;

    .left {
      width: 70%;
      display: flex;
      align-items: center;

      .logo {
        margin-right: 40px;

        .slogan {
          font-size: 21px;
        }

        .en_slogan {
          font-size: 12px;
          margin-top: 2px;
        }
      }

      .nav {
        > a {
          margin-right: 30px;
          font-size: 16px;
        }
      }

      .gvb_search {
        svg {
          cursor: pointer;
        }
      }
    }

    .right {
      width: 30%;
      display: flex;
      justify-content: end;

      .theme {
        margin-left: 20px;
        display: flex;
        align-items: center;

        svg {
          cursor: pointer;
        }
      }
    }
  }
}
</style>