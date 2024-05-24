<script setup>
import { useNavStore } from '@/stores/nav.store'
import {storeToRefs} from 'pinia';

const navStore = useNavStore()
const {currentPage: currentPage} = storeToRefs(navStore)

</script>

<template>
  <v-app-bar class="px-4" :elevation="0">
    <router-link to="/">
      <v-img width="250" height="60" :src="require('@/assets/img/logo.png')"></v-img>
    </router-link>
    <v-tabs
        color="grey-darken-2"
        v-model="currentPage"
        centered
    >
      <router-link to="/">
        <v-tab
            key="home"
            text="Home"
            value="home"
        ></v-tab>
      </router-link>

      <router-link to="/accommodations">
        <v-tab
            key="accommodations"
            text="Accommodations"
            value="accommodations"
        ></v-tab>
      </router-link>

      <router-link to="/activities">
        <v-tab
            key="activities"
            text="Activities"
            value="activities"
        ></v-tab>
      </router-link>

      <router-link to="/events">
        <v-tab
            key="events"
            text="Events"
            value="events"
        ></v-tab>
      </router-link>

    </v-tabs>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<style scoped>

</style>