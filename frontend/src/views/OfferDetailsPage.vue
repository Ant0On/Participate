<script setup>
import {computed, defineProps, onMounted, ref, toRef} from 'vue';
import {storeToRefs} from 'pinia';

import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import chipsMapper from "@/_helpers/chips";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import RoomDetail from "@/components/detail/RoomDetail.vue";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const props = defineProps({
  type: String,
  id: String,
  chatID: String
})
const priceAfterDiscount = computed(() => calculatePriceAfterDiscount(offer.price, offer?.discount))

const formFilled = ref(false)
const cardPage = ref('description')
const offerPage = ref('description')
const chosenPayment = ref(null)
const paymentType = ref([
  {name: 'paypal', url: 'https://paypal.com', img: '@/assets/img/paypal.png'},
  {name: 'credit_card', url: 'https://www.przelewy24.pl', img: '@/assets/img/credit_card.png'},
  {name: 'bitcoin', url: 'https://bitcoin.org/', img: '@/assets/img/bitcoin.png'},
])

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
  generalFacilities: ["Swimming Pool", "Gym", "Spa", "Restaurant"],
  numberOfRooms: 3,
  dateFrom: new Date().toLocaleDateString(),
  dateTo: new Date().toLocaleDateString(),
  date: new Date().toLocaleDateString(),
  skill: 'beginner',
  rooms: [
    {
      roomNumber: 1,
      roomName: 'Pierwszy',
      roomDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
      capacity: 2,
      roomFacilities: ["Swimming Pool", "Gym", "Spa", "Restaurant"],
      area: 'Zamosc'
    },
    {
      roomNumber: 2,
      roomName: 'Drugi',
      roomDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
      capacity: 2,
      roomFacilities: ["Swimming Pool", "Gym", "Spa", "Restaurant"],
      area: 'Zamosc'
    },
    {
      roomNumber: 3,
      roomName: 'Trzeci',
      roomDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
      capacity: 1,
      roomFacilities: ["Swimming Pool", "Gym", "Spa", "Restaurant"],
      area: 'Zamosc'
    },
    {
      roomNumber: 4,
      roomName: 'Czwarty',
      roomDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ",
      capacity: 3,
      roomFacilities: ["Swimming Pool", "Gym", "Spa", "Restaurant"],
      area: 'Zamosc'
    },
  ]

}
const chips = ref(chipsMapper(offer?.discount))

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
    <v-window v-model="offerPage" class="align-self-center" style="width: 60%; height: 600px;">
      <v-window-item value="description" elevation="4" class="h-100">
        <v-card class="d-flex flex-row rounded-xl h-100">
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
          <v-card class="w-50 d-flex flex-column">
            <v-card-title class="font-weight-black text-center">
              {{ offer.title }}
            </v-card-title>
            <v-card-subtitle class="d-flex align-center justify-center mb-2">

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
            <v-card-subtitle class="mx-0 d-flex ">
              <div class="font-weight-bold"
                   :class="(offer?.discount > 0)? 'text-decoration-line-through text-red-lighten-1': ''">
                {{ (type === "accommodation") ? `Price: ${offer.price} $/day` : `Price: ${offer.price} $` }}
              </div>

              <div class="font-weight-black ml-1" v-if="offer?.discount > 0">
                {{ (type === "accommodation") ? `${priceAfterDiscount} $/day` : `${priceAfterDiscount} $` }}
              </div>
            </v-card-subtitle>
            <div class="ma-4 text-subtitle-1">
              <v-chip :prepend-icon="chips.recommended.icon"
                      :color="chips.recommended.color"
                      variant="flat"
                      v-if="offer.isRecommended"
                      class="mr-1"
              >
                {{ chips.recommended.text }}
              </v-chip>
              <v-chip :prepend-icon="chips.discount.icon"
                      :color="chips.discount.color"
                      variant="flat"
                      v-if="offer?.discount > 0"
                      class="mr-1"
              >
                {{ chips.discount.text }}
              </v-chip>
              <v-chip :prepend-icon="chips?.[offer?.type].icon"
                      :color="chips?.[offer?.type].color"
                      variant="flat"
                      class="mr-1"
              >
                {{ chips?.[offer?.type].text }}
              </v-chip>
              <v-chip :prepend-icon="chips?.[offer?.skill].icon"
                      :color="chips?.[offer?.skill].color"
                      variant="flat"
                      class="mr-1"
                      v-if="type=== 'activity'"
              >
                {{ chips?.[offer?.skill].text }}
              </v-chip>
            </div>
            <v-divider></v-divider>
            <v-card-actions>
              <v-btn-toggle block rounded="lg">
                <v-btn @click="cardPage = 'description'">
                  <v-icon icon="mdi-menu"></v-icon>
                </v-btn>
                <v-btn @click="cardPage = 'info'">
                  <v-icon icon="mdi-information"></v-icon>
                </v-btn>
                <v-btn @click="cardPage = 'accommodation'">
                  <!--                <v-btn @click="cardPage = 'accommodation'" v-if="type === 'accommodation' && offer.type in ['hotel', 'hostel', 'guesthouse']">-->
                  <v-icon icon="mdi-home"></v-icon>
                </v-btn>
                <v-btn @click="cardPage = 'activity'">
                  <!--              <v-btn @click="cardPage = 'activity'" v-if="type === 'activity'">-->
                  <v-icon icon="mdi-run"></v-icon>
                </v-btn>
              </v-btn-toggle>
            </v-card-actions>

            <v-card-text v-if="cardPage === 'description'">
              {{ offer.description }}
            </v-card-text>
            <v-card-text v-else-if="cardPage === 'info'">
              <v-list class="h-100 w-100" style="overflow: hidden">
                <v-row cols="2">
                  <v-col>
                    <v-list-item
                        key="standard_price"
                        title="Standard Price"
                        :subtitle="`${offer.price} $ ${(type === 'accommodation')? ' per day': ''}`"
                    ></v-list-item>
                  </v-col>
                  <v-col>
                    <v-list-item
                        key="capacity"
                        title="Capacity"
                        :subtitle="offer.capacity"
                    ></v-list-item>
                  </v-col>
                </v-row>
                <v-row cols="2" v-if="type === 'accommodation'">
                  <v-col>
                    <v-list-item
                        key="number_of_rooms"
                        title="Number of rooms"
                        :subtitle="offer?.numberOfRooms"
                    ></v-list-item>
                  </v-col>
                  <v-col>
                    <v-list-item
                        key="general_facilities"
                        title="General Facilities"
                        :subtitle="offer?.generalFacilities.join(', ')"
                    ></v-list-item>
                  </v-col>
                </v-row>
                <v-row cols="2" v-if="type === 'event'">
                  <v-col>
                    <v-list-item
                        key="date_from"
                        title="Date From"
                        :subtitle="offer?.dateFrom"
                    ></v-list-item>
                  </v-col>
                  <v-col>
                    <v-list-item
                        key="date_to"
                        title="Date To"
                        :subtitle="offer?.dateTo"
                    ></v-list-item>
                  </v-col>
                </v-row>
                <v-row cols="2" v-if="type === 'activity'">
                  <v-col>
                    <v-list-item
                        key="date"
                        title="Date"
                        :subtitle="offer?.date"
                    ></v-list-item>
                  </v-col>
                  <v-col>
                    <v-list-item
                        key="duration"
                        title="Duration"
                        :subtitle="offer?.duration"
                    ></v-list-item>
                  </v-col>
                </v-row>
                <v-row>
                  <v-list-item
                      key="general_facilities"
                      title="General Facilities"
                      v-if="type === 'accommodation'"
                  >{{
                      offer?.generalFacilities.join(', ')
                    }}
                  </v-list-item>
                  <v-list-item
                      key="equipment"
                      title="Equipment"
                      v-if="type === 'activity'"
                  >
                    {{
                      offer?.equipment.join(', ')
                    }}
                  </v-list-item>
                </v-row>
              </v-list>
            </v-card-text>

            <v-card-text v-else-if="cardPage === 'accommodation'" class="h-100" style="overflow-y: scroll">
              <v-card-title>
                Rooms
              </v-card-title>
              <v-list style="overflow: hidden">
                <v-list-item v-for="room in offer.rooms" :key="room.roomNumber">
                  <RoomDetail :room="room"/>
                </v-list-item>
              </v-list>
            </v-card-text>

            <v-card-text v-else-if="cardPage === 'activity'">
              <v-list style="overflow: hidden;">
                <v-row>
                  <v-list-item
                      key="activity_type"
                      title="Type"
                  >
                    <template v-slot:prepend>
                      <v-icon icon="mdi-tree"></v-icon>
                    </template>
                    {{
                      offer.type
                    }}
                  </v-list-item>
                </v-row>
                <v-row>
                  <v-list-item
                      key="skill"
                      title="Skill"
                  >
                    <template v-slot:prepend>
                      <v-icon icon="mdi-bullseye"></v-icon>
                    </template>

                    {{
                      offer.skill
                    }}
                  </v-list-item>
                </v-row>
              </v-list>
            </v-card-text>

            <v-divider></v-divider>
            <v-card-actions class="align-self-end w-100">
              <v-btn
                  color="blue-darken-2 flex-grow"
                  text="Reserve"
                  block
                  border
                  @click="offerPage = 'summary'"
              ></v-btn>
            </v-card-actions>
          </v-card>
        </v-card>
      </v-window-item>
      <v-window-item value="summary">
        <v-card class="d-flex flex-row rounded-xl" :height="600">
          <v-card class="w-50 d-flex flex-column">
            <v-form>
<!--
-->
              <v-input></v-input>
            </v-form>
            <v-card-actions class="align-self-end w-100">
              <v-btn
                  color="grey-darken-2 flex-grow"
                  text="Back"
                  block
                  border
                  @click="offerPage = 'description'"
              ></v-btn>
            </v-card-actions>
          </v-card>
          <v-fade-transition>
            <v-card class="w-50 d-flex flex-column justify-space-between">
<!--            <v-card class="w-50 d-flex flex-column" v-if="formFilled">-->
              <v-card-title class="text-center">
                Choose form of payment
              </v-card-title>
              <v-card-text class="mt-4">
                <v-list @update:selected="(x) => console.log(x)">
                  <v-list-item value="paypal">
                    <template v-slot:prepend>
                      <v-img :src="require('@/assets/img/paypal.png')" :height="75" :width="75"></v-img>
                    </template>
                    Paypal
                  </v-list-item>
                  <v-list-item value="credit_card">
                    <template v-slot:prepend>
                      <v-img :src="require('@/assets/img/credit_card.png')" :height="75" :width="75"></v-img>
                    </template>
                    Credit card
                  </v-list-item>
                  <v-list-item value="bitcoin">
                    <template v-slot:prepend>
                      <v-img :src="require('@/assets/img/bitcoin.png')" :height="75" :width="75"></v-img>
                    </template>
                    Bitcoin
                  </v-list-item>
                </v-list>
              </v-card-text>
              <v-card-actions class="w-100">
                <v-btn
                    color="blue-darken-2 flex-grow"
                    text="Make reservation"
                    block
                    border
                    @click="offerPage = 'description'"
                ></v-btn>
              </v-card-actions>
            </v-card>
          </v-fade-transition>
        </v-card>
      </v-window-item>
    </v-window>

  </v-sheet>
</template>

<style scoped>

</style>