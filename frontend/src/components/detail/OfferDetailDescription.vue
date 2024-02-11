<script setup>
import {defineProps, defineEmits} from 'vue';
import OfferDetailHeader from "@/components/detail/OfferDetailHeader.vue";
import OfferDetailReserve from "@/components/detail/OfferDetailReserve.vue";
import OfferDetailHost from "@/components/detail/OfferDetailHost.vue";
import OfferDetailInformations from "@/components/detail/OfferDetailInformation.vue";

const props = defineProps({
  name: '',
  price: 0,
  location: '',
  numberOfPeople: 0,
  host_detail: '',
  imagePath:'',
  host_first_name: '',
  type: '',
  description: '',
  offer_id: 0,
})
const emits = defineEmits(['moveToSummary'])

</script>

<template>
  <div class="offer_detail_description">
    <OfferDetailHeader :name="name" :location="location" :type="type" :numberOfPeople="numberOfPeople"/>
    <div class="line"/>
    <OfferDetailReserve @move-to-summary="$emit('moveToSummary')"/>
    <div class="line"/>
    <OfferDetailInformations :description="description"/>
    <div class="line"/>
    <OfferDetailHost :offer_id="offer_id" :host_detail="host_detail" :host_first_name="host_first_name" :imagePath="imagePath"/>
  </div>
</template>

<style scoped>
div.offer_detail_description {
  display: flex;
  flex-direction: column;
}
div.line{
  display: flex;
  flex-direction: row;
  color: black;
  border-bottom: 1px solid;
  margin: 1%;
}


</style>