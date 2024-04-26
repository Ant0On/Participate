<script setup>
import {reactive, ref} from 'vue';
import * as Yup from 'yup';

import {useAuthStore} from "@/stores/auth.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";
import PasswordInput from "@/components/ui/PasswordInput.vue";

const auth = useAuthStore();
const user = auth.user;

const errors = reactive({
  apiError: "",
  oldPassword: "",
  newPassword: "",
  confirmNewPassword: ""
})

const passwordData = ref({
  OldPassword: '',
  NewPassword: '',
  ConfirmNewPassword: '',
})

async function onSubmit() {
  try {
    await schemaPassword.validate(passwordData.value)
    isSubmitMode.value = false;

    if (passwordData.value.OldPassword) {
      await fetchWrapper.put(`/api/customer/${user.ID}/change/password`, {
        old_password: passwordData.value.OldPassword.trim(),
        new_password: passwordData.value.NewPassword.trim(),
        confirm_password: passwordData.value.ConfirmNewPassword.trim(),
      });
    }
    passwordData.value.OldPassword = '';
    passwordData.value.NewPassword = '';
    passwordData.value.ConfirmNewPassword = '';

    errors.apiError = 'Password changed successfully!';

    user.Password = passwordData.value.NewPassword;
  } catch (error) {
    errors.apiError = formatErrors(error.errors);
  }
}

function formatErrors(validationErrors) {
  let errorMessage = "Incorrect data! Please check the following errors:";
  for (const field in validationErrors) {
    errorMessage += `\n ${validationErrors[field]}`;
  }
  return errorMessage;
}

function onCancel() {
  passwordData.value.OldPassword = '';
  passwordData.value.NewPassword = '';
  passwordData.value.ConfirmNewPassword = '';
  isSubmitMode.value = false
  errors.apiError = null
}

function onChangeData() {
  isSubmitMode.value = true;
}

const isSubmitMode = ref(false)

const schemaPassword = Yup.object().shape({
  OldPassword: Yup.string().required('Old password is required'),
  NewPassword: Yup.string()
      .required('Enter new password')
      .min(8, 'Password must be at least 8 characters')
      .notOneOf([Yup.ref('OldPassword'), null], 'New Password must be different from Old Password')
      .matches(/[a-zA-Z]/, 'Password must contain at least one letter')
      .matches(/\d/, 'Password must contain at least one digit')
      .matches(/[!@#$%^&*(),.?":{}|<>]/, 'Password must contain at least one special character'),
  ConfirmNewPassword: Yup.string().required('Confirm new password').oneOf([Yup.ref('NewPassword'), null], 'Passwords must match'),
});

</script>

<template>
  <div class="account_data_component">
    <div class="user_data">
      <div class="errors" v-if="errors.apiError">{{ errors.apiError }}</div>
      <div class="user_fields">
        <div class="customer_fields">
          <PasswordInput :label-text="'Old Password'" :isActive="isSubmitMode" v-model="passwordData.OldPassword"
                         type="password" width="100%"/>
          <PasswordInput :label-text="'New Password'" :isActive="isSubmitMode" v-model="passwordData.NewPassword"
                         type="password" width="100%"/>
          <PasswordInput :label-text="'Confirm New Password'" :isActive="isSubmitMode"
                         v-model="passwordData.ConfirmNewPassword" type="password" width="100%"/>
          <p class="error-message" v-if="errors.oldPassword">{{ errors.oldPassword }}</p>
          <p class="error-message" v-if="errors.newPassword">{{ errors.newPassword }}</p>
          <p class="error-message" v-if="errors.confirmNewPassword">{{ errors.confirmNewPassword }}</p>
        </div>
      </div>
      <div class="buttons">
        <button v-if="!isSubmitMode" class="button_basic" @click="onChangeData">
          Change data
        </button>
        <button v-if="isSubmitMode" class="button_basic button_cancel" @click="onCancel">
          Cancel
        </button>
        <button v-if="isSubmitMode" class="button_basic" @click="onSubmit">
          Submit
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
label {
  display: flex;
  border: 1px solid black;
  padding: 10%;
  border-radius: 5px;
  height: 50px;
  width: 100px;
  text-align: center;
}

div.account_data_component {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-grow: 1;
}

div.user_data {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: space-between;
}

div.user_fields {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

div.customer_fields {
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}

div.logout_button_container {
  margin-right: 10%;
  margin-left: 5%;
}

div.line {
  display: flex;
  flex-direction: row;
  color: black;
  border-right: 1px solid;
}

div.buttons {
  align-self: flex-end;
  display: flex;
  margin: 2%;
  column-gap: 5%;
}

.button_basic {
  all: unset;
  font-family: "IBM Plex Sans", Helvetica, serif !important;
  background-color: #efefef;
  color: black;
  border-radius: 6px;
  box-sizing: border-box;
  padding: 4px 16px;
  border: 1px solid #808080;
  border-radius: 6px;

  width: 150px;
  height: 35px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.button_basic:active {
  background-color: rgba(22, 89, 224, 0.5);
  color: var(--systemwhite)
}

.button_cancel {
  background-color: rgba(255, 0, 0, 60%);
}

div.errors {
  font-family: "Sarabun", Helvetica;
  border: 1px solid #ff0000;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 10px;
}

p.error-label {
  font-weight: bold;
  color: #ff0000;
}

p.error-message {
  color: #ff0000;
}
</style>