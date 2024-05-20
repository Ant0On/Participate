<script setup>
import {onMounted, ref, reactive} from 'vue';
import {useAuthStore} from "@/stores/auth.store";
import SwitchListPage from "@/components/common/SwitchListPage.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import MyOffersItem from "@/components/account/my_offers/MyOffersItem.vue";


const auth = useAuthStore();
const user = auth.user;
const pageSize = 5;
const errors = reactive({
  apiError: ""
})

const allEvents = ref([])
const myEvents = ref([])
const currentPage = ref(1)
const maxPage = ref(1)

async function getMyOffers() {
  console.log('Starting request to fetch offers'); // Log before the request
  await fetchWrapper.get(`/api/host/event/${user.ID}/offers`)
      .then((response) => {
        console.log('Request successful:', response); // Log the response
        console.log('Response Data:', response.data);
        try {
          const responseData = response.data;
          allEvents.value = responseData.map((data) => {
            // Log each data item being mapped
            console.log('Mapping data:', data);
            return {
              'location': data["country_name"] + ', ' + data["town_name"],
              'title': data["title"],
              'offerID': data['offer_id'],
              'discount': data['discount'],
              'capacity': data['capacity'],
              'eventType': data['event_type'],
              'offerType': 'event'
            };
          });
          console.log('HELLO:');
          console.log('Events:', allEvents.value);
          myEvents.value = allEvents.value.slice(0, pageSize);
          maxPage.value = Math.floor(allEvents.value.length / pageSize) + 1;
        } catch (error) {
          console.error('Error during mapping or assignment:', error);
        }
      })
      .catch((error) => {
        console.error('Request failed:', error); // Log the error
        errors.apiError = "Default Error";
      });
}


function pageBack() {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
    myEvents.value = allEvents.value.slice((currentPage.value - 1) * pageSize, currentPage.value * pageSize)
  }
}

function pageFroward() {
  if (currentPage.value < maxPage.value) {
    currentPage.value += 1;
    myEvents.value = allEvents.value.slice((currentPage.value - 1) * pageSize, currentPage.value * pageSize)

  }
}

onMounted(async () => await getMyOffers())
</script>

<template>

  <div class="my_offers">
    <p>My offers</p>
    <div class="items_list">
      <div class="my_offer_items">
        <MyOffersItem v-for="myOffer in myEvents" :title="myOffer.title" :offer-type="myOffer.offerType"
                      :discount="myOffer.discount" :offerType="myOffer.offerType" :offerID="myOffer.offerID"/>
      </div>
      <div class="navigation">
        <SwitchListPage v-if="maxPage !== 1" :currentPage="currentPage" :maxPage="maxPage" @page-back="pageBack"
                        @page-forward="pageFroward"/>
      </div>
    </div>
  </div>

</template>

<style scoped>
div.my_offers {
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

div.my_offer_items {
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

</style>