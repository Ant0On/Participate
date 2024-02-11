<script setup>
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
  type: String,
  id: String,
  isHost: Boolean,
  isChatAlreadyCreated: Boolean
})

const emits = defineEmits(['joinChat'])
</script>

<template>
  <button class="chat-button" :disabled=" (!isChatAlreadyCreated && !isHost)" @click="$emit('joinChat')">{{ (isChatAlreadyCreated && isHost) ? 'Create Chat' : 'Join Chat' }}</button>
</template>

<style scoped>
.chat-button {
  background-color: #4caf50;
  color: white;
  border: none;
  padding: 10px 20px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  border-radius: 10px;
}
</style>
