<script setup>
import {defineProps, computed} from 'vue'
import {router} from '@/router'

const props = defineProps({
  type: String,
  id: String,
  location: String,
  title: String,
  price: String,
  capacity: String,
})

function onItemClicked() {
  router.push({title: 'Offers', params:{type: props.type, id: props.id}})
}

const isImageSource = computed(() =>{
  try{
    require(`@/../images/offers/${props.type}/${props.id}/${props.id}_0.jpeg`)
    return true
  }
  catch{
    return false
  }
})
</script>

<template>
  <div class="offer_item" @click="onItemClicked">
    <img v-if="isImageSource" :src="require(`@/../images/offers/${type}/${id}/${id}_0.jpeg`)" alt="Image">
    <img v-else :src="require(`@/assets/img/image_placeholder.png`)" alt="Image">
    <div class="item_details">
      <div class="title">{{ title }}</div>
      <div class="price">Price: {{ price }} $</div>
      <div class="capacity">Capacity: {{ capacity }} people</div>
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

div.capacity {
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
  height: 250px;
  width: 250px;
  opacity: 0.9;
  flex-shrink: 0;
  margin: 2%;
}
</style>