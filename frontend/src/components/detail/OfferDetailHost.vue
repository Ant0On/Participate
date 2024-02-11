<template>
<div class="offer_detail_host">
  <div class="offer_detail_host_picture">
    <div class="about_host">
      About the host
    </div>
    <img v-if="imagePath" class="detail_host_image" :src="require(`@/../${imagePath}`)" alt="Host picture">
    <div class="host_name">
      {{host_first_name }}
    </div>
  </div>
  <div class="offer_detail_host_description">
    {{ host_detail }}
  </div>
  <transition name="fade" mode="out-in" :style="{ transitionDuration: `${transitionDuration}ms` }">
    <img :key="currentOfferImageIndex" class="detail_offer_image" :src="currentOfferImagePath" alt="Offer image">
  </transition></div>
</template>

<script setup>
import {computed, onMounted, ref} from 'vue';

const props = defineProps({
  host_first_name: '',
  offer_id: 0,
  host_detail: '',
  imagePath: '',
});

const currentOfferImageIndex = ref(0);
const transitionDuration = ref(1000);

const currentOfferImagePath = computed(() => {
  const imageIndex = currentOfferImageIndex.value;
  const imageName = `${props.offer_id}_${imageIndex}.jpeg`;
  const imagePath = require(`@/../images/offers/${props.offer_id}/${imageName}`);

  new Image().src = imagePath;
  return imagePath;
});

function changeImage() {
  if (!imageExists(currentOfferImageIndex.value + 1)) {
    currentOfferImageIndex.value = 0;
  } else {
    currentOfferImageIndex.value += 1;
  }
}

function imageExists(index) {
  const imageName = `${props.offer_id}_${index}.jpeg`;
  try {
    require(`@/../images/offers/${props.offer_id}/${imageName}`);
    return true;
  } catch (error) {
    return false;
  }
}

onMounted(() => {
  setInterval(changeImage, 3000);
});
</script>

<style scoped>
div.offer_detail_host{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  padding: 1%;
}
div.offer_detail_host_description{
  font-family: "Poppins-Regular", Helvetica,serif;
  font-size: 1.4rem;
}
div.about_host{
  font-family: "Poppins-SemiBold", Helvetica;
  font-size: 1.4rem;
  font-weight: 600;
  align-self: center;
  padding-bottom: 5%;
}
div.host_name{
  font-family: "jsMath-cmr10-cmr10", Helvetica, serif;
  font-size: 1.4rem;
  font-weight: 400;
  align-self: center;
  padding-top: 5%;
}
div.offer_detail_host_picture{
  display: flex;
  flex-direction: column;
  padding: 2%;
}
div.offer_detail_host_description{
  padding: 2% 5% 2% 5%;
  flex-grow: 1;
}
img.detail_host_image{
  border-radius: 50%;
  height: 250px;
  width: 250px;
  align-self: center;

}

img.detail_offer_image {
  padding: 1%;
  height: 250px;
  width: 250px;
  object-fit: cover;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>