<template>
  <div class="flex relative h-screen overflow-hidden">
    <div class="w-1/4 border-r-[1px] border-white/5 h-screen bg-midnight">
      <MainMenu/>
    </div>
    <div class="w-full  overflow-y-auto h-screen pt-6">
      <div>
      <slot/>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import MainMenu from "../navigations/MainMenu.vue";
</script>