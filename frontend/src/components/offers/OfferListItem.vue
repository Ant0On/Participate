<script setup>
import {computed, defineProps, ref} from 'vue'

const props = defineProps({
  type: String,
  offerItem: String,
})

const image = computed(() => {
  try {
    const image = require(`@/../images/offers/${props.type}/${props.offerItem.offerId}/${props.offerItem.offerId}_0.jpeg`)
    return image
  } catch {
    return undefined
  }
})

const cardPage = ref('main')
const chips = {
  "sports_event": {
    color: "deep-purple",
    icon: "mdi-stadium",
    text: "Sports event"
  },
  "festival": {
    color: "pink-darken-2",
    icon: "mdi-account-group",
    text: "Festival"
  },
  "concert": {
    color: "amber",
    icon: "mdi-music",
    text: "Concert"
  },
  "conference": {
    color: "indigo",
    icon: "mdi-library",
    text: "Conference"
  },
  "beginner": {
    color: "green",
    icon: "mdi-yoga",
    text: "Beginner"
  },
  "intermediate": {
    color: "orange",
    icon: "mdi-bullseye",
    text: "Intermediate"
  },
  "advanced": {
    color: "red",
    icon: "mdi-hiking",
    text: "Advanced"
  },
  "indoor": {
    color: "grey",
    icon: "mdi-home",
    text: "Indoor"
  },
  "outdoor": {
    color: "blue-lighten-1",
    icon: "mdi-cloud",
    text: "Outdoor"
  },
  "animal_friendly": {
    color: "blue",
    icon: "mdi-dog-side",
    text: "Animal friendly"
  },
  "guesthouse": {
    color: "light-blue-darken-4",
    icon: "mdi-home",
    text: "Guesthouse"
  },
  "villa": {
    color: "deep-purple-accent-4",
    icon: "mdi-warehouse",
    text: "Villa"
  },
  "apartment": {
    color: "green-accent-4",
    icon: "mdi-city-variant",
    text: "Apartment"
  },
  "hostel": {
    color: "blue-darken-4",
    icon: "mdi-office-building",
    text: "Hostel"
  },
  "hotel": {
    color: "deep-orange",
    icon: "mdi-domain",
    text: "Hotel"
  },
  "recommended": {
    color: "amber",
    icon: "mdi-star-david",
    text: "Recommended"
  },
  "discount": {
    color: "red",
    icon: "mdi-tag",
    text: `-${props.offerItem?.discount}%`
  }
}
</script>

<template>
  <v-card
      class="mx-auto my-12 "
  >
    <v-carousel v-if="image"
                :show-arrows="[image].length > 1"
                :hide-delimiters="[image].length === 1"
                cycle
    >
      <v-carousel-item
          :src="image"
          cover
      >

      </v-carousel-item>
    </v-carousel>
    <v-img v-else
           :src="require(`@/assets/img/image_placeholder.png`)"
           cover
    ></v-img>

    <v-card-item>
      <router-link :to="`/offers/${type}/${offerItem.offerId}`">
        <v-card-title class="font-weight-black">
          {{ offerItem.title }}
        </v-card-title>
      </router-link>

      <v-card-subtitle>
      <span class="me-1">
        {{ offerItem.location }}
      </span>
      </v-card-subtitle>

    </v-card-item>

    <v-card-text >

      <v-row
          align="center"
          class="mx-0"
      >
        <div class="font-weight-bold" :class="(offerItem?.discount > 0)? 'text-decoration-line-through text-red-lighten-1': ''">
          {{(type === "accommodation")? `${offerItem.price} $/day`: `${offerItem.price} $`}}
        </div>

        <div class="font-weight-black ml-1" v-if="offerItem?.discount > 0">
          {{ (type === "accommodation") ? `   ${offerItem.price - offerItem.price * offerItem.discount / 100} $/day`
            : `  ${offerItem.price * offerItem.discount / 100} $`}}
        </div>
      </v-row>

      <v-row
          align="center"
          class="mx-0"
          v-if="type === 'accommodation'"
      >
        <v-rating
            :model-value="offerItem.rating"
            color="amber"
            density="compact"
            size="small"
            half-increments
            readonly
        ></v-rating>

        <div class="text-grey ms-4">
          {{offerItem.rating}}
        </div>
      </v-row>

      <div class="my-4 text-subtitle-1">
        <v-chip :prepend-icon="chips.recommended.icon"
                :color="chips.recommended.color"
                variant="flat"
                v-if="offerItem.isRecommended"
                class="mr-1"
        >
          {{chips.recommended.text}}
        </v-chip>
        <v-chip :prepend-icon="chips.discount.icon"
                :color="chips.discount.color"
                variant="flat"
                v-if="offerItem?.discount > 0"
                class="mr-1"
                >
          {{chips.discount.text}}
        </v-chip>
        <v-chip :prepend-icon="chips?.[offerItem?.type].icon"
                :color="chips?.[offerItem?.type].color"
                variant="flat"
                class="mr-1"
        >
          {{chips?.[offerItem?.type].text}}
        </v-chip>
        <v-chip :prepend-icon="chips?.[offerItem?.skill].icon"
                :color="chips?.[offerItem?.skill].color"
                variant="flat"
                class="mr-1"
                v-if="type=== 'activity'"
        >
          {{chips?.[offerItem?.skill].text}}
        </v-chip>

      </div>

      <v-card class="" v-if="cardPage === 'main'" height="100" elevation="0">
        {{ offerItem.description }}
      </v-card>
      <v-card class="" v-if="cardPage === 'info'" height="100" elevation="0">
        <v-container cols="2">
          <v-row
              class="font-weight-bold"
          >
            <v-col>
              <v-icon
                  icon="mdi-home"
                  size="small"
              >
              </v-icon>
              <span class="me-1">
                Capacity
              </span>
            </v-col>
            <v-col>
               <span class="me-1">
                {{ offerItem.capacity}}
              </span>
            </v-col>
          </v-row>
          <v-row class=" font-weight-bold" v-if="type === 'activity'" >
            <v-col>
              <v-icon
                  icon="mdi-clock-outline"
                  size="small"
              >
              </v-icon>
              <span class="me-1">
                Duration
              </span>
            </v-col>
            <v-col>
               <span class="me-1">
                {{ offerItem.duration}}
              </span>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-card-text>

    <v-card-actions v-if="type !== 'event'">
      <v-btn-toggle block>
        <v-btn @click="cardPage = 'main'">
          <v-icon icon="mdi-menu"></v-icon>
        </v-btn>
        <v-btn @click="cardPage = 'info'">
          <v-icon icon="mdi-information"></v-icon>
        </v-btn>
      </v-btn-toggle>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
</style>