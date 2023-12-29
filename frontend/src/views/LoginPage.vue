<script setup>
import {ref} from 'vue';
import Footer from "@/components/layout/FooterBar.vue";
import NavBar from "@/components/nav/NavBar.vue";
import LoginSection from "@/components/login/LoginSection.vue";
import SignUpSection from "@/components/login/SignUpSection.vue";
import {useAuthStore} from "@/stores/auth.store";
import AccountNav from "@/components/account/AccountNav.vue";
import AccountData from "@/components/account/AccountData.vue";
import AccountHistory from "@/components/account/AccountHistory.vue";
import AccountNewOffer from "@/components/account/AccountNewOffer.vue";
import AccountCurrentOffer from "@/components/account/AccountCurrentOffer.vue";

const loginPage = ref(true)

const authStore = useAuthStore()
const user = authStore.user;

const currentPage = ref('Account Information')

function onPageChange(page){
  currentPage.value = page;
}

</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <NavBar currentPage="login"/>
    <div class="login_container" v-if="user === null">
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
      <p> Welcome again {{ user.first_name }}!</p>
      <div class="account_data">
        <AccountNav @page-changed="onPageChange"/>
        <div class="data">
          <AccountData v-if="currentPage === 'Account Information'"/>
          <AccountHistory v-else-if="currentPage === 'History'"/>
          <AccountNewOffer v-else-if="currentPage === 'New Offer'"/>
          <AccountCurrentOffer v-else/>
        </div>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<style scoped>
div.account_container{
  display: flex;
  flex-direction: column;
  margin-top: 3%;
  row-gap: 30px;
  height: 450px;
}
div.account_data{
  display: flex;
  flex-direction: row;
}
div.data{
  margin: 2%;
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
  justify-content: space-between;
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