<script setup lang="ts">

import Gvb_title from "@/components/common/gvb_title.vue";
import {reactive, ref} from "vue";
import type {sessionOptionType} from "@/api/big_model_api";
import {sessionOptionApi} from "@/api/big_model_api";
import {MdPreview} from "md-editor-v3";
import {useStore} from "@/stores";
import {sessionOptionUpdateApi} from "@/api/big_model_api";
import {Message} from "@arco-design/web-vue";

const store = useStore()
const form = reactive<sessionOptionType>({
  chat_scope: 0,
  session_scope: 0,
  day_scope: 0,
})


async function getData() {
  let res = await sessionOptionApi()
  Object.assign(form, res.data)
}


async function update() {
  let res = await sessionOptionUpdateApi(form)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  getData()
}

getData()


</script>

<template>
  <div class="big_model_base_option">
    <div class="left">
      <a-alert>用户使用大模型，每一个会话就是全新的聊天</a-alert>
      <a-form ref="formRef" :model="form">
        <a-form-item label="对话积分消耗">
          <a-input-number v-model="form.session_scope" placeholder="对话积分消耗"></a-input-number>
        </a-form-item>
        <a-form-item label="会话积分消耗">
          <a-input-number v-model="form.chat_scope" placeholder="会话积分消耗"></a-input-number>
        </a-form-item>
        <a-form-item label="每日积分领取">
          <a-input-number v-model="form.day_scope" placeholder="每日积分领取"></a-input-number>
        </a-form-item>
        <a-form-item>
          <a-button @click="update" type="primary"> 更新</a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="right">
      <gvb_title title="帮助"></gvb_title>
    </div>
  </div>
</template>

<style lang="scss">
.big_model_base_option {
  display: flex;
  padding: 20px;

  .left {
    width: 40%;

    .arco-form {
      margin-top: 20px;
    }
  }

  .right {
    width: 40%;
    margin-left: 20px;
  }
}
</style>