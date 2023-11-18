<script setup>
import Footer from "@/components/layout/Footer.vue";
import NavBar from "@/components/nav/NavBar.vue";
import LoginSection from "@/components/login/LoginSection.vue";
</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <div class="logging">
    <NavBar currentPage="login"/>
    <div class="login_section">
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
  justify-content: space-around;
}


.login_section {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 30px 9% 10px;
  background: white;
  max-width: 1200px;
}

</style>