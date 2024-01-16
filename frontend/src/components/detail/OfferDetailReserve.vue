<script setup>

import { defineEmits, reactive } from 'vue';
import * as Yup from 'yup';
import {storeToRefs} from 'pinia';

import NumberInput from "@/components/ui/NumberInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();

const {numberOfPeople, dateFrom, dateTo} = storeToRefs(searchStore)
const emit = defineEmits(['moveToSummary'])
const errors = reactive({
  apiError: ""
})
const schema = Yup.object().shape({
  dateFrom: Yup.date().required('Starting date is required'),
  dateTo: Yup.date().min(Yup.ref('dateFrom')),
  numberOfPeople: Yup.number().min(1).required('Number of people is required'),
});

function moveToSummary(){
  schema.validate({
    'dateFrom': dateFrom.value,
    'dateTo': dateTo.value,
    'numberOfPeople': numberOfPeople.value,
  })
      .then(()=>{
        emit('moveToSummary')
      })
      .catch(error =>{
          errors.apiError = "Invalid data!"

      })

}
</script>

<template>
<div class="offer_detail_reserve">
  <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
  <div class="offer_detail_reserve_fields">
    <NumberInput v-model="numberOfPeople" label-text="Number of people" />
    <DateInput v-model="dateFrom" label-text="Arrival date"/>
    <DateInput v-model="dateTo" label-text="Departure date"/>
  </div>
  <div class="reserve_button">
    <button class="reserve" @click="moveToSummary()">
      <span>Reserve</span>
    </button>
  </div>
</div>
</template>

<style scoped>
div.offer_detail_reserve{
  display: flex;
  flex-direction: column;
}
div.offer_detail_reserve_fields{
  display: flex;
  flex-direction: row;
  column-gap: 5%;
  padding: 2% 5% 2% 5%;
}
div.reserve_button{
  display: flex;
  width: 10%;
  height: 30px;
  align-self: center;
}
button.reserve{
  font-family: "Poppins-Regular", Helvetica;
  font-weight: 400;
  font-size: 1rem;
  flex-grow: 1;
  border-radius: 6px;
  background-color: var(--systemwhite);
  border: 1px solid black;

}
.reserve:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}
div.errors{
  font-family: "Sarabun", Helvetica;
  padding: 2% 5% 2% 5%;
}

</style>