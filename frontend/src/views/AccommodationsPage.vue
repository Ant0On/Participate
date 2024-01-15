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

const accommodations = ref([]);

const allAccommodations = ref([]);

async function getCurrentAccommodations() {
  const response = await fetchWrapper.get(`/api/offers?type=accommodation`)

  const responseData = response.data

  allAccommodations.value = responseData.map((data) => {
    return {
      'offerId': data["offer_id"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'name': data["name"],
      'price': data["price"],
      'max_people': data["max_people"]
    }
  })
}

function getNewAccommodations(location, dateFrom, dateTo, numberOfPeople){
  return allAccommodations.value.filter((data)=>{

    return data.location.startsWith(location) // W tym miejscu trzeba dobry warunek
  })
}
onMounted(async () => {
  await getCurrentAccommodations();
  accommodations.value = allAccommodations.value
});

watch(location, (newLocation) => {
  accommodations.value =  getNewAccommodations(newLocation, dateFrom, dateTo, numberOfPeople)
})
watch(dateFrom,  (newDateFrom) => {
  accommodations.value =  getNewAccommodations(location, newDateFrom, dateTo, numberOfPeople)
})
watch(dateTo, (newDateTo) => {
  accommodations.value =  getNewAccommodations(location, dateFrom, newDateTo, numberOfPeople)
})

watch(numberOfPeople, (newNumberOfPeople) => {
  accommodations.value =  getNewAccommodations(location, dateFrom, dateTo, newNumberOfPeople)
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
      <OfferListItem v-for="accommodation in accommodations"
                     :location="accommodation.location"
                     :description="accommodation.description" :image="accommodation.image" :name="accommodation.name"
                     :price="accommodation.price" :max_people="accommodation.max_people" type="accommodations" :id="accommodation.offerId"/>
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
