<script setup>
import {computed, defineProps, onMounted, ref, toRef} from 'vue';

import NavBar from "@/components/nav/NavBar.vue";
import OfferDetailSummary from "@/components/detail/OfferDetailSummary.vue";
import OfferDetailDescription from "@/components/detail/OfferDetailDescription.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import ChatButton from "@/components/detail/ChatButton.vue";
import ChatPopup from "@/components/detail/ChatPopup.vue";
import {useAuthStore} from "@/stores/auth.store";

const auth = useAuthStore();
const user = auth.user;

const props = defineProps({
  type: String,
  id: String,
  chatID: String
})
const isChatAlreadyCreated = ref(false)
const showChat = ref(false);

function createChatHost() {
  if (isChatAlreadyCreated.value) {
    openChatCustomer()
  }
  if (user.Role === 'host' && offer.value.hostID !== user.ID) {
    console.error('You are not the owner of this offer!')
  } else {
    fetchWrapper.post(`/api/host/${props.id}/chat/create`).then(() => {
          isChatAlreadyCreated.value = true;
          showChat.value = true
          window.location.reload();
        }
    ).catch()
  }
}

const email = ref(user.Email);
const userID = ref(user.ID);
const chatID = toRef(props.chatID)

function openChatCustomer() {
  showChat.value = true
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
const backgroundImage = computed(() => `url(${require(`@/../images/offers/${props.id}/${props.id}_0.jpeg`)})`);

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

  hostData.value = {
    'firstName': response_host["FirstName"],
    'detail': response_host["Description"],
    'imagePath': response_host["ImagePath"]
  }
}

async function doesChatExist() {
  return fetchWrapper.get(`/api/chat/offer/${props.id}`).then((response) => {
    if (!response) {
      isChatAlreadyCreated.value = false
    } else {
      chatID.value = response.data["ID"]
      isChatAlreadyCreated.value = true
    }
  }).catch(error => {
    console.log(error)
  })
}

onMounted(async () => {
  await getOfferDetails();
  await doesChatExist();
});

</script>

<template>
  <div class="event_page">
    <NavBar :currentPage="type"/>
    <div class="item_detail">
      <OfferDetailDescription :type="type" :name="offer.name" :price="offer.price" :location="offer.location"
                              :numberOfPeople="offer.numberOfPeople" :host_first_name="hostData.firstName"
                              :offer_id="id"
                              :host_detail="hostData.detail" :imagePath="hostData.imagePath"
                              v-if="isDescription" @move-to-summary="isDescription = !isDescription"/>
      <OfferDetailSummary :price="offer.price" :id="id" v-else/>
      <ChatButton v-if="user.Role === 'host' && type === 'events'" :is-host="true"
                  :is-chat-already-created=!isChatAlreadyCreated @click="createChatHost"/>
      <ChatButton v-else-if="user.Role === 'customer' && type === 'events'" :is-host="false"
                  :is-chat-already-created=isChatAlreadyCreated @join-chat="openChatCustomer"/>

      <ChatPopup v-if="showChat" :email="email" :userID="userID" :offerID="id" :chatID="chatID"
                 @close-chat="closeChat"/>
    </div>
  </div>
</template>

<style scoped>
div.item_detail {
  margin: 5% 5% 1% 5%;
  background-color: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(5px);
  border-radius: 10px;
}

div.item_detail:before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: v-bind(backgroundImage) center/cover no-repeat;
  opacity: 0.5;
  z-index: -1;
  border-radius: 10px;
}
</style>