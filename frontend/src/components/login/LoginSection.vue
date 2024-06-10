<script setup>
import {reactive, ref, defineEmits } from 'vue';
import * as Yup from 'yup';

import TextInput from "@/components/ui/TextInput.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import {useAuthStore} from "@/stores/auth.store";


const loginData = ref({
  Email: '',
  Password: ''
})
const errors = reactive({
  apiError: ""
})

async function onSubmit() {
  const authStore = useAuthStore();
  await schema.validate(loginData.value).then(() => {
    return authStore.login(loginData.value.Email, loginData.value.Password)
  }).catch(error =>{
    errors.apiError = "Incorrect sign up data! " + error;
  })
}

const schema = Yup.object().shape({
  Email: Yup.string().required('Email is required'),
  Password: Yup.string().when('Email', {
    is: (value) => value !== '' || typeof value !== 'undefined',
    then: schema => schema.required('Password is required')
  } )
});

function onSignUp(){
  emits('signUpClicked')
}

const emits = defineEmits(['signUpClicked'])
</script>

<template>
  <div class="d-flex flex-column justify-start h-100" style="min-width: 300px">
    <p>Welcome back. Please log in to your account</p>
    <v-alert v-if="errors.apiError" text tile="Login error" color="error">{{ errors.apiError }}</v-alert>
    <form class="login_form" @submit.prevent>
      <TextInput labelText="Your email" placeholder="Type your email" :isRequired="true" v-model="loginData.Email"/>
      <PasswordInput v-model="loginData.Password" :labelText="'Your password'" @keyup.enter="onSubmit"/>
      <div class="d-flex justify-space-evenly">
        <v-btn text="Sign up" @click="onSignUp"></v-btn>
        <v-btn text="Log in" @click="onSubmit"></v-btn>
      </div>
    </form>
  </div>
</template>

<style scoped>

.login_form {
  display: flex;
  flex-direction: column;
}

p {
  color: black;
  font-family: "Sarabun", Helvetica;
  font-size: 1.5rem;
  font-weight: 100;
  line-height: normal;
  margin-bottom: 20px;
  margin-top: 80px;
}
</style>