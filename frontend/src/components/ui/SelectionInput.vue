<script setup>
import {defineProps, defineEmits, ref} from 'vue';
import { v4 as uuidv4 } from 'uuid';

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: Boolean,
  modelValue: String,
  items: Array,
});

const emits = defineEmits(['update:modelValue']);
const inputUUID = uuidv4();

const rules = ref({
  required: value => !!value || 'Required.',
});
</script>

<template>
  <v-select
      :id="inputUUID"
      :label="labelText"
      :items="items"
      :placeholder="placeholder"
      @update:modelValue="$emit('update:modelValue', $event)"
      v-model="modelValue"
      class="w-100"
      clearable
      :rules="[isRequired && rules.required]"
  ></v-select>
</template>

<style scoped>
</style>
