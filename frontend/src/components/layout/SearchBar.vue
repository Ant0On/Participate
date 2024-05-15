<script setup>
import {defineEmits, defineProps, ref} from 'vue';
import { router } from "@/router";
import {storeToRefs} from 'pinia';
import {useSearchStore} from "@/stores/search.store";

const searchStore = useSearchStore()
const {previousSearch: previousSearch} = storeToRefs(searchStore)
const type = ref("accommodations")
const isSearchLocalization = ref(false)
const searchInput = ref(null)
const offerType = [
  {
    title: "Accommodation",
    value: "accommodations"
  },
  {
    title: "Events",
    value: "events"
  },
  {
    title: "Activities",
    value: "activities"
  }]
const props = defineProps({
  searchNameCallback: {
    type: Function,
    default: (item) => {
      console.log("default callback", item)
    }
  },
  searchLocalizationCallback: {
    type: Function,
    default: (item) => {
      console.log("default callback", item)
    }
  },
  main: {
    type: Boolean,
    default: false,
  }
})
const emits = defineEmits(['update:searchItems'])
const items = ref(previousSearch.value)

function onAppendIconClick() {
  searchStore.addPreviousSearch({
    prependIcon: 'mdi-clock-outline',
    title: searchInput.value,
  })
  if(props.main)
  {
    router.push(`/${type.value}`)
  }
  if (isSearchLocalization.value) {
    props.searchLocalizationCallback(searchInput.value)
    return
  }
  props.searchNameCallback(searchInput.value)
}

searchStore.$subscribe((mutation, state) => {
  items.value = mutation['payload']?.previousSearch
}, {detached: true})
</script>

<template>
  <div class="w-100 d-flex flex-row justify-center">

    <v-autocomplete style="max-width: 60%"
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
      <template v-slot:append>
        <v-switch v-model="isSearchLocalization" label="Search localization" color="primary" density="compact"
                  style="height: 40px">

        </v-switch>
      </template>
      <template v-slot:prepend>
        <v-select v-if="main" v-model="type" :items="offerType" style="height: 40px; width:181px"
                  density="compact"
                  bg-color="blue-grey-lighten-5"
                  variant="solo"
                  single-line
                  rounded="pill"
                  placeholder="Type"
                  flat
        >

        </v-select>
      </template>
    </v-autocomplete>
  </div>
</template>

<style scoped>

</style>