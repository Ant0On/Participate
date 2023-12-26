<script setup>
import {ref, watch} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();

const accommodations = []
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
    <OfferListItem v-for="accommodation in accommodations" :location="accommodation.location"
                   :description="accommodation.description" :image="accommodation.image" :title="accommodation.title"/>
  </div>
</template>

<style scoped>

</style>
