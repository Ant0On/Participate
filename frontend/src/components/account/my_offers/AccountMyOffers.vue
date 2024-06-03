<script setup>
import {onMounted, reactive, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {router} from "@/router";

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
      'generalFacilities': data?.GeneralFacilities?.map((generalFacility) => generalFacility.Name),


    };
  });
}

function mapActivities(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["ID"],
      'title': data["Title"],
      'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name|| 'city'}`,
      'description': data["Description"],
      'capacity': data["Capacity"],
      'price': data['Price'],
      'discount': data['Discount'],
      'skill': data['Skill'],
      'type': data['Type'],
      'duration': data['Duration'],
      'date': data?.Date?.split('T')?.[0],
      'equipment': data?.Equipment?.map((equipment)=> equipment?.Name),
      'rating': data['RatingAvg'] || 0,
      'ratingCount': data['RatingCount'] || 0,
    };
  });
}

function mapEvents(responseData) {
  return responseData.map((data) => {
    return {
      'offerId': data["ID"],
      'title': data["Title"],
      'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name|| 'city'}`,
      'description': data["Description"],
      'capacity': data["Capacity"],
      'price': data['Price'],
      'discount': data['Discount'],
      'type': data['Type'],
      'dateFrom': data?.DateFrom?.split('T')?.[0],
      'dateTo': data?.DateTo?.split('T')?.[0]
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

function editOffer(offerId, type){
  router.push(`/offers/${type}/edit/${offerId}`)
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
          <template v-for="event in allEvents.value" :key="event.offerId">
            <v-col cols="4">
              <OfferListItem type="event" :offerItem="event" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(event.offerId, 'event')">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(event.offerId, 'event')">Delete
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
          <template v-for="accommodation in allAccommodations.value" :key="accommodation.offerId">
            <v-col cols="4">
              <OfferListItem type="accommodation" :offerItem="accommodation" custom>
                <template v-slot:template>
                  <v-card elevation="0" class=" d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(accommodation.offerId, 'accommodation')">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(accommodation.offerId, 'accommodation')">Delete
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
          <template v-for="activity in allActivities.value" :key="activity.offerId">
            <v-col cols="4">
              <OfferListItem type="activity" :offerItem="activity" custom>
                <template v-slot:template>
                  <v-card elevation="0" class="d-flex justify-space-between align-center" height="100">
                    <v-btn elevation="0" color="blue-grey-lighten-2" rounded
                           @click="editOffer(activity.offerId, 'activity')">Edit
                    </v-btn>
                    <v-btn color="red-lighten-2" elevation="0" rounded
                           @click="confirmDeleteOffer(activity.offerId, 'activity')">Delete
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
