<script setup>
import { ref, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { useSearchStore } from "@/stores/search.store";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore);

const events = ref([]);

function mapEvents(responseData){

  return responseData.map((data) => {
    const priceAfterDiscount = calculatePriceAfterDiscount(data['price'], data['discount'])
    return {
      'offerId': data["offer_id"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'title': data["title"],
      'price': priceAfterDiscount,
      'capacity': data["capacity"]
    };
  });
}

const pagesGenerator = fetchPaginatedData('/api/offers/events', mapEvents)

onMounted(async () => {
  events.value = await pagesGenerator.next();
});

async function load({ done }){
  const response = await pagesGenerator.next();
  if(response?.done)
  {
    done('empty')
    return
  }
  events.value.push(...response.value)
  done('ok');
}
</script>

<template>
  <div class="event_page">
    <p>Unforgettable events</p>
    <OfferSearch v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div v-if="events.length > 0" class="offer_items">
      <v-infinite-scroll :height="300" :items="events" :onLoad="load">
        <template v-slot:empty>
          <p class="no_offer_placeholder">Currently there are no more offers to display!</p>
        </template>
        <template v-for="event in events" :key="item">
          <OfferListItem :location="event.location" :description="event.description"
                         :title="event.title" :price="event.price"
                         :capacity="event.capacity" type="event" :id="event.offerId"/>
        </template>
      </v-infinite-scroll>
    </div>
    <div v-else class="no_offers">
      <p class="no_offer_placeholder">Currently there are no offers of given type!</p>
    </div>
  </div>
</template>

<style scoped>
.event_page{
  height: 100%;
}
div.no_offers{
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 10%;
}
p.no_offer_placeholder{
  text-align: center;
}

div.event_page {
  display: flex;
  flex-direction: column;
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

.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 20px 0;
}

.pagination button {
  cursor: pointer;
  background-color: #3498db;
  color: #ffffff;
  border: none;
  padding: 10px 15px;
  border-radius: 5px;
  margin: 0 5px;
}

.pagination button:hover {
  background-color: #2980b9;
}

.pagination span {
  margin: 0 10px;
  font-size: 1.2rem;
  color: #333;
}
</style>
