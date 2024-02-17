<script setup>
import {onMounted, ref} from 'vue';
import {useAuthStore} from "@/stores/auth.store";
import SwitchListPage from "@/components/common/SwitchListPage.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import HistoryItem from "@/components/account/history/HistoryItem.vue";
import RateOfferModal from "@/components/account/history/rate_offer/RateOfferModal.vue";

const auth = useAuthStore();
const user = auth.user;
const pageSize = 5;

const allHistoryItems = ref([])
const historyItems = ref([])
const currentPage = ref(1)
const maxPage = ref(1)
const offerToGrade = ref([])

const error = ref(null);

async function getHistoryItems() {

  await fetchWrapper.get(`/api/customer/${user.ID}/reservations/history`)
      .then((response) => {
        const responseData = response.data
        allHistoryItems.value = responseData.map((data) => {
          return {
            'location': data["country_name"] + ', ' + data["town_name"],
            'name': data["name"],
            'price': data["price"],
            'dateFrom': data['date_from'],
            'dateTo': data['date_to'],
            'offerType': data['offer_type'],
            'withAnimals': data['is_animal_friendly'],
            'reservationState': data['reservation_state'],
            'reservationId': data['reservation_id'],
            'offerId': data["offer_id"],
            'gradeId': data["grade_id"]
          }
        })
        historyItems.value = allHistoryItems.value.slice(0, pageSize);
        maxPage.value = Math.floor(allHistoryItems.value.length / pageSize) + 1
        gradeOffers()
      })
      .catch((error) => {
        error.value = 'Failed to fetch history items. Please try again later. ' + error.message;
      })
}
function gradeOffers(){
  const allFinishedItems =  allHistoryItems.value.filter((item) => item.gradeId === 0 && item.reservationState === 'finished')
  offerToGrade.value =  JSON.parse(JSON.stringify(allFinishedItems.map((item) => JSON.parse(JSON.stringify(item)))))
}
function pageBack() {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
    historyItems.value = allHistoryItems.value.slice((currentPage.value - 1) * pageSize, currentPage.value * pageSize)
  }
}

function pageFroward() {
  if (currentPage.value < maxPage.value) {
    currentPage.value += 1;
    historyItems.value = allHistoryItems.value.slice((currentPage.value - 1) * pageSize, currentPage.value * pageSize)
  }
}

onMounted(async () => {
  await getHistoryItems()
})
</script>

<template>

  <div class="account_history">
    <p>Previously booked offers</p>
    <div class="items_list">
      <div v-if="error" class="error-message">{{ error }}</div>
      <div class="history_items">
        <HistoryItem v-for="historyItem in historyItems" :name="historyItem.name" :offer-type="historyItem.offerType"
                     :date-from="historyItem.dateFrom" :date-to="historyItem.dateTo"
                     :with-animals="historyItem.withAnimals" :offer-id="historyItem.offerId"
                     :reservation-state="historyItem.reservationState"/>
      </div>
      <div class="navigation">
        <SwitchListPage v-if="maxPage !== 1" :currentPage="currentPage" :maxPage="maxPage" @page-back="pageBack"
                        @page-forward="pageFroward"/>
      </div>
      <RateOfferModal v-if="offerToGrade.length > 0" :offers-to-grade="offerToGrade"/>
    </div>
  </div>

</template>

<style scoped>
div.account_history {
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

div.history_items {
  display: flex;
  flex-direction: column;
  padding-left: 2%;
}

div.navigation {
  padding-top: 2%;
}

p {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.4rem;
  padding-bottom: 2%;
  font-weight: 400;
  align-self: center;
}

div.error-message {
  color: red;
  padding: 1%;
  text-align: center;
}

</style>