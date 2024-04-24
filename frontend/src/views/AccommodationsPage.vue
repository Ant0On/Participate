<script setup>
import { ref, watch, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { useSearchStore } from "@/stores/search.store";
import { fetchWrapper } from "@/_helpers/fetch-wrapper";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore);

const accommodations = ref([]);
const currentPage = ref(1);
const totalPages = ref(0);

async function getCurrentAccommodations(page) {
  const response = await fetchWrapper.get(`/api/offers/accommodations?page=${page}`);
  const responseData = response?.data || [] ;

  if(responseData){
    accommodations.value = responseData.map((data) => {
      const priceAfterDiscount = calculatePriceAfterDiscount(data['price_per_day'], data['discount'])

      return {
        'offerId': data["offer_id"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'title': data["title"],
        'price': priceAfterDiscount,
        'capacity': data["capacity"]
      };
    });

    totalPages.value = response.totalPages;
  }
}

function getNewAccommodations(location, dateFrom, dateTo, numberOfPeople) {
  return accommodations.value.filter((data) => {
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
  await getCurrentAccommodations(currentPage.value);
});

watch(location, (newLocation) => {
  accommodations.value = getNewAccommodations(newLocation, dateFrom, dateTo, numberOfPeople);
});

watch(dateFrom, (newDateFrom) => {
  accommodations.value = getNewAccommodations(location, newDateFrom, dateTo, numberOfPeople);
});

watch(dateTo, (newDateTo) => {
  accommodations.value = getNewAccommodations(location, dateFrom, newDateTo, numberOfPeople);
});

watch(numberOfPeople, (newNumberOfPeople) => {
  accommodations.value = getNewAccommodations(location, dateFrom, dateTo, newNumberOfPeople);
});

watch(currentPage, (newPage) => {
  getCurrentAccommodations(newPage);
});
</script>

<template>
  <div class="accommodation_page">
    <p>Climatic places</p>
    <OfferSearch v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div v-if="accommodations.length > 0" class="offer_items">
      <OfferListItem v-for="accommodation in accommodations"
                     :location="accommodation.location"
                     :description="accommodation.description" :title="accommodation.title"
                     :price="accommodation.price" :capacity="accommodation.capacity" type="accommodation" :id="accommodation.offerId"/>
    </div>
    <div v-else class="no_offers">
      <v-spacer></v-spacer>
      <p class="no_offer_placeholder">Currently there are no offers of given type!</p>
      <v-spacer></v-spacer>
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
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
}
p.no_offer_placeholder{
  text-align: center;
}
div.accommodation_page {
  display: flex;
  flex-direction: column;
  height: 100%;
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
