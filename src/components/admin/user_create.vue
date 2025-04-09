<template>
  <div>
    <a-modal title="创建用户" :visible="props.visible" :on-before-ok="createUser" @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item field="user_name" label="用户名"
                     :rules="[{required:true,message:'请输入用户名'}]"
                     :validate-trigger="['blur']"
        >
          <a-input v-model="form.user_name" placeholder="用户名"></a-input>
        </a-form-item>
        <a-form-item field="nick_name" label="昵称"
                     :rules="[{required:true,message:'请输入昵称'}]"
                     :validate-trigger="['blur']"
        >
          <a-input v-model="form.nick_name" placeholder="昵称"></a-input>
        </a-form-item>
        <a-form-item field="password" label="密码"
                     :rules="[{required:true,message:'请输入密码'}]"
                     :validate-trigger="['blur']"
        >
          <a-input type="password" v-model="form.password" placeholder="密码"></a-input>
        </a-form-item>
        <a-form-item field="re_password" label="确认密码"
                     :rules="[{required:true,message:'请输入确认密码'}, {validator: rePasswordValidator}]"
                     :validate-trigger="['blur']"
        >
          <a-input type="password" v-model="form.re_password" placeholder="确认密码"></a-input>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.role" :options="roleOptions" placeholder="选择角色"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import {reactive, ref} from "vue";
import {userCreateApi} from "@/api/user_api";
import type {userCreateRequest} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";
import {roleOptions} from "@/global/global";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  nick_name: "",
  password: "",
  role: 2,
  user_name: "",
  re_password: "",
}

const form = reactive<userCreateRequest & { re_password: string }>({
  nick_name: "",
  password: "",
  role: 2,
  user_name: "",
  re_password: "",
})

const formRef = ref()

function rePasswordValidator(value: string, callback: (error?: string) => void): void {
  if (value !== form.password) return callback("确认密码不一致")
}

async function createUser() {
  let val = await formRef.value.validate()
  if (val) return false

  let res = await userCreateApi(form)
  if (res.code) {
    Message.error(res.msg)
    return false
  }
  Message.success(res.msg)
  emits("update:visible", false)
  emits("ok")
  Object.assign(form, defaultForm)
  return true
}


</script>
<style scoped lang="scss">

</style>