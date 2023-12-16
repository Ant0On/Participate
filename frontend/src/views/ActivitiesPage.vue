<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const activities = []
const searchOffer = ref({
  offer_type: 'Activities',
  location: '',
  dateFrom: '',
  dateTo: '',
  numberOfPeople: '',
})

// watch(searchOffer.value, (newOffer) => {
//     searchOffer.value = newOffer
//       console.log("Value has changed")
//     }
// )
</script>

<template>
  <div class="activities_page">
    <NavBar currentPage="activities"/>
    <p>Inspiring activities</p>
    <OfferSearch offer_type="Activities" v-model:location="searchOffer.location"
                 v-model:date-from="searchOffer.dateFrom" v-model:date-to="searchOffer.dateTo"
                 v-model:number-of-people="searchOffer.numberOfPeople"/>
    <OfferListItem v-for="activity in activities" :location="activity.location" :description="activity.description"
                   :image="activity.image" :title="activity.title"/>
  </div>
</template>

<style scoped>

</style>
