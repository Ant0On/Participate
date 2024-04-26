<script setup>
import {ref} from 'vue';
import {storeToRefs} from 'pinia';

import LoginSection from "@/components/login/LoginSection.vue";
import SignUpSection from "@/components/login/SignUpSection.vue";
import {useAuthStore} from "@/stores/auth.store";
import AccountData from "@/components/account/data/AccountData.vue";

const loginPage = ref(true)

const userStore = useAuthStore()
const {user: user} = storeToRefs(userStore);

const currentPage = ref('Account Information')


</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <div class="login_container" v-if="!user">
      <Transition name="slide-fade">
        <div class="login" v-if="loginPage">
          <LoginSection @sign-up-clicked="loginPage = !loginPage"/>
          <img alt="Image" src="../assets/img/login_image.png"/>
        </div>
        <div class="signup" v-else>
          <img alt="Image" src="../assets/img/login_image.png"/>
          <SignUpSection/>
        </div>
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
div.data{
  margin: 2%;
  display: flex;
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
.logging {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.login {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 30px 9% 10px 9%;
  background: white;
}
.signup {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 30px 9% 10px 9%;
  background: white;
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

@media(max-width: 900px){
  img{
    display: none;
  }
  .login_container{
    margin-right: 30px;
    margin-left: 30px;
  }
}

</style>