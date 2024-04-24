<script setup>
</script>

<template>
  <v-app-bar class="px-4">
      <router-link to="/">
          <v-img width="250" height="60" :src="require('@/assets/img/logo.png')"></v-img>
      </router-link>


    <v-tabs
        color="grey-darken-2"
        centered
    >
      <router-link to="/recommended">
        <v-tab
            key="recommended"
            text="Recommended"
        ></v-tab>
      </router-link>

      <router-link to="/accommodations">
        <v-tab
            key="accommodations"
            text="Accommodations"
        ></v-tab>
      </router-link>

      <router-link to="/activities">
        <v-tab
            key="activities"
            text="Activities"
        ></v-tab>
      </router-link>

      <router-link to="/events">
        <v-tab
            key="events"
            text="Events"
        ></v-tab>
      </router-link>

    </v-tabs>
    <v-spacer></v-spacer>

    <v-avatar
        class="hidden-sm-and-down"
        color="grey-darken-1"
        size="32"
    ></v-avatar>

  </v-app-bar>
</template>

<style scoped>

</style>