<template>
  <div class="gvb_promotion">
    <a-carousel auto-play
                indicator-type="dot"
                show-arrow="hover">
      <a-carousel-item v-for="item in promotionList">
        <a :href="item.href" target="_blank" class="link">
          <img :src="item.images" :alt="item.title"/>
        </a>
      </a-carousel-item>
    </a-carousel>
  </div>
</template>

<script setup lang="ts">
import {promotionListApi} from "@/api/promotion_api";
import type {promotionType} from "@/api/promotion_api";
import {ref} from "vue";

const promotionList = ref<promotionType[]>([])

async function getData(){
  let res = await promotionListApi()
  promotionList.value = res.data.list
}
getData()

</script>
<style lang="scss">
.gvb_promotion {
  height: 200px;

  .arco-carousel {
    height: 100%;

    .link {
      width: 100%;
      height: 100%;
      display: block;

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }
}
</style>