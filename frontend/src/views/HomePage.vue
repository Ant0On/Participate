<script setup>
import HomeOfferSearch from "@/components/home/HomeOfferSearch.vue";
import HomeWelcomeText from "@/components/home/HomeWelcomeText.vue";
</script>

<template>
  <div class="home_page">
    <div class="search_and_welcome">
      <div class="welcome">
        <HomeWelcomeText/>
        <HomeOfferSearch/>
      </div>
      <div class="image_container">
        <router-link :to="{ path: '/accommodations' }">
          <div
              class="image-wrapper"
              @mouseover="handleHover(true, 0)"
              @mouseout="handleHover(false, 0)"
          >
            <img
                class="image"
                src="../assets/img/accomodation_home_page.jpg"
                alt="Accommodation Image"
            />
            <div v-if="hoverStates[0]" class="hover-content">
              <p class="hover-text">{{ hoverContent[0].text }} <span class="hover-emoji">{{ hoverContent[0].emoji }}</span>
              </p>
            </div>
          </div>
        </router-link>
        <router-link :to="{ path: '/activities' }">
          <div
              class="image-wrapper"
              @mouseover="handleHover(true, 1)"
              @mouseout="handleHover(false, 1)"
          >
            <img
                class="image"
                src="../assets/img/activity_home_page.jpg"
                alt="Activity Image"
            />
            <div v-if="hoverStates[1]" class="hover-content">
              <p class="hover-text">{{ hoverContent[1].text }} <span class="hover-emoji">{{ hoverContent[1].emoji }}</span>
              </p>
            </div>
          </div>
        </router-link>
        <router-link :to="{ path: '/events' }">
          <div
              class="image-wrapper"
              @mouseover="handleHover(true, 2)"
              @mouseout="handleHover(false, 2)"
          >
            <img
                class="image"
                src="../assets/img/event_home_page.jpg"
                alt="Accommodation Image"
            />
            <div v-if="hoverStates[2]" class="hover-content">
              <p class="hover-text">{{ hoverContent[2].text }} <span class="hover-emoji">{{ hoverContent[2].emoji }}</span>
              </p>
            </div>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      hoverStates: [false, false, false],
      hoverContent: [
        {
          text: "Rest",
          emoji: "⛺",
        },
        {
          text: "Explore",
          emoji: "🧭",
        },
        {
          text: "Experience",
          emoji: "🪂",
        }
      ]
    };
  },
  methods: {
    handleHover(value, index) {
      this.hoverStates[index] = value;
    },
  },
};
</script>

<style scoped>

.home_page {
  display: flex;
  flex-direction: column;
  padding-top: 16px;
  padding-bottom: 16px;
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
  transition: transform 0.3s ease;
}

.image-wrapper {
  position: relative;
  overflow: hidden;
}

.image-wrapper:hover .image {
  transform: scale(1.1);
}

.hover-content {
  position: absolute;
  top: 5%;
  left: 50%;
  transform: translateX(-50%);
  background-color: white;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center;
  z-index: 1;
  color: #000;
  font-family: DM Serif Display, sans-serif;
  font-size: 20px;
  font-style: normal;
  font-weight: 800;
  line-height: normal;
}

</style>