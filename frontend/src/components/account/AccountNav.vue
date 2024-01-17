<script setup>
import { defineEmits, ref} from 'vue';
import AccountNavItem from "@/components/account/AccountNavItem.vue";

const emits = defineEmits(['pageChanged'])
const activeItem = ref('Account Information');

function onItemClicked(navItem){
  activeItem.value = navItem
  emits('pageChanged', navItem)
}


</script>

<template>
<div class="account_nav">
  <AccountNavItem text="Account Information" :is-active="activeItem === 'Account Information'" @item-clicked="onItemClicked" style="margin-top: 20px"/>
  <AccountNavItem text="History" :is-active="activeItem === 'History'" @item-clicked="onItemClicked"/>
  <AccountNavItem text="New Offer" :is-active="activeItem === 'New Offer'" @item-clicked="onItemClicked"/>
  <AccountNavItem text="Current Offers" :is-active="activeItem === 'Current Offers'" @item-clicked="onItemClicked"/>
  <AccountNavItem text="Become a host" :is-active="activeItem === 'Become a host'" @item-clicked="onItemClicked"/>
</div>
</template>

<style scoped>
div.account_nav{
  display: flex;
  flex-direction: column;
  margin-left: 5%;
  margin-top: 2%;
  padding-right: 20px;
  border-right: 1px solid;
  height: max(70%, 400px);
}
</style>