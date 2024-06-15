<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import OfferListItem from "@/components/offers/OfferListItem.vue";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import SearchBar from "@/components/layout/SearchBar.vue";
import {useOfferStore} from "@/stores/offers.store";

const offerStore = useOfferStore();
const {isLocalization: isLocalization, inputValue: inputValue} = storeToRefs(offerStore)

const accommodations = ref([]);

function mapAccommodation(responseData) {

  return responseData.map((data) => {
    return {
      'offerId': data["ID"],
      'title': data["Title"],
      'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name|| 'city'}`,
      'description': data["Description"],
      'capacity': data["Capacity"],
      'price': data['PricePerDay'],
      'discount': data['Discount'],
      'type': data['Type'],
      'animal_friendly': data['IsAnimalFriendly'],
      'rating': data['RatingAvg'] || 0,
      'ratingCount': data['RatingCount'] || 0,
      'numberOfRooms': data?.NumberOfRooms,
      'rooms': data?.Rooms?.map((room) => {return {...room}}),
      'generalFacilities': data?.GeneralFacilities?.map((generalFacility) => generalFacility.Name)
    };
  });
}

function getQuery() {
  if (inputValue.value) {
    return (isLocalization.value) ? `/?localisation=${inputValue.value}` : `/?title=${inputValue.value}`
  }
  return ''
}

let pagesGenerator = fetchPaginatedData(`/api/offers/accommodations${getQuery()}`, mapAccommodation)

onMounted(async () => {
  const response = await pagesGenerator.next();
  accommodations.value = response.value;
});

async function load({done}) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty')
    return
  }
  accommodations.value.push(...response.value)
  done('ok');
}

offerStore.$subscribe(async (mutation, state) => {
  pagesGenerator = fetchPaginatedData(`/api/offers/accommodations${getQuery()}`, mapAccommodation)
  const response = await pagesGenerator.next();
  accommodations.value = response.value;
})
</script>

<template>
  <div class="accommodation_page">
    <p>Climatic places</p>
    <SearchBar/>
    <div v-if="accommodations.length > 0">
      <v-infinite-scroll
          :items="accommodations"
          :onLoad="load"
          empty-text="Currently there are no more offers to display!"
          mode="manual"
          class="w-100"
      >
        <v-row class="w-100">
          <template v-for="accommodation in accommodations" :key="accommodation.offerId">
            <v-col cols="4">
              <OfferListItem type="accommodation" :offer-item="accommodation"/>
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

div.accommodation_page {
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
