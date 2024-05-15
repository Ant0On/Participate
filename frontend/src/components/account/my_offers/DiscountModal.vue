<script setup>
import { defineProps, toRef, watch, ref, defineEmits} from 'vue';
import NumberInput from "@/components/ui/NumberInput.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
const props = defineProps({
  id: String,
  isSetDiscountModal: Boolean,
  discount: String,
})
const emits = defineEmits(['discountChanged'])
const isModalVisible = toRef(props.isSetDiscountModal)
const discount = toRef(props.discount)

function setDiscount(){
  fetchWrapper.put(`/api/host/discount/${props.id}/`, {
    discount: Number(discount.value)
  }).then(()=>{
    emits('discountChanged', discount.value)
    isModalVisible.value = false
  }).catch(error => {
    isModalVisible.value = false
  })
}
watch(() => props.isSetDiscountModal, (value) => {
  isModalVisible.value = value
})
</script>
<template>
  <Teleport to="body">
    <Transition name="modal-fade">
      <div v-if="isModalVisible" class="modal_wrapper">
        <div class="data">
          <p> Change offer discount: </p>
          <NumberInput v-model="discount" label-text="Offer percentage discount" :min="1" :max="100" width="70" align-items="center"/>
          <div class="discount_buttons">
            <button class="button_basic" @click="isModalVisible = !isModalVisible " > Close</button>
            <button class="button_basic" @click="setDiscount"> Set discount </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
<style scoped>
div.discount_buttons{
  display: flex;
  flex-direction: row;
  column-gap: 10%;
}
p {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.4rem;
  padding: 2%;
  margin-top: 2%;
  font-weight: 400;
  text-align: center;
}
div.data{
  background-color: white;
  height: max(40%, 350px);
  width: 30%;
  display: flex;
  align-items: center;
  flex-direction: column;
  row-gap: 10%;
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
.button_basic {
  all: unset;
  font-family: "IBM Plex Sans", Helvetica, serif !important;
  background-color: #efefef;
  color: black;
  border-radius: 6px;
  box-sizing: border-box;
  padding: 4px 16px;
  border: 1px solid #808080;
  border-radius: 6px;
  width: 80px;
  height: 35px;
  display: flex;
  justify-content: center;
  align-self: center;
  align-items: center;
  text-align: center;
}
</style>