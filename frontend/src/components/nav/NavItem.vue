<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  text: String,
  navigateTo: {
    type: String,
    validator: (prop) => ['home', 'accommodation', 'activities', 'events', 'login'].indexOf(prop) !== -1
  },
  isActive: Boolean
})

function handleMouseOut(element) {
  element.style.backgroundColor = "#efefef";
}
function handleMouseOver(element) {
  element.style.backgroundColor = "#3fc1c9";
}

</script>

<template>
  <div @mouseover="handleMouseOver($el)" @mouseout="handleMouseOut($el)">
      <a :href="navigateTo" :text="text"/>
      <img v-if="isActive" class="vector" alt="Vector" src="../../assets/img/vector-1.svg"/>
  </div>
</template>

<style scoped>
div {
  background-color: #efefef;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  padding: 12px 26px;
}

a{
  background-color: inherit;
  color: black;
  font-family: "Poppins", Helvetica;
  font-size: 1rem;
  font-weight: 400;
  margin-bottom: 5px;
}
img {
  align-self: flex-end;
}
@media (max-width: 900px) {
  img{
    align-self: center;
  }

}
</style>
