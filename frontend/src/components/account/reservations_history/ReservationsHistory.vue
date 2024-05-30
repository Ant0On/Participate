<script setup>
import {onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferReservationListItem from "@/components/offers/OfferReservationListItem.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore);

const allEvents = ref([]);
const allActivities = ref([]);
const allAccommodations = ref([]);
const allRooms = ref([]);

const ratings = ref({});

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
      'ratingId': data['rating_id']
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
      'duration': data['duration'],
      'ratingId': data['rating_id']

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

function mapRooms(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["room_id"],
      'title': data["title"],
      'location': data["country_name"] + ', ' + data["town_name"],
      'state': data['reservation_state'],
      'capacity': data["capacity"],
      'dateFrom': data['date_from'],
      'dateTo': data['dateTo'],
      'price': data['price_per_day'],
      'reservationId': data['reservation_id'],
      'animal_friendly': data['is_animal_friendly'],
      'ratingId': data['rating_id']
    };
  });
}

const eventsGenerator = fetchPaginatedData(`/api/customer/${user.value.ID}/reservations/event/history`, mapEvents);
const activitiesGenerator = fetchPaginatedData(`/api/customer/${user.value.ID}/reservations/activity/history`, mapActivities);
const accommodationsGenerator = fetchPaginatedData(`/api/customer/${user.value.ID}/reservations/accommodation/history`, mapAccommodation);
const roomsGenerator = fetchPaginatedData(`/api/customer/${user.value.ID}/reservations/room/history`, mapRooms);

onMounted(async () => {
  allAccommodations.value = await accommodationsGenerator.next();
  allActivities.value = await activitiesGenerator.next();
  allEvents.value = await eventsGenerator.next();
  allRooms.value = await roomsGenerator.next();
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

async function rateReservation(reservation, type) {
  console.log(reservation)
  await fetchWrapper.post(`/api/customer/offer/${type}/${reservation.reservationId}/rate`, {
    Count: reservation.ratingId
  });
}
</script>

<template>
  <div class="event_page">
    <div>
      <v-infinite-scroll
          :items="allEvents"
          :onLoad="({ done }) => loadOffers(eventsGenerator, allEvents, done)"
          empty-text="Currently there are no more reservations to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Events</p>
        <v-row class="w-100">
          <template v-for="event in allEvents.value" :key="event.reservationId">
            <v-col cols="4">
              <OfferReservationListItem v-else type="event" :offerItem="event"/>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div>
      <v-infinite-scroll
          :items="allAccommodations"
          :onLoad="({ done }) => loadOffers(accommodationsGenerator, allAccommodations, done)"
          empty-text="Currently there are no more reservations to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Accommodations</p>
        <v-row class="w-100">
          <template v-for="accommodation in allAccommodations.value" :key="accommodation.reservationId">
            <v-col cols="4">
              <OfferReservationListItem v-if="accommodation.state === 'finished'" type="accommodation"
                                        :offerItem="accommodation" custom>
                <template v-slot:template>
                  <v-card v-if="! accommodation.ratingId"
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      Rate your reservation!
                    </v-card-title>
                    <v-card-text>
                      <v-rating
                          v-model="accommodation.ratingId"
                          @click="rateReservation(accommodation, 'accommodation')"
                          color="red-lighten-2"
                          hover
                          :disabled="accommodation.ratingId > 0"
                      ></v-rating>
                    </v-card-text>
                  </v-card>
                  <v-card v-else
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      You rated your reservation:
                    </v-card-title>
                    <v-card-text class="d-flex flex-row justify-space-between align-center">
                      <v-rating
                          v-model="accommodation.ratingId"
                          @click="rateReservation(accommodation, 'accommodation')"
                          color="red-lighten-2"
                          hover
                          readonly
                      ></v-rating>
                      <span class="ml-4 mt-1"> {{ accommodation.ratingId }} </span>
                    </v-card-text>
                  </v-card>
                </template>
              </OfferReservationListItem>
              <OfferReservationListItem v-else type="accommodation" :offerItem="accommodation"/>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div>
      <v-infinite-scroll
          :items="allActivities"
          :onLoad="({ done }) => loadOffers(activitiesGenerator, allActivities, done)"
          empty-text="Currently there are no more reservations to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Activities</p>
        <v-row class="w-100">
          <template v-for="activity in allActivities.value" :key="activity.reservationId">
            <v-col cols="4">
              <OfferReservationListItem v-if="activity.state === 'finished'" type="activity" :offerItem="activity"
                                        custom>
                <template v-slot:template>
                  <v-card v-if="! activity.ratingId"
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      Rate your reservation!
                    </v-card-title>
                    <v-card-text>
                      <v-rating
                          v-model="activity.ratingId"
                          @click="rateReservation(activity, 'activity')"
                          color="red-lighten-2"
                          hover
                          :disabled="activity.ratingId > 0"
                      ></v-rating>
                    </v-card-text>
                  </v-card>
                  <v-card v-else
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      You rated your reservation:
                    </v-card-title>
                    <v-card-text class="d-flex flex-row justify-space-between align-center">
                      <v-rating
                          v-model="activity.ratingId"
                          @click="rateReservation(activity, 'activity')"
                          color="red-lighten-2"
                          hover
                          readonly
                      ></v-rating>
                      <span class="ml-4 mt-1"> {{ activity.ratingId }} </span>
                    </v-card-text>
                  </v-card>

                </template>
              </OfferReservationListItem>
              <OfferReservationListItem v-else type="activity" :offerItem="activity"/>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <div>
      <v-infinite-scroll
          :items="allRooms"
          :onLoad="({ done }) => loadOffers(roomsGenerator, allRooms, done)"
          empty-text="Currently there are no more reservations to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Rooms</p>
        <v-row class="w-100">
          <template v-for="room in allRooms.value" :key="room.reservationId">
            <v-col cols="4">
              <OfferReservationListItem v-if="room.state === 'finished'" type="room" :offerItem="room" custom>
                <template v-slot:template>
                  <v-card v-if="! room.ratingId"
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      Rate your reservation!
                    </v-card-title>
                    <v-card-text>
                      <v-rating
                          v-model="room.ratingId"
                          @click="rateReservation(room, 'room')"
                          color="red-lighten-2"
                          hover
                          :disabled="room.ratingId > 0"
                      ></v-rating>
                    </v-card-text>
                  </v-card>
                  <v-card v-else
                          elevation="0"
                          class="d-flex flex-column justify-space-between align-center"
                          height="100">
                    <v-card-title>
                      You rated your reservation:
                    </v-card-title>
                    <v-card-text class="d-flex flex-row justify-space-between align-center">
                      <v-rating
                          v-model="room.ratingId"
                          @click="rateReservation(room, 'room')"
                          color="red-lighten-2"
                          hover
                          readonly
                      ></v-rating>
                      <span class="ml-4 mt-1"> {{ room.ratingId }} </span>
                    </v-card-text>
                  </v-card>
                </template>
              </OfferReservationListItem>
              <OfferReservationListItem v-else type="room" :offerItem="room"/>
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
