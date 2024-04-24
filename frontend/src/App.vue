<script setup>

import FooterBar from "@/components/layout/FooterBar.vue";
</script>

<template>
  <div id="app_main">
    <router-view/>
    <FooterBar/>
  </div>
</template>

<style scoped>
#app_main{
  height: 100%;

}
</style>