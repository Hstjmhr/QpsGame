<template>
  <div class="gvb_user_info_menu">
    <a-dropdown @select="select" class="gvb_dropdown">
      <div class="gvb_user_info_menu_dropdown">
        <img :src="store.userInfo.avatar" alt="">
        <span class="gvb_user_info_menu_dropdown_span">{{ store.userInfo.nick_name }}</span>
        <IconDown></IconDown>
      </div>
      <template #content>

        <template v-for="(item, index) in dopList" :key="index">
          <a-dgroup v-if="item.type === 'line'"></a-dgroup>
          <a-doption v-else :value="item.name">{{ item.title }}</a-doption>
        </template>
      </template>
    </a-dropdown>
  </div>
</template>
<script setup lang="ts">
import {IconDown} from "@arco-design/web-vue/es/icon";
import {useStore} from "@/stores";
import type {tabType} from "@/types";
import router from "@/router";

interface dopType extends tabType {
  type?: string
}

const store = useStore()

let dopList: dopType[] = [
  {
    name: "home",
    title: "管理系统",
  },
  {
    name: "",
    title: "",
    type: "line"
  },
  {
    name: "article_list",
    title: "文章列表",
  },
  {
    name: "user_list",
    title: "用户列表",
  },
  {
    name: "log_list",
    title: "系统日志",
  },
  {
    name: "",
    title: "",
    type: "line"
  },
  {
    name: "logout",
    title: "注销退出",
  }
]

if (store.isGeneral) {
  dopList = [
    {
      name: "home",
      title: "管理系统",
    },
    {
      name: "user_info",
      title: "个人信息",
    },
    {
      name: "messages",
      title: "我的消息",
    },
    {
      name: "logout",
      title: "注销退出",
    }
  ]
}

function select(value: any) {
  let val = (value as string)
  if (val === "logout") {
    store.logout()
    router.push({name: "index"})
    return
  }
  router.push({
    name: val
  })
}

</script>

<style lang="scss">

.gvb_user_info_menu {
  img {
    width: 30px;
    height: 30px;
    border-radius: 50%;
  }

  .gvb_user_info_menu_dropdown {
    display: flex;
    cursor: pointer;
    align-items: center;

    .gvb_user_info_menu_dropdown_span {
      margin: 0 5px;
    }
  }
}

.gvb_dropdown {
  .arco-dropdown-group-title {
    margin-top: 0;

    &::after {
      content: "";
      display: block;
      width: 100%;
      height: 1px;
      background-color: var(--bg);
      position: absolute;
      left: 0;
    }
  }
}
</style>