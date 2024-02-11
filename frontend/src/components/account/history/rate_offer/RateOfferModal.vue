<script setup>
import { defineProps, ref } from 'vue';
import RateOffer from "@/components/account/history/rate_offer/RateOffer.vue";

const props = defineProps({
  offersToGrade: {
    type: Array,
    default: []
  }
})
const isModalVisible = ref(false);
const offerToGradeIterator = ref(props.offersToGrade[Symbol.iterator]())
const offerData = ref(offerToGradeIterator.value.next())

function nextOffer(){
 offerData.value = offerToGradeIterator.value.next()
}

</script>
<template>
  <Teleport to="body">
    <Transition name="modal-fade">
      <div v-if="true" class="modal_wrapper">
        <div class="data">
          <RateOffer :offerData="offerData" @next-offer="nextOffer"/>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
<style scoped>
div.data{
  background-color: white;
  height: 40%;
  width: 30%;
}
div.modal_wrapper {
  position: fixed;
  left: 0;
  top: 0;

  z-index: 500;

  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.2);

  display: grid;
  place-items: center;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: 0.25s ease all;
}
</style>