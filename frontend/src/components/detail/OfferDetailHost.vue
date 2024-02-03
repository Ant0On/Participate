<script setup>

import {defineProps} from 'vue';

const props = defineProps({
  host_first_name: '',
  offer_id: 0,
  host_detail: '',
  imagePath: '',
})
console.log("tata", props.imagePath)
</script>

<template>
<div class="offer_detail_host">
  <div class="offer_detail_host_picture">
    <div class="about_host">
      About the host
    </div>
<!--    <img class="detail_host_image" :src="`@/../${host_image_path}`" alt="Host picture">-->
    <div class="host_name">
      {{host_first_name }}
    </div>
  </div>
  <div class="offer_detail_host_description">
    {{ host_detail }}
  </div>
  <img class="detail_offer_image" :src="require(`@/../images/offers/${offer_id}.jpeg`)" alt="Offer image">
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