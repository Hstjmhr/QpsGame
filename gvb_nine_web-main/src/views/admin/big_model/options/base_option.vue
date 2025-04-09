<script setup lang="ts">

import Gvb_title from "@/components/common/gvb_title.vue";
import {reactive, ref} from "vue";
import type {bigModelSettingsType, bigModelUsableType} from "@/api/big_model_api";
import {bigModelSettingsApi} from "@/api/big_model_api";
import {MdPreview} from "md-editor-v3";
import {useStore} from "@/stores";
import {bigModelSettingsUpdateApi} from "@/api/big_model_api";
import {Message} from "@arco-design/web-vue";
import {bigModelUsableApi} from "@/api/big_model_api";

const store = useStore()
const form = reactive<bigModelSettingsType>({
  name: "",
  enable: false,
  api_key: "",
  api_secret: "",
  title: "",
  prompt: "",
  slogan: "",
  help: "",
  order: 0,
})

const bigModelList = ref<bigModelUsableType[]>([])

async function getBigModelList() {
  let res = await bigModelUsableApi()
  bigModelList.value = res.data
}

getBigModelList()

async function getData() {
  let res = await bigModelSettingsApi()
  Object.assign(form, res.data)
}


async function update() {
  let res = await bigModelSettingsUpdateApi(form)
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
      <a-alert>使用大模型必须的配置项</a-alert>
      <a-form ref="formRef" :model="form">
        <a-form-item label="大模型选择">
          <a-select v-model="form.name" placeholder="大模型选择" :options="bigModelList"></a-select>
        </a-form-item>
        <a-form-item label="api_key">
          <a-input v-model="form.api_key" placeholder="api_key"></a-input>
        </a-form-item>
        <a-form-item label="api_secret">
          <a-input v-model="form.api_secret" placeholder="api_secret"></a-input>
        </a-form-item>
        <a-form-item label="大模型名称">
          <a-input v-model="form.title" placeholder="大模型名称"></a-input>
        </a-form-item>
        <a-form-item label="大模型提示词">
          <a-textarea v-model="form.prompt" :auto-size="{minRows:4, maxRows: 5}"
                      placeholder="大模型提示词"></a-textarea>
        </a-form-item>
        <a-form-item label="slogan">
          <a-textarea v-model="form.slogan" :auto-size="{minRows:3, maxRows: 5}" placeholder="slogan"></a-textarea>
        </a-form-item>
        <a-form-item label="菜单序号">
          <a-input-number v-model="form.order"></a-input-number>
        </a-form-item>
        <a-form-item label="是否启用">
          <a-switch v-model="form.enable"></a-switch>
        </a-form-item>
        <a-form-item>
          <a-button @click="update" type="primary"> 更新</a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="right">
      <gvb_title title="帮助"></gvb_title>
      <md-preview :model-value="form.help" :theme="store.themeString"></md-preview>
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