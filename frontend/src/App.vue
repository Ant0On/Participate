<script setup>

import FooterBar from "@/components/layout/FooterBar.vue";
import HeaderBar from "@/components/layout/HeaderBar.vue";
import NavigationDrawer from "@/components/layout/NavigationDrawer.vue";
</script>

<template>
  <v-app id="app_main">
    <NavigationDrawer/>
    <HeaderBar/>
    <div class="d-flex flex-0-1-100">
    <v-main class="pl-0">
      <v-container class="h-100" fluid>
        <router-view/>
      </v-container>
    </v-main>
    </div>
    <FooterBar />
  </v-app>
</template>

<style scoped>
#app_main {
  min-height: 100%;

}
</style>