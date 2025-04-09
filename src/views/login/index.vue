<template>
  <div class="gvb_login">
    <div class="gvb_login_mask">
      <gvb_login_form @ok="ok" :qq-redirect-path="back"></gvb_login_form>
    </div>
  </div>
</template>
<script setup lang="ts">
import Gvb_login_form from "@/components/common/gvb_login_form.vue";
import router from "@/router";
import {useRoute} from "vue-router";
import {qqLoginApi} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";

interface routerQuery {
  flag?: string
  code?: string
}
interface historyState {
  back: string
}

const back = (window.history.state as historyState).back


const store = useStore()
const route = useRoute()

function ok() {
  let back = window.history.state.back
  // 拿不到之前的路径，就跳转首页
  if (back === null) {
    router.push({name: "index"})
    return
  }
  router.push(back)
}

async function init(query: routerQuery) {
  if (!query.code || !query.flag) {
    return
  }
  let res = await qqLoginApi(query.code)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  store.setToken(res.data)

  // 重定向到点击登录的页面
  let path = localStorage.getItem("redirectPath")
  if (path === null) {
    path = "/"
  }
  router.push(path)
}

init(route.query)


</script>
<style lang="scss">
.gvb_login {
  background: url(/image/bg.png) 50%/cover no-repeat;
  width: 100%;
  height: 100vh;

  .gvb_login_mask {
    position: absolute;
    right: 0;
    width: 400px;
    height: 100%;
    background-color: var(--login_bg);
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 40px;


  }
}

</style>