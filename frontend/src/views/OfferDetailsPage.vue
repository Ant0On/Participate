<script setup>
import {computed, defineProps, onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import chipsMapper from "@/_helpers/chips";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import RoomDetail from "@/components/detail/RoomDetail.vue";
import {router} from "@/router";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)
const offer = ref(null)
const props = defineProps({
  type: String,
  id: String
})

const hostData = ref()

const userImageSource = computed(() => {
  return require(`@/../images/users/${hostData.value.hostId}.jpeg`);
})

const defaultImageSource = computed(() => {
  return require(`@/../images/users/default_image.png`);
})

const isUserImage = computed(() => {
  try {
    require(`@/../images/users/${hostData.value.hostId}.jpeg`);
    return true;
  } catch {
    return false;
  }
});

const priceAfterDiscount = computed(() => calculatePriceAfterDiscount(offer.value?.price, offer.value?.discount))
const cardPage = ref('description')
const offerPage = ref('description')
const chosenPayment = ref(null)
const paymentType = ref({
  'paypal': {id: 1, url: 'https://paypal.com', img: '@/assets/img/paypal.png'},
  'credit_card': {id: 2, url: 'https://www.przelewy24.pl', img: '@/assets/img/credit_card.png'},
  'bitcoin': {id: 3, url: 'https://bitcoin.org/', img: '@/assets/img/bitcoin.png'},
})
const reservationCreated = ref(null)
const reservation = ref({
  dateFrom: null,
  dateTo: null,
  animal: null,
  room: null,
  numberOfPeople: null,
})
const openRoomDialog = ref(false)
const typeLink = {
  'event': 'events',
  'accommodation': 'accommodations',
  'activity': 'activities',
}
const choosePaymentAlert = ref(false)

const form = ref()

async function makeReservation() {
  choosePaymentAlert.value = false;
  const {valid} = await form.value.validate()
  if (!valid) {
    return
  }
  if (!chosenPayment.value) {
    choosePaymentAlert.value = true;
    return
  }
  const reservationBody = {
    'reservation_state': 'pending',
    'number_of_people': Number(reservation.value.numberOfPeople),
    'user_id': Number(user.value.ID),
    'payment_id': paymentType.value[chosenPayment.value].id
  }
  if (props.type === "event") {
    reservationBody.date = new Date(reservation.value.dateTo).toISOString()
    reservationBody.event_id = Number(props.id)
  } else if (props.type === "activity") {
    reservationBody.date = new Date(reservation.value.dateTo).toISOString()
    reservationBody.activity_id = Number(props.id)
  } else if (props.type === "accommodation") {
    reservationBody.date_from = new Date(reservation.value.dateFrom).toISOString()
    reservationBody.date_to = new Date(reservation.value.dateTo).toISOString()
    if (['hotel', 'hostel', 'guesthouse'].includes(offer.value.type)) {
      reservationBody.room_id = Number(reservation.value.room)

    } else {
      reservationBody.accommodation_id = Number(props.id)
    }
  }

  fetchWrapper.post(`/api/reservation/${(props.type === 'accommodation'
      && ['hotel', 'hostel', 'guesthouse'].includes(offer.value.type) ? 'room' : props.type)}/add`,
      reservationBody).then((resp) => {
    router.push('/');
    window.open(paymentType.value[chosenPayment.value].url, '_blank');
  }).catch((error) => {

  })
}

const chips = computed(() => chipsMapper(offer.value?.discount))
const formatDecimalPlaces = (num) => (Math.round(num * 100) / 100).toFixed(2)

const image = computed(() => {
  const images = []
  let number = 0
  try {
    while (true) {
      const image = require(`@/../images/offers/${props.type}/${offer.value?.offerId}/${offer.value?.offerId}_${number}.jpeg`)
      images.push(image)
      number++
    }
  } catch (error) {
    return images
  }
})

async function getOfferDetails() {
  const response = await fetchWrapper.get(`/api/offers/${props.type}/${props.id}`)
  if (response?.data) {
    const data = response.data[0]
    if (props.type === "accommodation") {
      return {
        'offerId': data["ID"],
        'hostId': data["UserID"],
        'title': data["Title"],
        'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name || 'city'}`,
        'description': data["Description"],
        'capacity': data["Capacity"],
        'price': data['PricePerDay'],
        'discount': data['Discount'],
        'type': data['Type'],
        'animal_friendly': data['IsAnimalFriendly'],
        'rating': data['RatingAvg'] || 0,
        'ratingCount': data['RatingCount'] || 0,
        'numberOfRooms': data?.NumberOfRooms,
        'rooms': data?.Rooms?.map((room) => {
          return {
            area: Number(room.area),
            capacity: Number(room.capacity),
            description: room.description,
            roomFacilities: room.room_facilities.map((facility) => facility.Name),
            roomName: room.name,
            roomNumber: Number(room.number),
            ID: room.ID,
          }
        }),
        'generalFacilities': data?.GeneralFacilities?.map((generalFacility) => generalFacility.Name)

      }

    } else if (props.type === "event") {
      return {
        'offerId': data["ID"],
        'hostId': data["UserID"],
        'title': data["Title"],
        'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name || 'city'}`,
        'description': data["Description"],
        'capacity': data["Capacity"],
        'price': data['Price'],
        'discount': data['Discount'],
        'type': data['Type'],
        'dateFrom': data?.DateFrom?.split('T')?.[0],
        'dateTo': data?.DateTo?.split('T')?.[0]
      }
    } else if (props.type === "activity") {
      return {
        'offerId': data["ID"],
        'hostId': data["UserID"],
        'title': data["Title"],
        'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name || 'city'}`,
        'description': data["Description"],
        'capacity': data["Capacity"],
        'price': data['Price'],
        'discount': data['Discount'],
        'skill': data['Skill'],
        'type': data['Type'],
        'duration': data['Duration'],
        'date': data?.Date?.split('T')?.[0],
        'equipment': data?.Equipment?.map((equipment) => equipment?.Name),
        'rating': data['RatingAvg'] || 0,
        'ratingCount': data['RatingCount'] || 0,
      }
    }
  }
}

async function getHostDetails(hostId) {
  const response = await fetchWrapper.get(`/api/host/${hostId}`)
  if (response) {
    return {
      'hostId': response["ID"],
      'description': response["Description"],
      'firstName': response["FirstName"]
    }
  }
}

onMounted(async () => {
  offer.value = await getOfferDetails();
  hostData.value = await getHostDetails(offer.value.hostId)
});

</script>

<template>
  <v-progress-circular v-if="!offer"
                       color="primary"
                       indeterminate
  ></v-progress-circular>
  <v-sheet v-else class="d-flex flex-column">
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
                        :show-arrows="[...image].length > 1"
                        :hide-delimiters="[...image].length === 1"
                        height="600px"
                        cycle
            >
              <v-carousel-item
                  v-for="img in image"
                  :src="img"
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
                  v-if="type === 'accommodation' || type==='activity' "
              ></v-rating>
              <span class="text-grey me-1" v-if="type === 'accommodation' || type==='activity' ">
              {{ `${formatDecimalPlaces(offer?.rating || 0)} (${offer.ratingCount})` }}
            </span>
            </v-card-subtitle>
            <v-card-subtitle class="mx-0 d-flex ">
              <div class="font-weight-bold"
                   :class="(offer?.discount > 0)? 'text-decoration-line-through text-red-lighten-1': ''">
                {{
                  (type === "accommodation") ? `Price: ${formatDecimalPlaces(offer.price)} $/day` : `Price: ${formatDecimalPlaces(offer.price)} $`
                }}
              </div>

              <div class="font-weight-black ml-1" v-if="offer?.discount > 0">
                {{
                  (type === "accommodation") ? `${formatDecimalPlaces(priceAfterDiscount)} $/day` : `${formatDecimalPlaces(priceAfterDiscount)} $`
                }}
              </div>
            </v-card-subtitle>
            <div class="ma-4 text-subtitle-1">
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
                <v-btn @click="cardPage = 'host_info'">
                  <v-icon icon="mdi-account"></v-icon>
                </v-btn>
                <v-btn @click="cardPage = 'accommodation'"
                       v-if="type === 'accommodation' && ['hotel', 'hostel', 'guesthouse'].includes(offer.type)">
                  <v-icon icon="mdi-home"></v-icon>
                </v-btn>
                <v-btn @click="cardPage = 'activity'" v-if="type === 'activity'">
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
                        :subtitle="offer?.numberOfRooms || 1"
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
                      offer?.generalFacilities?.join(', ') || 'None'
                    }}
                  </v-list-item>
                  <v-list-item
                      key="equipment"
                      title="Equipment"
                      v-if="type === 'activity'"
                  >
                    {{
                      offer?.equipment?.join(', ') || 'None'
                    }}
                  </v-list-item>
                </v-row>
              </v-list>
            </v-card-text>

            <v-card-text v-else-if="cardPage === 'host_info'" class="h-100" style="overflow-y: scroll">
              <v-row cols="1">
                <v-col>
                  <v-card>
                    <v-avatar v-if="isUserImage"
                              size="150"
                              text="HELLO"
                              density="comfortable"
                              :image="userImageSource"
                    ></v-avatar>
                    <v-avatar v-else
                              size="150"
                              :text="hostData.firstName"
                              density="comfortable"
                              :image="defaultImageSource"
                    ></v-avatar>
                    <v-list-item
                        :title="hostData.firstName"
                    ></v-list-item>
                    <v-list-item
                        key="about"
                        title="About the host"
                        :subtitle="hostData.description"
                    ></v-list-item>
                  </v-card>
                </v-col>
              </v-row>

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
          <v-card class="w-50 d-flex flex-column justify-space-between">
            <v-card-title class="text-center">
              Fill necessary information
            </v-card-title>
            <v-card-text class="d-flex flex-column align-center justify-center">
              <v-form v-if="type === 'event' || type === 'activity' "
                      v-model="reservationCreated"
                      class="w-75 d-flex align-center justify-center flex-column"
                      ref="form"
              >
                <v-text-field
                    label="Number of people"
                    class="w-100"
                    clearable
                    type="number"
                    variant="outlined"
                    v-model="reservation.numberOfPeople"
                    :rules="[value => !!value && value > 0 || 'Number of people is required',
                    value => value <= offer?.capacity || 'Number of people must be lower than capacity',
                    ]"
                ></v-text-field>
                <v-text-field
                    label="Date"
                    class="w-100"
                    clearable
                    type="date"
                    variant="outlined"
                    v-model="reservation.dateTo"
                    :rules="[value => !!value || 'Date is required',
                    value => new Date(value) >= new Date() || 'Date can\'t be smaller than today\'s date',
                    ]"
                ></v-text-field>
              </v-form>
              <v-form v-else-if="type === 'accommodation'" v-model="reservationCreated"
                      class="w-75 d-flex flex-column align-center justify-center"
                      ref="form"
              >
                <v-text-field
                    label="Number of people"
                    class="w-100"
                    clearable
                    type="number"
                    variant="outlined"
                    v-model="reservation.numberOfPeople"
                    :rules="[value => !!value && value > 0 || 'Number of people is required',
                    value => value <= offer?.capacity || 'Number of people must be lower than capacity',
                    ]"
                ></v-text-field>
                <v-text-field
                    label="Date From"
                    clearable
                    type="date"
                    variant="outlined"
                    v-model="reservation.dateFrom"
                    class="w-100 mt-2"
                    :rules="[
                        value => !!value || 'Starting date is required',
                        value => value < reservation.dateTo || 'Date must be smaller than ending date',
                        value => new Date(value) >= new Date() || 'Date can\'t be smaller than today\'s date',
                    ]"
                ></v-text-field>
                <v-text-field
                    label="Date To"
                    clearable
                    type="date"
                    variant="outlined"
                    v-model="reservation.dateTo"
                    class="w-100 mt-2"
                    :rules="[
                        value => !!value || 'Ending date is required',
                        value => value > reservation.dateFrom || 'Date must be grater than starting date',
                        value => new Date(value) >= new Date() || 'Date can\'t be smaller than today\'s date'
                    ]"
                ></v-text-field>
                <v-select
                    label="Room name"
                    class="w-100 mt-2"
                    :rules="[
                        value => !!value || 'Room is required'
                    ]"
                    prepend-icon="mdi-home"
                    hint="Click on the house icon to see more!"
                    persistent-hint
                    :items="offer.rooms"
                    item-title="roomName"
                    item-value="ID"
                    v-model="reservation.room"
                    v-if="['hotel', 'hostel', 'guesthouse'].includes(offer.type)"
                    @click:prepend="() => openRoomDialog = true"
                ></v-select>
                <v-dialog
                    v-model="openRoomDialog"
                    width="auto"
                >
                  <v-card>
                    <v-card-title>
                      Choose a room
                    </v-card-title>
                    <v-card-text>
                      <v-card-title>
                        Rooms
                      </v-card-title>
                      <v-list v-model:selected="reservation.room">
                        <v-list-item v-for="room in offer.rooms"
                                     :key="room.roomNumber"
                                     :value="room.roomNumber"
                                     @click="openRoomDialog = false"
                                     color="blue-darken-2"
                                     rounded="shaped"
                        >
                          <RoomDetail :room="room"/>
                        </v-list-item>
                      </v-list>
                    </v-card-text>
                    <v-card-actions>
                      <v-btn
                          color="grey-darken-2 flex-grow"
                          text="Back"
                          block
                          border
                          @click="openRoomDialog = false"
                      ></v-btn>
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-form>
            </v-card-text>
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
          <v-card class="w-50 d-flex flex-column justify-space-between">
            <v-card-title class="text-center">
              Choose form of payment
            </v-card-title>
            <v-alert v-model="choosePaymentAlert"
                     type="error"
                     density="compact"
                     closable
            >
              Please choose payment method!
            </v-alert>
            <v-card-text class="mt-4">
              <v-list @update:selected="(payment) => chosenPayment = payment">
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
                  @click="makeReservation()"
              ></v-btn>
            </v-card-actions>
          </v-card>
        </v-card>
      </v-window-item>
    </v-window>
  </v-sheet>
</template>

<style scoped>

</style>