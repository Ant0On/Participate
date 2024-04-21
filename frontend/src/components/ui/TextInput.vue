<script setup>
import {defineEmits, defineProps, ref} from 'vue'
import {v4 as uuidv4} from 'uuid';

let inputUUID = uuidv4();

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: Boolean,
  modelValue: String,
  isActive: {
    type: Boolean,
    default: true
  },
  rules: []
})

const emits = defineEmits(['update:modelValue'])
const required = value => !!value || 'Required.'

</script>

<template>

  <v-text-field
      :label="labelText"
      :id="inputUUID"
      @input="$emit('update:modelValue', $event.target.value)"
      :disabled="!isActive"
      :rules="rules || [isRequired && required]"
      v-model="modelValue"
      clearable
      :placeholder="placeholder"
      class="w-100"
  >

  </v-text-field>
</template>

<style scoped>
</style>