<script setup>
import defineProps from 'vue'

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
  max-width: 450px;
  max-height: 450px;
  margin: 0 13% 5% 13%;
}

</style>