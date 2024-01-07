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

//const accommodations = ref(getCurrentAccommodations(location, dateFrom, dateTo, numberOfPeople))
const accommodations = ref([{
  id: 1,
  location: "Polska",
  image: "@/assets/img/test.jpg",
  title: "Domek w polsce",
  description: "Duzy domek wygodny w polsce tanio!!",
  price: "200 zlotych polskich",
  numberOfPeople: 3,
}])

function getCurrentAccommodations(location, dateFrom, dateTo, numberOfPeople){
  return fetchWrapper.get('/api/offers')
}

watch(location, (newLocation) => {
  accommodations.value = getCurrentAccommodations(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom, (newDateFrom) => {
  accommodations.value = getCurrentAccommodations(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  accommodations.value = getCurrentAccommodations(location, dateFrom, newDateTo, numberOfPeople)
})
watch(numberOfPeople, (newNumberOfPeople) => {
  accommodations.value = getCurrentAccommodations(location, dateFrom, dateTo, newNumberOfPeople)
})
</script>

<template>
  <div class="accommodation_page">
    <NavBar currentPage="accommodations"/>
    <p>Climatic places</p>
    <OfferSearch offer_type="Accommodation" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="accommodation in accommodations" type="accommodations" :location="accommodation.location" :id="accommodation.id"
                     :description="accommodation.description" :image="accommodation.image" :title="accommodation.title"
                     :price="accommodation.price" :number-of-people="accommodation.numberOfPeople"/>
    </div>
  </div>
</template>

<style scoped>
div.accommodation_page {
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
