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
  rules: [],
  min: Number,
  max: Number
})

const emits = defineEmits(['update:modelValue'])
const rules = ref({
  required: value => !!value || 'Required.',
  min: v => v.length >= props.min || `Min ${props.min} characters`,
  max: v => v.length <= props.max || `Max ${props.max} characters`,
})

</script>

<template>

  <v-text-field
      :label="labelText"
      :id="inputUUID"
      @input="$emit('update:modelValue', $event.target.value)"
      :disabled="!isActive"
      :rules="[isRequired && rules.required, min && rules.min, max && rules.max]"
      v-model="modelValue"
      clearable
      :placeholder="placeholder"
      class="w-100"
  >

  </v-text-field>
</template>

<style scoped>
</style>