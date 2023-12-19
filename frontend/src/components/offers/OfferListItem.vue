<script setup>
import {defineEmits, defineProps} from 'vue'

const props = defineProps({
  location: String,
  image: String,
  title: String,
  description: String,
})

const emit = defineEmits(['itemClicked'])

function onItemClicked() {
  emit('itemClicked')
}
</script>

<template>
  <div class="offer_item" @click="onItemClicked">
    <img :src="image">
    <div class="location">{{ location }}</div>
    <div class="title">{{ title }}</div>
    <div class="description"> {{ description }}</div>
  </div>
</template>

<style scoped>
div.offer_item {
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  position: relative;
  row-gap: 10px;
}

</style>