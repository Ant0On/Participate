<script setup>
import { onMounted, ref, reactive } from 'vue';
import { useAuthStore } from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferListItem from "@/components/offers/OfferListItem.vue";

const auth = useAuthStore();
const user = auth.user;
const errors = reactive({
  apiError: ""
})

const allEvents = ref([])

async function getMyOffers(responseData) {
  try {
    return responseData.map((data) => {
      return {
        'location': data["country_name"] + ', ' + data["town_name"],
        'title': data["title"],
        'offerID': data['offer_id'],
        'discount': data['discount'],
        'capacity': data['capacity'],
        'eventType': data['type'],
        'offerType': 'event'
      };
    });
  } catch (error) {
    console.error('Error during mapping or assignment:', error);
  }
}

const pagesGenerator = fetchPaginatedData(`/api/host/event/${user.ID}/offers`, getMyOffers);

onMounted(async () => {
  allEvents.value = await pagesGenerator.next();
});

async function load({ done }) {
  const response = await pagesGenerator.next();
  if (response?.done) {
    done('empty');
    return;
  }
  allEvents.value.push(...response.value);
  done('ok');
}

function setDiscount(offerID) {
  // Logic for setting discount
}

function changePrice(offerID) {
  // Logic for changing price
}

function editOffer(offerID) {
  // Logic for changing price
}

function deleteOffer(offerID) {
  // Logic for changing price
}

</script>

<template>
  <div class="event_page">
    <div>
      <v-infinite-scroll
          :items="allEvents"
          :onLoad="load"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <v-row class="w-100">
          <template v-for="event in allEvents.value" :key="event.offerID">
            <v-col cols="4">
              <OfferListItem type="event" :offerItem="event">
                <template #actions="{ index }">
                  <v-btn v-if="index === 1" @click="setDiscount(event.offerID)">Set discount</v-btn>
                  <v-btn v-if="index === 1" @click="changePrice(event.offerID)">Change price</v-btn>
                  <v-btn v-if="index === 0" @click="editOffer(event.offerID)">Edit</v-btn>
                  <v-btn v-if="index === 0" @click="deleteOffer(event.offerID)">Delete</v-btn>
                </template>
              </OfferListItem>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div v-else class="no_offers">
      <p class="text-center">Currently there are no offers!</p>
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
