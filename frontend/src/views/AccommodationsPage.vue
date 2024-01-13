<script setup>
import { ref, watch, onMounted } from 'vue';
import { storeToRefs } from 'pinia';

import NavBar from "@/components/nav/NavBar.vue";
import OfferSearch from "@/components/offers/OfferSearch.vue";
import OfferListItem from "@/components/offers/OfferListItem.vue";
import { useSearchStore } from "@/stores/search.store";
import { fetchWrapper } from "@/_helpers/fetch-wrapper";

const searchStore = useSearchStore();
const { location, dateFrom, dateTo, numberOfPeople } = storeToRefs(searchStore);

const accommodations = ref([]);

const getCurrentAccommodations = () => {
  fetchWrapper.get('/api/offers')
      .then(response => {
        const data = response.data;

        if (data && Array.isArray(data)) {
          // Fetch town information for each offer
          const promises = data.map(offer => {
            const townId = offer.town_id;

            if (!townId) {
              console.error('Town ID is missing for offer:', offer);
              return Promise.resolve(null);
            }

            return fetchWrapper.get(`/api/town/get/${townId}`)
                .then(townResponse => {
                  console.log('Town Response:', townResponse.data);
                  return townResponse.data;
                })
                .then(town => ({ ...offer, town }));
          });

          Promise.all(promises)
              .then(accommodationsWithTown => {
                console.log('Accommodations with Town:', accommodationsWithTown);
                accommodations.value = accommodationsWithTown.filter(item => item !== null);
              })
              .catch(error => {
                console.error('Error fetching town information:', error);
              });
        } else {
          console.error('Invalid response format:', response);
        }
      })
      .catch(error => {
        console.error('Error fetching accommodations:', error);
      });
};


onMounted(() => {
  getCurrentAccommodations();
});

watch(location, getCurrentAccommodations);
watch(dateFrom, getCurrentAccommodations);
watch(dateTo, getCurrentAccommodations);
watch(numberOfPeople, getCurrentAccommodations);

</script>

<template>
  <div class="accommodation_page">
    <NavBar currentPage="accommodations"/>
    <p>Climatic places</p>
    <OfferSearch offer_type="Accommodation" v-model:location="location"
                 v-model:date-from="dateFrom" v-model:date-to="dateTo"
                 v-model:number-of-people="numberOfPeople"/>
    <div class="offer_items">
      <OfferListItem v-for="accommodation in accommodations"
                     :location="accommodation.town.name"
                     :description="accommodation.description" :image="accommodation.image" :name="accommodation.name"
                     :price="accommodation.price" :max_people="accommodation.max_people"/>
    </div>
  </div>
</template>

<style scoped>
div.accommodation_page {
  display: flex;
  flex-direction: column;
  overflow: scroll;
}

p {
  color: #000000;
  font-family: "Poppins", Helvetica;
  font-size: 1.8rem;
  font-weight: 700;
  line-height: normal;
  align-self: center;
  margin: 1% 1% 1% 1%;
}

.offer_items {
  margin: 3% 5% 0 5%;
}
</style>
