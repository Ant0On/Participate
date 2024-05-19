<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import OfferListItem from "@/components/offers/OfferListItem.vue";
import {useOfferStore} from "@/stores/offers.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import SearchBar from "@/components/layout/SearchBar.vue";

const offerStore = useOfferStore();
const {isLocalization: isLocalization, inputValue: inputValue} = storeToRefs(offerStore)
const activities = ref([]);

function mapActivities(responseData) {

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
      'skill': data['skill_level'],
      'type': data['type'],
      'duration': data['duration']
    };
  });
}

function getQuery() {
  if (inputValue.value) {
    return (isLocalization) ? `/?localization=${inputValue.value}` : `/?name=${inputValue.value}`
  }
  return ''
}


let pagesGenerator = fetchPaginatedData(`/api/offers/activities${getQuery()}`, mapActivities)

onMounted(async () => {
  const response = await pagesGenerator.next();
  activities.value = response.value;
});

async function load({done}) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty')
    return
  }
  activities.value.push(...response.value)
  done('ok');
}

offerStore.$subscribe(async (mutation, state) => {
  pagesGenerator = fetchPaginatedData(`/api/offers/activities${getQuery()}`, mapActivities)
  const response = await pagesGenerator.next();
  activities.value = response.value;
})
</script>

<template>
  <div class="activities_page">
    <p>Inspiring activities</p>
    <SearchBar/>
    <div v-if="activities.length > 0">
      <v-infinite-scroll
          :items="activities"
          :onLoad="load"
          empty-text="Currently there are no more offers to display!"
          mode="manual"
          class="w-100"
      >
        <v-row class="w-100">
          <template v-for="activity in activities" :key="activity.offerId">
            <v-col cols="4">
              <OfferListItem type="activity" :offer-item="activity"/>
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

div.activities_page {
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

