<script setup>
import {ref} from 'vue';


import TextInput from "@/components/ui/TextInput.vue";
import LoginButtons from "@/components/login/LoginButtons.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";


const loginData = ref({
  login: '',
  password: ''
})

async function handleLogIn() {
  if (!validateLoginData(loginData.value)) {
    alert('Please fill all fields!')
    return
  }
  const response = await fetch('/api/login', createRequestOptions(loginData.value))
  if (response.status === 200) {
    response.json().then(data => console.log(data))
  }

}

function validateLoginData(loginData) {
  if (typeof loginData.login !== 'undefined' && typeof loginData.password !== 'undefined')
    if (loginData.login !== '' && loginData.password !== '')
      return true
  return false
}

function createRequestOptions(loginData) {
  return {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({email: loginData.login.trim(), password: loginData.password.trim(), table: 'customer'})
  }
}

</script>

<template>
  <div class="login_section">
    <p>Welcome back. Please log in to your account</p>
    <TextInput labelText="Your email" placeholder="Type your email" :isRequired="true" v-model="loginData.login"/>
    <PasswordInput v-model="loginData.password"/>
    <LoginButtons @log-in-clicked="handleLogIn"/>
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