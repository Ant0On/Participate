<script setup>
import {defineProps} from 'vue';

const props = defineProps({
  room: Object
})

</script>

<template>
  <v-card>
    <v-card-title>
      {{ room?.roomNumber + '. ' + room?.roomName }}
    </v-card-title>
    <v-card-subtitle>
      {{'Area: ' + room?.area + ', Capacity: '+ room?.capacity }}
    </v-card-subtitle>
    <v-card-text>
      {{ room?.roomDescription }}
      <div class="text-subtitle-1 my-1">Room facilities</div>
      <v-divider></v-divider>
      <div class="my-1">{{(room?.roomFacilities) ? room.roomFacilities.join(', '):''}}</div>
    </v-card-text>
  </v-card>
</template>

<style scoped>

</style>