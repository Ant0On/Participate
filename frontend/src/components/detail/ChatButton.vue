<!-- ChatButton.vue -->
<template>
  <button class="chat-button" @click="joinChat">Join Chat</button>
</template>

<script>
export default {
  props: {
    type: String,
    id: String,
  },
  methods: {
    joinChat() {
      // Emit an event to notify the parent component to open the chat
      this.$emit('join-chat', { type: this.type, id: this.id });
    },
  },
};
</script>

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
}
</style>
