<script setup>
import { defineProps, defineEmits } from 'vue'

function handleOnClick(){

}

const props = defineProps({
  text: String
})

</script>

<template>
  <button class="button" @click="handleOnClick">
    {{ text }}
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
  border-radius: 6px;
}

</style>