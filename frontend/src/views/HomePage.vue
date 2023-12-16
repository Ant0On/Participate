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
        <router-link :to="{ path: '/accommodations' }">
          <HomeOfferImage v-bind:src="require('../assets/img/accomodation_home_page.jpg')"/>
        </router-link>
        <router-link :to="{ path: '/activities' }">
          <HomeOfferImage v-bind:src="require('../assets/img/activity_home_page.jpg')"/>
        </router-link>
        <router-link :to="{ path: '/events' }">
          <HomeOfferImage v-bind:src="require('../assets/img/event_home_page.jpg')"/>
        </router-link>
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
}
.welcome {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  width: 50%;
}

.image {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-left: 7%;
  margin-top: 2%;
}

.offer_search {
  margin: 30px 5% 10px 5%;
  width: 30%;
}
</style>