<script setup>
import {defineProps, defineEmits} from 'vue'

import LoginButton from './LoginButton.vue'

const emits = defineEmits(['logInClicked', 'signUpClicked'])

function signUpHandler() {
  emits('signUpClicked')
}

function logInHandler() {
  emits('logInClicked')
}

</script>

<template>
  <div class="login_button_container">
    <div class="button_container">
      <LoginButton class="lgn-button" text="Sign up" @button-clicked="signUpHandler"/>
      <LoginButton class="lgn-button" text="Log in" @button-clicked="logInHandler"/>
    </div>
    <a href="" text="Forgot password?"/>
  </div>
</template>

<style scoped>
div.login_button_container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-start;
  margin-top: 20px;

}

div.button_container {
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  column-gap: 20px;
}

a {
  color: #0048d9;
  font-family: "IBM Plex Sans", Helvetica;
  font-weight: 400;
  line-height: normal;
  align-self: flex-end;
  padding-right: 10px;
  margin-top: 10px;
}
</style>
