<template>
  <div class="role_list_view">
    <a-drawer v-model:visible="visible" class="big_model_chat_drawer" placement="right" width="50%" :footer="false">
      <template #title>
        <div class="left">{{ drawerTitle }}</div>
        <div class="right">
          <a-button type="primary" status="danger" size="mini" v-if="useIDList.length" @click="chatRemove">删除
          </a-button>
          <a-checkbox v-if="isManage" @change="allIn">全选</a-checkbox>
          <a-checkbox v-model="isManage">管理模式</a-checkbox>
        </div>
      </template>
      <div class="big_model_chat_list_view">
        <a-checkbox-group v-model="useIDList">
          <div :class="{item: true, isCheck: isInList(item.id)}" v-for="item in data.list">
            <div class="user_item">
              <a-checkbox v-if="isManage" :value="item.id"></a-checkbox>
              <div class="avatar">
                <img :src="item.userAvatar" alt="">
              </div>
              <div class="container">
                <div class="date">{{ dateTimeFormat(item.created_at) }}</div>
                <div class="content">{{ item.userContent }}</div>
              </div>
            </div>
            <div class="bot_item">
              <div class="avatar">
                <img :src="item.botAvatar" alt="">
              </div>
              <div class="container">
                <div class="date">{{ dateTimeFormat(item.created_at) }}</div>
                <div class="content">
                  <MdPreview :model-value="item.botContent" :theme="store.themeString"></MdPreview>
                </div>
              </div>
            </div>
          </div>
        </a-checkbox-group>
      </div>
    </a-drawer>
    <gvb_table
        :url="bigModelSessionListApi"
        :columns="columns"
        default-delete
        ref="gvbTable"
        no-add
        no-check
        no-action-group
        edit-label="查看"
        search-placeholder="搜索会话名称"
        @edit="showSession">
    </gvb_table>
  </div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {h, reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {bigModelSessionListApi} from "@/api/big_model_api";
import type {bigModelchatListType} from "@/api/big_model_api";
import type {bigModelSessionType} from "@/api/big_model_api";
import {dateTimeFormat} from "@/utils/date";
import {MdPreview} from "md-editor-v3";
import "md-editor-v3/lib/preview.css"
import {useStore} from "@/stores";
import type {listDataType} from "@/api";
import {bigModelChatListApi} from "@/api/big_model_api";
import {bigModelChatDeleteApi} from "@/api/big_model_api";
import {idl} from "../../../../../dist/assets/idl-91cb7788";

const store = useStore()

const gvbTable = ref()
const columns = [
  {title: '用户名称', dataIndex: 'nickName'},
  {title: '会话', dataIndex: 'sessionName'},
  {title: '角色', dataIndex: 'roleName'},
  {title: '对话次数', dataIndex: 'chatCount'},
  {title: '最后一次聊天内容', dataIndex: 'lastContent'},
  {title: '创建时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]

const visible = ref(false)
const drawerTitle = ref("")
const activeSessionID = ref<number>(0)

const data = reactive<listDataType<bigModelchatListType>>({
  count: 0,
  list: [],
})

const allIDList = ref<number[]>([])

async function showSession(record: bigModelSessionType) {
  visible.value = true
  drawerTitle.value = `${record.nickName}-${record.roleName} 会话详情`
  activeSessionID.value = record.id
  getData()

}

async function getData() {
  let res = await bigModelChatListApi({sessionID: activeSessionID.value})
  data.count = res.data.count
  data.list = res.data.list
  allIDList.value = []
  for (const re of res.data.list) {
    allIDList.value.push(re.id)
  }
}

const isManage = ref<boolean>(false)
const useIDList = ref<number[]>([])

function isInList(id: number): boolean {
  for (const uid of useIDList.value) {
    if (uid === id) {
      return true
    }
  }
  return false
}

function allIn(value: (boolean | (string | number | boolean)[]), ev: Event) {
  let v = value as boolean
  if (v) {
    // 全选
    useIDList.value = allIDList.value
  } else {
    // 取消全选
    useIDList.value = []
  }
}


async function chatRemove() {
  let res = await bigModelChatDeleteApi(useIDList.value)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  useIDList.value = []
  getData()
}

</script>


<style lang="scss">
.big_model_chat_drawer {
  .arco-drawer-title {
    display: flex;
    justify-content: space-between;
    width: 100%;
    align-items: center;

    .right {
      display: flex;
      align-items: center;

      .arco-btn {
        margin-right: 5px;
      }
    }
  }

}

.big_model_chat_list_view {
  .item {
    .bot_item {
      display: flex;
      margin-bottom: 20px;

      .container {
        margin-left: 10px;

        .md-editor {
          background-color: transparent;
        }
      }
    }

    .user_item {
      display: flex;
      flex-direction: row-reverse;
      margin-bottom: 10px;

      .container {
        margin-right: 10px;
        width: auto;
        display: flex;
        flex-direction: column;
        align-items: end;

        .content {
          padding: 10px;
        }

        .date {
          text-align: right;
        }
      }
    }

    .container {
      width: calc(100% - 120px);

      .date {
        font-size: 12px;
      }

      .content {
        background-color: var(--bg);
        border-radius: 5px;
        margin-top: 5px;
        width: fit-content;
      }
    }

    .avatar {
      width: 50px;
      height: 50px;

      img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
      }
    }

    &.isCheck {
      background-color: var(--color-fill-3);
      border-radius: 5px;
      padding: 10px;
      margin-bottom: 20px;
    }
  }
}
</style>