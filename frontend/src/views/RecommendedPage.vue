<script setup>
import {ref, watch, onMounted} from 'vue';
import {storeToRefs} from 'pinia';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const searchStore = useSearchStore();
const {location, dateFrom, dateTo, numberOfPeople} = storeToRefs(searchStore);

const offers = ref([]);

const allOffers = ref([]);

async function getCurrentRecommendedOffers() {
  const response = await fetchWrapper.get(`/api/offers/recommended`)

  const responseData = response.data

  allOffers.value = responseData.map((data) => {
    return {
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'name': data["name"],
      'price': data["price"],
      'maxPeople': data["max_people"]
    }
  })
}

function getNewRecommendedOffers(location, dateFrom, dateTo, numberOfPeople) {
  return allOffers.value.filter((data) => {

    return data.location.startsWith(location)
  })
}

onMounted(async () => {
  await getCurrentRecommendedOffers();
  offers.value = allOffers.value
});

watch(location, (newLocation) => {
  offers.value = getNewRecommendedOffers(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom, (newDateFrom) => {
  offers.value = getNewRecommendedOffers(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  offers.value = getNewRecommendedOffers(location, dateFrom, newDateTo, numberOfPeople)
})

watch(numberOfPeople, (newNumberOfPeople) => {
  offers.value = getNewRecommendedOffers(location, dateFrom, dateTo, newNumberOfPeople)
})

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
</style>
