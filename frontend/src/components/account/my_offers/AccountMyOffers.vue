<script setup>
import {onMounted, reactive, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const userStore = useAuthStore();

const {user: user} = storeToRefs(userStore)
const errors = reactive({
  apiError: ""
});

const allEvents = ref([]);
const allActivities = ref([]);
const allAccommodations = ref([]);
const confirmDelete = ref(false);
const offerToDelete = ref(null);

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

const eventsGenerator = fetchPaginatedData(`/api/host/event/${user.value.ID}/offers`, mapEvents);
const activitiesGenerator = fetchPaginatedData(`/api/host/activity/${user.value.ID}/offers`, mapActivities);
const accommodationsGenerator = fetchPaginatedData(`/api/host/accommodation/${user.value.ID}/offers`, mapAccommodation);

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

async function deleteOffer(offerID, offerType) {
  try {
    const apiMap = {
      event: `/api/host/event/delete/${offerID}`,
      activity: `/api/host/activity/delete/${offerID}`,
      accommodation: `/api/host/accommodation/delete/${offerID}`
    };

    await fetchWrapper.delete(apiMap[offerType]);

    const offerListMap = {
      event: allEvents,
      activity: allActivities,
      accommodation: allAccommodations
    };

    offerListMap[offerType].value.value = offerListMap[offerType].value.value.filter(offer => offer.offerID !== offerID);
  } catch (error) {
    console.error('Error deleting the offer:', error);
  }
}

function confirmDeleteOffer(offerID, offerType) {
  offerToDelete.value = {offerID, offerType};
  confirmDelete.value = true;
}

function handleDeleteConfirmation(result) {
  if (result) {
    deleteOffer(offerToDelete.value.offerID, offerToDelete.value.offerType);
  }
  confirmDelete.value = false;
  offerToDelete.value = null;
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
              <OfferListItem type="event" :offerItem="event" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(event.offerID)">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(event.offerID, 'event')">Delete
                    </v-btn>
                  </v-card>
                </template>
              </OfferListItem>
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
              <OfferListItem type="accommodation" :offerItem="accommodation" custom>
                <template v-slot:template>
                  <v-card elevation="0" class=" d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(accommodation.offerID)">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(accommodation.offerID, 'accommodation')">Delete
                    </v-btn>
                  </v-card>
                </template>
              </OfferListItem>
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
              <OfferListItem type="activity" :offerItem="activity" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(activity.offerID)">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(activity.offerID, 'activity')">Delete
                    </v-btn>
                  </v-card>
                </template>
              </OfferListItem>
            </v-col>
          </template>
        </v-row>
      </v-infinite-scroll>
    </div>
    <v-dialog v-model="confirmDelete" max-width="400">
      <v-card>
        <v-card-title class="headline">Confirm Deletion</v-card-title>
        <v-card-text>Are you sure you want to delete this offer?</v-card-text>
        <v-card-actions>
          <v-btn color="primary" text @click="handleDeleteConfirmation(false)">Cancel</v-btn>
          <v-btn color="primary" @click="handleDeleteConfirmation(true)">Yes</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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
