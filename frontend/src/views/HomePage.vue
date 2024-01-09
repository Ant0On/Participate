<script setup>
import NavBar from "@/components/nav/NavBar.vue";
import HomeOfferSearch from "@/components/home/HomeOfferSearch.vue";
import HomeWelcomeText from "@/components/home/HomeWelcomeText.vue";
import HomeOfferImage from "@/components/home/HomeOfferImage.vue";

const getApiImageUrl = () => `http://localhost:8080/api/image/get`;
</script>

<template>
  <div class="home_page">
    <NavBar currentPage="home"/>
  <div class="search_and_welcome">
    <div class="welcome">
      <HomeWelcomeText/>
      <HomeOfferSearch/>
    </div>
    <div class="image">
      <router-link :to="{ path: '/accommodations' }">
        <HomeOfferImage v-bind:src="require('../assets/img/accomodation_home_page.jpg')"/>
      </router-link>
      <router-link :to="{ path: '/activities' }">
        <HomeOfferImage v-bind:src="require('../assets/img/activity_home_page.jpg')"/>
      </router-link>
      <router-link :to="{ path: '/events' }">
        <HomeOfferImage v-bind:src="getApiImageUrl()"/>
      </router-link>
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
  flex-direction: column;
}
.welcome {
  display: flex;
  flex-direction: row;
  gap: 15%;
  flex-grow: 1;
  width: 95%;
}

.image {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-top: 3%;
  column-gap: 2%;
  width: 95%;
}

</style>