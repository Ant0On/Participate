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

const activities = ref([]);

const allActivities = ref([]);

async function getCurrentActivities() {
  const response = await fetchWrapper.get(`/api/offers?type=activity`)

  const responseData = response.data

  allActivities.value = responseData.map((data) => {
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

function getNewActivities(location, dateFrom, dateTo, numberOfPeople){
  return allActivities.value.filter((data)=>{

    return data.location.startsWith(location)
  })
}
onMounted(async () => {
  await getCurrentActivities();
  activities.value = allActivities.value
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
</script>

<template>
  <div class="activities_page">
    <NavBar currentPage="activities"/>
    <p>Inspiring activities</p>
    <OfferSearch offer_type="Activities" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="activity in activities" :location="activity.location" :description="activity.description"
                     :image="activity.image" :name="activity.name" :price="activity.price"
                     :max_people="activity.max_people" type="activities" :id="activity.offerId"/>
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
