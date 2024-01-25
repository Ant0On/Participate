<script setup>
import {computed, ref, onMounted, watch} from 'vue';

import NumberInput from "@/components/ui/NumberInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import CheckButtonInput from "@/components/ui/CheckButtonInput.vue";
import SelectionInput from "@/components/ui/SelectionInput.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const offerTypes = ['Accommodation', 'Activities', 'Events']

async function onSubmit() {

}

const newOffer = ref({
  offerType: '',
  name: '',
  description: '',
  price: '',
  maxPeople: '',
  isAnimalFriendly: false,
  country: '',
  city: '',
});

const isOfferTypeFilled = computed(() => newOffer.value.offerType !== '');
const isOfferInfoFilled = computed(() => checkIfOfferInfoIsFilled())
const isOfferCountryFilled = computed(() => newOffer.value.country !== '' && newOffer.value.city !== '')
const isOfferImageFilled = computed(() => newOffer.value.country !== '' && newOffer.value.city !== '')

function checkIfOfferInfoIsFilled() {
  let offerValues = newOffer.value;
  return offerValues.name !== '' && offerValues.description !== '' && offerValues.price !== ''
      && offerValues.maxPeople !== '' && offerValues.isAnimalFriendly !== ''

}
const countries = ref([])
const countriesId = ref([])

onMounted(async ()=>{
  const response = await fetchWrapper.get('/api/country/get/all')

  countries.value = response.data.map((country) => country.Name)

  let countriesId = {}
  response.data.forEach((country)=>{
    countriesId[country.Name] = country.ID
  })

  countriesId.value = countriesId
})

</script>

<template>
  <div class="new_offer">
    <div class="new_offer_choose_type">
      <SelectionInput v-model="newOffer.offerType" label-text="Offer type" placeholder="Type" :items="offerTypes"/>
    </div>
    <Transition>
      <div v-if="isOfferTypeFilled" class="new_offer_info">
        <TextInput v-model="newOffer.name" label-text="Name"/>
        <TextInput v-model="newOffer.description" label-text="Description"/>
        <NumberInput v-model="newOffer.price" label-text="Price"/>
        <NumberInput v-model="newOffer.maxPeople" label-text="Max number of people"/>
        <CheckButtonInput :model-value="newOffer.isAnimalFriendly" @changed-value="newOffer.isAnimalFriendly = !newOffer.isAnimalFriendly" label-text="Is animal friendly?"/>
      </div>
    </Transition>
    <Transition>
      <div v-if="isOfferInfoFilled" class="new_offer_location">
        <SelectionInput v-model="newOffer.country" label-text="Country" :items="countries"/>
        <TextInput v-model="newOffer.city" label-text="City"/>
      </div>
    </Transition>
    <Transition>
      <div v-if="isOfferCountryFilled" class="new_offer_image">
        tutaj bedzie dodanie obrazka...
      </div>
    </Transition>
    <Transition>
      <button v-if="isOfferImageFilled" class="button_basic" @click="onSubmit">
        Create offer
      </button>
    </Transition>

  </div>
</template>

<style scoped>

</style>