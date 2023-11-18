<script setup>
import LoginButton from './LoginButton.vue'
</script>

<template>
<div>
  <LoginButton class="lgn-button" text="Sign up"/>
  <LoginButton class="lgn-button" text="Log in"/>
  <div class="text-wrapper-3">Forgot password?</div>
</div>
</template>

<style scoped>
div {
  display: flex;
  justify-content: space-between;
}
.lgn-button {
  height: 100px;
  width: 200px;
}
</style>