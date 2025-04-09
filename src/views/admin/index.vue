<template>
  <div class="gvb_admin">
    <aside :class="{collapsed: store.collapsed}">
      <Gvb_logo></Gvb_logo>
      <Gvb_menu></Gvb_menu>
    </aside>
    <main :class="{isLaptops: isLaptops1}">
      <div class="gvb_head">
        <Gvb_bread_crumb></Gvb_bread_crumb>
        <div class="gvb_function_area">
          <IconHome class="action_icon" @click="goIndex"></IconHome>
          <Gvb_theme></Gvb_theme>
          <Gvb_user_info_menu></Gvb_user_info_menu>
        </div>
      </div>
      <Gvb_tabs></Gvb_tabs>
      <div class="gvb_container">
        <router-view v-slot="{Component}">
          <transition name="fade" mode="out-in">
            <component :is="Component"></component>
          </transition>
        </router-view>
      </div>
    </main>

  </div>
</template>
<script setup lang="ts">
import Gvb_menu from "@/components/admin/gvb_menu.vue";
import {IconHome, IconDown,} from '@arco-design/web-vue/es/icon';
import type {Component} from "vue";
import {useRouter} from "vue-router";
import {useRoute} from "vue-router";
import Gvb_bread_crumb from "@/components/admin/gvb_bread_crumb.vue";
import Gvb_logo from "@/components/admin/gvb_logo.vue";
import Gvb_tabs from "@/components/admin/gvb_tabs.vue";
import Gvb_theme from "@/components/common/gvb_theme.vue";
import {useStore} from "@/stores";
import Gvb_user_info_menu from "@/components/common/gvb_user_info_menu.vue";
const isLaptops1 = isLaptops

const store = useStore()
const route = useRoute()
const router = useRouter()

function goIndex() {
  router.push({
    name: "index"
  })
}


</script>
<style lang="scss">
.gvb_admin {
  display: flex;
  color: var(--color-text-1);
  height: 100vh;
  overflow-y: hidden;

  aside {
    width: 240px;
    border-right: 1px solid var(--bg);
    height: 100vh;
    background-color: var(--color-bg-1);
    transition: all .3s;
    position: relative;

    .gvb_menu{
      height: calc(100vh - 90px);
      overflow-y: auto;
      overflow-x: hidden;
    }
  }

  aside.collapsed {
    width: 48px;

    & ~ main {
      width: calc(100% - 48px);
    }
  }

  main {
    width: calc(100% - 240px);
    transition: all .3s;

    .gvb_head {
      width: 100%;
      height: 60px;
      border-bottom: 1px solid var(--bg);
      display: flex;
      justify-content: space-between;
      padding: 0 20px;
      align-items: center;
      background-color: var(--color-bg-1);

      .gvb_function_area {
        display: flex;
        align-items: center;

        .action_icon {
          margin-right: 10px;
          cursor: pointer;
          font-size: 16px;
          transition: color .3s;

          &:hover {
            color: var(--active);
          }
        }
      }
    }


    .gvb_container {
      background-color: var(--bg);
      height: calc(100vh - 90px);
      padding: 20px;
      overflow-y: auto;
      overflow-x: hidden;
    }

  }
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

.fade-enter-to {
  transform: translateX(0px);
  opacity: 1;
}

.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}

</style>