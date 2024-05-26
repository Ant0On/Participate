<script setup>
import {computed, onMounted, reactive, ref} from 'vue';

import NumberInput from "@/components/ui/NumberInput.vue";
import TextInput from "@/components/ui/TextInput.vue";
import CheckButtonInput from "@/components/ui/CheckButtonInput.vue";
import SelectionInput from "@/components/ui/SelectionInput.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import CheckboxesInput from "@/components/ui/CheckboxesInput.vue";
import DateInput from "@/components/ui/DateInput.vue";
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'

const newOffer = ref({
  offerType: '',
  title: '',
  description: '',
  capacity: '',
  country: '',
  city: '',
  images: [],
});

const rooms = ref([]);

const newAccommodation = ref({
  numberOfRooms: '',
  accommodationType: '',
  isAnimalFriendly: false,
  pricePerDay: '',
  generalFacilities: [],
})

const newActivity = ref({
  skillLevel: '',
  price: '',
  activityType: '',
  dateRange: [],
  equipment: '',
})

const newEvent = ref({
  dateFrom: '',
  dateTo: '',
  price: '',
  eventType: ''
})

const newRoom = ref({
  number: '',
  name: '',
  description: '',
  capacity: '',
  area: '',
  roomFacilities: [],
})

const errors = reactive({
  apiError: '',
  offerType: '',
  offerInfo: '',
  country: '',
  image: '',
  dateRange: '',
});

const userStore = useAuthStore();
const user = userStore.user;
const isOfferTypeFilled = computed(() => newOffer.value.offerType !== '')
const isOfferTypeAccommodation = computed(() => newOffer.value.offerType === 'Accommodation')
const isAccommodationTypeWithSelectableRooms = computed(() =>
    newAccommodation.value.accommodationType === 'Hotel' ||
    newAccommodation.value.accommodationType === 'Hostel' ||
    newAccommodation.value.accommodationType === 'Guesthouse')
const isOfferTypeActivity = computed(() => newOffer.value.offerType === 'Activity')
const isOfferTypeEvent = computed(() => newOffer.value.offerType === 'Event')
const isOfferInfoFilled = computed(() => checkIfOfferInfoIsFilled())
const isRoomInfoFilled = computed(() => checkIfRoomInfoIsFilled())
const isOfferCountryFilled = computed(() => newOffer.value.country !== '' && newOffer.value.city !== '')
const isOfferImageFilled = computed(() => newOffer.value.images.length !== 0)
const isAddingNewOffer = ref(false)
const addedNewOffer = ref(false)
const countries = ref([])
const offerTypes = ['Accommodation', 'Activity', 'Event']
const accommodationTypes = ['Hotel', 'Hostel', 'Apartment', 'Villa', 'Guesthouse']
const generalFacilities = ["Swimming Pool", "Gym", "Spa", "Restaurant", "Bar", "Lounge", "Conference Room",
  "Business Center", "WiFi", "Parking", "24-Hour Front Desk", "Fitness Center", "Laundry Service", "Room Service",
  "Concierge Service", "Outdoor Pool", "Children's Playground", "Tennis Court", "Library", "Garden", "Sauna",
  "Jacuzzi", "Billiards Room", "Cinema", "Karaoke Room", "Bowling Alley", "BBQ Area", "Shuttle Service"]
const roomFacilities = ["Television", "Air Conditioning", "Mini Fridge", "Safe", "Coffee Maker", "Microwave",
  "Kettle", "Iron and Ironing Board", "Hair Dryer", "Desk", "Ocean View", "Mountain View", "Balcony", "Bathtub",
  "Shower", "WiFi", "Room Service", "Breakfast Included", "In-Room Safe", "Telephone", "DVD Player", "Alarm Clock",
  "Robes", "Slippers", "Toiletries", "Work Desk", "Sofa Bed", "Fireplace", "Refrigerator", "Dining Area"
];

const skillLevels = ['Beginner', 'Intermediate', 'Advanced']
const activityTypes = ['Indoor', 'Outdoor']
const equipmentList = ["Life Jacket", "Kayak", "Paddle", "Helmet", "Snowboard", "Sled", "Snowshoes", "Tent",
  "Sleeping Bag", "Backpack", "Hiking Boots", "Compass", "Map", "Binoculars", "Flashlight", "First Aid Kit",
  "Water Bottle", "Climbing Harness", "Climbing Rope", "Carabiners", "Rock Climbing Shoes", "Fishing Rod",
  "Bait", "Camera", "Swimsuit", "Snorkel Gear", "Tennis Racket", "Golf Clubs", "Bicycle"];
const eventTypes = ['Conference', 'Concert', 'Festival', 'Sports event']

async function getCountryId(countryName) {
  const response = await fetchWrapper.get('/api/country/get/all')
  return response.data.filter((country) => country.Name === countryName)[0].ID
}

async function onSubmit() {
  try {
    const townData = await fetchWrapper.post('/api/town/add', {
      name: newOffer.value.city,
      country_id: await getCountryId(newOffer.value.country),
    });

    const imageFiles = dataURLsToFiles(newOffer.value.images, 'image');

    const offerTypeHandlers = {
      'Accommodation': handleAccommodationOffer,
      'Activity': handleActivityOffer,
      'Event': handleEventOffer
    };

    if (offerTypeHandlers[newOffer.value.offerType]) {
      await offerTypeHandlers[newOffer.value.offerType](townData.town.ID, imageFiles);
    } else {
      throw new Error('Invalid offer type');
    }

    resetForm();

  } catch (error) {
    errors.apiError = "Something went wrong - " + error;
  }
}

async function handleAccommodationOffer(townID, imageFiles) {
  const accommodationData = await fetchWrapper.post('/api/host/accommodation/create', {
    title: newOffer.value.title,
    description: newOffer.value.description,
    capacity: newOffer.value.capacity,
    town_id: townID,
    user_id: user.ID,
    number_of_rooms: newAccommodation.value.numberOfRooms,
    type: newAccommodation.value.accommodationType.toLowerCase(),
    is_animal_friendly: newAccommodation.value.isAnimalFriendly,
    price_per_day: newAccommodation.value.pricePerDay,
    images: imageFiles
  }, "multipart/form-data");

  await fetchWrapper.post(`/api/host/accommodation/${accommodationData.offer.ID}/facilities/add`, {
    facilities: newAccommodation.value.generalFacilities
  });

  rooms.value.forEach(room => {
    room.accommodation_id = accommodationData.offer.ID;
  });

  if (newAccommodation.value.accommodationType !== 'Villa' && newAccommodation.value.accommodationType !== 'Apartment') {
    await fetchWrapper.post(`/api/host/room/create`, rooms.value);
  }
}

async function handleActivityOffer(townID, imageFiles) {
  const activityData = await fetchWrapper.post('/api/host/activity/create', {
    title: newOffer.value.title,
    description: newOffer.value.description,
    capacity: newOffer.value.capacity,
    town_id: townID,
    user_id: user.ID,
    date: new Date(newActivity.value.dateRange[0]).toJSON(),
    skill_level: newActivity.value.skillLevel.toLowerCase(),
    activity_type: newActivity.value.activityType.toLowerCase(),
    price: newActivity.value.price,
    duration: calculateDurationInHours(newActivity.value.dateRange[0], newActivity.value.dateRange[1]) + 'h',
    images: imageFiles
  }, "multipart/form-data");

  await fetchWrapper.post(`/api/host/activity/${activityData.offer.ID}/equipment/add`, {
    equipment: newActivity.value.equipment
  });
}

async function handleEventOffer(townID, imageFiles) {
  await fetchWrapper.post('/api/host/event/create', {
    title: newOffer.value.title,
    description: newOffer.value.description,
    capacity: newOffer.value.capacity,
    town_id: townID,
    user_id: user.ID,
    date_from: new Date(newEvent.value.dateFrom).toJSON(),
    date_to: new Date(newEvent.value.dateTo).toJSON(),
    price: newEvent.value.price,
    event_type: newEvent.value.eventType.toLowerCase(),
    images: imageFiles
  }, "multipart/form-data");
}

function resetForm() {
  newOffer.value = {
    offerType: '',
    title: '',
    description: '',
    capacity: '',
    country: '',
    city: '',
    images: [],
  };

  newAccommodation.value = {
    numberOfRooms: '',
    accommodationType: '',
    isAnimalFriendly: '',
    pricePerDay: '',
    generalFacilities: []
  };

  newActivity.value = {
    dateRange: [],
    skillLevel: '',
    price: '',
    activityType: '',
    equipment: [],
  };

  newEvent.value = {
    dateFrom: '',
    dateTo: '',
    price: '',
    eventType: ''
  };

  isAddingNewOffer.value = false;
  addedNewOffer.value = true;
  errors.apiError = null;
}

async function onAddRoom() {
  const room = {
    number: Number(newRoom.value.number),
    name: newRoom.value.name,
    description: newRoom.value.description,
    capacity: Number(newRoom.value.capacity),
    area: Number(newRoom.value.area),
    room_facilities: newRoom.value.roomFacilities.map(facility => ({name: facility})),
  };

  rooms.value.push(room);

  newRoom.value = {
    number: '',
    name: '',
    description: '',
    capacity: '',
    area: '',
    roomFacilities: [],
  };
}

function calculateDurationInHours(startDate, endDate) {
  const start = new Date(startDate);
  const end = new Date(endDate);
  const diff = end - start;
  const hours = diff / (1000 * 60 * 60);
  return hours;
}

function checkIfOfferInfoIsFilled() {
  errors.offerInfo = '';
  errors.dateRange = '';

  if (newOffer.value.title === '' || newOffer.value.description === '' || newOffer.value.capacity === '') {
    errors.offerInfo = 'Please fill out all required offer information.';
    return false;
  }

  const now = new Date();

  if (isOfferTypeActivity.value) {
    if (newActivity.value.skillLevel === '' || newActivity.value.price === '' || newActivity.value.activityType === '' || newActivity.value.dateRange.length === 0) {
      errors.offerInfo = 'Please fill out all required activity information.';
      return false;
    }

    const [startDate, endDate] = newActivity.value.dateRange;
    if (!startDate || !endDate) {
      errors.dateRange = 'Date range is required.';
      return false;
    }

    if (new Date(startDate) <= now) {
      errors.dateRange = 'Start date must be after the current date.';
      return false;
    }

    if (new Date(startDate) >= new Date(endDate)) {
      errors.dateRange = 'Start date must be before end date.';
      return false;
    }
  }

  if (isOfferTypeEvent.value) {
    if (newEvent.value.dateFrom === '' || newEvent.value.dateTo === '' || newEvent.value.price === '' || newEvent.value.eventType === '') {
      errors.offerInfo = 'Please fill out all required event information.';
      return false;
    }

    if (new Date(newEvent.value.dateFrom) <= now) {
      errors.dateRange = 'Start date must be after the current date.';
      return false;
    }

    if (new Date(newEvent.value.dateFrom) >= new Date(newEvent.value.dateTo)) {
      errors.dateRange = 'Start date must be before end date.';
      return false;
    }
  }

  return true;
}

function checkIfRoomInfoIsFilled() {
  let accommodationValues = newAccommodation.value
  let roomValues = newRoom.value
  if (accommodationValues.accommodationType === 'Villa' || accommodationValues.accommodationType === 'Apartment') {
    return
  }
  if (
      roomValues.number === '' ||
      roomValues.name === '' ||
      roomValues.description === '' ||
      roomValues.capacity === '' ||
      roomValues.area === ''
  ) {
    return false;
  } else {
    errors.offerInfo = '';
    return true;
  }
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

onMounted(async () => {
  const response = await fetchWrapper.get('/api/country/get/all')
  countries.value = response.data.map(country => country.Name)
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
                        :isRequired="true"/>
      </div>
    </Transition>
    <Transition name="bounce">
      <div v-if="isOfferTypeFilled" class="new_offer_info">
        <TextInput v-model="newOffer.title" label-text="Title" :isRequired="true" :min="2" :max="100"/>
        <TextInput v-model="newOffer.description" label-text="Description" :isRequired="true" :min="30" :max="300"/>
        <NumberInput v-model="newOffer.capacity" label-text="Capacity" :isRequired="true" :min="1"/>
        <div v-if="isOfferTypeAccommodation" class="w-100">
          <NumberInput v-model="newAccommodation.numberOfRooms" label-text="Number of rooms" :isRequired="true"
                       :min="1"/>
          <SelectionInput v-model="newAccommodation.accommodationType" label-text="Accommodation type"
                          placeholder="Type" :items="accommodationTypes" :isRequired="true"/>
          <CheckButtonInput :model-value="newAccommodation.isAnimalFriendly"
                            @changed-value="newAccommodation.isAnimalFriendly = !newAccommodation.isAnimalFriendly"
                            label-text="Is animal friendly?" width="100%"/>
          <NumberInput v-model="newAccommodation.pricePerDay" label-text="Price per day in $" :isRequired="true" :min="1"/>
          <CheckboxesInput v-model="newAccommodation.generalFacilities" label-text="Select general facilities"
                           placeholder="Facilities" :items="generalFacilities"/>
          <div v-if="isAccommodationTypeWithSelectableRooms" class="w-100 room-info-container">
            <p class="new_offer_text">Add rooms to your accommodation</p>
            <v-text-field v-model="newRoom.number" label="Room number" clearable placeholder="Number" class="w-100"
                          type="number"/>
            <v-text-field v-model="newRoom.name" label="Room name" clearable placeholder="Name" class="w-100"/>
            <v-text-field v-model="newRoom.description" label="Room description" clearable placeholder="Description"
                          class="w-100"/>
            <v-text-field v-model="newRoom.capacity" label="Room capacity" clearable placeholder="Capacity"
                          class="w-100" type="number"/>
            <v-text-field v-model="newRoom.area" label="Room area in m2" clearable placeholder="Area"
                          class="w-100" type="number"/>
            <v-select v-model="newRoom.roomFacilities" label="Select room facilities"
                      class="w-100" clearable chips multiple :items="roomFacilities"/>
            <button :disabled="!isRoomInfoFilled" class="button_basic" :class="{ 'disabled': !isRoomInfoFilled }"
                    @click="onAddRoom">
              Add room
            </button>
          </div>
        </div>
        <div v-if="isOfferTypeActivity" class="w-100">
          <SelectionInput v-model="newActivity.skillLevel" label-text="Skill level" placeholder="Level"
                          :items="skillLevels" :isRequired="true"/>
          <SelectionInput v-model="newActivity.activityType" label-text="Activity type" placeholder="Type"
                          :items="activityTypes" :isRequired="true"/>
          <NumberInput v-model="newActivity.price" label-text="Price in $" :isRequired="true" :min="1"/>
          <VueDatePicker v-model="newActivity.dateRange" placeholder="Select date range" range :model-config="{ format: 'YYYY-MM-DD' }"
                         :isRequired="true"/>
          <br>
          <CheckboxesInput v-model="newActivity.equipment" label-text="Select equipment needed" placeholder="Equipment"
                           :items="equipmentList"/>
        </div>
        <div v-if="isOfferTypeEvent" class="w-100">
          <DateInput v-model="newEvent.dateFrom" label-text="Date from" :isRequired="true"/>
          <DateInput v-model="newEvent.dateTo" label-text="Date to" :isRequired="true"/>
          <SelectionInput v-model="newEvent.eventType" label-text="Event type" placeholder="Type" :items="eventTypes"
                          :isRequired="true"/>
          <NumberInput v-model="newEvent.price" label-text="Price in $" :isRequired="true" :min="1"/>
        </div>
      </div>
    </Transition>
    <Transition name="bounce">
      <div v-if="isOfferInfoFilled" class="new_offer_info">
        <SelectionInput v-model="newOffer.country" label-text="Country" :items="countries" :isRequired="true"/>
        <TextInput v-model="newOffer.city" label-text="City" :isRequired="true" :min="2" :max="50"/>
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
        <div class="errors" v-if="errors.dateRange">{{ errors.dateRange }}</div>
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

.disabled {
  background-color: #cccccc;
  color: #666666;
  cursor: not-allowed;
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

.room-info-container {
  border: 2px solid #333;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1);
}

</style>