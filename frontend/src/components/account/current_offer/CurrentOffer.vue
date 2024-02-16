<script setup>
import {ref, onMounted, reactive} from 'vue';
import {useAuthStore} from "@/stores/auth.store";
import SwitchListPage from "@/components/common/SwitchListPage.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import CurrentOfferItem from "@/components/account/current_offer/CurrentOfferItem.vue";

const auth = useAuthStore();
const user = auth.user;
const pageSize = 5;

const allCurrentOffers = ref([])
const currentOffers = ref([])
const currentPage = ref(1)
const maxPage = ref(1)

const errors = reactive({
  apiError: ""
})

async function getCurrentOffers() {
  fetchWrapper.get(`/api/host/${user.ID}/reservations/pending`)
      .then((response) => {
        const responseData = response.data

        allCurrentOffers.value = responseData.map((data) => {
          return {
            'reservationID': data["reservation_id"],
            'location': data["country_name"] + ', ' + data["town_name"],
            'name': data["name"],
            'price': data["price"],
            'dateFrom': data['date_from'],
            'dateTo': data['date_to'],
            'offerType': data['offer_type'],
            'withAnimals': data['is_animal_friendly'],
            'offerId': data['offer_id']
          }
        })
        currentOffers.value = allCurrentOffers.value.slice(0, pageSize);
        maxPage.value = Math.floor(allCurrentOffers.value.length / pageSize) + 1
      })
      .catch((error) =>{
        errors.apiError = "Failed to fetch current offers. Please try again later. " + error.message;
      })
}

function pageBack() {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
    currentOffers.value = allCurrentOffers.value.slice((currentPage.value - 1) *pageSize, currentPage.value * pageSize )
  }
}

function pageFroward() {
  if (currentPage.value < maxPage.value) {
    currentPage.value += 1;
    currentOffers.value = allCurrentOffers.value.slice((currentPage.value - 1) *pageSize, currentPage.value * pageSize )

  }
}

onMounted(async () => await getCurrentOffers())
</script>

<template>

  <div class="current_offers">
    <p>Current offers</p>
    <div class="items_list">
      <div class="offer_items">
        <CurrentOfferItem v-for="currentOffer in currentOffers" :name="currentOffer.name" :offer-type="currentOffer.offerType"
                     :date-from="currentOffer.dateFrom" :date-to="currentOffer.dateTo" :id="currentOffer.reservationID"
                     :with-animals="currentOffer.withAnimals" :offer-id="currentOffer.offerId"/>
      </div>
      <div class="navigation">
        <SwitchListPage v-if="maxPage !== 1" :currentPage="currentPage" :maxPage="maxPage" @page-back="pageBack"
                        @page-forward="pageFroward"/>
      </div>
      <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    </div>
  </div>

</template>

<style scoped>
div.current_offers{
  display: flex;
  flex-direction: column;
  height: max(500px, 80%);
  flex-grow: 1;
}
div.items_list {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: space-between;
  padding-top: 2%;

}
div.offer_items{
  display: flex;
  flex-direction: column;
  padding-left: 2%;
}
div.navigation{
  padding-top: 2%;
}
p{
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.4rem;
  padding-bottom: 2%;
  font-weight: 400;
  align-self: center;
}

div.errors {
  font-family: "Sarabun", Helvetica;
  color: red;
  padding-top: 2%;
}

</style>