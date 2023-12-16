<script setup>
import {defineEmits, defineProps} from 'vue'

const props = defineProps({
  src: String
})

</script>

<template>
    <img class="image" :src="src" alt="Image"/>
</template>

<style scoped>

.image {
  height: auto;
  width: 100%;
  max-width: 300px;
  max-height: 300px;
  margin: 0 13% 120px 13%;
}

</style>