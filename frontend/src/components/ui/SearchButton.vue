<script setup>
import {defineEmits, defineProps} from 'vue'

const emit = defineEmits(['buttonClicked'])

function buttonClicked() {
  emit('buttonClicked')
}

const props = defineProps({
  text: String,
  width: String,
  height: String,
})

</script>

<template>
  <button class="button" @click="buttonClicked()">
    <span>{{ text }}</span>
  </button>
</template>


<style scoped>
.button {
  all: unset;
  font-family: "IBM Plex Sans", Helvetica, serif !important;
  background-color: var(--systemwhite);
  color: var(--text-link);
  border-radius: 6px;
  box-sizing: border-box;
  padding: 4px 16px;
  border: 1px solid;
  border-color: var(--text-link);
  width: 100px;
  height: 35px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.button:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}

</style>