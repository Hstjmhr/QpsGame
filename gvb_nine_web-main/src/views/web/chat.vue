<template>
  <div class="chat_view">
    <gvb_nav no-scroll></gvb_nav>
    <main>
      <div class="container">
        <gvb_chat_group></gvb_chat_group>
      </div>
    </main>
  </div>
</template>
<script setup lang="ts">
import Gvb_nav from "@/components/web/gvb_nav.vue";
import Gvb_chat_group from "@/components/common/gvb_chat_group.vue";
</script>
<style lang="scss">
.chat_view {
  height: 100vh;

  main {
    width: 100%;
    display: flex;
    justify-content: center;
    padding-top: 80px;

    .container {
      width: var(--container_width);
    }
  }
}
</style>