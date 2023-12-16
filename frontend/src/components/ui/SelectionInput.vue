<script setup>
import {defineEmits, defineProps} from 'vue'
import {v4 as uuidv4} from 'uuid';

let inputUUID = uuidv4();

const props = defineProps({
  labelText: String,
  placeholder: String,
  isRequired: Boolean,
  modelValue: String,
  items: Array,
  width: {
    type: String,
    default: "300px"
  }
})

const emits = defineEmits(['update:modelValue'])
</script>

<template>
  <div class="text_input">
    <label :for="inputUUID">{{ labelText }}<span v-if="isRequired">*</span></label>
    <select :id="inputUUID" :placegolder="placeholder" :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)">
      <option v-for="item in items"> {{ item }} </option>
    </select>
  </div>
</template>

<style scoped>
select, option{
  width: v-bind(width);
  height: 40px;
  background-color: var(--surfacelight);
  border: 1px;
}
.text_input {
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  position: relative;
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