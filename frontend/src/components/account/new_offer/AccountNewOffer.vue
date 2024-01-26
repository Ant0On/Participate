<script setup>
import {computed, onMounted, ref} from 'vue';

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
const isAddingNewOffer = ref(false)
const countries = ref([])
const countriesId = ref([])

onMounted(async () => {
  const response = await fetchWrapper.get('/api/country/get/all')

  countries.value = response.data.map((country) => country.Name)

  let countriesId = {}
  response.data.forEach((country) => {
    countriesId[country.Name] = country.ID
  })

  countriesId.value = countriesId
})

</script>

<template>
  <div class="new_offer">
    <Transition>
      <div v-if="!isAddingNewOffer" class="new_offer_start">
        <p class="new_offer_text">Add new amazing experience!</p>
        <button v-if="!isAddingNewOffer" class="button_basic" @click="isAddingNewOffer = !isAddingNewOffer">
          Start now!
        </button>
      </div>
    </Transition>
    <Transition>
      <div v-if="isAddingNewOffer" class="new_offer_info">
        <SelectionInput v-model="newOffer.offerType" label-text="Offer type" placeholder="Type" :items="offerTypes"
                        width="100%"/>
      </div>
    </Transition>
    <Transition>
      <div v-if="isOfferTypeFilled" class="new_offer_info">
        <TextInput v-model="newOffer.name" label-text="Name" width="100%"/>
        <TextInput v-model="newOffer.description" label-text="Description" width="100%"/>
        <NumberInput v-model="newOffer.price" label-text="Price"/>
        <NumberInput v-model="newOffer.maxPeople" label-text="Max number of people" width="100%"/>
        <CheckButtonInput :model-value="newOffer.isAnimalFriendly"
                          @changed-value="newOffer.isAnimalFriendly = !newOffer.isAnimalFriendly"
                          label-text="Is animal friendly?" width="100%"/>
      </div>
    </Transition>
    <Transition>
      <div v-if="isOfferInfoFilled" class="new_offer_info">
        <SelectionInput v-model="newOffer.country" label-text="Country" :items="countries" width="100%"/>
        <TextInput v-model="newOffer.city" label-text="City" width="100%"/>
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
div.new_offer {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  align-items: center;
  row-gap: 15px;
}

div.new_offer_info {
  display: flex;
  align-items: center;
  flex-direction: column;
  flex-grow: 1;
  width: 60%;
  row-gap: 15px;

}
div.new_offer_start{
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  flex-grow: 1;
  width: 60%;
  row-gap: 15px;
}

p.new_offer_text {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.6rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
  margin-top: 5%;
  margin-bottom: 5%;
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

  width: 150px;
  height: 35px;
  display: flex;
  justify-content: center;
  align-self: center;
  align-items: center;
}

.button_basic:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}
</style>