<script setup>
import TextInput from "@/components/ui/TextInput.vue";
import LoginButtons from "@/components/login/LoginButtons.vue";
import PasswordInput from "@/components/ui/PasswordInput.vue";
</script>

<template>
  <div class="login_section">
    <p>Welcome back. Please log in to your account</p>
    <TextInput labelText="Your email" placeholder="Type your email" isRequired="True"/>
    <PasswordInput/>
    <LoginButtons/>
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
@media(max-width: 900px){
  div.login_section{
    margin-left: 5%;
  }
}
</style>