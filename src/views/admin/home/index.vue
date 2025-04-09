<template>
  <div class="home_view">
    <gvb_home_welcome></gvb_home_welcome>
    <div class="main">
      <div class="left">
        <gvb_card title="快捷入口" class="quick_entry">
          <div class="item" v-for="item in quickEntryList">
            <div class="icon" @click="goLink(item)">
              <component :is="item.font"></component>
            </div>
            <div class="text">{{ item.text }}</div>
          </div>
        </gvb_card>
        <gvb_card class="statistics" title="数据统计">
          <login_data_charts></login_data_charts>
        </gvb_card>
      </div>
      <div class="right">
        <gvb_card title="更新日志" class="update_log">
          <div class="item" v-for="(item, index) in updateLogList">
            <div class="line1">
                   <span>
             <span class="index">{{ index + 1 }}.</span>
             <span class="content"><a-typography-paragraph :ellipsis="{
        rows: 1,
        showTooltip: true,
        css: true
      }">{{ item.title }}</a-typography-paragraph></span>
           </span>
              <span class="date" :title="dateFormat(item.created_at)">{{ relativeCurrentTime(item.created_at) }}</span>
            </div>
            <ul class="line2" v-if="item.items">
              <li v-for="li in item.items">{{ li }}</li>
            </ul>
          </div>
        </gvb_card>
      </div>
    </div>

  </div>
</template>

<script lang="ts" setup>
import Gvb_card from "@/components/common/gvb_card.vue";
import type {Component} from "vue";
import {IconSettings, IconUser, IconMessage, IconBook, IconUserGroup} from "@arco-design/web-vue/es/icon";
import router from "@/router";
import {relativeCurrentTime, dateFormat} from "@/utils/date";
import Login_data_charts from "@/components/charts/login_data_charts.vue";
import Gvb_home_welcome from "@/components/admin/gvb_home_welcome.vue";
import {h} from "vue";
import {defineComponent} from "vue";
import {useStore} from "@/stores";


const store = useStore()
interface quickEntryType {
  font: Component
  text: string // 文字
  name?: string // 路由的名称
  link?: string // 可以跳外部链接
}

let quickEntryList: quickEntryType[] = [
  {
    font: IconUser,
    text: "个人信息",
    name: "user_info"
  },
  {
    font: IconMessage,
    text: "群聊管理",
    name: "chat_list"
  },
  {
    font: IconSettings,
    text: "系统日志",
    name: "log_list"
  },
  {
    font: defineComponent({
      render: () => {
        return h("i", {class: "fa fa-comments"})
      }
    }),
    text: "评论列表",
    name: "comment_list"
  },
  {
    font: IconBook,
    text: "文章列表",
    name: "article_list"
  },
  {
    font: IconUserGroup,
    text: "用户列表",
    name: "user_list"
  },
]

if (store.isGeneral){
  quickEntryList = [
    {
      font: IconUser,
      text: "个人信息",
      name: "user_info"
    }
  ]
}


function goLink(item: quickEntryType) {
  if (item.name) {
    router.push({
      name: item.name
    })
    return
  }
  if (item.link) {
    window.open(item.link, "_blank")
  }

}


interface updateLogType {
  id?: number
  title: string
  created_at: string
  items?: string[]
}

const updateLogList: updateLogType[] = [
  {
    title: " 八代博客重磅更新，日志功能优化，群聊功能大更新",
    created_at: "2023-9-29",
    items: [
      "技术栈：gin vue3 TypeScript ArcoDesign",
      "经过原型设计，ui设计的全新博客系统",
      "全新的后台管理系统",
      "丰富可玩的配置项功能",
      "聊天室可发图片",
      "系统设置新增帮助",
      "前端mock数据，可独立运行",
      "移动端单独开发",
    ]
  },
  {
    title: "七代博客教程发布，在go圈里比较火的博客教程",
    created_at: "2023-2-28",
    items: [
      "技术栈：gin vue3 JavaScript AntDesign"
    ]
  },
  {
    title: "六代博客发布，基于gin+vue的前后端不分离项目",
    created_at: "2023-1-16",
  },
  {
    title: "五代博客教程发布，个人最受欢迎的的前后端不分离博客教程（基于django）",
    created_at: "2022-1-12",
  },
  {
    title: "四代个人博客发布（基于django）",
    created_at: "2022-1-10"
  },
  {
    title: "二代博客论坛发布（基于django）",
    created_at: "2021-10-26"
  }
]


</script>

<style lang="scss">
.home_view {

  .main {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;

    .left {
      width: 70%;

      .quick_entry {
        .body {
          display: flex;
          overflow: hidden;
        }

        .item {
          display: flex;
          flex-direction: column;
          justify-content: center;
          margin-right: 50px;
          align-items: center;

          &:last-child {
            margin-right: 0;
          }

          .icon {
            background-color: var(--color-fill-2);
            border-radius: 10px;
            width: 60px;
            height: 60px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 24px;
            transition: all .3s;
            cursor: pointer;

            svg {
              transition: all .3s;
            }

            &:hover {
              transform: scale(1.05);

              svg {
                transform: scale(1.20);
              }
            }
          }

          .text {
            margin-top: 5px;
            color: var(--color-text-2);
          }
        }
      }

      .statistics {
        margin-top: 20px;
      }
    }

    .right {
      width: calc(30% - 20px);

      .update_log {
        .item {
          color: var(--color-text-2);
          margin-bottom: 15px;

          .line1 {
            display: flex;
            align-items: center;
            justify-content: space-between;

            > span {
              display: flex;
              align-items: center;
              white-space: nowrap;

              .index {
                margin-right: 10px;
              }
            }

            .content {
              display: flex;

              .arco-typography {
                margin-bottom: 0;
                display: inline-block;
              }
            }
          }

          .line2 {
            margin: 5px 0;
            line-height: 1.5rem;
          }

        }
      }
    }
  }
}
</style>