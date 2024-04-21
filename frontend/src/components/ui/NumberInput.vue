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
})

const emits = defineEmits(['update:modelValue'])
const required = value => !!value || 'Required.'
const rules = ref({
  required: value => !!value || 'Required.',
  min: v => v >= props.min || `Minimal value must be higher than ${props.min - 1}`,
  max: v => v <= props.max || `Minimal value must be lower than ${props.min + 1}`,
})
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
      class="w-100"
      type="number"
  ></v-text-field>

</template>

<style scoped>
input[type=number]{
  width: v-bind(width);
  height: 40px;
  background-color: var(--surfacelight);
  border: 1px;
  resize: vertical;
}
.text_input {
  display: flex;
  align-items: v-bind(alignItems);
  flex-direction: column;
  position: relative;
  width: 100%;
}

label {
  font-family: "IBMPlex Sans-Regular", Helvetica;
  font-style: normal;
  font-weight: 500;
  color: var(--text-secondary-grey2);
  line-height: 150%;
}

span {
  font-family: "IBMPlex Sans-Regular", Helvetica;
  font-style: normal;
  color: var(--systemred);
}

input:focus {
  border: 1px;
  color: var(--text-link)
}
</style>