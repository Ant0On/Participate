<script setup>

import {defineProps, reactive} from 'vue';
import {storeToRefs} from 'pinia';
import * as Yup from 'yup';



import DateInput from "@/components/ui/DateInput.vue";
import NumberInput from "@/components/ui/NumberInput.vue";
import {useSearchStore} from "@/stores/search.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const searchStore = useSearchStore();

const {numberOfPeople, dateFrom, dateTo} = storeToRefs(searchStore)

const errors = reactive({
  apiError: ""
})

const props = defineProps({
  price: '',
  image: '',
})
const schema = Yup.object().shape({
  dateFrom: Yup.date().required('Starting date is required'),
  dateTo: Yup.date().min(Yup.ref('dateFrom')),
  numberOfPeople: Yup.number().required('Number of people is required'),
});

function submitOffer(){

  schema.validate({
    'dateFrom': dateFrom.value,
    'dateTo': dateTo.value,
    'numberOfPeople': numberOfPeople.value,
  })
      .then(()=>{
        fetchWrapper.post()
      })
      .catch(error =>{
        errors.apiError = "Invalid data!"
      })
}
</script>

<template>
<div class="offer_detail_summary">
  <div class="offer_detail_summary_header">
    <img src="@/assets/img/magnifying_glass.jpg" alt="Magnifying glass">
    <div class="offer_detail_summary_title">
      <div class="title">
        Details
      </div>
      <img class="magnifying_glass_mirror_small" src="@/assets/img/magnifying_glass.jpg" alt="Magnifying glass">
    </div>
    <img class="magnifying_glass_mirror" src="@/assets/img/magnifying_glass.jpg" alt="Magnifying glass">
  </div>
  <div class="offer_detail_summary_info">
    <div class="offer_detail_summary_fields">
      <div class="offer_detail_summary_data">
        <NumberInput v-model="numberOfPeople" label-text="Number of people"/>
        <DateInput v-model="dateFrom" label-text="Arrival date"/>
        <DateInput v-model="dateTo" label-text="Departure date"/>
        <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
      </div>
      <div class="offer_detail_payment">
        <div class="offer_detail_price">
          Price: {{ price }}
        </div>
        <div class="offer_detail_summary_book_button">
          <button class="book" @click="submitOffer()">
            <span>Book</span>
          </button>
        </div>
      </div>
    </div>
    <div class="offer_detail_summary_image">
      <img class="offer_detail_summary_image" :src="image" alt="Offer image">
    </div>
  </div>
</div>
</template>

<style scoped>
div.errors{
  font-family: "Sarabun", Helvetica;

}
div.offer_detail_summary{
  display: flex;
  flex-direction: column;
  padding: 1% 3% 1% 3%;
}
div.offer_detail_summary_header{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 5%;

}
div.offer_detail_summary_title{
  font-weight: 800;
  font-family: "Poppins-ExtraBold", Helvetica;
  font-size: 2.5rem;
  display: flex;
  flex-direction: row;
  column-gap: 5%;
  align-self: flex-end;
}
.magnifying_glass_mirror_small{
  transform: scaleX(-1);
  height: 2.5rem;
  width: 2.5rem;
}
.magnifying_glass_mirror{
  transform: scaleX(-1);
}
div.offer_detail_summary_info{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
div.offer_detail_summary_fields{
  flex-grow: 1;
}
div.offer_detail_summary_image{
  flex-grow: 1;
  display: flex;
}
img.offer_detail_summary_image{
  align-self: center;
  width: 100%;
  height: 100%;
  padding: 0 10% 5% 20%;
}
div.offer_detail_summary_data{
  display: flex;
  flex-direction: column;
  row-gap: 30px;
  padding: 0 5% 0 5%;
}
div.offer_detail_payment{
  display: flex;
  flex-direction: column;
  padding: 5% 5% 0 5%;
}
div.offer_detail_price{
  font-family: "Playfair Display-SemiBold", Helvetica;
  font-weight: 600;
  font-size: 1.5rem;
  padding-bottom: 5%;
}
div.offer_detail_summary_book_button{
  align-self: center;
  width: 15%;
  height: 30px;
  display: flex;
}
button.book{
  font-family: "Poppins-Regular", Helvetica;
  font-weight: 400;
  font-size: 1rem;
  border-radius: 6px;
  background-color: var(--systemwhite);
  border: 1px solid black;
  flex-grow: 1;
}
.book:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}
</style>