<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();

const accommodations = ref([
])

const searchOffer = ref({
  offer_type: 'Accommodation',
  location: searchStore.location || '',
  dateFrom: searchStore.dateFrom || '',
  dateTo: searchStore.dateTo || '',
  numberOfPeople: searchStore.numberOfPeople || '',
})

watch(searchOffer.value, (newOffer) => {
  searchStore.setSearchValues(newOffer)
    }
)
</script>

<template>
  <div class="accommodation_page">
    <NavBar currentPage="accommodations"/>
    <p>Climatic places</p>
    <OfferSearch offer_type="Accommodation" v-model:location="searchOffer.location"
                 v-model:date-from="searchOffer.dateFrom" v-model:date-to="searchOffer.dateTo"
                 v-model:number-of-people="searchOffer.numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="accommodation in accommodations" :location="accommodation.location"
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
