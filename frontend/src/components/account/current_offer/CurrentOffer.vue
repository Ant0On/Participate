<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferReservationListItem from "@/components/offers/OfferReservationListItem.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const allEvents = ref([]);
const allActivities = ref([]);
const allAccommodations = ref([]);
const stateMap = {
  pending: {
    color: "gray",
    icon: "mdi-state-machine",
    text: "Pending"
  },
  finished: {
    color: "grey",
    icon: "mdi-state-machine",
    text: "Finished",
  },
  accepted: {
    color: "green",
    text: "Accepted",
    icon: "mdi-state-machine",
  },
  rejected: {
    color: "red",
    text: "Rejected",
    icon: "mdi-state-machine",
  },
  ongoing: {
    color: "blue",
    text: "Ongoing",
    icon: "mdi-state-machine",
  }
}

function mapAccommodation(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["accommodation_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'state': data['reservation_state'],
      'type': data['type'],
      'capacity': data["capacity"],
      'dateFrom': data['date_from'],
      'dateTo': data['dateTo'],
      'price': data['price_per_day'],
      'reservationId': data['reservation_id'],
      'animal_friendly': data['is_animal_friendly'],

    };
  });
}

function mapActivities(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["activity_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'date': data["date"],
      'capacity': data["capacity"],
      'price': data['price'],
      'state': data['reservation_state'],
      'reservationId': data['reservation_id'],
      'skill': data['skill_level'],
      'type': data['type'],
      'duration': data['duration']
    };
  });
}

function mapEvents(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["event_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'state': data["reservation_state"],
      'capacity': data["capacity"],
      'price': data['price'],
      'reservationId': data['reservation_id'],
      'date': data['date'],
      'type': data['type']
    };
  });
}

const eventsGenerator = fetchPaginatedData(`/api/host/event/${user.value.ID}/reservations`, mapEvents);
const activitiesGenerator = fetchPaginatedData(`/api/host/activity/${user.value.ID}/reservations`, mapActivities);
const accommodationsGenerator = fetchPaginatedData(`/api/host/accommodation/${user.value.ID}/reservations`, mapAccommodation);

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

function changeListReservationState(reservations, reservationId, state){
  reservations.value.map((reservation) => {
    if(reservation.reservationId === reservationId)
      reservation.state = state
    return reservation
  })
}

async function acceptReservation(reservationId, type, reservations){
  fetchWrapper.post(`/api/reservation/${type}/${reservationId}/accepted`, {}).then(() => {
    changeListReservationState(reservations, reservationId, "accepted")
  }).catch((err) => {

  })
}
async function rejectReservation(reservationId, type, reservations){
  fetchWrapper.post(`/api/reservation/${type}/${reservationId}/rejected`, {}).then(() => {
    changeListReservationState(reservations, reservationId, "rejected")
  }).catch((err) => {

  })
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
          <template v-for="event in allEvents.value" :key="event.reservationId">
            <v-col cols="4">
              <OfferReservationListItem v-if="event.state === 'pending' "  type="event" :offerItem="event" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="green-lighten-2" rounded
                           @click="acceptReservation(event.reservationId, 'event', allEvents)">Accept
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="rejectReservation(event.reservationId, 'event', allEvents)">Reject
                    </v-btn>
                  </v-card>
                </template>
              </OfferReservationListItem>
              <OfferReservationListItem v-else  type="event" :offerItem="event"/>
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
          <template v-for="accommodation in allAccommodations.value" :key="accommodation.reservationId">
            <v-col cols="4">
              <OfferReservationListItem type="accommodation" :offerItem="accommodation" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="green-lighten-2" rounded
                           @click="acceptReservation(accommodation.reservationId, 'accommodation', allAccommodations)">Accept
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="rejectReservation(accommodation.reservationId, 'accommodation', allAccommodations)">Reject
                    </v-btn>
                  </v-card>
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
          <template v-for="activity in allActivities.value" :key="activity.reservationId">
            <v-col cols="4">
              <OfferReservationListItem type="activity" :offerItem="activity" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="green-lighten-2" rounded
                           @click="acceptReservation(activity.reservationId, 'activity', allActivities)">Accept
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="rejectReservation(activity.reservationId, 'activity', allActivities)">Reject
                    </v-btn>
                  </v-card>
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
