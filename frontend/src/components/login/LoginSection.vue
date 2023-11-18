<script setup>

import LoginButtons from "@/components/login/LoginButtons.vue";
</script>

<template>

  <p class="p">Welcome back. Please log in to your account</p>

  <text_input_component class="text-input-instance" state="default" text="Your email" text1="Type your email"/>
  <text_input_component
      class="design-component-instance-node"
      state="password"
      text="Your password"
      text1="Type your password"
  />
  <LoginButtons/>

</template>

<style scoped>

</style>