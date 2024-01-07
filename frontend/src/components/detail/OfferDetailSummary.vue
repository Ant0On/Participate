<script setup>

import {storeToRefs} from 'pinia';


import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore();

const {numberOfPeople, dateFrom, dateTo} = storeToRefs(searchStore)

</script>

<template>
<div class="offer_detail_summary">
  <NumberInput v-model="numberOfPeople"/>
  <DateInput v-model="dateFrom"/>
  <DateInput v-model="dateTo"/>

</div>
</template>

<style scoped>

</style>