<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const activities = ref([
])
const searchOffer = ref({
  offer_type: 'Activities',
  location: '',
  dateFrom: '',
  dateTo: '',
  numberOfPeople: '',
})

watch(searchOffer.value, (newOffer) => {
      searchOffer.value = newOffer
    }
)
</script>

<template>
  <div class="activities_page">
    <NavBar currentPage="activities"/>
    <p>Inspiring activities</p>
    <OfferSearch offer_type="Activities" v-model:location="searchOffer.location"
                 v-model:date-from="searchOffer.dateFrom" v-model:date-to="searchOffer.dateTo"
                 v-model:number-of-people="searchOffer.numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="activity in activities" :location="activity.location" :description="activity.description"
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
