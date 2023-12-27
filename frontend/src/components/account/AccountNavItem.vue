<script setup>
import { defineProps, defineEmits, ref, watch } from 'vue'

const props = defineProps({
  text: String,
  isActive: Boolean,
})

const emits = defineEmits(['itemClicked'])
function handleMouseOut(element) {
  if(!props.isActive)
  element.style.backgroundColor = "#efefef";
}
function handleMouseOver(element) {
  element.style.backgroundColor = "#c8c8c8";
}

const color = ref((props.isActive)? "#c8c8c8":"#efefef")

function onClick(){
  emits('itemClicked', props.text)
}

watch(props, (newValue, oldValue) =>{
  color.value = (newValue) ? "#c8c8c8":"#efefef"
})
</script>

<template>
  <div @mouseover="handleMouseOver($el)" @mouseout="handleMouseOut($el)" @click="onClick" :style="{'background-color': color}">
    {{ text }}
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  padding: 12px 26px;
  color: black;
  font-family: "Poppins", Helvetica;
  font-size: 1rem;
  font-weight: 400;
  margin-bottom: 5px;
}


</style>
