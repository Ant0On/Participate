<script setup>
import NavBar from "@/components/nav/NavBar.vue";
import HomeOfferSearch from "@/components/home/HomeOfferSearch.vue";
import HomeWelcomeText from "@/components/home/HomeWelcomeText.vue";
import HomeOfferImage from "@/components/home/HomeOfferImage.vue";
</script>

<template>
  <div class="home_page">
    <NavBar currentPage="home"/>
  <div class="search_and_welcome">
    <div class="welcome">
      <HomeWelcomeText/>
      <div class="image">
        <HomeOfferImage/>
      </div>
    </div>
    <div class="offer_search">
      <HomeOfferSearch/>
    </div>
  </div>
    </div>
</template>

<style scoped>

.home_page {
  display: flex;
  flex-direction: column;
}
.search_and_welcome {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  margin-left: 7%;
}
.welcome {
  display: flex;
  flex-direction: column;
  width: 50%;

}

.offer_search {
  //display: flex;
  //flex-direction: row-reverse;
  margin: 30px 8% 10px 16%;
}
</style>