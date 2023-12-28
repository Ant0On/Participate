<script setup>

import {useAuthStore} from "@/stores/auth.store";
import PasswordInput from "@/components/ui/PasswordInput.vue";
import TextInput from "@/components/ui/TextInput.vue";

const auth = useAuthStore();
const user = auth.user;

</script>

<template>
<div class="account_data_component">
  <TextInput label-text="Name"/>
  <TextInput label-text="Last Name"/>
  <TextInput label-text="Email"/>
  <PasswordInput/>
</div>
</template>

<style scoped>

</style>