<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const events = []

const accommodations = []
const searchOffer = ref({
  offer_type: 'Events',
  location: '',
  dateFrom: '',
  dateTo: '',
  numberOfPeople: '',
})

watch(searchOffer.value, (newOffer) => {
      searchOffer.value = newOffer
      console.log("Value has changed")
    }
)
</script>

<template>
  <div class="event_page">
    <NavBar currentPage="events"/>
    <p>Events</p>
    <OfferSearch offer_type="Events" v-model:location="searchOffer.location"
                 v-model:date-from="searchOffer.dateFrom" v-model:date-to="searchOffer.dateTo"
                 v-model:number-of-people="searchOffer.numberOfPeople"/>
    <OfferListItem v-for="event in events" :location="event.location" :description="event.description"
                   :image="event.image" :title="event.title"/>
  </div>
</template>

<style scoped>

</style>
