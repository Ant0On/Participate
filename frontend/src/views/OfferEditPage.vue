<script setup>
import {computed, defineProps, onMounted, ref} from 'vue';
import {storeToRefs} from 'pinia';

import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import chipsMapper from "@/_helpers/chips";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
import RoomDetail from "@/components/detail/RoomDetail.vue";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const offer = ref(null)
const props = defineProps({
  type: String,
  id: String,
  chatID: String
})
const priceAfterDiscount = computed(() => calculatePriceAfterDiscount(offer.value?.price, offer.value?.discount))
const cardPage = ref('description')
const chips = ref(chipsMapper(offer.value?.discount))

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

onMounted(async () => {
  offer.value = await getOfferDetails();
});

</script>

<template>
  <v-progress-circular v-if="!offer"
                       color="primary"
                       indeterminate
  ></v-progress-circular>
  <v-sheet v-else class="d-flex flex-column">
    <v-breadcrumbs>
      <v-breadcrumbs-item :to="'/account/current_offers'" title="My offers"></v-breadcrumbs-item>
      <v-breadcrumbs-divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </v-breadcrumbs-divider>
      <v-breadcrumbs-item :title="offer.title"></v-breadcrumbs-item>
    </v-breadcrumbs>
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
      </v-card>
    </v-card>
  </v-sheet>
</template>

<style scoped>

</style>