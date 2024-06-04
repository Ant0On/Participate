<script setup>
import {computed, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useAuthStore} from "@/stores/auth.store";

const drawer = ref(true);
const rail = ref(true);

const userStore = useAuthStore()

const {
  user: user
} = storeToRefs(userStore);

const userImage = computed(() => {
  try {
    return require(`@/../images/users/${user.value?.ID}.jpeg`)
  } catch {
    return undefined
  }
})
const userName = computed(() => {
  return (user.value) ? user.value.FirstName + " " + user.value.LastName : "Guest"
})

</script>

<template>
  <v-navigation-drawer
      v-model="drawer"
      :rail="rail"
      permanent
      @click="rail = false"
      location="right"
      color="blue-grey-darken-1"
  >
    <v-list-item
        :title="userName"
        nav
    >
      <template v-slot:prepend>
        <v-avatar v-if="userImage" :image="userImage"></v-avatar>
        <v-avatar v-else>
          <v-icon icon="mdi-account-circle"/>
        </v-avatar>
      </template>

      <template v-slot:append>
        <v-btn
            icon="mdi-chevron-left"
            variant="text"
            @click.stop="rail = !rail"
        ></v-btn>
      </template>
    </v-list-item>

    <v-divider></v-divider>

    <v-list density="compact" nav>
      <div v-if="user">
        <router-link to="/account">
          <v-list-item prepend-icon="mdi-account" title="My account" value="account" @click=""></v-list-item>
        </router-link>
        <router-link to="/account/change_password">
          <v-list-item prepend-icon="mdi-account-key" title="Change password" value="change_password" @click=""></v-list-item>
        </router-link>
        <div v-if="user.Role === 'customer'">
          <router-link to="/account/reservations_history">
            <v-list-item prepend-icon="mdi-account-box" title="Reservations history" value="reservations_history"></v-list-item>
          </router-link>
          <router-link to="/account/become_host">
            <v-list-item prepend-icon="mdi-account-plus" title="Become a host" value="become_host"
                         @click=""></v-list-item>
          </router-link>
        </div>
        <div v-else-if="user.Role === 'host'">
          <router-link to="/account/my_offers">
            <v-list-item prepend-icon="mdi-account-details" title="My offers" value="my_offers" @click=""></v-list-item>
          </router-link>
          <router-link to="/account/new_offer">
            <v-list-item prepend-icon="mdi-invoice-text-plus" title="New offer" value="new_offer"
                         @click=""></v-list-item>
          </router-link>
          <router-link to="/account/current_reservations">
            <v-list-item prepend-icon="mdi-account-group-outline" title="Current reservations" value="current_reservations"
                         @click=""></v-list-item>
          </router-link>
        </div>
        <v-list-item prepend-icon="mdi-account-arrow-left" title="Log out" value="logout"
                     @click="userStore.logout()"></v-list-item>
      </div>
      <router-link v-else to="/login">
        <v-list-item v-else prepend-icon="mdi-account-arrow-left" title="Log in" value="login"></v-list-item>
      </router-link>
    </v-list>
  </v-navigation-drawer>
</template>

<style scoped>

</style>