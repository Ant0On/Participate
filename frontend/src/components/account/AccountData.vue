<script setup>

import {useAuthStore} from "@/stores/auth.store";

const auth = useAuthStore();
const user = auth.user;

</script>

<template>
<p>ALA</p>
</template>

<style scoped>

</style>