<script setup>

import {defineProps, defineEmits} from 'vue';

const props = defineProps({
  maxPage: Number,
  currentPage: Number,

})

const emits = defineEmits(['pageBack', 'pageForward'])

function pageGoBack(){
  emits('pageBack')
}
function pageGoForward(){
  emits('pageForward')
}

</script>

<template>

  <div class="switch_list_page">
    <div class="page_go_back" @click="pageGoBack">
      <p class="arrow_left"> &#10132; </p>
    </div>
    <div class="page_number_max_page">
      <p class="page_number"> {{ currentPage }} of {{ maxPage }} </p>
    </div>
    <div class="page_go_forward" @click="pageGoForward">
      <p class="arrow_right"> &#10132; </p>
    </div>
  </div>

</template>

<style scoped>
div.switch_list_page{
  display: flex;
  flex-grow: 1;
  flex-direction: row;
  justify-content: center;
  font-size: 1.5rem;
}
p.arrow_left{
  transform: scaleX(-1);
  padding-right: 1%;
}
p.arrow_right{
  padding-left: 1%;
}
div.page_number_max_page{
  margin-left: 2%;
  margin-right: 2%;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
p.page_number{

}

</style>