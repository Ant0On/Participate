<script setup>
import {ref, watch} from 'vue';
import {storeToRefs} from 'pinia';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore)

const activities = ref(getCurrentActivities(location, dateFrom, dateTo, numberOfPeople))


function getCurrentActivities(location, dateFrom, dateTo, numberOfPeople){
  return fetchWrapper.get('/api/offers')
}

watch(location, (newLocation) => {
  activities.value = getCurrentActivities(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom, (newDateFrom) => {
  activities.value = getCurrentActivities(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  activities.value = getCurrentActivities(location, dateFrom, newDateTo, numberOfPeople)
})
watch(numberOfPeople, (newNumberOfPeople) => {
  activities.value = getCurrentActivities(location, dateFrom, dateTo, newNumberOfPeople)
})
</script>

<template>
  <div class="activities_page">
    <NavBar currentPage="activities"/>
    <p>Inspiring activities</p>
    <OfferSearch offer_type="Activities" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="activity in activities" type="activities" :id="activity.id" :location="activity.location" :description="activity.description"
                     :image="activity.image" :title="activity.title" :price="activity.price"
                     :number-of-people="activity.numberOfPeople"/>
    </div>
  </div>
</template>

<style scoped>
div.activities_page {
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
