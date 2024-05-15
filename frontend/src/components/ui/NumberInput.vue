<script setup>
import {defineEmits, defineProps, ref} from 'vue'
import {v4 as uuidv4} from 'uuid';

let inputUUID = uuidv4();

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: Boolean,
  modelValue: String | Number,
  isActive: {
    type: Boolean,
    default: true
  },
  min: {
    type: Number,
    default: 1
  },
  max: {
    type: Number,
  },
  width: {
    default: "100",
    type: String,
  }
})

const emits = defineEmits(['update:modelValue'])
const required = value => !!value || 'Required.'
const rules = ref({
  required: value => !!value || 'Required.',
  min: v => v >= props.min || `Minimal value must be higher than ${props.min - 1}`,
  max: v => v <= props.max || `Minimal value must be lower than ${props.min + 1}`,
})

const widthClass = ref(`w-${props.width}`)
</script>

<template>

  <v-text-field
      :label="labelText"
      :id="inputUUID"
      @input="$emit('update:modelValue', $event.target.value)"
      :disabled="!isActive"
      :rules="[isRequired && rules.required, rules.min, max && rules.max]"
      v-model="modelValue"
      :placeholder="placeholder"
      :class="widthClass"
      type="number"
  ></v-text-field>

</template>

<style scoped>
</style>