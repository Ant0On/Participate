<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import SearchBar from "@/components/layout/SearchBar.vue";
import {useOfferStore} from "@/stores/offers.store";

const offerStore = useOfferStore();
const {isLocalization: isLocalization, inputValue: inputValue} = storeToRefs(offerStore)

const events = ref([]);

function mapEvents(responseData) {

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

let pagesGenerator = fetchPaginatedData(`/api/offers/events${getQuery()}`, mapEvents)

onMounted(async () => {
  const response = await pagesGenerator.next();
  events.value = response.value
});

async function load({done}) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty')
    return
  }
  events.value.push(...response.value)
  done('ok');
}
offerStore.$subscribe(async (mutation, state) => {
  pagesGenerator = fetchPaginatedData(`/api/offers/events${getQuery()}`, mapEvents)
  const response = await pagesGenerator.next();
  events.value = response.value

})
</script>

<template>
  <div class="event_page">
    <p>Unforgettable events</p>
    <SearchBar />
    <div v-if="events?.length > 0" >
            <v-infinite-scroll
                :items="eventItem"
                :onLoad="load"
                empty-text="Currently there are no more offers to display!"
                mode="manual"
                class="w-100"
            >
              <v-row class="w-100">
                <template v-for="event in eventItem" :key="event.offerId">
                  <v-col cols="4">
                    <OfferListItem type="event" :offer-item="event"/>
                  </v-col>
                </template>
              </v-row>
            </v-infinite-scroll>

    </div>
    <div v-else class="no_offers">
      <p class="text-center">Currently there are no offers of given type!</p>
    </div>
  </div>
</template>

<style scoped>
.event_page {
  height: 100%;
}

div.no_offers {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 10%;
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

</style>
