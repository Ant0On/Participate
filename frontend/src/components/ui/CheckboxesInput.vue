<script setup>
import { defineProps, defineEmits, ref, watchEffect } from 'vue';
import { v4 as uuidv4 } from 'uuid';

const props = defineProps({
  labelText: String,
  modelValue: Array,
  items: Array,
});

const emits = defineEmits(['update:modelValue']);

const selectedItems = ref(props.modelValue || []);

const inputUUID = uuidv4();

function generateCheckboxId(index) {
  return `checkbox-${inputUUID}-${index}`;
}

watchEffect(() => {
  emits('update:modelValue', selectedItems.value);
});
</script>

<template>
  <div>
    <label :for="inputUUID">{{ labelText }}</label>
    <div>
      <input
          type="checkbox"
          :id="generateCheckboxId(index)"
          v-for="(item, index) in items"
          :key="index"
          :value="item"
          v-model="selectedItems"
      />
      <label :for="generateCheckboxId(index)" v-for="(item, index) in items" :key="'label-' + index">{{ item }}</label>
    </div>
  </div>
</template>

<style scoped>
</style>
