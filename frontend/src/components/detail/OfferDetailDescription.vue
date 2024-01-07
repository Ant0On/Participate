<script setup>
import {defineProps, defineEmits} from 'vue';
import OfferDetailHeader from "@/components/detail/OfferDetailHeader.vue";
import OfferDetailReserve from "@/components/detail/OfferDetailReserve.vue";
import OfferDetailHost from "@/components/detail/OfferDetailHost.vue";

const props = defineProps({
  name: '',
  price: 0,
  location: '',
  numberOfPeople: 0,
  rooms: 0,
  hostId: 0,
  type: '',
})
const emits = defineEmits(['moveToSummary'])
</script>

<template>
  <div class="offer_detail_description">
    <OfferDetailHeader :name="name" location="location" :type="type" :numberOfPeople="numberOfPeople" :rooms="rooms"/>
    <OfferDetailReserve @move-to-summary="$emit('moveToSummary')"/>
    <OfferDetailHost />
  </div>
</template>

<style scoped>
div.offer_detail_description {
  display: flex;
}
</style>