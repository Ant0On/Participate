<script setup>
import {ref, defineProps} from 'vue';
import {storeToRefs} from 'pinia';
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore()
const {previousSearch: previousSearch} = storeToRefs(searchStore)

const searchInput = ref(null)
const props = defineProps({
  searchCallback: {
    type: Function,
    default: (item) => {console.log("default callback", item)}
  },
})
const items = ref(previousSearch.value)

function onAppendIconClick(){
  searchStore.addPreviousSearch( {
    prependIcon: 'mdi-clock-outline',
    title: searchInput.value,
  })
  props.searchCallback(searchInput.value)
}

searchStore.$subscribe((mutation, state)=>{
  items.value = mutation['payload']?.previousSearch
}, {detached: true})
</script>

<template>
  <v-card class="w-100 d-flex flex-row justify-center" elevation="0">
    <v-autocomplete style="max-width: 50%;"
                  bg-color="blue-grey-lighten-5"
                  class="mx-auto"
                  density="compact"
                  clearable
                  placeholder="Search..."
                  auto-select-first
                  item-props
                  variant="solo"
                  hide-details
                  single-line
                  elevation="0"
                  rounded="pill"
                  flat
                  menu-icon=""
                  :items="items"
                  append-inner-icon="mdi-magnify"
                  @click:append-inner="onAppendIconClick"
                  @keyup.enter="onAppendIconClick"
                  @update:search="(input) => searchInput = input"

    >
    </v-autocomplete>

  </v-card>
</template>

<style scoped>

</style>