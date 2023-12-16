<script setup>

import {ref} from 'vue';

import SelectionInput from "@/components/ui/SelectionInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";
import SearchButton from "@/components/ui/SearchButton.vue";

const offer_types = ['Accommodation', 'Activities', 'Events']

const searchOffer= ref({
  offer_type: '',
  location: '',
  date_from: '',
  date_to: '',
  number_of_people: '',
})

function searchOffers(){
  console.log(searchOffer.value)
}

</script>

<template>
  <div class="home_offer_search">
    <p>What are you looking for?</p>
    <SelectionInput v-model="searchOffer.offer_type" label-text="Offer type" placeholder="Type" :is-required="true" :items="offer_types"/>
    <TextInput v-model="searchOffer.location" width="300px" label-text="Location" placeholder="Location" :is-required="true" />
    <DateInput v-model="searchOffer.date_from" label-text="Date from" :is-required="true"/>
    <DateInput v-model="searchOffer.date_to" label-text="Date to" :is-required="false"/>
    <NumberInput v-model="searchOffer.number_of_people" label-text="Number of people" placeholder="People" :is-required="true"/>
    <SearchButton text="Search" width="100px" height="35px" @button-clicked="searchOffers" style="align-self: center"/>
  </div>
</template>

<style scoped>
div.home_offer_search {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-start;
  row-gap: 10px;
  width: 335px;
  border: 1px solid;
  background-color: #E6E6E6;
  padding: 10px 15px;
}
p{
  font-family: "Playfair Display-Bold", Helvetica !important;
  font-weight: 700;
  font-size: 1rem;
}

</style>