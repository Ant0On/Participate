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
}
p {
  color: black;
  font-family: "Sarabun", Helvetica;
  font-size: 1.5rem;
  font-weight: 100;
  line-height: normal;
}
</style>