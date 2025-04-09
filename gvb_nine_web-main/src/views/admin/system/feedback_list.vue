<template>
  <div class="feedback_list_view">
    <gvb_table
        :url="feedbackListApi"
        :columns="columns"
        default-delete
        no-confirm
        ref="gvbTable"
        search-placeholder="搜索邮箱|反馈内容"
        no-add
        @edit="edit">
    </gvb_table>
  </div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {reactive, ref} from "vue";
import {feedbackListApi} from "@/api/fade_back_api";
import type {fadeBackType} from "@/api/fade_back_api";

const gvbTable = ref()


const columns = [
  {title: 'ID', dataIndex: 'id'},
  {title: '邮箱', dataIndex: 'email'},
  {title: '反馈内容', dataIndex: 'content'},
  {title: '更新时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]

const visible = ref(false)



function edit(record: fadeBackType) {
  visible.value = true
}


</script>
