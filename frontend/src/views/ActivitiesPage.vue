<script setup>
import { ref, watch, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { useSearchStore } from "@/stores/search.store";
import { fetchWrapper } from "@/_helpers/fetch-wrapper";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore);

const activities = ref([]);
const currentPage = ref(1);
const totalPages = ref(0);

async function getCurrentActivities(page) {
  const response = await fetchWrapper.get(`/api/offers/activities?page=${page}`)
  const responseData = response?.data || [] ;

  if(responseData){
    activities.value = responseData.map((data) => {
      const priceAfterDiscount = calculatePriceAfterDiscount(data['price'], data['discount'])

      return {
        'offerId': data["offer_id"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'title': data["title"],
        'price': priceAfterDiscount,
        'capacity': data["capacity"]
      }
    })

    totalPages.value = response.totalPages;
  }
}

function getNewActivities(location, dateFrom, dateTo, numberOfPeople) {
  return activities.value.filter((data) => {
    return checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople)
  })
}

function checkIfOfferMatchesSearch(data, location, dateFrom, dateTo, numberOfPeople) {
  let isMatchingSearch = true;
  if (typeof location !== "undefined" && location !== "") {
    isMatchingSearch = data.location.includes(location)
  }

  if (typeof numberOfPeople !== "undefined" && numberOfPeople !== 0) {
    isMatchingSearch = numberOfPeople <= data.maxPeople
  }

  return isMatchingSearch
}

onMounted(async () => {
  await getCurrentActivities(currentPage.value);
});

watch(location, (newLocation) => {
  activities.value = getNewActivities(newLocation, dateFrom, dateTo, numberOfPeople)
})

watch(dateFrom, (newDateFrom) => {
  activities.value = getNewActivities(location, newDateFrom, dateTo, numberOfPeople)
})

watch(dateTo, (newDateTo) => {
  activities.value = getNewActivities(location, dateFrom, newDateTo, numberOfPeople)
})

watch(numberOfPeople, (newNumberOfPeople) => {
  activities.value = getNewActivities(location, dateFrom, dateTo, newNumberOfPeople)
})

watch(currentPage, (newPage) => {
  getCurrentActivities(newPage);
})
</script>

<template>
  <div class="activities_page">
    <NavBar currentPage="activities"/>
    <p>Inspiring activities</p>
    <OfferSearch v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div v-if="activities.length > 0" class="offer_items">
      <OfferListItem v-for="activity in activities" :location="activity.location" :description="activity.description"
                     :title="activity.title" :price="activity.price"
                     :capacity="activity.capacity" type="activity" :id="activity.offerId"/>
    </div>
    <div v-else class="no_offers">
      <p class="no_offer_placeholder">Currently there are no offers of given type!</p>
    </div>
    <div v-if="totalPages > 1" class="pagination">
      <button @click="currentPage > 1 && (currentPage -= 1)">Previous</button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
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

