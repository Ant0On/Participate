<script setup>
import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const events = []


</script>

<template>
  <div class="event_page">
    <NavBar currentPage="events"/>
    <p>Events</p>
    <OfferSearch offer_type="Events"/>
    <OfferListItem v-for="event in events" :location="event.location" :description="event.description"
                   :image="event.image" :title="event.title"/>
  </div>
</template>

<style scoped>

</style>
