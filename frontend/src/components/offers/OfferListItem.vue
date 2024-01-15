<script setup>
import {defineProps} from 'vue'
import {router} from '@/router'

const props = defineProps({
  type: String,
  id: String,
  location: String,
  image: String,
  name: String,
  description: String,
  price: String,
  max_people: String,
})


function onItemClicked() {
  router.push({name: 'Offers', params:{type: props.type, id: props.id}})
}
</script>

<template>
  <div class="offer_item" @click="onItemClicked">
    <img :src="image" alt="Image">
    <div class="item_details">
      <div class="title">{{ name }}</div>
      <div class="description"> {{ description }}</div>
      <div class="price">Price: {{ price }}</div>
      <div class="number_of_people">Number of people: {{ max_people }}</div>
      <div class="location">Location: {{ location }}</div>
    </div>

  </div>
</template>

<style scoped>
div.offer_item {
  display: flex;
  flex-direction: row;
  row-gap: 10px;
  margin-top: 1%;
  background-color: #E6E6E6;
  border-radius: 10px;
}

div.item_details {
  margin: 1% 5% 1% 5%;
  display: flex;
  flex-direction: column;
  flex-grow: 1;

}

div.title {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.8rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
  padding: 1%;
}

div.price {
  color: #7a7a7a;
  font-family: "Poppins", Helvetica;
  font-size: 1.2rem;
  font-weight: 500;
  line-height: normal;
  padding: 1%;
}

div.number_of_people {
  color: #7a7a7a;
  font-family: "Poppins", Helvetica;
  font-size: 1.2rem;
  font-weight: 500;
  line-height: normal;
  padding: 1%;
}

div.description {
  color: #7a7a7a;
  font-family: "Poppins", Helvetica;
  font-size: 1.2rem;
  font-weight: 500;
  line-height: normal;
  padding: 1%;
}

div.location {
  color: #7a7a7a;
  font-family: "Poppins", Helvetica;
  font-size: 1.2rem;
  font-weight: 500;
  line-height: normal;
  align-self: flex-end;
  padding: 1%;
}

img {
  border-radius: 10px;
  overflow: hidden;
  height: 300px;
  width: 300px;
  opacity: 0.9;
  flex-shrink: 0;
}
</style>