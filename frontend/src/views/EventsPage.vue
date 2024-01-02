<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const events = ref([
])

const searchOffer = ref({
  offer_type: 'Events',
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
  <div class="event_page">
    <NavBar currentPage="events"/>
    <p>Unforgettable events</p>
    <OfferSearch offer_type="Events" v-model:location="searchOffer.location"
                 v-model:date-from="searchOffer.dateFrom" v-model:date-to="searchOffer.dateTo"
                 v-model:number-of-people="searchOffer.numberOfPeople"/>
    <div class="offer_items">
    <OfferListItem v-for="event in events" :location="event.location" :description="event.description"
                   :image="event.image" :title="event.title" :price="event.price"
                   :number-of-people="event.numberOfPeople"/>
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
