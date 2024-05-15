<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useSearchStore} from "@/stores/search.store";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import SearchBar from "@/components/layout/SearchBar.vue";

const searchStore = useSearchStore();
const {location, dateFrom, dateTo, numberOfPeople} = storeToRefs(searchStore);

const events = ref([]);

function mapEvents(responseData) {

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

async function load({done}) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty')
    return
  }
  events.value.push(...response.value)
  done('ok');
}

const eventItem = [{
  offerId: 1,
  title: "Wakacyjne pierdolenie",
  location: "Zamosc, Polska",
  description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
  capacity: 10,
  isRecommended: true,
  discount: 10,
  price: 30,
  type: "festival",
  rating: 4.5,
  duration: 10,

},
  {
    offerId: 1,
    title: "Wakacyjne pierdolenie",
    location: "Zamosc, Polska",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
    capacity: 10,
    isRecommended: true,
    discount: 10,
    price: 30,
    type: "festival",
    rating: 4.5,
    duration: 10,

  },  {
    offerId: 1,
    title: "Wakacyjne pierdolenie",
    location: "Zamosc, Polska",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
    capacity: 10,
    isRecommended: true,
    discount: 10,
    price: 30,
    type: "festival",
    rating: 4.5,
    duration: 10,

  },  {
    offerId: 1,
    title: "Wakacyjne pierdolenie",
    location: "Zamosc, Polska",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
    capacity: 10,
    isRecommended: true,
    discount: 10,
    price: 30,
    type: "festival",
    rating: 4.5,
    duration: 10,

  },  {
    offerId: 1,
    title: "Wakacyjne pierdolenie",
    location: "Zamosc, Polska",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
    capacity: 10,
    isRecommended: true,
    discount: 10,
    price: 30,
    type: "festival",
    rating: 4.5,
    duration: 10,

  },  {
    offerId: 1,
    title: "Wakacyjne pierdolenie",
    location: "Zamosc, Polska",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
    capacity: 10,
    isRecommended: true,
    discount: 10,
    price: 30,
    type: "festival",
    rating: 4.5,
    duration: 10,

  },
  ]
</script>

<template>
  <div class="event_page">
    <p>Unforgettable events</p>
    <SearchBar />
    <div v-if="eventItem.length > 0" >
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
