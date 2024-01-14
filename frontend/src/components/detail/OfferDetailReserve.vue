<script setup>

import { defineEmits } from 'vue';
import {storeToRefs} from 'pinia';

import NumberInput from "@/components/ui/NumberInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();

const {numberOfPeople, dateFrom, dateTo} = storeToRefs(searchStore)
const emits = defineEmits(['moveToSummary'])
</script>

<template>
<div class="offer_detail_reserve">
  <div class="offer_detail_reserve_fields">
    <NumberInput v-model="numberOfPeople" label-text="Number of people" />
    <DateInput v-model="dateFrom" label-text="Arrival date"/>
    <DateInput v-model="dateTo" label-text="Departure date"/>
  </div>
  <div class="reserve_button">
    <button class="reserve" @click="$emit('moveToSummary')">
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
</style>