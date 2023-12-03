<script setup>
import {reactive, ref } from 'vue';
import * as Yup from 'yup';

import TextInput from "@/components/ui/TextInput.vue";
import LoginButtons from "@/components/login/LoginButtons.vue";
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
        errors.apiError = "Incorrect login data!"
      }
      else{
        errors.apiError = error.message
      }
  }
  )
}

const schema = Yup.object().shape({
  Login: Yup.string().required('Login is required'),
  Password: Yup.string().required('Password is required')
});



</script>

<template>
  <div class="login_section">
    <p>Welcome back. Please log in to your account</p>
    <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
    <form class="login_form" @submit.prevent>
      <TextInput labelText="Your email" placeholder="Type your email" :isRequired="true" v-model="loginData.Login"/>
      <PasswordInput v-model="loginData.Password" @keyup.enter="onSubmit"/>
      <LoginButtons @log-in-clicked="onSubmit"/>
    </form>
  </div>
</template>

<style scoped>
div.login_section {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  row-gap: 40px;
  margin-left: 50px;
  margin-top: 50px;
}
div.errors{
  font-family: "Sarabun", Helvetica;

}

.login_form {
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
  div.login_section {
    margin-left: 5%;
  }
}
</style>