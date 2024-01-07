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
  <NumberInput v-model="numberOfPeople"/>
  <DateInput v-model="dateFrom"/>
  <DateInput v-model="dateTo"/>

  <button class="button" @click="$emit('moveToSummary')">
    <span>Reserve</span>
  </button>
</div>
</template>

<style scoped>

</style>