<script setup>
import {computed, defineProps, onMounted, ref, toRef} from 'vue';
import {storeToRefs} from 'pinia';

import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const reserve = () => {
  console.log("reserve")
  step.value++;
}

const props = defineProps({
  type: String,
  id: String,
  chatID: String
})

const step = ref(1)
const typeLink = {
  'event': 'events',
  'accommodation': 'accommodations',
  'activity': 'activities',
  'recommended': 'recommended'
}

const offer = {
  offerId: 1,
  title: "Wakacyjne pierdolenie",
  location: "Zamosc, Polska",
  description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
  capacity: 10,
  isRecommended: true,
  discount: 10,
  price: 30,
  type: "festival",
  rating: 4.5,
  duration: 10,

}

const image = computed(() => {
  try {
    const image = require(`@/../images/offers/${props.type}/${offer.offerId}/${offer.offerId}_0.jpeg`)
    return image
  } catch {
    return undefined
  }
})


const isChatAlreadyCreated = ref(false)
const showChat = ref(false);

function createChatHost() {
  if (isChatAlreadyCreated.value) {
    openChatCustomer()
  }
  if (user.Role === 'host' && offer.value.userID !== user.ID) {
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

const chatID = toRef(props.chatID)

function openChatCustomer() {
  showChat.value = true
}

function closeChat() {
  showChat.value = false;
}

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
    'userID': responseData["user_id"],
    'location': responseData["country_name"] + ', ' + responseData["town_name"],
    'description': responseData["description"],
    'name': responseData["name"],
    'price': responseData["price"],
    'numberOfPeople': responseData["max_people"],
  }

  const response_host = await fetchWrapper.get(`/api/host/${offer.value.userID}`)

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
  // await getOfferDetails();
  // await doesChatExist();
});

</script>

<template>
  <v-sheet class="d-flex flex-column">
    <v-breadcrumbs>
      <v-breadcrumbs-item :to="`/${typeLink[type]}`"
                          :title="type[0].toUpperCase() + type.slice(1)"></v-breadcrumbs-item>
      <v-breadcrumbs-divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </v-breadcrumbs-divider>
      <v-breadcrumbs-item :title="offer.title"></v-breadcrumbs-item>
    </v-breadcrumbs>
    <v-card class="d-flex flex-row align-self-center h-100 rounded-xl" style="width: 60%">
      <v-card class="w-50">
        <v-carousel v-if="image"
                    :show-arrows="[image].length > 1"
                    :hide-delimiters="[image].length === 1"
                    height="600px"
                    cycle
        >
          <v-carousel-item
              :src="image"
              style="min-height: 100%"
              cover
          >
          </v-carousel-item>
        </v-carousel>
        <v-img v-else
               :src="require(`@/assets/img/image_placeholder.png`)"
               style="min-height: 100%"
               cover
        ></v-img>
      </v-card>
      <v-card class="w-50">
        <v-card-title class="font-weight-black text-center">
          {{ offer.title }}
        </v-card-title>
        <v-card-subtitle class="d-flex align-center justify-center">

          <span class="me-1">
            {{ offer.location }}
          </span>
          <v-spacer></v-spacer>

          <v-rating
              :model-value="offer.rating"
              color="amber"
              density="compact"
              size="small"
              half-increments
              readonly
              class="me-1"
              v-if="type === 'accommodation' "
          ></v-rating>
          <span class="text-grey me-1" v-if="type === 'accommodation' ">
              {{ offer.rating }}
            </span>
        </v-card-subtitle>

        <v-card-text>

        </v-card-text>

        <v-card-actions>
          <v-btn
              color="blue-darken-2"
              text="Reserve"
              block
              border
              @click="reserve"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </v-card>
  </v-sheet>
</template>

<style scoped>

</style>