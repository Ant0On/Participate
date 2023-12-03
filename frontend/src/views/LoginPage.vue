<script setup>
import Footer from "@/components/layout/FooterBar.vue";
import NavBar from "@/components/nav/NavBar.vue";
import LoginSection from "@/components/login/LoginSection.vue";
</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <NavBar currentPage="login"/>
    <div class="login_container">
      <LoginSection/>
      <img alt="Image" src="../assets/img/image-3.png"/>
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

.login_container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 30px 9% 10px 9%;
  background: white;
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