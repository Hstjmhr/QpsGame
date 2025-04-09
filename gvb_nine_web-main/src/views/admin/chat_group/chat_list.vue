<template>
  <div class="chat_list_view">
    <gvb_chat_group class="record_list_view"></gvb_chat_group>
    <div class="config">
      <chat_group_config></chat_group_config>
    </div>
  </div>
</template>

<script setup lang="ts">
import Chat_group_config from "@/components/admin/chat_group_config.vue";
import Gvb_chat_group from "@/components/common/gvb_chat_group.vue";

</script>

<style lang="scss">
.chat_list_view {
  display: flex;
  height: calc(100vh - 130px);
  justify-content: space-between;

  .record_list_view {
    width: calc(100% - 400px);
    height: 100%;
  }

  .config {
    width: 380px;
    height: 100%;
    background-color: var(--color-bg-1);
    border-radius: 5px;
    padding: 20px;
    overflow-y: auto;
  }
}
</style>