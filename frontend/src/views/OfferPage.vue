<script setup>
import {defineProps, ref, onMounted} from 'vue';

import OfferDetailSummary from "@/components/detail/OfferDetailSummary.vue";
import OfferDetailDescription from "@/components/detail/OfferDetailDescription.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import calculatePriceAfterDiscount from "@/_helpers/calculate-price-after-discount";

const props = defineProps({
  type: String,
  id: String,
})


const offer = ref({
  name: '',
  price: 0,
  location: '',
  numberOfPeople: 0,
  userID: 0,
})

const isDescription = ref(true)

async function getOfferDetails() {
  const response = await fetchWrapper.get(`/api/offers/${props.id}`)

  const responseData = response.data

  const priceAfterDiscount = calculatePriceAfterDiscount(responseData['price'], responseData['discount'])
  offer.value = {
      'userID': responseData["user_id"],
      'location': responseData["country_name"] + ', ' + responseData["town_name"],
      'description': responseData["description"],
      'name': responseData["name"],
      'price': priceAfterDiscount,
      'numberOfPeople': responseData["max_people"],
    }

}

onMounted(async () => {
  await getOfferDetails();
});

</script>

<template>
  <div class="event_page">
    <div class="item_detail">
      <OfferDetailDescription :type="type" :name="offer.name" :price="offer.price" :location="offer.location"
                              :numberOfPeople="offer.numberOfPeople" :userId="offer.appUserID" :description="offer.description"
                              v-if="isDescription" @move-to-summary="isDescription = !isDescription"/>
      <OfferDetailSummary :price="offer.price" :image="offer.image" :id="id" v-else/>
    </div>
  </div>
</template>

<style scoped>
div.item_detail{
  margin: 5% 5% 1% 5%;
  background-color: white;
}
</style>