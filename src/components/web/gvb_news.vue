<template>
  <div class="gvb_news">
    <span class="index">{{ data.index }}</span>
    <a class="title" :href="data.link" target="_blank">{{ data.title }}</a>
    <span class="hot">{{ data.hotValue }}</span>
  </div>
</template>

<script setup lang="ts">
import type {newsType} from "@/api/news_api";
import {reactive} from "vue";
import {newsListApi} from "@/api/news_api";

const data = reactive<newsType>({
  hotValue: "192.3万",
  index: "1",
  link: "http://www.fengfengzhidao.com",
  title: "枫枫知道发布八代博客了",
})

async function getData(){
  let res = await newsListApi({
    id: "KqndgxeLl9",
    size: 1,
  })
  if (res.code){
    return
  }
  Object.assign(data, res.data[0])
}
getData()

</script>
<style lang="scss">
.gvb_news {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .title {
    color: var(--active);
    text-decoration: none;
    width: 80%;
  }
}
</style>