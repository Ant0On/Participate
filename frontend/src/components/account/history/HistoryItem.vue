<script setup>
import {defineProps, computed} from 'vue';
const props = defineProps({
  name: String,
  offerType: String,
  dateFrom: String,
  dateTo: String,
  withAnimals: Boolean,
  reservationState: String,
  offerId: String,
});

const isAccepted = computed(() => props.reservationState === "accepted")
const isRejected = computed(() => props.reservationState === "rejected")
const isImageSource = computed(() =>{
  try{
    require(`@/../images/offers/${props.offerId}.jpeg`)
    return true
  }
  catch{
    return false
  }
})
</script>

<template>
  <div class="account_history_item">
    <img v-if="isImageSource" :src="require(`@/../images/offers/${offerId}.jpeg`)" alt="Image">
    <img v-else :src="require(`@/assets/img/image_placeholder.png`)" alt="Image">
    <div class="account_history_item_details">
      <div class="title">{{ name }}</div>
      <div class="summary_data">
        <div class="details">
          <div class="field">Type: {{ offerType }}</div>
          <div class="field">Animals: {{ (withAnimals)? 'Yes': 'No'}}</div>
        </div>
        <div class="details">
          <div class="field">Date from: {{ dateFrom.split('T')[0] }}</div>
          <div class="field">Date to: {{ dateTo.split('T')[0] }}</div>
        </div>
        <div class="accepted_offer" v-if="isAccepted">
          <p class="accepted">Accepted</p>
        </div>
        <div class="rejected_offer" v-else-if="isRejected">
          <p class="rejected">Rejected</p>
        </div>
        <div class="finished_offer" v-else>
          <p class="finished">Finished</p>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>
div.summary_data{
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
}
div.details{
  display: flex;
  flex-direction: column;
}
div.account_history_item {
  display: flex;
  flex-direction: row;
  row-gap: 10px;
  margin-top: 1%;
  background-color: #E6E6E6;
  border-radius: 10px;
}

div.account_history_item_details {
  margin: 1% 2% 1% 2%;
  display: flex;
  flex-direction: column;
  flex-grow: 1;

}

div.title {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.3rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
  padding: 1%;
}

div.field {
  color: #7a7a7a;
  font-family: "Poppins", Helvetica;
  font-size: 1rem;
  font-weight: 500;
  line-height: normal;
  padding: 1%;
}

img {
  border-radius: 50%;
  height: 100px;
  width: 100px;
  opacity: 0.9;
  flex-shrink: 0;
  align-self: center;
  margin-right: 2%;
  margin-left: 2%;
}
p.accepted{
  color: green;
  font-family: "Poppins", Helvetica;
  font-size: 1.3rem;
  font-weight: 700;
}
p.rejected{
  color: red;
  font-family: "Poppins", Helvetica;
  font-size: 1.3rem;
  font-weight: 700;
}
p.finished{
  color: mediumpurple;
  font-family: "Poppins", Helvetica;
  font-size: 1.3rem;
  font-weight: 700;
}
</style>