<script setup>
import {ref} from 'vue';

import NumberInput from "@/components/ui/NumberInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import CheckButtonInput from "@/components/ui/CheckButtonInput.vue";
import SelectionInput from "@/components/ui/SelectionInput.vue";


async function onSubmit(){

}

</script>

<template>
<div class="new_offer">
  <div class="new_offer_choose_type">
    <SelectionInput label-text="Offer type"/>
  </div>
  <div class="new_offer_info">
    <TextInput label-text="Name"/>
    <TextInput label-text="Description"/>
    <NumberInput label-text="Price"/>
    <NumberInput label-text="Max number of people"/>
    <CheckButtonInput label-text="Is animal friendly?"/>
  </div>
  <div class="new_offer_location">
    <SelectionInput label-text="Country"/>
    <TextInput label-text="City"/>
  </div>
  <div class="new_offer_image">

  </div>
  <button class="button_basic" @click="onSubmit">
    Create offer
  </button>
</div>
</template>

<style scoped>

</style>