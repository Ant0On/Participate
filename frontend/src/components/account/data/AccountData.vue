<script setup>
import {ref, reactive} from 'vue';
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

const customerData = ref({
  name: user.first_name,
  lastName: user.last_name,
  email: user.email,
})

const hostData = ref({
  description: user.description ? user.description : '',
  phoneNumber: user.phoneNumber ? user.phoneNumber : '',
  bankAccount: user.bankAccount ? user.bankAccount : '',
})

async function onSubmit() {
  await schemaCustomer.validate(customerData.value).then(async () => {
    if(user.Role === "host"){
      await schemaHost.validate(hostData.value).then(async () => {
        await fetchWrapper.post(`/api/host/${user.ID}/change/description`, {
          description: hostData.value.description
        })
        await fetchWrapper.post(`/api/host/${user.ID}/change/phone_number`, {
          phone_number: hostData.value.phoneNumber
        })
        await fetchWrapper.post(`/api/host/${user.ID}/change/bank_account`, {
          bank_account: hostData.value.bankAccount
        })

        user.description = hostData.value.description
        user.phone_number = hostData.value.phone_number
        user.bank_account = hostData.value.bank_account
      })

    }
    isSubmitMode.value = false;
    await fetchWrapper.post(`/api/customer/${user.ID}/change/first_name`, {
      first_name: customerData.value.name.trim()
    })
    await fetchWrapper.post(`/api/customer/${user.ID}/change/last_name`, {
      last_name: customerData.value.lastName.trim()
    })
    await fetchWrapper.post(`/api/customer/${user.ID}/change/email`,{
      email: customerData.value.email.trim()
    })
    user.first_name = hostData.value.first_name
    user.last_name = hostData.value.last_name
    user.email = hostData.value.email
  }).catch(error =>{
        errors.apiError = "Incorrect data!"
      }
  )
}
function onCancel() {
  isSubmitMode.value = false
}

function onChangeData() {
  isSubmitMode.value = true;
}

const isSubmitMode = ref(false)

const schemaCustomer = Yup.object().shape({
  name: Yup.string().min(5).required('Name is required'),
  lastName: Yup.string().min(5).required('Last name is required'),
  email: Yup.string().email().required('Email is required'),
});

const schemaHost = Yup.object().shape({
  description: Yup.string().required('Description is required'),
  phoneNumber: Yup.number().min(9),
  bankAccount: Yup.number().min(16),
});

</script>

<template>
  <div class="account_data_component">
    <div class="user_data">
      <div class="user_fields">
        <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
        <div class="customer_fields">
          <TextInput label-text="Name" :is-active="isSubmitMode" v-model="customerData.name" width="100%"/>
          <TextInput label-text="Last Name" :is-active="isSubmitMode" v-model="customerData.lastName" width="100%"/>
          <TextInput label-text="Email" :is-active="isSubmitMode" v-model="customerData.email" width="100%"/>
        </div>
        <div class="host_fields" v-if="userRole === 'host'">
          <TextInput label-text="Description" :is-active="isSubmitMode" v-model="hostData.description" width="100%"/>
          <TextInput label-text="Phone number" :is-active="isSubmitMode" v-model="hostData.phoneNumber" width="100%"/>
          <TextInput label-text="Bank account" :is-active="isSubmitMode" v-model="hostData.bankAccount" width="100%"/>
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
div.account_data_component{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-grow: 1;
}
div.user_data{
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: space-between;
}
div.user_fields{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
div.customer_fields{
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}
div.host_fields{
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}
div.logout_button_container{
  margin-right: 10%;
  margin-left: 5%;
}
div.line{
  display: flex;
  flex-direction: row;
  color: black;
  border-right: 1px solid;
}
div.buttons{
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
div.errors{
  font-family: "Sarabun", Helvetica;
}
</style>