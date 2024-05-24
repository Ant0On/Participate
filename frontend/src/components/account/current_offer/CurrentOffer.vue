<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferReservationListItem from "@/components/offers/OfferReservationListItem.vue";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const allEvents = ref([]);
const allActivities = ref([]);
const allAccommodations = ref([]);

function mapAccommodation(responseData) {

  return responseData.map((data) => {
    return {
      'offerId': data["offer_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'description': data["description"],
      'capacity': data["capacity"],
      'price': data['price_per_day'],
      'isRecommended': data['is_recommended'],
      'discount': data['discount'],
      'type': data['type'],
      'animal_friendly': data['is_animal_friendly'],
      'rating': data['rating']
    };
  });
}

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

const eventsGenerator = fetchPaginatedData(`/api/host/events/${user.value.ID}/reservations/pending`, mapEvents);
const activitiesGenerator = fetchPaginatedData(`/api/host/activity/${user.value.ID}/reservations/pending`, mapActivities);
const accommodationsGenerator = fetchPaginatedData(`/api/host/accommodation/${user.value.ID}/reservations/pending`, mapAccommodation);

onMounted(async () => {
  allAccommodations.value = await accommodationsGenerator.next();
  allActivities.value = await activitiesGenerator.next();
  allEvents.value = await eventsGenerator.next();
});

async function loadOffers(generator, offerList, done) {
  const response = await generator.next();
  if (response?.done) {
    done('empty');
    return;
  }
  offerList.value.push(...response.value);
  done('ok');
}

</script>

<template>
  <div class="event_page">
    <div>
      <v-infinite-scroll
          :items="allEvents"
          :onLoad="({done}) => loadOffers(eventsGenerator, allEvents, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Events</p>
        <v-row class="w-100">
          <template v-for="event in allEvents.value" :key="event.offerID">
            <v-col cols="4">
              <OfferReservationListItem type="event" :offerItem="event" custom>
                <template v-slot:template>
                </template>
              </OfferReservationListItem>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div>
      <v-infinite-scroll
          :items="allAccommodations"
          :onLoad="({done}) => loadOffers(accommodationsGenerator, allAccommodations, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Accommodations</p>
        <v-row class="w-100">
          <template v-for="accommodation in allAccommodations.value" :key="accommodation.offerID">
            <v-col cols="4">
              <OfferReservationListItem type="accommodation" :offerItem="accommodation" custom>
                <template v-slot:template>
                </template>
              </OfferReservationListItem>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div>
      <v-infinite-scroll
          :items="allActivities"
          :onLoad="({done}) => loadOffers(activitiesGenerator, allActivities, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Activities</p>
        <v-row class="w-100">
          <template v-for="activity in allActivities.value" :key="activity.offerID">
            <v-col cols="4">
              <OfferReservationListItem type="activity" :offerItem="activity" custom>
                <template v-slot:template>
                </template>
              </OfferReservationListItem>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
  </div>
</template>

<style scoped>
.event_page {
  height: 100%;
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
