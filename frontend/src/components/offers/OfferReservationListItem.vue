<script setup>
import {computed, defineProps, ref} from 'vue'
import chipsMapper from "@/_helpers/chips";

const props = defineProps({
  type: String,
  offerItem: Object,
  custom: {
    type: Boolean,
    default: false,
  }
});
const chips = ref(chipsMapper(props.offerItem?.discount));
const image = computed(() => {
  const images = []
  let number = 0
  try {
    while(true){
      const image = require(`@/../images/offers/${props.type}/${props.offerItem.offerId}/${props.offerItem.offerId}_${number}.jpeg`)
      images.push(image)
      number++
    }
  } catch (error){
    return images
  }
});
const stateMap = {
  pending: {
    color: "gray",
    icon: "mdi-state-machine",
    text: "Pending"
  },
  finished: {
    color: "grey",
    icon: "mdi-state-machine",
    text: "Finished",
  },
  accepted: {
    color: "green",
    text: "Accepted",
    icon: "mdi-state-machine",
  },
  rejected: {
    color: "red",
    text: "Rejected",
    icon: "mdi-state-machine",
  },
  ongoing: {
    color: "blue",
    text: "Ongoing",
    icon: "mdi-state-machine",
  }
}
const cardPage = ref((props.custom)? 'custom' : 'info');
</script>

<template>
  <v-card
      class="mx-auto my-12 "
  >
    <v-carousel v-if="image"
                :show-arrows="[...image].length > 1"
                :hide-delimiters="[...image].length === 1"
                cycle
    >
      <v-carousel-item
          v-for="img in image"
          :src="img"
          cover
      >

      </v-carousel-item>
    </v-carousel>
    <v-img v-else-if="type !== 'room'"
           :src="require(`@/assets/img/image_placeholder.png`)"
           cover
    ></v-img>
    <v-img v-else-if="type !== 'room'"
           :src="require(`@/assets/img/room_placeholder.png`)"
           cover
    ></v-img>

    <v-card-item>
      <v-card-title class="font-weight-black">{{ offerItem.title }}</v-card-title>
      <v-card-subtitle v-if="type !== 'room'">
        <span class="me-1">{{ offerItem.location }}</span>
      </v-card-subtitle>
    </v-card-item>

    <v-card-text>
      <v-row align="center" class="mx-0">
        <div class="font-weight-bold">
          {{ type === "accommodation" ? `${offerItem.price} $/day` : `${offerItem.price} $` }}
        </div>
      </v-row>

      <v-row align="center" class="mx-0" v-if="type === 'accommodation'">
        <v-rating v-if="offerItem.rating" :model-value="offerItem.rating" color="amber" density="compact" size="small" half-increments readonly></v-rating>
        <div class="text-grey ms-4">{{ offerItem.rating }}</div>
      </v-row>

      <div class="my-4 text-subtitle-1">
        <v-chip :prepend-icon="stateMap[offerItem.state]?.icon" :color="stateMap[offerItem.state]?.color"  variant="flat" class="mr-1">
          {{
                stateMap[offerItem.state]?.text
          }}
        </v-chip>
      </div>
      <slot v-if="cardPage === 'custom'" name="template">
      </slot>
      <v-card v-if="cardPage === 'info'" height="100" elevation="0">
        <v-container cols="2">
          <v-row class="font-weight-bold">
            <v-col>
              <v-icon icon="mdi-home" size="small"></v-icon>
              <span class="me-1">Capacity</span>
            </v-col>
            <v-col><span class="me-1">{{ offerItem.capacity }}</span></v-col>
          </v-row>
          <v-row class="font-weight-bold" v-if="type === 'activity'">
            <v-col>
              <v-icon icon="mdi-clock-outline" size="small"></v-icon>
              <span class="me-1">Duration</span>
            </v-col>
            <v-col><span class="me-1">{{ offerItem.duration }}</span></v-col>
          </v-row>
          <v-row class="font-weight-bold" v-if="type !== 'activity' && type !== 'event'">
            <v-col>
              <v-icon icon="mdi-clock-outline" size="small"></v-icon>
              <span class="me-1">Date From</span>
            </v-col>
            <v-col><span class="me-1">{{ offerItem.dateFrom }}</span></v-col>
          </v-row>
          <v-row class="font-weight-bold" v-if="type !== 'activity' && type !== 'event'">
            <v-col>
              <v-icon icon="mdi-clock-outline" size="small"></v-icon>
              <span class="me-1">Date To</span>
            </v-col>
            <v-col><span class="me-1">{{ offerItem.dateTo }}</span></v-col>
          </v-row>
          <v-row class="font-weight-bold" v-if="type === 'activity' || type === 'event'">
            <v-col>
              <v-icon icon="mdi-clock-outline" size="small"></v-icon>
              <span class="me-1">Date</span>
            </v-col>
            <v-col><span class="me-1">{{ offerItem.date }}</span></v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-card-text>

    <v-card-actions>
      <v-row>
        <v-col cols="12">
          <v-btn-toggle block>
            <v-btn v-if="custom" @click="cardPage = 'custom'"><v-icon icon="mdi-cog"></v-icon></v-btn>
            <v-btn @click="cardPage = 'info'"><v-icon icon="mdi-information"></v-icon></v-btn>
          </v-btn-toggle>
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
</style>
