<script setup>
import {ref} from 'vue';


const icons = ref([
  'mdi-facebook',
  'mdi-twitter',
  'mdi-linkedin',
  'mdi-instagram',
])
</script>

<template>
  <v-layout>
    <v-footer class="d-flex flex-column" style="background-color: #efefef; padding: 16px 0 0 0">
      <div class="d-flex w-100 align-center px-4">
        <strong>Get connected with us on social networks!</strong>

        <v-spacer></v-spacer>

        <v-btn
            v-for="icon in icons"
            :key="icon"
            :icon="icon"
            class="mx-4"
            size="small"
            variant="plain"
        ></v-btn>
      </div>

      <div class="px-4 py-2 bg-white text-center w-100">
        {{ new Date().getFullYear() }} — <strong>Participate</strong>
      </div>
    </v-footer>
  </v-layout>
</template>

<style scoped>

</style>