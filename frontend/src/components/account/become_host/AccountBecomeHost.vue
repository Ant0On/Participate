<script setup>
import {ref, reactive} from 'vue';
import * as Yup from 'yup';

import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import {useAuthStore} from "@/stores/auth.store";
import TextInput from "@/components/ui/TextInput.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";

const isJoining = ref(false);
const auth = useAuthStore();
const user = auth.user;

const hostData = ref({
  description: '',
  phoneNumber: '',
  bankAccount: '',
  password: '',
})

const errors = reactive({
  apiError: ""
})

const schemaHost = Yup.object().shape({
  description: Yup.string().required('Description is required'),
  phoneNumber: Yup.string().min(9, 'Phone number must be at least 9 characters').max(12, 'Phone number must be at most 12 characters').matches(/^\d+$/, 'Phone number must contain only digits'),
  bankAccount: Yup.string().min(23, 'Bank account must be at least 23 characters').max(31, 'Bank account must be at most 31 characters').matches(/^\d+$/, 'Bank account must contain only digits'),
  password: Yup.string().required('Password is required')
});

async function onSubmit(){
  await schemaHost.validate(hostData.value).then(async () => {
    await fetchWrapper.post(`/api/customer/${user.ID}/promote`, {
      description: hostData.value.description,
      phone_number: hostData.value.phoneNumber,
      bank_account: hostData.value.bankAccount,
      password: hostData.value.password,
    })
    alert("Soon you will be logged out. Please log in again!")
    await auth.logout()
  }).catch(error =>{
    console.log(error)
    if (typeof error === 'string') {
      if (error.includes('SQLSTATE 23505')) {
        errors.apiError = "Bank account or phone number is already in use!";
      }
      else {
        errors.apiError = "Incorrect data! Please check the following errors: " + error;
      }
    } else {
      errors.apiError = "Incorrect data! Please check the following errors: " + error
    }
  });
}
</script>

<template>
<div class="become_host">
  <div v-if="!isJoining" class="join_info">
    <p class="join_info"> Do you want to become a host and participate in unique and thrilling journey?</p>
    <button class="button_basic" @click="isJoining = !isJoining">
      Join now!
    </button>
  </div>
  <div v-else class="host_info">
    <p class="host_info_description"> Please fill all your data to become a host!</p>
    <v-alert v-if="errors.apiError" text tile="Form error" color="error">{{ errors.apiError }}</v-alert>
    <div class="host_info_fields">
      <v-textarea v-model="hostData.description" label="Description"
                  :rules="[
                        value=> !!value,
                        value => value.length > 30 || 'Description too short',
                        value => value.length < 300 || 'Description too long'
         ]"
                  clearable
                  class="w-100"
                  counter
      />
      <TextInput label-text="Phone number" v-model="hostData.phoneNumber" :min="9" :max="12" is-required/>
      <TextInput label-text="Bank account" v-model="hostData.bankAccount" :min="23" :max="31" is-required/>
      <PasswordInput :labelText="'Enter your password'" v-model="hostData.password"/>
    </div>
    <button class="button_basic" @click="onSubmit">
      Join now!
    </button>
  </div>
</div>
</template>

<style scoped>
div.become_host{
  flex-grow: 1;
}
div.host_info{
  display: flex;
  flex-direction: column;
  row-gap: 50px;
  align-items: center;
  flex-grow: 1;
}
div.host_info_fields{
  display: flex;
  flex-direction: column;
  row-gap: 30px;
  flex-grow: 1;
  width: 40%;
}
div.join_info{
  display: flex;
  flex-direction: column;
  row-gap: 60px;
  align-items: center;
  padding-top: 5%;
  padding-left: 2%;
  text-align: center;
  flex-grow: 1;
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
p.join_info{
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.6rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
}
p.host_info_description{
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.4rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
}

</style>