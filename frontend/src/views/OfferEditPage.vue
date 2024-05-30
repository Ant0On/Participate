<script setup>
import {computed, defineProps, onMounted, ref, watch} from 'vue';
import {storeToRefs} from 'pinia';


import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import chipsMapper from "@/_helpers/chips";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import RoomDetail from "@/components/detail/RoomDetail.vue";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const offer = ref(null)
const offerChange = ref(null)
const offerCountries = ref(null)
const isEdit = ref(false)
const offerFilled = ref(true)
const isAddRoom = ref(false)
const offerChangeRoom = ref({})
const form = ref()

const props = defineProps({
  type: String,
  id: String,
  chatID: String
})
const priceAfterDiscount = computed(() => calculatePriceAfterDiscount(offer.value?.price, offer.value?.discount))
const cardPage = ref('description')
const chips = computed(() => chipsMapper(offer.value?.discount))

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
        'title': data["Title"],
        'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name || 'city'}`,
        'description': data["Description"],
        'capacity': data["Capacity"],
        'price': data['PricePerDay'],
        'discount': data['Discount'],
        'type': data['Type'],
        'animal_friendly': data['IsAnimalFriendly'],
        'rating': data['Rating'] || 0,
        'numberOfRooms': data?.NumberOfRooms,
        'rooms': data?.Rooms?.map((room) => {
          return {...room}
        }),
        'generalFacilities': data?.GeneralFacilities?.map((generalFacility) => generalFacility.Name)

      }

    } else if (props.type === "event") {
      return {
        'offerId': data["ID"],
        'title': data["Title"],
        'location': `${data?.Town?.Country?.CountryName || 'Country'}, ${data?.Town?.name || 'city'}`,
        'description': data["Description"],
        'capacity': data["Capacity"],
        'price': data['Price'],
        'discount': data['Discount'],
        'type': data['Type'],
        'dateFrom': data?.DateFrom?.split('T')?.[0],
        'dateTo': data?.DateTo?.split('T')?.[0],
        'townId': data?.TownID,
        'townName': data?.Town?.name,
      }
    } else if (props.type === "activity") {
      return {
        'offerId': data["ID"],
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
        'townId': data?.TownID
      }
    }
  }
}

async function getCountries() {
  const response = await fetchWrapper.get('/api/country/get/all')
  return response.data.map((country) => {
    return {id: country.ID, name: country.CountryName}
  })
}

const formatDecimalPlaces = (num) => (Math.round(num * 100) / 100).toFixed(2)

function getOfferChange(offer) {
  let offerChange = {...offer};
  delete offerChange.location;
  offerChange.country = offerCountries.value.find(country => country.name === offer.location.split(',')[0])
  offerChange.town = offer.location.split(',')[1].trim()
  return offerChange
}

async function onSave() {
  const {valid} = await form.value.validate()
  isEdit.value = !isEdit.value
  if (valid) {
    if(props.type === 'event'){
      await saveEvent()
    }
    if(props.type === 'activity'){
      await saveActivity()
    }
    offer.value = offerChange.value
    offer.value.location = `${offerCountries.value.filter(country =>
        country.id === offerChange.value.country || country.id === offerChange.value.country?.id)?.[0]?.name}, ${offerChange.value.town}`
  }
}

async function saveEvent(){
  await fetchWrapper.put(`/api/host/event/update/${offerChange.value.offerId}`, {
    Title: offerChange.value.title,
    Description: offerChange.value.description,
    Capacity: Number(offerChange.value.capacity),
    Town: {
      name: offerChange.value.town,
      country_id: Number(offerChange.value?.country?.id) || Number(offerChange.value.country),
    },
    UserID: user.value?.ID,
    DateFrom: new Date(offerChange.value.dateFrom).toJSON(),
    DateTo: new Date(offerChange.value.dateTo).toJSON(),
    Price: Number(offerChange.value.price),
    Discount: Number(offerChange.value.discount),
    Type: offerChange.value.type,
  })
}

async function saveActivity(){
  await fetchWrapper.put(`/api/host/activity/update/${offerChange.value.offerId}`, {
    Title: offerChange.value.title,
    Description: offerChange.value.description,
    Capacity: Number(offerChange.value.capacity),
    Town: {
      name: offerChange.value.town,
      country_id: Number(offerChange.value?.country?.id) || Number(offerChange.value.country),
    },
    UserID: user.value?.ID,
    Date: new Date(offerChange.value.date).toJSON(),
    Price: Number(offerChange.value.price),
    Discount: Number(offerChange.value.discount),
    event_type: offerChange.value.type,
    skill_level: offerChange.value.skill,
    activity_type: offerChange.value.type,
    duration: offerChange.value.duration,
    images: '',
  })
}

async function saveRoom(){
  const {valid} = await form.value.validate()
  if(valid){
    offerChange.value.rooms = [... offerChange.value?.rooms || [], {... offerChangeRoom.value}]
    isAddRoom.value = false
    offerChange.value.numberOfRooms = offerChange.value.rooms?.length || 1
  }
}

function deleteRoom(roomNumber){
  offerChange.value.rooms = offerChange.value.rooms?.filter((room) => room.roomNumber !== roomNumber) || offerChange.value.rooms
}

onMounted(async () => {
  offer.value = await getOfferDetails();
  offerCountries.value = await getCountries();
  offerChange.value = getOfferChange(offer.value);
});

const accommodationTypes = ['Hotel', 'Hostel', 'Apartment', 'Villa', 'Guesthouse']
const generalFacilities = ["Swimming Pool", "Gym", "Spa", "Restaurant", "Bar", "Lounge", "Conference Room",
  "Business Center", "WiFi", "Parking", "24-Hour Front Desk", "Fitness Center", "Laundry Service", "Room Service",
  "Concierge Service", "Outdoor Pool", "Children's Playground", "Tennis Court", "Library", "Garden", "Sauna",
  "Jacuzzi", "Billiards Room", "Cinema", "Karaoke Room", "Bowling Alley", "BBQ Area", "Shuttle Service"]
const roomFacilities = ["Television", "Air Conditioning", "Mini Fridge", "Safe", "Coffee Maker", "Microwave",
  "Kettle", "Iron and Ironing Board", "Hair Dryer", "Desk", "Ocean View", "Mountain View", "Balcony", "Bathtub",
  "Shower", "WiFi", "Room Service", "Breakfast Included", "In-Room Safe", "Telephone", "DVD Player", "Alarm Clock",
  "Robes", "Slippers", "Toiletries", "Work Desk", "Sofa Bed", "Fireplace", "Refrigerator", "Dining Area"
];

const skillLevels = ['Beginner', 'Intermediate', 'Advanced']
const activityTypes = ['Indoor', 'Outdoor']
const equipmentList = ["Life Jacket", "Kayak", "Paddle", "Helmet", "Snowboard", "Sled", "Snowshoes", "Tent",
  "Sleeping Bag", "Backpack", "Hiking Boots", "Compass", "Map", "Binoculars", "Flashlight", "First Aid Kit",
  "Water Bottle", "Climbing Harness", "Climbing Rope", "Carabiners", "Rock Climbing Shoes", "Fishing Rod",
  "Bait", "Camera", "Swimsuit", "Snorkel Gear", "Tennis Racket", "Golf Clubs", "Bicycle"];
const eventTypes = ['Conference', 'Concert', 'Festival', 'Sports event']
</script>

<template>
  <v-progress-circular v-if="!offer"
                       color="primary"
                       indeterminate
  ></v-progress-circular>
  <v-sheet v-else class="d-flex flex-column align-center justify-center">
    <v-breadcrumbs class="align-self-start">
      <v-breadcrumbs-item :to="'/account/my_offers'" title="My offers"></v-breadcrumbs-item>
      <v-breadcrumbs-divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </v-breadcrumbs-divider>
      <v-breadcrumbs-item :title="offer.title"></v-breadcrumbs-item>
    </v-breadcrumbs>
    <v-form v-model="offerFilled" class="h-100 w-100 d-flex align-center justify-center" ref="form">
      <v-card class="d-flex flex-row rounded-xl h-100" width="60%">
        <v-card class="w-50">
          <v-carousel v-if="image"
                      :show-arrows="[...image].length > 1"
                      :hide-delimiters="[...image].length === 1"
                      :height="(isEdit)? '700px': '600px'"
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
          <v-card-title v-if="!isEdit" class="font-weight-black text-center">
            {{ offer.title }}
          </v-card-title>
          <v-card-title v-else class="font-weight-black text-center">
            <v-text-field v-model="offerChange.title" flat
                          label="Title"
                          density="compact"
                          :rules="[value => !!value, value => value.length > 3]"
            ></v-text-field>
          </v-card-title>

          <v-card-subtitle class="d-flex mb-2">

          <span v-if="!isEdit" class="me-1">
            {{ offer.location }}
          </span>
            <div v-else class="d-flex justify-space-between w-100">
              <v-select label="Coutnry"
                        v-model="offerChange.country"
                        :items="offerCountries"
                        class="w-100"
                        density="compact"
                        item-title="name"
                        item-value="id"
              ></v-select>
              <v-text-field label="Town"
                            v-model="offerChange.town"
                            class="w-100 ma-2"
                            density="compact"
                            :rules="[(value) => !!value, (value) => value.length > 3]"></v-text-field>
            </div>
          </v-card-subtitle>
          <v-card-subtitle class="mx-0 d-flex " v-if="!isEdit">
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
          <v-card-subtitle class="d-flex flex-space-between w-100" v-else>
            <v-text-field v-model="offerChange.price"
                          width="40px"
                          density="compact"
                          label="Price"
                          type="number"
                          :rules="[(value) => !!value, (value) => value > 0]"
            ></v-text-field>
            <v-text-field v-model="offerChange.discount"
                          density="compact"
                          class="pl-1"
                          label="Discount"
                          type="number"
                          :rules="[(value) => value >= 0 && value < 100]"
            ></v-text-field>
          </v-card-subtitle>
          <div class="ma-4 text-subtitle-1" v-if="!isEdit">
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
          <v-card-subtitle class="d-flex justify-space-between w-100" v-else>
            <v-select v-model="offerChange.type"
                      :items="(type === 'event')? eventTypes.map((eventType) => eventType.toLowerCase())
                    : (type === 'accommodation')? accommodationTypes.map((accommodationType) => accommodationType.toLowerCase())
                     : activityTypes.map((activityType) => activityType.toLowerCase())"
                      label="Type"
                      density="compact"
                      @update:menu="() => {
                        if(['hotel', 'hostel', 'guesthouse'].includes(offerChange.type)){
                          offerChange.numberOfRooms = offerChange.rooms?.length || 1
                        }
                      }"
            ></v-select>
            <v-select v-model="offerChange.skill" v-if="type === 'activity'"
                      label="Skill"
                      class="pl-1"
                      density="compact"
                      :items="skillLevels.map((skillLevel) => skillLevel.toLowerCase())"
            >
            </v-select>
          </v-card-subtitle>
          <v-divider></v-divider>
          <v-card-actions>
            <v-btn-toggle block rounded="lg">
              <v-btn @click="cardPage = 'description'">
                <v-icon icon="mdi-menu"></v-icon>
              </v-btn>
              <v-btn @click="cardPage = 'info'">
                <v-icon icon="mdi-information"></v-icon>
              </v-btn>
              <v-btn @click="cardPage = 'accommodation'"
                     v-if="type === 'accommodation' && ['hotel', 'hostel', 'guesthouse'].includes((isEdit)? offerChange.type : offer.type)">
                <v-icon icon="mdi-home"></v-icon>
              </v-btn>
              <v-btn @click="cardPage = 'activity'" v-if="type === 'activity'">
                <v-icon icon="mdi-run"></v-icon>
              </v-btn>
            </v-btn-toggle>
          </v-card-actions>

          <v-card-text v-if="cardPage === 'description'">
            <div v-if="!isEdit">{{ offer.description }}</div>
            <v-textarea v-else
                        density="compact"
                        label="Description"
                        v-model="offerChange.description"
                        :rules="[(values) => values.length > 30]"
            ></v-textarea>
          </v-card-text>
          <v-card-text v-else-if="cardPage === 'info'">
            <v-list class="h-100 w-100" style="overflow: hidden">
              <v-row cols="2">
                <v-col>
                  <v-list-item v-if="!isEdit"
                               key="capacity"
                               title="Capacity"
                               :subtitle="offer.capacity"
                  ></v-list-item>
                  <v-list-item v-else>
                    <v-text-field
                        v-model="offerChange.capacity"
                        density="compact"
                        label="Capacity"
                        type="number"
                        :rules="[(value) => !!value, (value) => value > 0]"
                    ></v-text-field>
                  </v-list-item>
                </v-col>
              </v-row>
              <v-row cols="2" v-if="type === 'accommodation'">
                <v-col>
                  <v-list-item
                      key="number_of_rooms"
                      title="Number of rooms"
                      :subtitle="offer?.numberOfRooms || 1"
                      v-if="!isEdit"
                  ></v-list-item>
                  <v-list-item
                      key="number_of_rooms"
                      title="Number of rooms"
                      :subtitle="offerChange?.numberOfRooms || 1"
                      v-else-if="isEdit && ['hotel', 'hostel', 'guesthouse'].includes(offerChange.type)"
                  ></v-list-item>
                  <v-list-item
                      key="number_of_rooms"
                      title="Number of rooms"
                      v-else
                  >
                    <v-text-field
                        v-model="offerChange.numberOfRooms"
                        density="compact"
                        label="Number of rooms"
                        type="number"
                        :rules="[(value) => !!value, (value) => value > 0]"
                    ></v-text-field>
                  </v-list-item>
                </v-col>
              </v-row>
              <v-row cols="2" v-if="type === 'event'">
                <v-col>
                  <v-list-item
                      key="date_from"
                      title="Date From"
                      :subtitle="offer?.dateFrom"
                      v-if="!isEdit"
                  ></v-list-item>
                  <v-list-item v-else>
                    <v-text-field v-model="offerChange.dateFrom"
                                  label="Date From"
                                  type="date"
                                  denisty="compact"
                                  :rules="[
                                  value => !!value || 'Ending date is required',
                                  value => !offerChange.dateTo || value <= offerChange.dateTo || 'Date must be lower than end date'
                               ]"
                    >
                    </v-text-field>
                  </v-list-item>
                </v-col>
                <v-col>
                  <v-list-item
                      key="date_to"
                      title="Date To"
                      :subtitle="offer?.dateTo"
                      v-if="!isEdit"

                  ></v-list-item>
                  <v-list-item v-else>
                    <v-text-field v-model="offerChange.dateTo"
                                  label="Date To"
                                  type="date"
                                  denisty="compact"
                                  :rules="[
                                  value => !!value || 'Ending date is required',
                                  value => !offerChange.dateFrom || value >= offerChange.dateFrom || 'Date must be grater than starting date',
                                  value => new Date(value) > new Date()
                               ]"
                    >
                    </v-text-field>
                  </v-list-item>
                </v-col>
              </v-row>
              <v-row cols="2" v-if="type === 'activity'">
                <v-col>
                  <v-list-item
                      key="date"
                      title="Date"
                      :subtitle="offer?.date"
                      v-if="!isEdit"
                  ></v-list-item>
                  <v-list-item v-else>
                    <v-text-field v-model="offerChange.date"
                                  label="Date"
                                  type="date"
                                  denisty="compact"
                                  :rules="[
                                  value => !!value || 'Ending date is required',
                                  value => new Date(value) > new Date()
                               ]"
                    >
                    </v-text-field>
                  </v-list-item>
                </v-col>
                <v-col>
                  <v-list-item
                      key="duration"
                      title="Duration"
                      :subtitle="`${offer?.duration} ${(offer?.duration === 1)? 'Hour': 'Hours'} `"
                      v-if="!isEdit"
                  ></v-list-item>
                  <v-list-item v-else>
                    <v-text-field type="number" v-model="offerChange.duration"
                                  label="Duration"
                                  hint="Time in hours"
                                  persistent-hint
                                  ></v-text-field>
                  </v-list-item>
                </v-col>
              </v-row>
              <v-row>
                <v-list-item
                    key="general_facilities"
                    title="General Facilities"
                    v-if="type === 'accommodation' && !isEdit"
                >{{
                    offer?.generalFacilities?.join(', ') || 'None'
                  }}
                </v-list-item>
                <v-list-item v-else-if="type === 'accommodation' && isEdit" class="w-100">
                  <v-select v-else v-model="offerChange.generalFacilities"
                            :items="generalFacilities"
                            label="General facilities"
                            multiple
                            chips
                            clearable
                            density="compact"
                            class="w-100"
                  >
                  </v-select>
                </v-list-item>
                <v-list-item
                    key="equipment"
                    title="Equipment"
                    v-if="type === 'activity' && !isEdit"
                >
                  {{
                    offer?.equipment?.join(', ') || 'None'
                  }}
                </v-list-item>
                <v-list-item v-else-if="type === 'activity' && isEdit" class="w-100">
                  <v-select v-else v-model="offerChange.equipment"
                            :items="equipmentList"
                            label="Equipment"
                            multiple
                            chips
                            clearable
                            density="compact"
                            class="w-100"
                  >
                  </v-select>
                </v-list-item>
              </v-row>
            </v-list>
          </v-card-text>

          <v-card-text v-else-if="cardPage === 'accommodation'" style="overflow-y: scroll; height: 295px;">
            <v-card-title>
              Rooms
            </v-card-title>
            <v-list style="overflow-y: scroll" v-if="!isEdit">
              <v-list-item v-for="room in offer.rooms" :key="room.roomNumber">
                <RoomDetail :room="room"/>
              </v-list-item>
            </v-list>
            <v-list style="overflow-y: scroll; " v-if="isEdit">
              <v-list-item v-for="room in offerChange.rooms" :key="room.roomNumber">
                <RoomDetail :room="room"/>
                <template v-slot:append>
                  <v-btn
                      color="grey-lighten-1"
                      icon="mdi-close"
                      variant="text"
                      @click="deleteRoom(room.roomNumber)"
                  ></v-btn>
                </template>
              </v-list-item>
              <v-list-item>
                <v-card class="border-s" flat color="grey-lighten-5">
                  <v-card-title>
                    Add a room
                  </v-card-title>
                  <v-card-actions class="d-flex align-center justify-center">
                    <v-btn text="Add a room" variant="outlined" @click="isAddRoom = true">
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-list-item>
            </v-list>
            <v-dialog
                v-model="isAddRoom"
                width="60%"
            >
              <v-card>
                <v-card-title>
                  Add a room
                </v-card-title>
                <v-card-text>
                  <v-card>
                    <v-form v-model="form"></v-form>
                    <v-card-title>
                      <v-text-field v-model="offerChangeRoom.roomNumber"
                                    label="Room number"
                                    type="number"
                                    flat
                                    denisty="compact"
                                    :rules="[value => !!value, value => value > 0,
                                    value => !offerChange.rooms || !offerChange.rooms?.some((room) => room.roomNumber === value)]"
                      ></v-text-field>
                      <v-text-field v-model="offerChangeRoom.roomName"
                                    label="Room name"
                                    flat
                                    denisty="compact"
                                    :rules="[value => !!value, value => value?.length > 3]"
                      ></v-text-field>
                    </v-card-title>
                    <v-card-subtitle>
                      <v-text-field v-model="offerChangeRoom.area"
                                    label="Room area"
                                    type="number"
                                    flat
                                    denisty="compact"
                                    :rules="[value => !!value, value => value > 0]"></v-text-field>
                      <v-text-field v-model="offerChangeRoom.capacity"
                                    label="Room capacity"
                                    type="number"
                                    flat
                                    denisty="compact"
                                    :rules="[value => !!value, value => value > 0]"
                      ></v-text-field>
                    </v-card-subtitle>
                    <v-card-text>
                      <v-textarea
                          density="compact"
                          label="Description"
                          v-model="offerChangeRoom.description"
                          :rules="[(values) => values?.length > 10]"
                      ></v-textarea>
                      <div class="text-subtitle-1 my-1">Room facilities</div>
                      <v-divider></v-divider>
                      <div class="my-1">
                        <v-select v-else v-model="offerChangeRoom.roomFacilities"
                                  :items="roomFacilities"
                                  label="Room facilities"
                                  multiple
                                  chips
                                  clearable
                                  density="compact"
                                  class="w-100"
                        ></v-select>
                      </div>
                    </v-card-text>
                  </v-card>
                </v-card-text>
                <v-card-actions class="d-flex justify-space-between">
                  <v-btn
                      color="grey-darken-2"
                      text="Back"
                      border
                      class="w-25"
                      @click="isAddRoom = false"
                  ></v-btn>
                  <v-btn
                      color="grey-darken-2"
                      class="w-25"
                      text="Save"
                      border
                      @click="saveRoom"
                  ></v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
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
          <v-card-actions class="w-100">
            <v-btn v-if="!isEdit"
                   color="blue-darken-2"
                   text="Edit"
                   block
                   border
                   @click="isEdit = !isEdit"
            ></v-btn>
            <v-btn v-else-if="isEdit"
                   color="blue-darken-2"
                   text="Save"
                   block
                   border
                   @click="onSave"
            ></v-btn>
          </v-card-actions>
        </v-card>
      </v-card>
    </v-form>
  </v-sheet>
</template>

<style scoped>

</style>