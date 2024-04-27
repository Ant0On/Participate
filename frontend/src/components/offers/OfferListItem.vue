<script setup>
import {computed, defineProps, ref} from 'vue'

const props = defineProps({
  type: String,
  offerItem: String,
})

function onItemClicked() {
  // router.push({title: 'Offers', params: {type: props.type, id: props.offerItem.offerId}})

}

const image = computed(() => {
  try {
    const image = require(`@/../images/offers/${props.type}/${props.offerItem.offerId}/${props.offerItem.offerId}_0.jpeg`)
    return image
  } catch {
    return undefined
  }
})

const cardPage = ref('main')

/*
  Common fields:
  OfferID       uint    `json:"offer_id" binding:"required"`
	Title         string  `json:"title" binding:"required,min=2"`
	Description   string  `json:"description" binding:"required,min=30"`
	Capacity      int     `json:"capacity" binding:"required,gt=0"`
	IsRecommended bool    `json:"is_recommended"`
	TownName      string  `json:"town_name" binding:"required,min=2"`
	CountryName   string  `json:"country_name" binding:"required,min=3"`
	UserID        uint    `json:"user_id" binding:"required"`
	Discount      float64 `json:"discount" binding:"required,min=0,max=100"`

	Event
	Price float64          `json:"price" binding:"required,gt=0"`
	Type  models.EventType `json:"type" binding:"required,oneof=conference concert festival 'sports event'"`

	Activity
  Price    float64             `json:"price" binding:"required,gt=0"`
	Skill    models.SkillLevel   `json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type     models.ActivityType `json:"type" binding:"required,oneof=indoor outdoor"`
	Duration time.Duration       `json:"duration" binding:"required"`

	Accommodation
  PricePerDay      float64                  `json:"price_per_day" binding:"required,gt=0"`
	IsAnimalFriendly bool                     `json:"is_animal_friendly"`
	Type             models.AccommodationType `json:"type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	Rating           int                      `json:"rating" binding:"required,min=1,max=5"`

 */

// "require(`@/assets/img/image_placeholder.png`)"

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
}
</script>

<template>
  <v-card
      class="mx-auto my-12 "
  >
    <v-carousel v-if="image"
                :show-arrows="[image].length > 1"
                :hide-delimiters="[image].length === 1"
                @click="onItemClicked"
                cycle
    >
      <v-carousel-item
          :src="image"
          contain
      >

      </v-carousel-item>
    </v-carousel>
    <v-img v-else
           :src="require(`@/assets/img/image_placeholder.png`)"
           height="100"
           @click="onItemClicked"
           cover
    ></v-img>
    <v-card-item>
      <v-card-title>
        {{ offerItem.title }}
      </v-card-title>

      <v-card-subtitle>

      </v-card-subtitle>

    </v-card-item>

    <v-card-text v-if="cardPage === 'main'">
      <v-chip :prepend-icon="chips.villa.icon"
              :color="chips.villa.color"
              variant="flat"
      >
        {{ chips.villa.text }}
      </v-chip>

      {{ offerItem.description }}


    </v-card-text>
    <v-card-text v-else-if="cardPage === 'accommodation'">


      {{ offerItem.capacity }}


    </v-card-text>
    <v-card-text v-else-if="cardPage === 'activity'">


      {{ offerItem.capacity }}


    </v-card-text>


    <v-card-actions v-if="type !== 'event'">
      <v-btn-toggle block>
        <v-btn @click="cardPage = 'main'">
          <v-icon icon="mdi-menu"></v-icon>
        </v-btn>
        <v-btn v-if="type === 'accommodation'" @click="cardPage = 'accommodation'">
          <v-icon icon="mdi-home"></v-icon>
        </v-btn>
        <v-btn v-if="type === 'activity' " @click="cardPage = 'activity'">
          <v-icon icon="mdi-run"></v-icon>
        </v-btn>
      </v-btn-toggle>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
</style>