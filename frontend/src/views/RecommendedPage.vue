<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import OfferListItem from "@/components/offers/OfferListItem.vue";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import SearchBar from "@/components/layout/SearchBar.vue";
import {useOfferStore} from "@/stores/offers.store";

const offerStore = useOfferStore();
const {isLocalization: isLocalization, inputValue: inputValue} = storeToRefs(offerStore)

const offers = ref([]);

function mapRecommended(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["offer_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'capacity': data["capacity"],
      'price': data['price'],
      'isRecommended': data['is_recommended'],
      'discount': data['discount'],
      'type': data['type']
    };
  });
}

function getQuery() {
  if (inputValue.value) {
    return (isLocalization) ? `/?localization=${inputValue.value}` : `/?name=${inputValue.value}`
  }
  return ''
}

let pagesGenerator = fetchPaginatedData(`/api/offers/recommended${getQuery()}`, mapRecommended)

onMounted(async () => {
  const response = await pagesGenerator.next();
  offers.value = response.value;
});

async function load({done}) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty')
    return
  }
  offers.value.push(...response.value)
  done('ok');
}

offerStore.$subscribe(async (mutation, state) => {
  pagesGenerator = fetchPaginatedData(`/api/offers/recommended${getQuery()}`, mapRecommended)
  const response = await pagesGenerator.next();
  offers.value = response.value;
})
</script>

<template>
  <div class="recommended_page">
    <p>Recommended offers</p>
    <SearchBar/>
    <div v-if="offers.length > 0">
      <v-infinite-scroll
          :items="offers"
          :onLoad="load"
          empty-text="Currently there are no more offers to display!"
          mode="manual"
          class="w-100"
      >
        <v-row class="w-100">
          <template v-for="offer in offers" :key="offer.offerId">
            <v-col cols="4">
              <OfferListItem type="event" :offer-item="offer"/>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div v-else class="no_offers">
      <p class="no_offer_placeholder">Currently there are no offers of given type!</p>
    </div>
  </div>
</template>

<style scoped>
div.no_offers {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 10%;
}

p.no_offer_placeholder {
  text-align: center;
}

div.recommended_page {
  display: flex;
  flex-direction: column;
  height: 100%;

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
</style>
