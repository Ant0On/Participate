<script setup>
import { ref, watch, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { useSearchStore } from "@/stores/search.store";
import { fetchWrapper } from "@/_helpers/fetch-wrapper";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore);

const events = ref([]);

const allEvents = ref([]);

async function getCurrentEvents() {
  const response = await fetchWrapper.get(`/api/offers?type=event`)

  const responseData = response.data

  allEvents.value = responseData.map((data) => {
    return {
      'offerId': data["offer_id"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'name': data["name"],
      'price': data["price"],
      'maxPeople': data["max_people"]
    }
  })
}

function getNewActivities(location, dateFrom, dateTo, numberOfPeople) {
  return allEvents.value.filter((data) => {

    return checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople)
  })
}

function checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople){
  let isMatchingSearch = true;
  if (typeof location !== "undefined" && location !== "")
  {
    isMatchingSearch = data.location.includes(location)
  }

  if (typeof numberOfPeople !== "undefined" && numberOfPeople !== 0)
  {
    isMatchingSearch = numberOfPeople <= data.maxPeople
  }

  return isMatchingSearch
}

onMounted(async () => {
  await getCurrentEvents();
  events.value = allEvents.value
});

watch(location, (newLocation) => {
  events.value = getNewActivities(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom, (newDateFrom) => {
  events.value = getNewActivities(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  events.value = getNewActivities(location, dateFrom, newDateTo, numberOfPeople)
})
watch(numberOfPeople, (newNumberOfPeople) => {
  events.value = getNewActivities(location, dateFrom, dateTo, newNumberOfPeople)
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
                     :name="event.name" :price="event.price"
                     :max_people="event.maxPeople" type="events" :id="event.offerId"/>
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
