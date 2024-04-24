<script setup>

import FooterBar from "@/components/layout/FooterBar.vue";
import HeaderBar from "@/components/layout/HeaderBar.vue";
</script>

<template>
  <v-app id="app_main">
    <HeaderBar/>
    <v-main class="fill-height">
      <v-container class="h-100" fluid>
        <router-view/>
      </v-container>
    </v-main>
    <FooterBar/>
  </v-app>
</template>

<style scoped>
#app_main {
  min-height: 100%;

}
</style>