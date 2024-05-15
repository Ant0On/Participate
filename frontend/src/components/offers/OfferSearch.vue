<script setup>
import {defineEmits, defineProps} from 'vue';

import TextInput from "@/components/ui/TextInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";

const props = defineProps({
  location: String,
  dateFrom: String,
  dateTo: String,
  numberOfPeople: Number | String,
})

const emit = defineEmits(['update:location', 'update:dateFrom', 'update:dateTo', 'update:numberOfPeople'])

let inputWidth = window.innerWidth / 5;

window.addEventListener("resize", () => {
  inputWidth = window.innerWidth / 5;
});

</script>

<template>
  <div class="offer_search">
    <TextInput :model-value="location" @input="$emit('update:location', $event.target.value)"
               label-text="Location" placeholder="Location" :is-required="true"/>
    <DateInput :model-value="dateFrom" @input="$emit('update:dateFrom', $event.target.value)" label-text="Date from"
               :is-required="true" :max="dateTo"/>
    <DateInput :model-value="dateTo" @input="$emit('update:dateTo', $event.target.value)" label-text="Date to"
               :is-required="false" :min="dateFrom"/>
    <NumberInput :model-value="numberOfPeople" @input="$emit('update:numberOfPeople', $event.target.value)"
                 label-text="Number of people" placeholder="People" :is-required="true"/>
  </div>
</template>

<style scoped>
div.offer_search {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
  column-gap: 10px;
  border: 1px solid;
  background-color: #E6E6E6;
  padding: 10px 30px;
  margin-left: 5%;
  margin-right: 5%;
  border-radius: 25px;
}

</style>