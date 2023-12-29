<script setup>
import {defineProps} from 'vue';
const props = defineProps({
  name: String,
  offerType: String,
  image: String,
  dateFrom: String,
  dateTo: String,
  withAnimals: Boolean,
  customer: Object,
});

</script>

<template>
  <div class="account_history_item">
    <img :src="image">
    <div>
      <span>{{ name }}</span>
      <span>Type: {{ offerType }}</span>
      <span>Date from: {{ dateFrom }}</span>
      <span>Date to: {{ dateTo }}</span>
      <span>Animals: {{ (withAnimals)? 'Yes': 'No'}}</span>
      <span><a href="link_to_customer/link_to_chat">Ordered by: {{ customer.first_name + " " + customer.last_name }}<a/></span>

    </div>

  </div>
</template>

<style scoped>

</style>