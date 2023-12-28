<script setup>
import {defineEmits, defineProps} from 'vue'

const props = defineProps({
  text: String,
  isActive: Boolean,
})

const emits = defineEmits(['itemClicked'])

function handleMouseOut(element) {
  if (!props.isActive)
  {
    element.classList.remove('active')
    element.classList.add('not_active')
  }

}

function handleMouseOver(element) {
  element.classList.remove('not_active')
  element.classList.add('active')
}

function onClick() {
  emits('itemClicked', props.text)
}

</script>

<template>
  <div @mouseover="handleMouseOver($el)" @mouseout="handleMouseOut($el)"
       :class="{'active':isActive, 'not_active':!isActive}" @click="onClick">
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
