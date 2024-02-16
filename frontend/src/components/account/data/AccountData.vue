<script setup>
import {computed, reactive, ref} from 'vue';
import * as Yup from 'yup';

import {useAuthStore} from "@/stores/auth.store";
import TextInput from "@/components/ui/TextInput.vue";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const auth = useAuthStore();
const user = auth.user;

const userRole = user.Role

const errors = reactive({
  apiError: ""
})

const isUserImage = computed(() => {
  try {
    require(`@/../images/customers/${user.ID}.jpeg`)
    return true
  } catch {
    return false
  }
})

const isImageUploaded = ref(false)
const uploadedImage = ref({})


const customerData = ref({
  Name: user.FirstName,
  LastName: user.LastName,
  Email: user.Email,
})

const hostData = ref({
  Description: user.Description ? user.Description : '',
  PhoneNumber: user.PhoneNumber ? user.PhoneNumber : '',
  BankAccount: user.BankAccount ? user.BankAccount : '',
})

const currentCustomerData = ref({
  Name: user.FirstName,
  LastName: user.LastName,
  Email: user.Email,
})
const currentHostData = ref({
  Description: user.Description ? user.Description : '',
  PhoneNumber: user.PhoneNumber ? user.PhoneNumber : '',
  BankAccount: user.BankAccount ? user.BankAccount : '',
})

async function onSubmit() {
  await schemaCustomer.validate(customerData.value).then(async () => {
    if (user.Role === "host") {
      await schemaHost.validate(hostData.value).then(async () => {
        if (hostData.value.Description !== currentHostData.value.Description)
          await fetchWrapper.put(`/api/host/${user.ID}/change/description`, {
            description: hostData.value.Description
          })
        if (hostData.value.PhoneNumber !== currentHostData.value.PhoneNumber)
          await fetchWrapper.put(`/api/host/${user.ID}/change/phone_number`, {
            phone_number: hostData.value.PhoneNumber
          })
        if (hostData.value.BankAccount !== currentHostData.value.BankAccount)
          await fetchWrapper.put(`/api/host/${user.ID}/change/bank_account`, {
            bank_account: hostData.value.BankAccount
          })

        user.Description = hostData.value.Description
        user.PhoneNumber = hostData.value.PhoneNumber
        user.BankAccount = hostData.value.BankAccount
      }).catch((hostErrors) => {
        errors.apiError = formatErrors(hostErrors.errors);
      });

    }
    isSubmitMode.value = false;
    if (customerData.value.Name !== currentCustomerData.value.Name)
      await fetchWrapper.put(`/api/customer/${user.ID}/change/first_name`, {
        first_name: customerData.value.Name.trim()
      })
    if (customerData.value.LastName !== currentCustomerData.value.LastName)
      await fetchWrapper.put(`/api/customer/${user.ID}/change/last_name`, {
        last_name: customerData.value.LastName.trim()
      })
    if (customerData.value.Email !== currentCustomerData.value.Email)
      await fetchWrapper.put(`/api/customer/${user.ID}/change/email`, {
        email: customerData.value.Email.trim()
      })
    user.FirstName = hostData.value.FirstName
    user.LastName = hostData.value.LastName
    user.Email = hostData.value.Email

    currentHostData.value = hostData.value
    currentCustomerData.value = customerData.value
    errors.apiError =null

    if (isImageUploaded.value) {
      const imageFile = dataURLtoFile(uploadedImage.value, 'image.jpeg');
      uploadedImage.value = null
      isImageUploaded.value = false
      if (user.Role === "customer") {
        await fetchWrapper.put(`/api/customer/${user.ID}/change/picture`, {
          image: imageFile
        }, "multipart/form-data")
      } else if (user.Role === "host") {
        await fetchWrapper.put(`/api/host/${user.ID}/change/picture`, {
          image: imageFile
        }, "multipart/form-data")
      }
    }
  }).catch((customerErrors) => {
    errors.apiError = formatErrors(customerErrors.errors);
  })
}

function formatErrors(validationErrors) {
  let errorMessage = "Incorrect data! Please check the following errors:";
  for (const field in validationErrors) {
    errorMessage += `\n ${validationErrors[field]}`;
  }
  return errorMessage;
}

function onCancel() {
  isSubmitMode.value = false
  customerData.value = currentCustomerData.value
  hostData.value = currentHostData.value
  errors.apiError = null
  uploadedImage.value = null
  isImageUploaded.value = false
}

function onChangeData() {
  isSubmitMode.value = true;
}

const isSubmitMode = ref(false)

const schemaCustomer = Yup.object().shape({
  Name: Yup.string().min(5).required('Name is required'),
  LastName: Yup.string().min(5).required('Last name is required'),
  Email: Yup.string().email().required('Email is required'),
});

const schemaHost = Yup.object().shape({
  Description: Yup.string().required('Description is required'),
  PhoneNumber: Yup.number().min(9),
  BankAccount: Yup.number().min(16),
});

async function uploadImage(imageInput) {
  const image = imageInput.target.files[0];
  const reader = new FileReader();
  reader.readAsDataURL(image);
  reader.onload = source => {
    uploadedImage.value = source.target.result;
    isImageUploaded.value = true;
  };
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
</script>

<template>
  <div class="account_data_component">
    <div class="user_data">
      <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
      <div class="user_fields">
        <div class="customer_fields">
          <TextInput label-text="Name" :is-active="isSubmitMode" v-model="customerData.Name" width="100%"/>
          <TextInput label-text="Last Name" :is-active="isSubmitMode" v-model="customerData.LastName" width="100%"/>
          <TextInput label-text="Email" :is-active="isSubmitMode" v-model="customerData.Email" width="100%"/>
        </div>
        <div class="host_fields" v-if="userRole === 'host'">
          <TextInput label-text="Description" :is-active="isSubmitMode" v-model="hostData.Description" width="100%"/>
          <TextInput label-text="Phone number" :is-active="isSubmitMode" v-model="hostData.PhoneNumber" width="100%"/>
          <TextInput label-text="Bank account" :is-active="isSubmitMode" v-model="hostData.BankAccount" width="100%"/>
        </div>
      </div>
      <div id="upload_image" class="customer_fields">
        <div class="image_container">
          <p class="image_input">User image</p>
          <img v-if="isImageUploaded" :src="uploadedImage" class="preview_image" alt="offerImage"/>
          <img v-else-if="isUserImage" :src="require(`@/../images/customers/${user.ID}.jpeg`)" class="preview_image"
               alt="offerImage"/>
          <img v-else :src="require(`@/../images/customers/default_image.png`)" class="preview_image" alt="offerImage"/>
        </div>
        <div id="image_input" v-if="isSubmitMode">
          <label for="image_upload">
            Add a photo
          </label>
          <input id="image_upload" type="file" accept="image/jpeg, image/png, image/jpg"
                 @change=uploadImage>
        </div>
      </div>
      <div class="buttons">
        <button v-if="!isSubmitMode" class="button_basic" @click="onChangeData">
          Change data
        </button>
        <button v-if="isSubmitMode" class="button_basic button_cancel" @click="onCancel">
          Cancel
        </button>
        <button v-if="isSubmitMode" class="button_basic" @click="onSubmit">
          Submit
        </button>
      </div>
    </div>
    <div class="line"></div>
    <div class="logout_button_container">
      <button class="button_basic logout_button" @click="auth.logout()">
        Log out
      </button>
    </div>
  </div>
</template>

<style scoped>
div#upload_image {
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  column-gap: 30px;
}

p.image_input {
  font-family: "IBMPlex Sans-Regular", Helvetica;
  font-style: normal;
  font-weight: 500;
  color: var(--text-secondary-grey2);
  line-height: 150%;
}

label {
  display: flex;
  border: 1px solid black;
  padding: 10%;
  border-radius: 5px;
  height: 50px;
  width: 100px;
  text-align: center;
}

input#image_upload {
  position: absolute;
  left: -99999rem
}

img.preview_image {
  padding-top: 5%;
  padding-bottom: 5%;
  border-radius: 50%;
  width: 100px;
  height: 100px;
  align-self: flex-start;

}

div.account_data_component {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-grow: 1;
}

div.user_data {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: space-between;
}

div.user_fields {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

div.customer_fields {
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}

div.host_fields {
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}

div.logout_button_container {
  margin-right: 10%;
  margin-left: 5%;
}

div.line {
  display: flex;
  flex-direction: row;
  color: black;
  border-right: 1px solid;
}

div.buttons {
  align-self: flex-end;
  display: flex;
  margin: 2%;
  column-gap: 5%;
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
  align-items: center;
}

.button_basic:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}

.button_cancel {
  background-color: rgba(255, 0, 0, 60%);
}

div.errors {
  font-family: "Sarabun", Helvetica;
}
</style>