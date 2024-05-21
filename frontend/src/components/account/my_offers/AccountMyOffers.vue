<script setup>
import { onMounted, ref, reactive } from 'vue';
import { useAuthStore } from "@/stores/auth.store";
import fetchPaginatedData from "@/_helpers/fetchPaginatedData";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { fetchWrapper } from "@/_helpers/fetch-wrapper";

const auth = useAuthStore();
const user = auth.user;
const errors = reactive({
  apiError: ""
});

const allEvents = ref([]);
const allActivities = ref([]);
const allAccommodations = ref([]);
const confirmDelete = ref(false);
const offerToDelete = ref(null);

const discountDialog = ref(false);
const discountValue = ref(0);
const priceDialog = ref(false);
const priceValue = ref(0);
const offerToModify = ref(null);

async function getMyOffers(responseData) {
  try {
    return responseData.map((data) => {
      return {
        location: data["country_name"] + ', ' + data["town_name"],
        title: data["title"],
        offerID: data['offer_id'],
        discount: data['discount'],
        capacity: data['capacity'],
        price: data['price'] !== undefined ? data['price'] : data['price_per_day']
      };
    });
  } catch (error) {
    console.error('Error during mapping or assignment:', error);
  }
}

const eventsGenerator = fetchPaginatedData(`/api/host/event/${user.ID}/offers`, getMyOffers);
const activitiesGenerator = fetchPaginatedData(`/api/host/activity/${user.ID}/offers`, getMyOffers);
const accommodationsGenerator = fetchPaginatedData(`/api/host/accommodation/${user.ID}/offers`, getMyOffers);

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

async function setDiscount(offerID, offerType) {
  offerToModify.value = { offerID, offerType };
  discountDialog.value = true;
}

async function changePrice(offerID, offerType) {
  offerToModify.value = { offerID, offerType };
  priceDialog.value = true;
}

async function applyDiscount() {
  try {
    const { offerID, offerType } = offerToModify.value;
    const apiUrl = `/api/host/${offerType}/discount/${offerID}`;
    await fetchWrapper.put(apiUrl, { discount: Number(discountValue.value) });

    const offerListMap = {
      event: allEvents,
      activity: allActivities,
      accommodation: allAccommodations
    };
    const offer = offerListMap[offerType].value.value.find(offer => offer.offerID === offerID);
    if (offer) {
      offer.discount = discountValue.value;
    }
  } catch (error) {
    console.error('Error setting discount:', error);
  } finally {
    discountDialog.value = false;
  }
}

async function applyPrice() {
  try {
    const { offerID, offerType } = offerToModify.value;
    const apiUrl = `/api/host/${offerType}/price/${offerID}`;
    await fetchWrapper.put(apiUrl, { price: Number(priceValue.value) });

    const offerListMap = {
      event: allEvents,
      activity: allActivities,
      accommodation: allAccommodations
    };
    const offer = offerListMap[offerType].value.value.find(offer => offer.offerID === offerID);
    if (offer) {
      offer.price = priceValue.value;
    }
  } catch (error) {
    console.error('Error changing price:', error);
  } finally {
    priceDialog.value = false;
  }
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
  offerToDelete.value = { offerID, offerType };
  confirmDelete.value = true;
}

function handleDeleteConfirmation(result) {
  if (result) {
    deleteOffer(offerToDelete.value.offerID, offerToDelete.value.offerType);
  }
  confirmDelete.value = false;
  offerToDelete.value = null;
}

function editOffer() {
  // TODO
}
</script>

<template>
  <div class="event_page">
    <div>
      <v-infinite-scroll
          :items="allEvents"
          :onLoad="(done) => loadOffers(eventsGenerator, allEvents, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Events</p>
        <v-row class="w-100">
          <template v-for="event in allEvents.value" :key="event.offerID">
            <v-col cols="4">
              <OfferListItem type="event" :offerItem="event">
                <template #actions="{ index }">
                  <v-btn v-if="index === 1" @click="setDiscount(event.offerID, 'event')">Set discount</v-btn>
                  <v-btn v-if="index === 1" @click="changePrice(event.offerID, 'event')">Change price</v-btn>
                  <v-btn v-if="index === 0" @click="editOffer(event.offerID)">Edit</v-btn>
                  <v-btn v-if="index === 0" @click="confirmDeleteOffer(event.offerID, 'event')">Delete</v-btn>
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
          :onLoad="(done) => loadOffers(accommodationsGenerator, allAccommodations, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Accommodations</p>
        <v-row class="w-100">
          <template v-for="accommodation in allAccommodations.value" :key="accommodation.offerID">
            <v-col cols="4">
              <OfferListItem type="accommodation" :offerItem="accommodation">
                <template #actions="{ index }">
                  <v-btn v-if="index === 1" @click="setDiscount(accommodation.offerID, 'accommodation')">Set discount</v-btn>
                  <v-btn v-if="index === 1" @click="changePrice(accommodation.offerID, 'accommodation')">Change price</v-btn>
                  <v-btn v-if="index === 0" @click="editOffer(accommodation.offerID)">Edit</v-btn>
                  <v-btn v-if="index === 0" @click="confirmDeleteOffer(accommodation.offerID, 'accommodation')">Delete</v-btn>
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
          :onLoad="(done) => loadOffers(activitiesGenerator, allActivities, done)"
          empty-text="Currently there are no offers to display!"
          mode="manual"
          class="w-100"
      >
        <p class="text-center">Activities</p>
        <v-row class="w-100">
          <template v-for="activity in allActivities.value" :key="activity.offerID">
            <v-col cols="4">
              <OfferListItem type="activity" :offerItem="activity">
                <template #actions="{ index }">
                  <v-btn v-if="index === 1" @click="setDiscount(activity.offerID, 'activity')">Set discount</v-btn>
                  <v-btn v-if="index === 1" @click="changePrice(activity.offerID, 'activity')">Change price</v-btn>
                  <v-btn v-if="index === 0" @click="editOffer(activity.offerID)">Edit</v-btn>
                  <v-btn v-if="index === 0" @click="confirmDeleteOffer(activity.offerID, 'activity')">Delete</v-btn>
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
    <v-dialog v-model="discountDialog" max-width="400">
      <v-card>
        <v-card-title class="headline">Set Discount</v-card-title>
        <v-card-text>
          <v-text-field v-model="discountValue" label="Discount Percentage" type="number" min="0" max="100"></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" text @click="discountDialog = false">Cancel</v-btn>
          <v-btn color="primary" @click="applyDiscount">Apply</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="priceDialog" max-width="400">
      <v-card>
        <v-card-title class="headline">Change Price</v-card-title>
        <v-card-text>
          <v-text-field v-model="priceValue" label="New Price" type="number" min="0"></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" text @click="priceDialog = false">Cancel</v-btn>
          <v-btn color="primary" @click="applyPrice">Apply</v-btn>
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
