<script setup>
import {ref} from 'vue';

import {useAuthStore} from "@/stores/auth.store";
import TextInput from "@/components/ui/TextInput.vue";

const auth = useAuthStore();
const user = auth.user;

const userRole = user.Role

function onSubmit() {

}

function onCancel() {
  isSubmitMode.value = false
}

function onChangeData() {
  isSubmitMode.value = true;
}

const isSubmitMode = ref(false)


</script>

<template>
  <div class="account_data_component">
    <div class="user_data">
      <div class="user_fields">
        <div class="customer_fields">
          <TextInput label-text="Name" :is-active="isSubmitMode" :model-value="user.first_name" width="100%"/>
          <TextInput label-text="Last Name" :is-active="isSubmitMode" :model-value="user.last_name" width="100%"/>
          <TextInput label-text="Email" :is-active="isSubmitMode" :model-value="user.email" width="100%"/>
        </div>
        <div class="host_fields" >
<!--        <div class="host_fields" v-if="userRole === 'host'">-->
          <TextInput label-text="Description" :is-active="isSubmitMode" :model-value="user.email" width="100%"/>
          <TextInput label-text="Phone number" :is-active="isSubmitMode" :model-value="user.email" width="100%"/>
          <TextInput label-text="Bank account" :is-active="isSubmitMode" :model-value="user.email" width="100%"/>
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
    <div class="line"></div>
    <div class="logout_button_container">
      <button class="button_basic logout_button" @click="auth.logout()">
        Log out
      </button>
    </div>
  </div>
</template>

<style scoped>
div.account_data_component{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-grow: 1;
}
div.user_data{
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: space-between;
}
div.user_fields{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
div.customer_fields{
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}
div.host_fields{
  display: flex;
  flex-direction: column;
  width: 50%;
  padding: 2% 5%;
}
div.logout_button_container{
  margin-right: 10%;
  margin-left: 5%;
}
div.line{
  display: flex;
  flex-direction: row;
  color: black;
  border-right: 1px solid;
}
div.buttons{
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

</style>