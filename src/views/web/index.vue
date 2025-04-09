<template>
  <div class="index_view">
    <gvb_nav></gvb_nav>
    <gvb_banner></gvb_banner>
    <main>
      <div class="container">
        <div class="left">
          <gvb_card title="今日热搜">
            <template #head-right>
              <router-link :to="{name: 'news'}">更多
                <IconDoubleRight></IconDoubleRight>
              </router-link>
            </template>
            <gvb_news/>
          </gvb_card>
          <gvb_card title="文章日历">
            <article_calendar></article_calendar>
          </gvb_card>
          <gvb_card title="文章列表" class="article_card">
            <template #head-right>
              <a-input-search v-model="key" @search="search" @keydown.enter="search"
                              placeholder="搜索文章标题"></a-input-search>
            </template>
            <gvb_article_list ref="gvbArticleList"></gvb_article_list>
          </gvb_card>
        </div>
        <div class="right">
          <gvb_card title="独家推广">
            <gvb_promotion></gvb_promotion>
          </gvb_card>
          <gvb_card title="标签云">
            <gvb_tags></gvb_tags>
          </gvb_card>
          <gvb_recommend></gvb_recommend>
          <gvb_card title="个人名片">
            <gvb_user_card></gvb_user_card>
          </gvb_card>
          <gvb_card title="意见反馈">
            <gvb_fade_back></gvb_fade_back>
          </gvb_card>
        </div>
      </div>
    </main>
    <gvb_footer></gvb_footer>
  </div>
</template>
<script setup lang="ts">
import Gvb_nav from "@/components/web/gvb_nav.vue";
import Gvb_banner from "@/components/web/gvb_banner.vue";
import Gvb_card from "@/components/common/gvb_card.vue";

import {IconDoubleRight} from "@arco-design/web-vue/es/icon";
import Gvb_footer from "@/components/web/gvb_footer.vue";
import Gvb_news from "@/components/web/gvb_news.vue";
import Gvb_promotion from "@/components/web/gvb_promotion.vue";
import Gvb_user_card from "@/components/web/gvb_user_card.vue";
import Gvb_fade_back from "@/components/web/gvb_fade_back.vue";
import Article_calendar from "@/components/charts/article_calendar.vue";
import Gvb_article_list from "@/components/web/gvb_article_list.vue";
import {ref} from "vue";
import Gvb_tags from "@/components/web/gvb_tags.vue";
import Gvb_recommend from "@/components/web/gvb_recommend.vue";


const gvbArticleList = ref()

const key = ref("")

function search() {
  gvbArticleList.value.getData({key: key.value})
}

</script>
<style lang="scss">
.index_view {
  main {
    width: 100%;
    display: flex;
    justify-content: center;
    margin-top: 20px;

    .container {
      width: var(--container_width);
      display: flex;
      justify-content: space-between;

      > .right {
        width: 380px;
      }

      > .left {
        width: calc(100% - 400px);


        .article_card {
          background: none;

          .head {
            background-color: var(--color-bg-1);
            border-radius: 5px 5px 0 0;
          }

          .body {
            padding: 0;
          }
        }
      }

      .gvb_card {
        margin-bottom: 20px;
      }
    }
  }
}
</style>