<script setup>
import {defineEmits, defineProps, ref} from 'vue'
import {v4 as uuidv4} from 'uuid';

let inputUUID = uuidv4();

const props = defineProps({
  labelText: String,
  isRequired: Boolean,
  modelValue: String,
  min: Date,
  max: Date
})

const emits = defineEmits(['update:modelValue'])
const required = value => !!value || 'Required.'
const rules = ref({
  required: value => !!value || 'Required.',
  min: v => v.length >= 8 || 'Date must ',
  max: v => v.length >= 8 || 'Date must ',
})

</script>

<template>

  <v-text-field
      :id="inputUUID"
      :label="labelText"
      @input="$emit('update:modelValue', $event.target.value)"
      v-model="modelValue"
      class="w-100"
      clearable
      :rules="[isRequired && rules.required, min && rules.min, max && rules.max]"
      type="date"
  ></v-text-field>
</template>

<style scoped>
</style>