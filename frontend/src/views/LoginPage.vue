<script setup>
import {ref} from 'vue';
import {storeToRefs} from 'pinia';
import { useWindowSize } from '@vueuse/core'

import LoginSection from "@/components/login/LoginSection.vue";
import SignUpSection from "@/components/login/SignUpSection.vue";
import {useAuthStore} from "@/stores/auth.store";
import AccountData from "@/components/account/data/AccountData.vue";

const loginPage = ref(true)

const userStore = useAuthStore()
const {user: user} = storeToRefs(userStore);

const currentPage = ref('Account Information')

const { width, height } = useWindowSize()
</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div>
    <div v-if="!user" class="d-flex justify-center align-center">
      <Transition name="slide-fade">
        <v-card v-if="loginPage" class="d-flex flex-row w-75 h-75">
          <LoginSection @sign-up-clicked="loginPage = !loginPage" class="pa-5"/>
          <v-img v-if="width > 600" :src="require('@/assets/img/login_image.png')" class="w-50" cover height="700"></v-img>
        </v-card>
        <v-card v-else class="d-flex flex-row h-75 w-75">
          <v-img v-if="width > 600" :src="require('@/assets/img/login_image.png')" class="w-50" height="700" cover></v-img>
          <SignUpSection class="w-50"/>
        </v-card>
      </Transition>
    </div>
    <div class="account_container" v-else>
      <p> Welcome again {{ user.FirstName }}!</p>
      <div class="account_data">
          <AccountData/>
      </div>
    </div>
  </div>
</template>

<style scoped>
div.account_container{
  display: flex;
  flex-direction: column;
  margin-top: 3%;
  row-gap: 30px;
}
div.account_data{
  display: flex;
  flex-direction: row;
  flex-grow: 1;
}
p{
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.8rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;

}

.slide-fade-enter-active {
  transition: all 0.4s ease-in-out;
}

.slide-fade-leave-active {
  transition: all 0.5s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(-100%);
  opacity: 0;
}

</style>