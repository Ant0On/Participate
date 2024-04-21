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
select, option{
  width: v-bind(width);
  height: 40px;
  background-color: var(--surfacelight);
  border: 1px;
  resize: vertical;
}
.text_input {
  display: flex;
  align-items: flex-start;
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