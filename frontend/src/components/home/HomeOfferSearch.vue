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
  schema.validate(searchOffer.value, { abortEarly: false })
      .then((value) => {
        location.value = value.location;
        dateFrom.value = value.dateFrom.toISOString().split('T')[0];
        dateTo.value = value.dateTo.toISOString().split('T')[0];
        numberOfPeople.value = value.numberOfPeople;
        router.push(offerRoutes[searchOffer.value.offerType]);
      })
      .catch(validationError => {
        errors.apiError = "Please fix the following issues:";
        if (validationError.inner) {
          validationError.inner.forEach(error => {
            switch (error.path) {
              case 'offerType':
                errors.apiError += `\n- ${error.message}`;
                break;
              case 'location':
                errors.apiError += `\n- ${error.message}`;
                break;
              case 'dateFrom':
                errors.apiError += `\n- ${error.message}`;
                break;
              case 'dateTo':
                errors.apiError += `\n- ${error.message}`;
                break;
              case 'numberOfPeople':
                errors.apiError += `\n- ${error.message}`;
                break;
              default:
                break;
            }
          });
        } else {
          errors.apiError += `\n- ${validationError.message}`;
        }
      });
}

const schema = Yup.object().shape({
  offerType: Yup.string().required('Please select an offer type'),
  location: Yup.string().min(3, 'Location must be at least 5 characters').required('Location is required'),
  dateFrom: Yup.date().required('Please select a starting date'),
  dateTo: Yup.date().required('Please select an ending date').min(Yup.ref('dateFrom'), 'Ending date must be after starting date'),
  numberOfPeople: Yup.number().min(1, 'Number of people must be at least 1').required('Number of people is required'),
});

</script>

<template>
  <div class="home_offer_search">
    <p>What are you looking for?</p>
    <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    <SelectionInput v-model="searchOffer.offerType" label-text="Offer type" placeholder="Type" :is-required="true" :items="offerTypes"/>
    <TextInput v-model="searchOffer.location" label-text="Location" placeholder="At least 5 characters" :is-required="true" :min="5" />
    <DateInput v-model="searchOffer.dateFrom" label-text="Date from" :is-required="true"/>
    <DateInput v-model="searchOffer.dateTo" label-text="Date to" :is-required="true"/>
    <NumberInput v-model="searchOffer.numberOfPeople" label-text="Number of people" placeholder="People" :is-required="true"/>
    <SearchButton text="Search" width="100px" height="35px" @button-clicked="searchOffers" style="align-self: center"/>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;500;700&display=swap');

div.errors {
  font-family: "Sarabun", Helvetica;
  color: #d9534f;
  margin-top: 10px;
  font-size: 0.9rem;
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