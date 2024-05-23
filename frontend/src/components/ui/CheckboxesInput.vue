<script setup>
import {defineProps, defineEmits, ref} from 'vue';
import { v4 as uuidv4 } from 'uuid';

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: {
    type: Boolean,
    default: false
  },
  modelValue: String,
  items: Array,
});

const emits = defineEmits(['update:modelValue']);
const inputUUID = uuidv4();

</script>

<template>
  <v-select
      chips
      :id="inputUUID"
      :label="labelText"
      :items="items"
      :placeholder="placeholder"
      @update:modelValue="$emit('update:modelValue', $event)"
      v-model="modelValue"
      class="w-100"
      clearable
      multiple
  ></v-select>
</template>

<style scoped>
</style>
