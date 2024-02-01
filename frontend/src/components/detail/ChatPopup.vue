<!-- ChatPopup.vue -->
<template>
  <div class="chat-popup">
    <div class="chat-header">
      <h2>Chat Room</h2>
      <button @click="closeChat">Close Chat</button>
    </div>
    <div class="chat-messages">
      <!-- Display chat messages here -->
      <div v-for="(message, index) in messages" :key="index" class="chat-message">
        {{ message }}
      </div>
    </div>
    <div class="chat-input">
      <textarea v-model="newMessage" placeholder="Type your message..." @keydown.enter.prevent="sendMessage"></textarea>
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [],       // Array to store chat messages
      newMessage: '',     // Input field for new messages
    };
  },
  methods: {
    closeChat() {
      // Emit an event to notify the parent component to close the chat
      this.$emit('close-chat');
    },
    sendMessage() {
      if (this.newMessage.trim() !== '') {
        this.messages.push(`User: ${this.newMessage}`);
        this.newMessage = '';
      }
    },
  },
};
</script>

<style scoped>
.chat-popup {
  position: fixed;
  bottom: 15px;
  right: 15px;
  border: 1px solid #ccc;
  background-color: #fff;
  padding: 15px;
  z-index: 1;
  width: 300px; /* Adjust the width as needed */
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chat-messages {
  max-height: 200px; /* Set a max height for the chat messages area */
  overflow-y: auto;
}

.chat-message {
  margin-bottom: 5px;
}

.chat-input {
  margin-top: 10px;
}

textarea {
  width: 100%;
  padding: 5px;
  margin-bottom: 5px;
}

button {
  background-color: #4caf50;
  color: white;
  border: none;
  padding: 8px 16px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 14px;
  cursor: pointer;
}
</style>
