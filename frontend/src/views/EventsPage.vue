<script setup>
import {ref, watch} from 'vue';
import {storeToRefs} from 'pinia';


import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore)

const events = ref(getCurrentEvents(location, dateFrom, dateTo, numberOfPeople))


function getCurrentEvents(location, dateFrom, dateTo, numberOfPeople){
  return fetchWrapper.get('/api/offers')
}

watch(location, (newLocation) => {
  events.value = getCurrentEvents(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom, (newDateFrom) => {
  events.value = getCurrentEvents(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  events.value = getCurrentEvents(location, dateFrom, newDateTo, numberOfPeople)
})
watch(numberOfPeople, (newNumberOfPeople) => {
  events.value = getCurrentEvents(location, dateFrom, dateTo, newNumberOfPeople)
})
</script>

<template>
  <div class="event_page">
    <NavBar currentPage="events"/>
    <p>Unforgettable events</p>
    <OfferSearch offer_type="Events" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
    <OfferListItem v-for="event in events" :location="event.location" :description="event.description"
                   :image="event.image" :title="event.title" :price="event.price"
                   :number-of-people="event.numberOfPeople"/>
    </div>
  </div>
</template>

<style scoped>
div.event_page {
  display: flex;
  flex-direction: column;
  overflow: scroll;
}

p {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.8rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
  margin: 1% 1% 1% 1%;
}

.offer_items {
  margin: 3% 5% 0 5%;
}
</style>
