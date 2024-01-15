<script setup>

import {defineProps, ref} from 'vue';
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const props = defineProps({
  hostId: 0,
  offerImage: '@/assets/img/test.jpg',
})

const hostData = ref(fetchHostData())

function fetchHostData(){
  //const response = fetchWrapper.get(`/api/host/${props.hostId}`)

  return ""
}
</script>

<template>
<div class="offer_detail_host">
  <div class="offer_detail_host_picture">
    <div class="about_host">
      About the host
    </div>
    <img class="detail_host_image" :src="hostData.image" alt="Host picture">
    <div class="host_name">
      {{ hostData.firstName }}
    </div>
  </div>
  <div class="offer_detail_host_description">
    {{ hostData.detail }}
  </div>
  <img class="detail_offer_image" :src="props.offerImage" alt="Offer image">
</div>
</template>

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
  height: 70%;
  width: 70%;
  align-self: center;

}
img.detail_offer_image{
  padding: 1%;
  height: 20%;
  width: 20%;
}
</style>