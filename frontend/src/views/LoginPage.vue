<script setup>
import {ref} from 'vue';
import Footer from "@/components/layout/FooterBar.vue";
import NavBar from "@/components/nav/NavBar.vue";
import LoginSection from "@/components/login/LoginSection.vue";
import SignUpSection from "@/components/login/SignUpSection.vue";

const loginPage = ref(true)



</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <NavBar currentPage="login"/>
    <div class="login_container">
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
    <Footer/>
  </div>
</template>

<style scoped>
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