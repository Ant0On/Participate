<script setup>

import FooterBar from "@/components/layout/FooterBar.vue";
</script>

<template>
  <v-app id="app_main">
    <router-view/>
    <FooterBar/>
  </v-app>
</template>

<style scoped>
#app_main{
  min-height: 100%;
  background-color: #efefef;

}
</style>