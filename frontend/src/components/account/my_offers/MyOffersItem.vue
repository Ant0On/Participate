<script setup>
import {defineProps, ref, computed, toRef, watch} from 'vue';
import DiscountModal from "@/components/account/my_offers/DiscountModal.vue";
import PriceModal from "@/components/account/my_offers/PriceModal.vue";

const props = defineProps({
  id: String,
  name: String,
  offerType: String,
  withAnimals: Boolean,
  offerId: String,
  discount: String,
  price: String
});

const discountRef = toRef(props.discount)
const priceRef = toRef(props.price)

const isSetDiscountModal = ref(false)
const isSetPriceModal = ref(false)


function onDiscountChanged(discount){
  discountRef.value = discount
}

function onPriceChanged(price){
  priceRef.value = price
}

const isImageSource = computed(() =>{
  try{
    require(`@/../images/offers/${props.offerId}/${props.offerId}_0.jpeg`)
    return true
  }
  catch{
    return false
  }
})
</script>

<template>
  <div class="my_offer_item">
    <img v-if="isImageSource" :src="require(`@/../images/offers/${props.offerId}/${props.offerId}_0.jpeg`)" alt="Image">
    <img v-else :src="require(`@/assets/img/image_placeholder.png`)" alt="Image">
    <div class="my_offer_item_details">
      <div class="title">{{ name }}</div>
      <div class="summary_data">
        <div class="details">
          <div class="field">Type: {{ offerType }}</div>
          <div class="field">Animals: {{ (withAnimals) ? 'Yes' : 'No' }}</div>
        </div>
        <div class="details">
          <div class="field">Price: {{ price }}</div>
          <div class="field">Discount: {{ discountRef }}%</div>
        </div>
        <div class="change_data">
          <button class="button_basic" @click="isSetDiscountModal = !isSetDiscountModal">
            Set discount
          </button>
          <DiscountModal :id="offerId" :discount="discount" :is-set-discount-modal="isSetDiscountModal" @discount-changed="onDiscountChanged"/>
        </div>
        <div class="change_data">
          <button class="button_basic" @click="isSetPriceModal = !isSetPriceModal">
            Set price
          </button>
          <PriceModal :id="offerId" :price="price" :is-set-price-modal="isSetPriceModal" @price-change="onPriceChanged"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.button_basic {
   all: unset;
   font-family: "IBM Plex Sans", Helvetica, serif !important;
   background-color: #efefef;
   color: black;
   border-radius: 6px;
   box-sizing: border-box;
   padding: 4px 16px;
   border: 1px solid #808080;
   border-radius: 6px;
   width: 80px;
   height: 35px;
   display: flex;
   justify-content: center;
   align-self: center;
   align-items: center;
   text-align: center;
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

div.my_offer_item {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  row-gap: 10px;
  margin-top: 1%;
  background-color: #E6E6E6;
  border-radius: 10px;
}

div.my_offer_item_details {
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