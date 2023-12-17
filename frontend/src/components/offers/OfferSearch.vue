<script setup>
import {defineEmits, defineProps, reactive} from 'vue';

import TextInput from "@/components/ui/TextInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";

const props = defineProps({
  location: String,
  dateFrom: String,
  dateTo: String,
  numberOfPeople: String,
})

const emit = defineEmits(['update:location', 'update:dateFrom', 'update:dateTo', 'update:numberOfPeople'])

let inputWidth = reactive({
  inputWidth: window.innerWidth / 5
})

window.addEventListener("resize", () => {
    inputWidth.inputWidth = window.innerWidth / 5;
});



</script>

<template>
  <div class="offer_search">
    <TextInput v-model="location" @input="$emit('update:location', $event.target.value)" :width="inputWidth.inputWidth+'px'"
               label-text="Location" placeholder="Location" :is-required="true"/>
    <DateInput v-model="dateFrom" @input="$emit('update:dateFrom', $event.target.value)" label-text="Date from"
               :is-required="true" :width="inputWidth.inputWidth+'px'"/>
    <DateInput v-model="dateTo" @input="$emit('update:dateTo', $event.target.value)" label-text="Date to"
               :is-required="false" :width="inputWidth.inputWidth+'px'"/>
    <NumberInput v-model="numberOfPeople" @input="$emit('update:numberOfPeople', $event.target.value)"
                 label-text="Number of people" placeholder="People" :is-required="true" :width="inputWidth.inputWidth+'px'"/>
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
}


</style>