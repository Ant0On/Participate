<script setup>
import { defineEmits, defineProps, computed, ref } from 'vue';
import StarRating from 'vue-star-rating';
import { fetchWrapper } from "@/_helpers/fetch-wrapper";

const props = defineProps({
  offerData: Object
});

const emits = defineEmits(['nextOffer']);
const error = ref(null);

function updateRating(rating){
  fetchWrapper.post(`/api/customer/offer/${props.offerData.value.reservationId}/grade`, {
    count: rating
  })
      .then(() => {
        emits('nextOffer');
      })
      .catch(errors => {
        error.value = "An error occurred while updating the rating. " + errors.message;
        emits('nextOffer');
      });
}

const isImageSource = computed(() =>{
  try{
    require(`@/../images/offers/${props.offerData.value.offerId}.jpeg`);
    return true;
  }
  catch{
    return false;
  }
});
</script>

<template>
  <Transition name="fade">
    <div :key="offerData.value.offerId" class="offer_rate">
      <p class="header">Rate this offer:</p>
      <div class="offer_rate_description">
        <p class="offer_name">{{ offerData.value.name }}</p>
        <img v-if="isImageSource" :src="require(`@/../images/offers/${offerData.value.offerId}.jpeg`)" alt="Image"/>
        <img v-else :src="require(`@/assets/img/image_placeholder.png`)" alt="Image">
        <star-rating :show-rating="false" @update:rating="updateRating"></star-rating>
        <p v-if="error">{{ error }}</p>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
img{
  height: 40%;
  width: 40%;
  border-radius: 50%;
}
div.offer_rate_description{
  display: flex;
  flex-direction: column;
  align-items: center;
}
p {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.4rem;
  padding: 2%;
  margin-top: 2%;
  font-weight: 400;
  text-align: center;

}
p.offer_name {
  font-size: 1rem;
}
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>