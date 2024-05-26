<script setup>
import {computed, defineProps, onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';
import VueDatePicker from '@vuepic/vue-datepicker';


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
    const data = response.data
    if (props.type === "accommodation") {
      return {
        'offerId': data["offer_id"],
        'title': data["title"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'capacity': data["capacity"],
        'price': data['price_per_day'],
        'isRecommended': data['is_recommended'],
        'discount': data['discount'],
        'type': data['type'],
        'animal_friendly': data['is_animal_friendly'],
        'rating': data['rating']
      }

    } else if (props.type === "event") {
      return {
        'offerId': data["offer_id"],
        'title': data["title"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'capacity': data["capacity"],
        'price': data['price'],
        'isRecommended': data['is_recommended'],
        'discount': data['discount'],
        'type': data['type']
      }
    } else if (props.type === "activity") {
      return {
        'offerId': data["offer_id"],
        'title': data["title"],
        'location': data["country_name"] + ', ' + data["town_name"],
        'description': data["description"],
        'capacity': data["capacity"],
        'price': data['price'],
        'isRecommended': data['is_recommended'],
        'discount': data['discount'],
        'skill': data['skill_level'],
        'type': data['type'],
        'duration': data['duration']
      }
    }
  }
}

async function getCountries() {
  const response = await fetchWrapper.get('/api/country/get/all')
  return response.data.map((country) => {
    return {id: country.ID, name: country.Name}
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

async function onSave(){
  const {valid} = await form.value.validate()
  isEdit.value = !isEdit.value

  if(valid)
  {
    offer.value = offerChange.value
    offer.value.location = `${offerChange.value.country.name}, ${offerChange.value.town}`
  }
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
            {{ (type === "accommodation") ? `Price: ${formatDecimalPlaces(offer.price)} $/day` : `Price: ${formatDecimalPlaces(offer.price)} $` }}
          </div>

          <div class="font-weight-black ml-1" v-if="offer?.discount > 0">
            {{ (type === "accommodation") ? `${formatDecimalPlaces(priceAfterDiscount)} $/day` : `${formatDecimalPlaces(priceAfterDiscount)} $` }}
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
                   v-if="type === 'accommodation' && ['hotel', 'hostel', 'guesthouse'].includes(offer.type)">
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
              <v-col>

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
                    :subtitle="offer?.generalFacilities?.join(', ') || 'None'"
                    v-if="!isEdit"
                ></v-list-item>
                <v-select v-else v-model="offerChange.generalFacilities"
                          :items="generalFacilities"
                          label="General facilities"
                          multiple
                          chips
                          clearable
                          density="compact"
                >
                </v-select>
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
                    :subtitle="offer?.duration"
                    v-if="!isEdit"
                ></v-list-item>
                <v-list-item v-else>
                  <v-text-field type="time" v-model="offerChange.duration"></v-text-field>
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
        <v-card-actions class="w-100">
          <v-btn v-if="!isEdit"
                 color="blue-darken-2 flex-grow"
                 text="Edit"
                 block
                 border
                 @click="isEdit = !isEdit"
          ></v-btn>
          <v-btn v-else-if="isEdit"
                 color="blue-darken-2 flex-grow"
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