<script setup>
import { computed, defineProps, ref } from 'vue';
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";
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
const priceAfterDiscount = computed(() => calculatePriceAfterDiscount(props.offerItem.price, props.offerItem?.discount));
const image = computed(() => {
  try {
    const image = require(`@/../images/offers/${props.type}/${props.offerItem.offerID}/${props.offerItem.offerID}_0.jpeg`);
    return image;
  } catch {
    return undefined;
  }
});

const cardPage = ref((props.custom)? 'custom' : 'main');
</script>

<template>
  <v-card class="mx-auto my-12">
    <v-carousel v-if="image" :show-arrows="[image].length > 1" :hide-delimiters="[image].length === 1" cycle>
      <v-carousel-item :src="image" cover></v-carousel-item>
    </v-carousel>
    <v-img v-else :src="require(`@/assets/img/image_placeholder.png`)" cover></v-img>

    <v-card-item>
      <v-card-title v-if="custom" class="font-weight-black">{{ offerItem.title }}</v-card-title>
      <router-link v-else :to="`/offers/${type}/${offerItem.offerID}`">
        <v-card-title class="font-weight-black">{{ offerItem.title }}</v-card-title>
      </router-link>
      <v-card-subtitle>
        <span class="me-1">{{ offerItem.location }}</span>
      </v-card-subtitle>
    </v-card-item>

    <v-card-text>
      <v-row align="center" class="mx-0">
        <div class="font-weight-bold" :class="offerItem?.discount > 0 ? 'text-decoration-line-through text-red-lighten-1' : ''">
          {{ type === "accommodation" ? `${offerItem.price} $/day` : `${offerItem.price} $` }}
        </div>
        <div class="font-weight-black ml-1" v-if="offerItem?.discount > 0">
          {{ type === "accommodation" ? `   ${priceAfterDiscount} $/day` : `  ${priceAfterDiscount} $` }}
        </div>
      </v-row>

      <v-row align="center" class="mx-0" v-if="type === 'accommodation'">
        <v-rating :model-value="offerItem.rating" color="amber" density="compact" size="small" half-increments readonly></v-rating>
        <div class="text-grey ms-4">{{ offerItem.rating }}</div>
      </v-row>

      <div class="my-4 text-subtitle-1">
        <v-chip v-if="offerItem?.discount > 0" :prepend-icon="chips.discount?.icon" :color="chips.discount?.color" variant="flat" class="mr-1">
          {{ chips.discount?.text }}
        </v-chip>
        <v-chip :prepend-icon="chips?.[offerItem?.type]?.icon" :color="chips?.[offerItem?.type]?.color" variant="flat" class="mr-1">
          {{ chips?.[offerItem?.type]?.text }}
        </v-chip>
        <v-chip v-if="type === 'activity'" :prepend-icon="chips?.[offerItem?.skill]?.icon" :color="chips?.[offerItem?.skill]?.color" variant="flat" class="mr-1">
          {{ chips?.[offerItem?.skill]?.text }}
        </v-chip>
      </div>
      <slot v-if="cardPage === 'custom'" name="template">
      </slot>
      <v-card v-if="cardPage === 'main'" height="100" elevation="0">{{ offerItem.description }}</v-card>
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
        </v-container>
      </v-card>
    </v-card-text>

    <v-card-actions>
      <v-row>
        <v-col cols="12">
          <v-btn-toggle block>
            <v-btn v-if="custom" @click="cardPage = 'custom'"><v-icon icon="mdi-cog"></v-icon></v-btn>
            <v-btn @click="cardPage = 'main'"><v-icon icon="mdi-menu"></v-icon></v-btn>
            <v-btn @click="cardPage = 'info'"><v-icon icon="mdi-information"></v-icon></v-btn>
          </v-btn-toggle>
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
</style>
