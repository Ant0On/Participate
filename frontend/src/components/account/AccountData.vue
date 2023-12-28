<script setup>

import {useAuthStore} from "@/stores/auth.store";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import TextInput from "@/components/ui/TextInput.vue";

const auth = useAuthStore();
const user = auth.user;

</script>

<template>
<div class="account_data_component">
  <TextInput label-text="Name" :is-active="false" :model-value="user.first_name"/>
  <TextInput label-text="Last Name" :is-active="false" :model-value="user.last_name"/>
  <TextInput label-text="Email" :is-active="false" :model-value="user.email"/>
</div>
</template>

<style scoped>

</style>