<script setup>
import {ref, onMounted} from 'vue';
import {useAuthStore} from "@/stores/auth.store";
import SwitchListPage from "@/components/common/SwitchListPage.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import HistoryItem from "@/components/account/history/HistoryItem.vue";

const auth = useAuthStore();
const user = auth.user;
const pageSize = 5;

const allHistoryItems = ref([])
const historyItems = ref([])

async function getHistoryItems() {
  fetchWrapper.get(`/api/reservation/customer/${user.ID}/finished`)
      .then((response)=>{
        const responseData = response.data
        allHistoryItems.value = responseData.map((data) => {
          return {
            'location': data["country_name"] + ', ' + data["town_name"],
            'name': data["name"],
            'price': data["price"],
            'dateFrom': data['date_from'],
            'dateTo': data['date_to'],
            'offerType': data['offer_type'],
            'withAnimals': data['is_animal_friendly']
          }
        })
        historyItems.value = allHistoryItems.value.slice(0, pageSize);
      })
      .catch((error)=>{
      })
}

let currentPage = 1;

let maxPage = Math.floor(allHistoryItems.value.length / pageSize) + 1

function pageBack() {
  if (currentPage > 1) {
    currentPage -= 1;
    historyItems.value = allHistoryItems.value.slice((currentPage - 1) *pageSize, currentPage * pageSize )
  }
}

function pageFroward() {
  if (currentPage < maxPage) {
    currentPage += 1;
    historyItems.value = allHistoryItems.value.slice((currentPage - 1) *pageSize, currentPage * pageSize )

  }
}

onMounted(async () => getHistoryItems())
</script>

<template>

  <div class="account_history">
    <p>Previously booked offers</p>
    <div class="items_list">
      <div class="history_items">
        <HistoryItem v-for="historyItem in historyItems" :name="historyItem.name" :offer-type="historyItem.offerType"
                     :date-from="historyItem.dateFrom" :date-to="historyItem.dateTo"
                     :with-animals="historyItem.withAnimals"/>
      </div>
      <div class="navigation">
        <SwitchListPage v-if="maxPage !== 1" :currentPage="currentPage" :maxPage="maxPage" @page-back="pageBack"
                        @page-forward="pageFroward"/>
      </div>
    </div>
  </div>

</template>

<style scoped>
div.account_history{
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
div.history_items{
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

</style>