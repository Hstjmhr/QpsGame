<template>
  <div class="role_list_view">
    <a-modal title="创建角色" width="45%" v-model:visible="visible" :on-before-ok="ok">
      <div class="role_add_modal">
        <div class="left">
          <a-form ref="formRef" :model="form">
            <a-form-item field="name" label="角色名称" :rules="[{required:true,message:'请输入角色名称'}]"
                         :validate-trigger="['blur']">
              <a-input v-model="form.name" placeholder="角色名称"></a-input>
            </a-form-item>
            <a-form-item field="icon" label="角色图标" :rules="[{required:true,message:'请输入角色图标'}]"
                         :validate-trigger="['blur']">
              <a-select v-model="form.icon"  placeholder="角色图标">
                <a-option v-for="item in iconOptions" :value="item.value" :key="item.value">
                  <div class="banner_image_div" style="display: flex; align-items: center">
                    <img height="30" width="30" style="border-radius: 5px; margin-right: 10px" :src="item.value as string" alt=""> {{ item.label }}
                  </div>
                </a-option>
              </a-select>
              <a-image width="30px" height="30px" v-if="form.icon !== ''" style="border-radius: 5px; margin-left: 5px"
                       :src="form.icon"></a-image>
            </a-form-item>
            <a-form-item  label="消耗积分">
              <a-input-number v-model="form.scope" placeholder="消耗积分"></a-input-number>
            </a-form-item>
            <a-form-item field="tagList" label="角色标签" :rules="[{required:true,message:'请选择角色标签'}]"
                         :validate-trigger="['blur']">
              <a-select v-model="form.tagList" multiple placeholder="角色标签" :options="tagOptions"></a-select>
            </a-form-item>
            <a-form-item field="abstract" label="角色简介" :rules="[{required:true,message:'请输入角色简介'}]"
                         :validate-trigger="['blur']">
              <a-textarea v-model="form.abstract" placeholder="角色简介"
                          :auto-size="{minRows: 3, maxRows: 4}"></a-textarea>
            </a-form-item>
            <a-form-item field="prologue" label="开场白" :rules="[{required:true,message:'请输入开场白'}]"
                         :validate-trigger="['blur']">
              <a-textarea v-model="form.prologue" placeholder="开场白"
                          :auto-size="{minRows: 3, maxRows: 4}"></a-textarea>
            </a-form-item>
            <a-form-item field="prompt" label="角色设定" :rules="[{required:true,message:'请输入角色设定'}]"
                         :validate-trigger="['blur']">
              <a-textarea v-model="form.prompt" placeholder="角色设定"
                          :auto-size="{minRows: 5, maxRows: 8}"></a-textarea>
            </a-form-item>
            <a-form-item label="自动回复">
              <a-switch v-model="form.autoReply"></a-switch>
            </a-form-item>
            <a-form-item label="是否启用">
              <a-switch v-model="form.enable"></a-switch>
            </a-form-item>
          </a-form>
        </div>
        <div class="right">
          <big_model_role_preview v-if="form.name !== '' && form.icon !== '' && form.prologue !== ''"
                                  :data="form"></big_model_role_preview>
        </div>
      </div>
    </a-modal>
    <gvb_table
        :url="bigModelRoleListApi"
        :columns="columns"
        default-delete
        add-label="创建角色"
        no-confirm
        ref="gvbTable"
        no-check
        no-action-group
        search-placeholder="搜索角色名称"
        @add="add"
        @edit="edit">
      <template #enable="{record}:{record: bigModelRoleType}">
        <a-switch :model-value="record.enable"></a-switch>
      </template>
      <template #autoReply="{record}:{record: bigModelRoleType}">
        <a-switch :model-value="record.autoReply"></a-switch>
      </template>
      <template #icon="{record}:{record: bigModelRoleType}">
        <a-image width="40px" height="40px" :src="record.icon"></a-image>
      </template>
      <template #tags="{record}:{record: bigModelRoleType}">
        <div class="col_tag">
          <a-tag v-for="item in record.tags" :color="item.color">
            {{ item.title }}
          </a-tag>
        </div>
      </template>
      <template #abstract="{record}:{record: bigModelRoleType}">
        <a-typography-text :ellipsis="{rows: 1}">{{ record.abstract }}</a-typography-text>
      </template>
      <template #prologue="{record}:{record: bigModelRoleType}">
        <a-typography-text :ellipsis="{rows: 1}">{{ record.prologue }}</a-typography-text>
      </template>
      <template #prompt="{record}:{record: bigModelRoleType}">
        <a-typography-text :ellipsis="{rows: 1}">{{ record.prompt }}</a-typography-text>
      </template>
    </gvb_table>
  </div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {h, reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {bigModelRoleListApi, bigModelRoleCreateApi, bigModelRoleUpdateApi} from "@/api/big_model_api";
import type {bigModelRoleType} from "@/api/big_model_api";
import {bigModelRoleTagOptionsApi} from "@/api/big_model_api";
import type {baseResponse, optionsResponse} from "@/api";
import Big_model_role_preview from "@/components/admin/big_model_role_preview.vue";
import {bigModelRoleIconOptionsApi} from "@/api/big_model_api";
import type {optionType} from "@/types";

const iconOptions = ref<optionType[]>([])

async function getIcons() {
  let res = await bigModelRoleIconOptionsApi()
  iconOptions.value = res.data
}

getIcons()

const form = reactive<bigModelRoleType>({
  id: 0,
  created_at: "",
  name: "",
  enable: false,
  icon: "",
  abstract: "",
  tags: [],
  scope: 2,
  prologue: "",
  prompt: "",
  autoReply: false,
  tagList: []
})

const emptyForm = {
  id: 0,
  name: "",
  enable: false,
  icon: "",
  abstract: "",
  tags: [],
  scope: 2,
  prologue: "",
  prompt: "",
  autoReply: false,
  tagList: []
}

const gvbTable = ref()
const columns = [
  {title: '角色名称', dataIndex: 'name'},
  {title: '是否启用', slotName: 'enable'},
  {title: '角色图标', slotName: 'icon'},
  {title: '角色简介', slotName: 'abstract'},
  {title: '角色标签', slotName: 'tags'},
  {title: '消耗积分', dataIndex: 'scope'},
  {title: '对话开场白', slotName: 'prologue'},
  {title: '角色设定', slotName: 'prompt'},
  {title: '是否启用自动回复', slotName: 'autoReply'},
  {title: '添加时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]

const visible = ref(false)

function edit(record: bigModelRoleType) {
  formRef.value.clearValidate(Object.keys(form))
  form.id = record.id
  form.name = record.name
  form.enable = record.enable
  form.icon = record.icon
  form.abstract = record.abstract
  form.scope = record.scope
  form.prologue = record.prologue
  form.prompt = record.prompt
  form.autoReply = record.autoReply
  let tagList = []
  for (const tag of record.tags) {
    tagList.push(tag.id)
  }
  form.tagList = tagList
  visible.value = true
}

function add() {
  formRef.value.clearValidate(Object.keys(form))
  Object.assign(form, emptyForm)
  visible.value = true
}

const tagOptions = ref<optionsResponse<number>[]>([])

async function getTagOptions() {
  let res = await bigModelRoleTagOptionsApi()
  tagOptions.value = res.data
}

getTagOptions()
const formRef = ref()

async function ok() {
  let val = await formRef.value.validate()
  if (val) return false
  let res: baseResponse<string>
  if (form.id === 0) {
    // 添加
    res = await bigModelRoleCreateApi(form)
  } else {
    res = await bigModelRoleUpdateApi(form)
  }
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  visible.value = false
  gvbTable.value.getList()
}

</script>


<style lang="scss">
.role_list_view {
  .arco-typography {
    margin-bottom: 0;
  }

  .col_tag {
    display: flex;

    .arco-tag {
      margin-left: 10px;

      &:first-child {
        margin-left: 0;
      }
    }
  }
}

.role_add_modal {
  display: flex;

  .left {
    width: 50%;
  }

  .right {
    width: 50%;
    padding-left: 20px;
  }
}
</style>