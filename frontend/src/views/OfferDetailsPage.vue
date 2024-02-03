<script setup>
import {defineProps, ref, onMounted} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferDetailSummary from "@/components/detail/OfferDetailSummary.vue";
import OfferDetailDescription from "@/components/detail/OfferDetailDescription.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import ChatButton from "@/components/detail/ChatButton.vue";
import ChatPopup from "@/components/detail/ChatPopup.vue";

const props = defineProps({
  type: String,
  id: String,
})

const showChat = ref(false);

function openChat({ type, id }) {
  // Add your logic to join the chat based on the type and id
  // For now, just toggle the showChat ref
  showChat.value = true;
}

function closeChat() {
  showChat.value = false;
}

const offer = ref({
  name: '',
  price: 0,
  location: '',
  numberOfPeople: 0,
  hostID: 0,
  offer_id: 0,
})
const hostData = ref({
  firstName: '',
  detail: '',
  imagePath: '',
})
const isDescription = ref(true)

async function getOfferDetails() {
  const response = await fetchWrapper.get(`/api/offers/${props.id}`)

  const responseData = response.data
  offer.value = {
      'hostID': responseData["host_id"],
      'location': responseData["country_name"] + ', ' + responseData["town_name"],
      'description': responseData["description"],
      'name': responseData["name"],
      'price': responseData["price"],
      'numberOfPeople': responseData["max_people"],
    }

  const response_host = await fetchWrapper.get(`/api/host/${offer.value.hostID}`)

  console.log('ala', response_host)

  hostData.value = {
    'firstName': response_host["FirstName"],
    'detail': response_host["Description"],
    'imagePath': response_host["ImagePath"]
  }

  console.log('kut', hostData.value)

}

onMounted(async () => {
  await getOfferDetails();
});

</script>

<template>
  <div class="event_page">
    <NavBar :currentPage="type"/>
    <div class="item_detail">
      <OfferDetailDescription :type="type" :name="offer.name" :price="offer.price" :location="offer.location"
                              :numberOfPeople="offer.numberOfPeople" :host_first_name="hostData.firstName" :offer_id="id"
                             :host_detail="hostData.detail" :imagePath="hostData.imagePath"
                              v-if="isDescription" @move-to-summary="isDescription = !isDescription"/>
      <OfferDetailSummary :price="offer.price" :image="require(`@/../images/offers/6.jpeg`)" :id="id" v-else/>
      <ChatButton :type="type" :id="id" @join-chat="openChat"/>
      <ChatPopup v-if="showChat" @close-chat="closeChat"/>
    </div>
  </div>
</template>

<style scoped>
div.item_detail{
  margin: 5% 5% 1% 5%;
  background-color: white;
}
</style>