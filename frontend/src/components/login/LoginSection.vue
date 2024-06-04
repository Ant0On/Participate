<script setup>
import {reactive, ref, defineEmits } from 'vue';
import * as Yup from 'yup';

import TextInput from "@/components/ui/TextInput.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import {useAuthStore} from "@/stores/auth.store";


const loginData = ref({
  Login: '',
  Password: ''
})
const errors = reactive({
  apiError: ""
})

async function onSubmit() {
  const authStore = useAuthStore();
  await schema.validate(loginData.value).then(() => {
    return authStore.login(loginData.value.Login, loginData.value.Password)
  }).catch(error =>{
        if(error === "Bad Request")
        {
          errors.apiError = "Incorrect login data! " + error
        }
        else{
          errors.apiError = error.message
        }
      }
  )
}

const schema = Yup.object().shape({
  Login: Yup.string().required('Login is required'),
  Password: Yup.string().when('Login', {
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
    <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    <form class="login_form" @submit.prevent>
      <TextInput labelText="Your email" placeholder="Type your email" :isRequired="true" v-model="loginData.Login"/>
      <PasswordInput v-model="loginData.Password" :labelText="'Your password'" @keyup.enter="onSubmit"/>
      <div class="d-flex justify-space-evenly">
        <v-btn text="Sign up" @click="onSignUp"></v-btn>
        <v-btn text="Log in" @click="onSubmit"></v-btn>
      </div>
    </form>
  </div>
</template>

<style scoped>
div.errors{
  font-family: "Sarabun", Helvetica;

}

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