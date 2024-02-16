<script setup>
import {defineProps} from 'vue';

const props = defineProps({
  description: '',
})
</script>

<template>
  <div class="offer_detail_information">
    <p class="information_header">Description</p>
    <p class="information">{{ description }}</p>
  </div>
</template>

<style scoped>

p.information_header {
  font-family: "Poppins-SemiBold", Helvetica;
  font-size: 1.4rem;
  font-weight: 600;
  padding: 2%;
}

p.information {
  padding-left: 3%;
  padding-top: 0;
  padding-bottom: 2%;
  font-size: 1.1rem;
  font-weight: 400;
  font-family: "Nokora-Regular", Helvetica;
}
</style>