<script setup>
import {computed, onMounted, reactive, ref} from 'vue';

import NumberInput from "@/components/ui/NumberInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import CheckButtonInput from "@/components/ui/CheckButtonInput.vue";
import SelectionInput from "@/components/ui/SelectionInput.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";


const newOffer = ref({
  offerType: '',
  name: '',
  description: '',
  price: '',
  maxPeople: '',
  isAnimalFriendly: false,
  country: '',
  city: '',
  images: [],
});

const errors = reactive({
  apiError: '',
  offerType: '',
  offerInfo: '',
  country: '',
  image: '',
});

const userStore = useAuthStore();
const user = userStore.user;
const isOfferTypeFilled = computed(() => newOffer.value.offerType !== '');
const isOfferInfoFilled = computed(() => checkIfOfferInfoIsFilled())
const isOfferCountryFilled = computed(() => newOffer.value.country !== '' && newOffer.value.city !== '')
const isOfferImageFilled = computed(() => newOffer.value.images.length !== 0)
const isAddingNewOffer = ref(false)
const addedNewOffer = ref(false)
const countries = ref([])
const offerTypes = ['Accommodation', 'Activities', 'Events']
const offerTypesId = {
  'Accommodation': 'accommodation',
  'Events': 'event',
  'Activities': 'activity'
}

async function getCountryId(countryName) {
  const response = await fetchWrapper.get('/api/country/get/all')
  return response.data.filter((country) => country.Name === countryName)[0].ID
}

async function onSubmit() {
  await fetchWrapper.post('/api/town/add', {
    name: newOffer.value.city,
    country_id: await getCountryId(newOffer.value.country),
  }).then((data) => {
        const imageFiles = dataURLsToFiles(newOffer.value.images, 'image');

        fetchWrapper.post('/api/host/create', {
              name: newOffer.value.name,
              description: newOffer.value.description,
              price: newOffer.value.price,
              max_people: newOffer.value.maxPeople,
              is_animal_friendly: newOffer.value.isAnimalFriendly,
              is_recommended: false,
              offer_type: offerTypesId[newOffer.value.offerType],
              host_id: user.ID,
              town_id: data.town.ID,
              images: imageFiles
            },
            "multipart/form-data")
            .then(() => {
              newOffer.value = {
                offerType: '',
                name: '',
                description: '',
                price: '',
                maxPeople: '',
                isAnimalFriendly: false,
                country: '',
                city: '',
                images: [],
              }
              isAddingNewOffer.value = false;
              addedNewOffer.value = true;
              errors.apiError = null
            }).catch(error => {
          errors.apiError = "Something went wrong - " + error
        })
      }
  ).catch((error) => {
    errors.apiError = "A problem occurred during addition of an offer! " + error.message;
  })
}

async function uploadImage(imageInput) {
  const images = imageInput.target.files;
  const promises = [];

  for (const image of images) {
    const reader = new FileReader();
    promises.push(
        new Promise((resolve) => {
          reader.onload = (source) => {
            resolve(source.target.result);
          };
          reader.readAsDataURL(image);
        })
    );
  }

  Promise.all(promises).then((imageDataArray) => {
    newOffer.value.images = imageDataArray;
  });
}

function dataURLtoFile(dataURL, fileName) {
  const arr = dataURL.split(',');
  const mime = arr[0].match(/:(.*?);/)[1];
  const bstr = atob(arr[1]);
  let n = bstr.length;
  const u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new File([u8arr], fileName, {type: mime});
}

function dataURLsToFiles(dataURLs, fileNameBase) {
  return dataURLs.map((dataURL, index) => {
    const fileName = `${fileNameBase}_${index}.jpeg`;
    return dataURLtoFile(dataURL, fileName);
  });
}

function checkIfOfferInfoIsFilled() {
  let offerValues = newOffer.value;
  if (
      offerValues.name === '' ||
      offerValues.description === '' ||
      offerValues.price === '' ||
      offerValues.maxPeople === '' ||
      offerValues.isAnimalFriendly === ''
  ) {
    errors.offerInfo = 'Please fill in all offer information fields.';
    return false;
  } else {
    errors.offerInfo = '';
    return true;
  }
}

onMounted(async () => {
  const response = await fetchWrapper.get('/api/country/get/all')
  countries.value = response.data.map((country) => country.Name)
})

</script>

<template>
  <div class="new_offer">
    <div v-if="!isAddingNewOffer && !addedNewOffer" class="new_offer_start">
      <p class="new_offer_text">Add new amazing experience!</p>
      <button v-if="!isAddingNewOffer" class="button_basic" @click="isAddingNewOffer = !isAddingNewOffer">
        Start now!
      </button>
    </div>
    <Transition name="bounce">
      <div v-if="isAddingNewOffer" class="new_offer_info">
        <p class="new_offer_text">Fill all information below to add a new experience!</p>
        <SelectionInput v-model="newOffer.offerType" label-text="Offer type" placeholder="Type" :items="offerTypes"
                        width="100%"/>
      </div>
    </Transition>
    <Transition name="bounce">
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
    <Transition name="bounce">
      <div v-if="isOfferInfoFilled" class="new_offer_info">
        <SelectionInput v-model="newOffer.country" label-text="Country" :items="countries" width="100%"/>
        <TextInput v-model="newOffer.city" label-text="City" width="100%"/>
      </div>
    </Transition>
    <Transition name="bounce">
      <div v-if="isOfferCountryFilled" class="new_offer_image">
        <div class="upload_image">
          <div v-for="(image, index) in newOffer.images" :key="index" class="image_preview">
            <img :src="image" class="preview_image" alt="offerImage"/>
          </div>
          <div id="image_input">
            <label for="image_upload">Add photos</label>
            <input id="image_upload" type="file" accept="image/jpeg, image/png, image/jpg" multiple
                   @change="uploadImage"/>
          </div>
        </div>
      </div>
    </Transition>
    <Transition name="bounce">
      <div class="submit_container">
        <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
        <div class="errors" v-if="errors.offerType">{{ errors.offerType }}</div>
        <div class="errors" v-if="errors.offerInfo">{{ errors.offerInfo }}</div>
        <div class="errors" v-if="errors.country">{{ errors.country }}</div>
        <div class="errors" v-if="errors.image">{{ errors.image }}</div>
        <button v-if="isOfferImageFilled" class="button_basic" @click="onSubmit">
          Create an offer
        </button>
      </div>
    </Transition>
    <Transition name="bounce">
      <div v-if="addedNewOffer" class="new_offer_start">
        <p class="new_offer_text">Add another amazing experience!</p>
        <button class="button_basic" @click="addedNewOffer = false">
          Add experience!
        </button>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}

.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}

@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.10);
  }
  100% {
    transform: scale(1);
  }
}

div.submit_container {
  display: flex;
  align-items: center;
  flex-direction: column;
}

div.new_offer_image {
  display: flex;
  flex-direction: column;
}

div.image_preview {
  margin-bottom: 10px;
}

div#image_input {
  display: flex;
  flex-grow: 1;
}

label {
  border: 1px solid black;
  padding: 5%;
  border-radius: 5px;
  width: 150px;
  text-align: center;
}

input#image_upload {
  position: absolute;
  left: -99999rem
}

img.preview_image {
  padding-top: 5%;
  padding-bottom: 5%;
  width: 300px;
  height: 300px;
  align-self: center;

}

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

div.new_offer_start {
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
  text-align: center;
}

.button_basic:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}

div.errors {
  font-family: "Sarabun", Helvetica;
  margin-bottom: 5%;
  margin-top: 5%;
  text-align: center;
}
</style>