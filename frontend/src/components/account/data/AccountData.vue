<script setup>
import {computed, reactive, ref} from 'vue';
import * as Yup from 'yup';
import {storeToRefs} from 'pinia';

import {useAuthStore} from "@/stores/auth.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const userStore = useAuthStore();
const {user: user} = storeToRefs(userStore)

const errors = reactive({
  apiError: ""
})

const userImageSource = computed(() => {
    return require(`@/../images/users/${user.value.ID}.jpeg`);
})

const defaultImageSource = computed(() => {
  return require(`@/../images/users/default_image.png`);
})

const isUserImage = computed(() => {
    try {
      require(`@/../images/users/${user.value.ID}.jpeg`);
      return true;
    } catch {
      return false;
  }
});

const isImageUploaded = ref(false)
const uploadedImage = ref({})

const customerData = ref({
  Name: user.value.FirstName,
  LastName: user.value.LastName,
  Email: user.value.Email
})

const hostData = ref({
  Description: user.value.Description ? user.value.Description : '',
  PhoneNumber: user.value.PhoneNumber ? user.value.PhoneNumber : '',
  BankAccount: user.value.BankAccount ? user.value.BankAccount : '',
})

const currentCustomerData = ref({
  Name: user.value.FirstName,
  LastName: user.value.LastName,
  Email: user.value.Email
})

const currentHostData = ref({
  Description: user.value.Description ? user.value.Description : '',
  PhoneNumber: user.value.PhoneNumber ? user.value.PhoneNumber : '',
  BankAccount: user.value.BankAccount ? user.value.BankAccount : '',
})

async function onSubmit() {
  await schemaCustomer.validate(customerData.value).then(async () => {
    if (user.value.Role === "host") {
      await schemaHost.validate(hostData.value).then(async () => {
        if (hostData.value.Description !== currentHostData.value.Description)
          await fetchWrapper.put(`/api/host/${user.value.ID}/change/description`, {
            value: hostData.value.Description
          })
        if (hostData.value.PhoneNumber !== currentHostData.value.PhoneNumber)
          await fetchWrapper.put(`/api/host/${user.value.ID}/change/phone_number`, {
            value: hostData.value.PhoneNumber
          })
        if (hostData.value.BankAccount !== currentHostData.value.BankAccount)
          await fetchWrapper.put(`/api/host/${user.value.ID}/change/bank_account`, {
            value: hostData.value.BankAccount
          })

        user.value.Description = hostData.value.Description
        user.value.PhoneNumber = hostData.value.PhoneNumber
        user.value.BankAccount = hostData.value.BankAccount
      }).catch((hostErrors) => {
        if (typeof hostErrors === 'string' && hostErrors.includes('SQLSTATE 23505')) {
          errors.apiError = "Incorrect data! Please check the following errors: Phone number or bank account is already in use!"
        }
        else {
          errors.apiError = formatErrors(hostErrors.errors)
        }
      });

    }
    if (customerData.value.Name !== currentCustomerData.value.Name)
      await fetchWrapper.put(`/api/customer/${user.value.ID}/change/first_name`, {
        value: customerData.value.Name.trim()
      })
    if (customerData.value.LastName !== currentCustomerData.value.LastName)
      await fetchWrapper.put(`/api/customer/${user.value.ID}/change/last_name`, {
        value: customerData.value.LastName.trim()
      })
    if (customerData.value.Email !== currentCustomerData.value.Email)
      await fetchWrapper.put(`/api/customer/${user.value.ID}/change/email`, {
        value: customerData.value.Email.trim()
      })
    user.value.FirstName = customerData.value.Name
    user.value.LastName = customerData.value.LastName
    user.value.Email = customerData.value.Email

    currentHostData.value = hostData.value
    currentCustomerData.value = customerData.value
    errors.apiError = null
    isSubmitMode.value = false;
    await userStore.saveToLocalStorage()
    if (isImageUploaded.value) {
      const imageFile = dataURLtoFile(uploadedImage.value, 'image.jpeg');
      uploadedImage.value = null
      isImageUploaded.value = false
        await fetchWrapper.put(`/api/customer/${user.value.ID}/change/picture`, {
          image: imageFile
        }, "multipart/form-data")
    }
    window.location.reload();
  }).catch((customerErrors) => {
    if (typeof customerErrors === 'string' && customerErrors.includes('SQLSTATE 23505')) {
      errors.apiError = "Incorrect data! Please check the following errors: Email is already in use!"
    }
    else {
      errors.apiError = formatErrors(customerErrors.errors)
    }
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
  Name: Yup.string().min(2).required('Name is required'),
  LastName: Yup.string().min(2).required('Last name is required'),
  Email: Yup.string().email().required('Email is required')
});

const schemaHost = Yup.object().shape({
  Description: Yup.string().required('Description is required'),
  phoneNumber: Yup.string().min(9, 'Phone number must be at least 9 characters').max(15, 'Phone number must be at most 15 characters').matches(/^\d+$/, 'Phone number must contain only digits'),
  bankAccount: Yup.string().min(16, 'Bank account must be at least 16 characters').max(40, 'Bank account must be at most 40 characters').matches(/^\d+$/, 'Bank account must contain only digits'),
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
      <v-alert v-if="errors.apiError" text tile="Edit error" color="error">{{ errors.apiError }}</v-alert>
      <div class="user_fields">
        <div class="customer_fields">
          <v-text-field label="Name"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Name is required!',
                            (value) => value.length >= 2 && value.length <= 30 || 'Name must be between 2 and 30 characters!'
                        ]"
                        v-model="customerData.Name"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
          <v-text-field label="Last Name"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Last name is required!',
                            (value) => value.length >= 2 && value.length <= 100 || 'Name must be between 2 and 100 characters!'
                        ]"
                        v-model="customerData.LastName"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
          <v-text-field label="Email"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Email is required!',
                            (value) => /.+@.+/.test(value) || 'Invalid Email address'
                        ]"
                        v-model="customerData.Email"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
        </div>
        <div class="host_fields" v-if="user.Role === 'host'">
          <v-text-field label="Description"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Description is required!',
                        ]"
                        v-model="hostData.Description"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
          <v-text-field label="Phone number"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Phone number is required!',
                            (value) => value.length >= 9 && value.length <= 12 || 'Phone number has to be between 9 to 12 digits!',
                            (value) =>  /^\d+$/.test(value) || 'Phone number can contain only numbers!'
                        ]"
                        v-model="hostData.PhoneNumber"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
          <v-text-field label="Bank account"
                        :disabled="!isSubmitMode"
                        :rules="[
                            (value) => !!value || 'Bank account is required!',
                            (value) => value.length >= 23 && value.length <= 31 || 'Bank account has to be between 23 to 31 digits!',
                            (value) =>  /^\d+$/.test(value) || 'Bank account can contain only numbers!'
                        ]"
                        v-model="hostData.BankAccount"
                        clearable
                        class="w-100"
                        counter
          ></v-text-field>
        </div>
      </div>
      <div id="upload_image" class="customer_fields">
        <div class="image_container">
          <p class="image_input">User image</p>
          <img v-if="isImageUploaded" :src="uploadedImage" class="preview_image" alt="offerImage"/>
          <img v-else-if="isUserImage" :src="userImageSource" class="preview_image" alt="offerImage"/>
          <img v-else :src="defaultImageSource" class="preview_image" alt="offerImage"/>
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

</style>