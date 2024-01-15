<script setup>
import NavBar from "@/components/nav/NavBar.vue";
import HomeOfferSearch from "@/components/home/HomeOfferSearch.vue";
import HomeWelcomeText from "@/components/home/HomeWelcomeText.vue";

</script>

<template>
  <div class="home_page">
    <NavBar currentPage="home"/>
  <div class="search_and_welcome">
    <div class="welcome">
      <HomeWelcomeText/>
      <HomeOfferSearch/>
    </div>
    <div class="image_container">
      <router-link :to="{ path: '/accommodations' }">
        <img class="image" src="../assets/img/accomodation_home_page.jpg" alt="Accommodation Image"/>
      </router-link>
      <router-link :to="{ path: '/activities' }">
        <img class="image" src="../assets/img/activity_home_page.jpg" alt="Activity Image"/>
      </router-link>
      <router-link :to="{ path: '/events' }">
        <img class="image" src="../assets/img/event_home_page.jpg" alt="Event Image"/>
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

.image_container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-top: 3%;
  column-gap: 2%;
  width: 95%;
}

.image {
  height: auto;
  width: 100%;
  max-width: 450px;
  max-height: 450px;
  margin: 0 13% 5% 13%;
}

</style>