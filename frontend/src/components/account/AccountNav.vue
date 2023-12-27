<script setup>
import { defineEmits, ref, computed } from 'vue';
import AccountNavItem from "@/components/account/AccountNavItem.vue";

const emits = defineEmits(['pageChanged'])
const activeItem = ref('Account Information');

function onItemClicked(navItem){
  activeItem.value = navItem
  emits('pageChanged', navItem)
}
const accountInfo = computed(()=> activeItem.value === 'Account Information')
const history = computed(()=> activeItem.value === 'History')
</script>

<template>
<div class="account_nav">
  <AccountNavItem text="Account Information" :is-active="accountInfo" @item-clicked="onItemClicked"/>
  <AccountNavItem text="History" :is-active="history" @item-clicked="onItemClicked"/>
</div>
</template>

<style scoped>
div.account_nav{
  margin-left:;
}
</style>