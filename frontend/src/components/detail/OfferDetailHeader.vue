<script setup>
import {defineProps, computed} from 'vue';

const props = defineProps({
  name: '',
  type: '',
  location: '',
  numberOfPeople: '',
  rooms: '',
})
const roomsText = computed(()=> (props.rooms > 1)? "rooms": "room")
const numberOfPeopleText = computed(()=> (props.numberOfPeople > 1)? "people": "person")
</script>

<template>
<div class="offer_detail_header">
  <div class="offer_detail_title">
    {{name}}
  </div>
  <div class="offer_detail_location">
    {{location}}
  </div>
  <div class="offer_detail_brief_info">
    <span class="offer_detail_type">{{type.charAt(0).toUpperCase() + type.slice(1)}}</span>
    <div class="delimiter"/>
    <span class="offer_detail_people">Max {{numberOfPeople}} {{ numberOfPeopleText}}</span>
    <div class="delimiter"/>
    <span v-if="rooms" class="offer_detail_rooms">{{rooms}} {{roomsText}}</span>
  </div>
</div>
</template>

<style scoped>
div.offer_detail_header{
  display: flex;
  flex-direction: column;
  padding: 2%;
  font-family: "Nokora-Bold", Helvetica;
}
div.offer_detail_title{
  font-weight: 600;
  font-size: 2rem;
}
div.offer_detail_location{
  font-family: "Nokora-Regular", Helvetica;
  font-weight: 400;
  font-size: 1.4rem;
}
div.offer_detail_brief_info{
  display: flex;
  flex-direction: row;
  font-family: "Nokora-Regular", Helvetica;
  font-weight: 400;
  font-size: 1.2rem;
}
div.delimiter{
  margin: 3px;
  height: 2px;
  width: 2px;
  color: black;
  background-color: black;
  align-self: center;

}

</style>