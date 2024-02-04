<script setup>

import {ref, reactive} from 'vue';
import * as Yup from 'yup';
import { router } from "@/router";
import { storeToRefs } from 'pinia';

import {useSearchStore} from "@/stores/search.store";
import SelectionInput from "@/components/ui/SelectionInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";
import SearchButton from "@/components/ui/SearchButton.vue";

const offerTypes = ['Accommodation', 'Activities', 'Events']
const offerRoutes = {
  'Accommodation': '/accommodations',
  'Activities': '/activities',
  'Events': '/events'
}

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore)
const searchOffer = ref({
  offerType: '',
  location: '',
  dateFrom: '',
  dateTo: '',
  numberOfPeople: '',
})

const errors = reactive({
  apiError: ""
})

function searchOffers() {
  schema.validate(searchOffer.value).then((value) => {
    location.value = value.location
    dateFrom.value = value.dateFrom.toISOString().split('T')[0]
    dateTo.value = value.dateTo.toISOString().split('T')[0]
    numberOfPeople.value = value.numberOfPeople
    router.push(offerRoutes[searchOffer.value.offerType])


  }).catch(error => {
    if(error === "Bad Request")
    {
      errors.apiError = "Incorrect data!"
    }
    else{
      errors.apiError = "Fill all required fields!"
    }

  })
}

const schema = Yup.object().shape({
  offerType: Yup.string().required('Offer type is required'),
  location: Yup.string().min(5).required('Location is required'),
  dateFrom: Yup.date().required('Date is required'),
  dateTo: Yup.date().min(Yup.ref('dateFrom')),
  numberOfPeople: Yup.number().min(1).required('Number of people is required'),
});


</script>

<template>
  <div class="home_offer_search">
    <p>What are you looking for?</p>
    <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    <SelectionInput v-model="searchOffer.offerType" label-text="Offer type" placeholder="Type" :is-required="true" :items="offerTypes"/>
    <TextInput v-model="searchOffer.location" width="100%" label-text="Location" placeholder="At least 5 characters" :is-required="true" />
    <DateInput v-model="searchOffer.dateFrom" label-text="Date from" :is-required="true"/>
    <DateInput v-model="searchOffer.dateTo" label-text="Date to" :is-required="true"/>
    <NumberInput v-model="searchOffer.numberOfPeople" label-text="Number of people" placeholder="People" :is-required="true"/>
    <SearchButton text="Search" width="100px" height="35px" @button-clicked="searchOffers" style="align-self: center"/>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;500;700&display=swap');

div.errors{
  font-family: "Sarabun", Helvetica;

}

.home_offer_search {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-start;
  row-gap: 10px;
  width: 100%;
  height: auto;
  border: 1px solid;
  background-color: #E6E6E6;
  padding: 10px 15px;
  border-radius: 25px;
}

p{
  font-family: "Playfair Display-Bold", serif;
  font-weight: 700;
  font-size: 1rem;
}

</style>