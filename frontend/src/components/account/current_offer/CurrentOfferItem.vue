<script setup>
import {defineProps, ref, watch} from 'vue';
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const props = defineProps({
  id: String,
  name: String,
  offerType: String,
  dateFrom: String,
  dateTo: String,
  withAnimals: Boolean,
  offerId: String,
});

const isAccepted = ref(false);
const isRejected = ref(false);

function onAccept() {
  isAccepted.value = true;
  fetchWrapper.post(`/api/reservation/${props.id}/accepted`, {})
}

function onReject() {
  isRejected.value = true;
  fetchWrapper.post(`/api/reservation/${props.id}/rejected`, {})
}
console.log(props)
watch(() => props.id, () =>{
  isAccepted.value = false
  isRejected.value = false;
})
</script>

<template>
  <div class="current_offer_item">
    <img :src="require(`@/../images/offers/${offerId}.jpeg`)" alt="Image">
    <div class="current_offer_item_details">
      <div class="title">{{ name }}</div>
      <div class="summary_data">
        <div class="details">
          <div class="field">Type: {{ offerType }}</div>
          <div class="field">Animals: {{ (withAnimals) ? 'Yes' : 'No' }}</div>
        </div>
        <div class="details">
          <div class="field">Date from: {{ dateFrom.split('T')[0] }}</div>
          <div class="field">Date to: {{ dateTo.split('T')[0] }}</div>
        </div>
        <div class="accept_reject_buttons" v-if="!isAccepted && !isRejected">
          <button class="accept_button" @click="onAccept">
            <p class="button_image">&#10003;</p>
          </button>
          <button class="reject_button" @click="onReject">
            <p class="button_image">&#10006;</p>
          </button>
        </div>
        <div class="accepted_offer" v-else-if="isAccepted">
          <p class="accepted">Accepted</p>
        </div>
        <div class="rejected_offer" v-else-if="isRejected">
          <p class="rejected">Rejected</p>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>
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
button.accept_button{
  border: 0px;
  height: 50px;
  width: 50px;
  font-size: 2rem;
  background-color: #E6E6E6;
}
button.reject_button{
  border: 0px;
  height: 50px;
  width: 50px;
  font-size: 2rem;
  background-color: #E6E6E6;

}
div.summary_data {
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
}

div.details {
  display: flex;
  flex-direction: column;
}

div.current_offer_item {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  row-gap: 10px;
  margin-top: 1%;
  background-color: #E6E6E6;
  border-radius: 10px;
}

div.current_offer_item_details {
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
  margin-left: 5%;
}
</style>