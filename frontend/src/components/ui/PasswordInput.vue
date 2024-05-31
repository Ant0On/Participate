<script setup>
import {ref, defineProps} from "vue";

const { modelValue, labelText, isActive } = defineProps({
  modelValue: String,
  labelText: String,
  isActive: {
    type: Boolean,
    default: true,
  },
});

const show = ref(false);
const rules = ref({
  required: value => !!value || 'Required.',
  min: v => v.length >= 8 || 'Min 8 characters',
})
</script>

<template>
  <v-text-field
      v-model="modelValue"
      :append-inner-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
      :rules="[rules.required, rules.min]"
      :type="show ? 'text' : 'password'"
      hint="At least 8 characters"
      :label="labelText"
      name="input-10-1"
      counter
      clearable
      @click:append-inner="show = !show"
      :disabled="!isActive"
      @input="$emit('update:modelValue', $event.target.value)"
  ></v-text-field>
</template>

<style scoped>
</style>