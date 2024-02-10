<script setup>
import {defineProps} from 'vue';
import NavItem from "./NavItem.vue";
import {useAuthStore} from "@/stores/auth.store";

const props = defineProps({
  currentPage: String
})

function isActive(text, currentPage) {
  return text === currentPage
}
const auth = useAuthStore();
const user = auth.user;

const navigateTo = (auth.user === null)? '/login' : '/account'
const text = (auth.user === null)? 'Log in': auth.user.FirstName + ' ' + auth.user.LastName;
</script>

<template>
  <nav>
    <img alt="Image" src="../../assets/img/logo.png"/>
    <div>
      <NavItem text="Home" :isActive="isActive('home', currentPage)" navigateTo="/"/>
      <NavItem text="Recommended" :isActive="isActive('recommended', currentPage)" navigateTo="/recommended"/>
      <NavItem text="Accommodations" :isActive="isActive('accommodations', currentPage)" navigateTo="/accommodations"/>
      <NavItem text="Activities" :isActive="isActive('activities', currentPage)" navigateTo="/activities"/>
      <NavItem text="Events" :isActive="isActive('events', currentPage)" navigateTo="/events"/>
      <NavItem :text="text" :isActive="isActive('login', currentPage)" :navigateTo="navigateTo"/>
    </div>
  </nav>
</template>

<style scoped>
img {
  object-fit: fill;
  max-height: 100%;
  max-width: 100%;
  height: 100px;
  width: 300px;
}
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 0.3rem;
  padding: 0.3rem;
}

div {
  display: flex;
  justify-content: flex-end;
  max-height: 3rem;
  column-gap: 10px;
}

@media(max-width: 900px){
  nav{
    flex-direction: column;
    margin: auto;
  }
  div{
    display: flex;
    flex-direction: column;
    max-height: fit-content;
  }

}
</style>