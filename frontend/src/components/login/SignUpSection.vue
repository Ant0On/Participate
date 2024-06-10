<script setup>
import {reactive, ref} from 'vue';
import * as Yup from 'yup';

import TextInput from "@/components/ui/TextInput.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import {useAuthStore} from "@/stores/auth.store";


const signUpData = ref({
  FirstName: '',
  LastName: '',
  Login: '',
  Password: '',
  ConfirmPassword: '',
  Image: null
})

const errors = reactive({
  apiError: ""
})

async function onSubmit() {
  const authStore = useAuthStore();
  await schema.validate(signUpData.value).then(() => {
    return authStore.signUp(signUpData.value.FirstName, signUpData.value.LastName, signUpData.value.Login,
        signUpData.value.Password, signUpData.value.ConfirmPassword, signUpData.value.Image)
  }).catch(error => {
    if (typeof error === 'string') {
      if (error.includes('SQLSTATE 23505')) {
        errors.apiError = "Email is already taken!";
      }
      else {
        errors.apiError = "Incorrect sign up data! " + error;
      }
    } else {
      errors.apiError = "Incorrect sign up data! " + error;
    }
  })
}

const schema = Yup.object().shape({
  FirstName: Yup.string()
      .min(2, "First name must be at least 2 characters")
      .max(50, "First name must be maximum 50 characters")
      .required('First name is required'),
  LastName: Yup.string()
      .min(2, "Last name must be at least 2 characters")
      .max(50, "Last name must be maximum 50 characters")
      .required('Last name is required'),
  Login: Yup.string().email("This is not a valid email address").required('Email is required'),
  Password: Yup.string()
      .required('Password is required')
      .min(8, 'Password must be at least 8 characters')
      .matches(/[a-zA-Z]/, 'Password must contain at least one letter')
      .matches(/\d/, 'Password must contain at least one digit')
      .matches(/[!@#$%^&*(),.?":{}|<>]/, 'Password must contain at least one special character'),
  ConfirmPassword: Yup.string().oneOf([Yup.ref('Password'), null], 'Passwords must match')
});


</script>

<template>
  <div class="signup_section">
    <v-alert v-if="errors.apiError" text tile="Email error" color="error">{{ errors.apiError }}</v-alert>
    <p>Create new account</p>
    <form class="signup_form" @submit.prevent>
      <TextInput labelText="Your first name" placeholder="Type your first name" :isRequired="true"
                 v-model="signUpData.FirstName"/>
      <TextInput labelText="Your last name" placeholder="Type your last name" :isRequired="true"
                 v-model="signUpData.LastName"/>
      <TextInput  labelText="Your email" placeholder="Type your email" :isRequired="true"
                 v-model="signUpData.Login"/>
      <PasswordInput v-model="signUpData.Password" :labelText="'Your password'" @keyup.enter="onSubmit" :isRequired="true"/>
      <PasswordInput v-model="signUpData.ConfirmPassword" :labelText="'Confirm your password'"
                     :placeholder="'Type your password again'" @keyup.enter="onSubmit" :isRequired="true"/>
      <v-file-input v-model="signUpData.Image" label="Upload profile image" filled prepend-icon="mdi-camera"/>
      <v-btn text="Sign up" @click="onSubmit" @keyup.enter="onSubmit"></v-btn>
    </form>
  </div>
</template>

<style scoped>
div.signup_section {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  flex-grow: 1;
  padding: 5%;
}

div.errors {
  font-family: "Sarabun", Helvetica;
}

.signup_form {
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

@media (max-width: 900px) {
  div.signup_section {
    margin-left: 5%;
  }
}
</style>