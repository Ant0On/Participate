<script setup>
import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const accommodations = []
</script>

<template>
  <div class="accommodation_page">
    <NavBar currentPage="accommodations"/>
    <p>Climatic places</p>
    <OfferSearch offer_type="Accommodation"/>
    <OfferListItem v-for="accommodation in accommodations" :location="accommodation.location"
                   :description="accommodation.description" :image="accommodation.image" :title="accommodation.title"/>
  </div>
</template>

<style scoped>

</style>
