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
const currentPage = ref(1);
const totalPages = ref(0);

async function getCurrentEvents(page) {
  const response = await fetchWrapper.get(`/api/offers?type=event&page=${page}`);
  const responseData = response?.data || [] ;
  if (responseData){
    events.value = responseData.map((data) => {
      return {
        'offerId': data["offer_id"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'name': data["name"],
        'price': data["price"],
        'maxPeople': data["max_people"]
      };
    });

    totalPages.value = response.totalPages;
  }
}

function getNewEvents(location, dateFrom, dateTo, numberOfPeople) {
  return events.value.filter((data) => {
    return checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople);
  });
}

function checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople) {
  let isMatchingSearch = true;
  if (typeof location !== "undefined" && location !== "") {
    isMatchingSearch = data.location.includes(location);
  }

  if (typeof numberOfPeople !== "undefined" && numberOfPeople !== 0) {
    isMatchingSearch = numberOfPeople <= data.maxPeople;
  }

  return isMatchingSearch;
}

onMounted(async () => {
  await getCurrentEvents(currentPage.value);
});

watch(location, (newLocation) => {
  events.value = getNewEvents(newLocation, dateFrom, dateTo, numberOfPeople);
});

watch(dateFrom, (newDateFrom) => {
  events.value = getNewEvents(location, newDateFrom, dateTo, numberOfPeople);
});

watch(dateTo, (newDateTo) => {
  events.value = getNewEvents(location, dateFrom, newDateTo, numberOfPeople);
});

watch(numberOfPeople, (newNumberOfPeople) => {
  events.value = getNewEvents(location, dateFrom, dateTo, newNumberOfPeople);
});

watch(currentPage, (newPage) => {
  getCurrentEvents(newPage);
});
</script>

<template>
  <div class="event_page">
    <NavBar currentPage="events"/>
    <p>Unforgettable events</p>
    <OfferSearch offer_type="Events" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div v-if="events.length > 0" class="offer_items">
      <OfferListItem v-for="event in events" :location="event.location" :description="event.description"
                     :name="event.name" :price="event.price"
                     :max_people="event.maxPeople" type="events" :id="event.offerId"/>
    </div>
    <div v-else class="no_offers">
      <p class="no_offer_placeholder"> Currently there are no offers of given type!  </p>
    </div>
    <div v-if="totalPages > 1" class="pagination">
      <button @click="currentPage > 1 && (currentPage -= 1)">Previous</button>
      <span >Page {{ currentPage }} of {{ totalPages }}</span>
      <button @click="currentPage < totalPages && (currentPage += 1)">Next</button>
    </div>
  </div>
</template>

<style scoped>
div.no_offers{
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 10%;
}
p.no_offer_placeholder{
  text-align: center;
}

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

.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 20px 0;
}

.pagination button {
  cursor: pointer;
  background-color: #3498db;
  color: #ffffff;
  border: none;
  padding: 10px 15px;
  border-radius: 5px;
  margin: 0 5px;
}

.pagination button:hover {
  background-color: #2980b9;
}

.pagination span {
  margin: 0 10px;
  font-size: 1.2rem;
  color: #333;
}
</style>
