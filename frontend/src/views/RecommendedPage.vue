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

const offers = ref([]);
const currentPage = ref(1);
const totalPages = ref(0);

async function getCurrentRecommendedOffers(page) {
  const response = await fetchWrapper.get(`/api/offers/recommended?page=${page}`);

  const responseData = response.data;

  offers.value = responseData.map((data) => {
    return {
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'name': data["name"],
      'price': data["price"],
      'maxPeople': data["max_people"]
    };
  });

  totalPages.value = response.totalPages;
}

function getNewRecommendedOffers(location, dateFrom, dateTo, numberOfPeople) {
  return offers.value.filter((data) => {
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
  await getCurrentRecommendedOffers(currentPage.value);
});

watch(location, (newLocation) => {
  offers.value = getNewRecommendedOffers(newLocation, dateFrom, dateTo, numberOfPeople);
});

watch(dateFrom, (newDateFrom) => {
  offers.value = getNewRecommendedOffers(location, newDateFrom, dateTo, numberOfPeople);
});

watch(dateTo, (newDateTo) => {
  offers.value = getNewRecommendedOffers(location, dateFrom, newDateTo, numberOfPeople);
});

watch(numberOfPeople, (newNumberOfPeople) => {
  offers.value = getNewRecommendedOffers(location, dateFrom, dateTo, newNumberOfPeople);
});

watch(currentPage, (newPage) => {
  getCurrentRecommendedOffers(newPage);
});
</script>

<template>
  <div class="recommended_page">
    <NavBar currentPage="recommended"/>
    <p>Recommended offers</p>
    <OfferSearch v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="offer in offers" :location="offer.location"
                     :description="offer.description"
                     :name="offer.name" :price="offer.price"
                     :max_people="offer.maxPeople"/>
    </div>
    <div class="pagination">
      <button @click="currentPage > 1 && (currentPage -= 1)">Previous</button>
      <span v-if="totalPages > 0">Page {{ currentPage }} of {{ totalPages }}</span>
      <button @click="currentPage < totalPages && (currentPage += 1)">Next</button>
    </div>
  </div>
</template>

<style scoped>
div.recommended_page {
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
