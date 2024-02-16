<script setup>
import {reactive, ref} from 'vue';
import * as Yup from 'yup';

import TextInput from "@/components/ui/TextInput.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import {useAuthStore} from "@/stores/auth.store";
import LoginButton from "@/components/login/LoginButton.vue";


const signUpData = ref({
  Name: '',
  Login: '',
  Password: ''
})

const errors = reactive({
  apiError: ""
})

async function onSubmit() {
  const authStore = useAuthStore();
  await schema.validate(signUpData.value).then(() => {
    return authStore.signUp(signUpData.value.Name, signUpData.value.Login, signUpData.value.Password)
  }).catch(error => {
          errors.apiError = "Incorrect sign up data! " + error.message
      }
  )
}

const schema = Yup.object().shape({
  Name: Yup.string().required('Name is required'),
  Login: Yup.string().email().required('Email is required'),
  Password: Yup.string().when('Login', {
    is: (value) => value !== '' || typeof value !== 'undefined',
    then: schema => schema.required('Password is required')
  } )
});


</script>

<template>
  <div class="signup_section">
    <p>Create new account</p>
    <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    <form class="signup_form" @submit.prevent>
      <TextInput width="440px" labelText="Your name" placeholder="Type your name" :isRequired="true"
                 v-model="signUpData.Name"/>
      <TextInput width="440px" labelText="Your email" placeholder="Type your email" :isRequired="true"
                 v-model="signUpData.Login"/>
      <PasswordInput v-model="signUpData.Password" @keyup.enter="onSubmit"/>
      <LoginButton text="Sign up" @button-clicked="onSubmit" @keyup.enter="onSubmit"/>
    </form>
  </div>
</template>

<style scoped>
div.signup_section {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  row-gap: 40px;
  margin-top: 50px;
  margin-right: 8%;
}

div.errors {
  font-family: "Sarabun", Helvetica;
}

.signup_form {
  display: flex;
  flex-direction: column;
  row-gap: 40px;
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

@media (max-width: 900px) {
  div.signup_section {
    margin-left: 5%;
  }
}
</style>