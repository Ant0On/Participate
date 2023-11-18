<script setup>
import {computed, ref} from "vue";

const hidePassword = ref(true);
const password = ref("");

const passwordFieldIcon = computed(() => hidePassword.value ? "fa-eye" : "fa-eye-slash");
const passwordFieldType = computed(() => hidePassword.value ? "password" : "text");
const passwordPlaceholder = computed(() => hidePassword.value ? "**********" : "Type your password")
</script>

<template>
  <div class="password_input">
    <label for="password">Your password<span>*</span></label>
    <div class="password_input_field">
      <input :type="passwordFieldType" id="password" v-model="password" :placeholder="passwordPlaceholder">
      <i class="fas" :class="[passwordFieldIcon]" @click="hidePassword = !hidePassword"></i>
    </div>
  </div>
</template>

<style scoped>
.password_input {
  display: flex;
  justify-content: flex-start;
  flex-direction: column;
  row-gap: 10px;
}

input {
  width: 440px;
  height: 40px;
  background-color: var(--surfacelight);
  border: 1px;
}

span {
  font-family: "IBMPlex Sans-Regular", Helvetica;
  font-style: normal;
  color: var(--systemred);
}

label {
  font-family: "IBMPlex Sans-Regular", Helvetica;
  font-style: normal;
  font-weight: 500;
  color: var(--text-secondary-grey2);
  line-height: 150%;

}

.password_input_field {
  display: flex;
  flex-direction: row;
  column-gap: 10px;
}

</style>