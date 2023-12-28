<script setup>
import {defineEmits, defineProps} from 'vue'

const props = defineProps({
  text: String,
  isActive: Boolean,
})

const emits = defineEmits(['itemClicked'])

function onClick() {
  emits('itemClicked', props.text)
}

</script>

<template>
  <div :class="{'active':isActive, 'not_active':!isActive}" @click="onClick">
    <p>{{ text }}</p>
  </div>
</template>

<style scoped>
div {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  height: max(10%, 50px);
  padding: 10px;
}

.active {
  background-color: #c8c8c8;
}

.not_active {
  background-color: #efefef;
}

p {
  color: black;
  font-family: "Poppins", Helvetica;
  font-size: 1rem;
  font-weight: 400;
}


</style>
