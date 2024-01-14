<script setup>
import {defineProps, ref} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferDetailSummary from "@/components/detail/OfferDetailSummary.vue";
import OfferDetailDescription from "@/components/detail/OfferDetailDescription.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const props = defineProps({
  type: String,
  id: String,
})


const offer = ref({
  name: '',
  price: 0,
  location: '',
  numberOfPeople: 0,
  rooms: 0,
  hostId: 0,
})

const isDescription = ref(true)


function getOfferDetails() {
  offer.value = fetchWrapper.get('/offer/detail/' + props.type + '/' + props.id)

}
</script>

<template>
  <div class="event_page">
    <NavBar :currentPage="type"/>
    <div class="item_detail">
      <OfferDetailDescription :type="type" :name="offer.name" :price="offer.price" :location="offer.location"
                              :numberOfPeople="offer.numberOfPeople" :rooms="offer.rooms" :hostId="offer.hostId"
                              v-if="isDescription" @move-to-summary="isDescription = !isDescription"/>
      <OfferDetailSummary :price="offer.price" :image="offer.image" v-else/>
    </div>
  </div>
</template>

<style scoped>
div.item_detail{
  margin: 5% 5% 1% 5%;
  background-color: white;
}
</style>