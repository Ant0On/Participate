<script setup>
import {ref, onMounted} from 'vue';
import {useAuthStore} from "@/stores/auth.store";
import SwitchListPage from "@/components/common/SwitchListPage.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import MyOffersItem from "@/components/account/my_offers/MyOffersItem.vue";


const auth = useAuthStore();
const user = auth.user;
const pageSize = 5;

const allMyOffers = ref([])
const myOffers = ref([])
const currentPage = ref(1)
const maxPage = ref(1)

async function getMyOffers() {
  fetchWrapper.get(`/api/host/${user.ID}/offers`)
      .then((response) => {
        const responseData = response.data
        console.log(responseData)
        allMyOffers.value = responseData.map((data) => {
          return {
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
        myOffers.value = allMyOffers.value.slice(0, pageSize);
        maxPage.value = Math.floor(allMyOffers.value.length / pageSize) + 1
      })
      .catch((error) =>{

      })
}


function pageBack() {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
    myOffers.value = allMyOffers.value.slice((currentPage.value - 1) *pageSize, currentPage.value * pageSize )
  }
}

function pageFroward() {
  if (currentPage.value < maxPage.value) {
    currentPage.value += 1;
    myOffers.value = allMyOffers.value.slice((currentPage.value - 1) *pageSize, currentPage.value * pageSize )

  }
}

onMounted(async () => await getMyOffers())
</script>

<template>

  <div class="my_offers">
    <p>My offers</p>
    <div class="items_list">
      <div class="my_offer_items">
        <MyOffersItem v-for="myOffer in myOffers" :name="myOffer.name" :offer-type="myOffer.offerType" discount="10%"
                          :date-from="myOffer.dateFrom" :date-to="myOffer.dateTo" :id="myOffer.offer_id"
                          :with-animals="myOffer.withAnimals" :offer-id="myOffer.offerId"/>
      </div>
      <div class="navigation">
        <SwitchListPage v-if="maxPage !== 1" :currentPage="currentPage" :maxPage="maxPage" @page-back="pageBack"
                        @page-forward="pageFroward"/>
      </div>
    </div>
  </div>

</template>

<style scoped>
div.my_offers{
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
div.my_offer_items{
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