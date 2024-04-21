<script setup>
import {defineEmits, defineProps, ref} from 'vue'
import {v4 as uuidv4} from 'uuid';

let inputUUID = uuidv4();

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: Boolean,
  modelValue: String,
  items: Array,
})

const emits = defineEmits(['update:modelValue'])
const rules = ref({
  required: value => !!value || 'Required.',
})
</script>

<template>
  <v-select
      :id="inputUUID"
      :label="labelText"
      :items="items"
      :placeholder="placeholder"
      @input="$emit('update:modelValue', $event.target.value)"
      v-model="modelValue"
      class="w-100"
      clearable
      :rules="[isRequired && rules.required]"
  ></v-select>
</template>

<style scoped>
</style>