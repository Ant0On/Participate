<script setup>
import Footer from "@/components/layout/Footer.vue";
import NavBar from "@/components/nav/NavBar.vue";
import LoginSection from "@/components/login/LoginSection.vue";
</script>

<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css">
  <NavBar currentPage="login"/>
  <div class="overlap-group">
    <LoginSection/>
    <img class="image" alt="Image" src="../assets/img/image-3.png"/>
  </div>
  <Footer/>
</template>

<style scoped>
.logging .div-2 {
  background-color: #efefef;
  height: 1024px;
  position: relative;
  width: 1440px;
}

.logging .overlap-group {
  background-image: url(../assets/img/rectangle-16.png);
  background-size: 100% 100%;
  height: 700px;
  left: 120px;
  position: absolute;
  top: 217px;
  width: 1200px;
}

.logging .p {
  color: #000000;
  font-family: "Sarabun", Helvetica;
  font-size: 25px;
  font-weight: 100;
  left: 49px;
  letter-spacing: 0;
  line-height: normal;
  position: absolute;
  top: 109px;
  width: 328px;
}

.logging .text-input-instance {
  left: 49px !important;
  position: absolute !important;
  top: 234px !important;
}

.logging .design-component-instance-node {
  height: 78px !important;
  left: 49px !important;
  position: absolute !important;
  top: 371px !important;
}

.logging .button-large {
  align-items: flex-start;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  height: 43px;
  position: relative;
  width: 226px;
}

.logging .button-large-filled {
  background-color: var(--systemwhite) !important;
  display: flex !important;
  height: 43px !important;
  width: 226px !important;
}

.logging .button-instance {
  align-self: stretch !important;
  color: var(--text-link) !important;
  font-family: "IBM Plex Sans", Helvetica !important;
  font-size: 18px !important;
  font-style: unset !important;
  font-weight: 500 !important;
  letter-spacing: 0 !important;
  line-height: normal !important;
  white-space: unset !important;
}

.logging .div-wrapper {
  align-items: flex-start;
  border: 1px solid;
  border-color: var(--text-link);
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  height: 43px;
  left: 49px;
  overflow: hidden;
  position: absolute;
  top: 512px;
  width: 224px;
}

.logging .button-large-filled-wrapper {
  align-items: flex-start;
  background-color: #ffffff;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  height: 43px;
  position: relative;
  width: 224px;
}

.logging .text-wrapper-3 {
  color: #0048d9;
  font-family: "IBM Plex Sans", Helvetica;
  font-size: 18px;
  font-weight: 500;
  height: 21px;
  left: 355px;
  letter-spacing: 0;
  line-height: normal;
  position: absolute;
  text-align: center;
  top: 567px;
  white-space: nowrap;
  width: 171px;
}

.logging .image {
  height: 700px;
  left: 600px;
  object-fit: cover;
  position: absolute;
  top: 0;
  width: 600px;
}

.logging .img {
  height: 89px;
  left: 120px;
  object-fit: cover;
  position: absolute;
  top: 36px;
  width: 275px;
}

.logging .overlap {
  height: 54px;
  left: 561px;
  position: absolute;
  top: 36px;
  width: 755px;
}


</style>