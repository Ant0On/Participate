<script setup>
import {defineEmits, defineProps, onMounted} from 'vue';
import StarRating from 'vue-star-rating';
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const props = defineProps({
  offerData: Object
})
const emits = defineEmits(['nextOffer'])

function updateRating(rating){
  fetchWrapper.put(`/api/reservation/${props.offerData.ID}/grade/${rating}`).then(()=>{
    emits('nextOffer')
  }).catch(errors => {

  })
}
onMounted(async ()=>{

})
</script>

<template>
<div class="offer_rate">
  <p class="header">Rate this offer:</p>
  <div class="offer_rate_description">

  </div>
  <star-rating :show-rating="false" @update:rating="updateRating"></star-rating>
</div>
</template>

<style scoped>

</style>