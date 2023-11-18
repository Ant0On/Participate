<script setup>
import LoginButton from './LoginButton.vue'
</script>

<template>
<div>
  <LoginButton class="lgn-button" text="Sign up"/>
  <LoginButton class="lgn-button" text="Log in"/>
  <a href="" text="Forgot password?"/>
</div>
</template>

<style scoped>
div {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
a {
  color: #0048d9;
  font-family: "IBM Plex Sans", Helvetica;
  font-weight: 400;
  line-height: normal;
  align-self: flex-end;
}
</style>