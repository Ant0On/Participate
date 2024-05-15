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
  min: v => v > props.min || `Date must be grater than ${props.min} `,
  max: v => v < props.max || `Date must be smaller than ${props.max}`,
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
      type="date"
      :rules="[isRequired && rules.required]"
  ></v-text-field>
<!--  :rules="[isRequired && rules.required, min && rules.min, max && rules.max]"-->
<!--  TODO: Create proper validation -->
</template>

<style scoped>
</style>