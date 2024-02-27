<script setup>
import { defineProps, defineEmits, onMounted, toRef, ref, nextTick } from 'vue';
import { fetchWrapper } from "@/_helpers/fetch-wrapper";

const props = defineProps({
  messages: {
    type: Array,
    default: () => []
  },
  newMessage: '',
  email: '',
  offerID: 0,
  userID: 0,
  chatID: 0
});

const messages = toRef(props.messages);
const newMessage = toRef(props.newMessage);
const chatMessagesRef = ref(null);
const emits = defineEmits(['closeChat']);
let errorMessage = null;

function sendMessage() {
  if (newMessage.value.trim() !== '') {
    const newMessageContent = `${props.email}: ${newMessage.value}`;
    messages.value.push(newMessageContent);

    fetchWrapper.post(`/api/customer/${props.userID}/${props.chatID}/message/send`, {
      'app_user_id': props.userID,
      'email': props.email,
      'content': newMessage.value,
      'chat_id': props.chatID
    }).catch(() => {
      errorMessage = 'Error sending message. Please try again.';
      const index = messages.value.indexOf(newMessageContent);
      if (index !== -1) {
        messages.value.splice(index, 1);
      }
    });

    newMessage.value = '';

    nextTick(() => {
      scrollToBottom();
    });
  }
}

function scrollToBottom() {
  if (chatMessagesRef.value) {
    chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight;
  }
}

onMounted(async () => {
  try {
    const response = await fetchWrapper.get(`/api/chat/${props.offerID}/messages`);
    messages.value = response.data.map(record => `${record.email}: ${record.content}`) || [];

    nextTick(() => {
      scrollToBottom();
    });
  } catch (error) {
    errorMessage = 'Error fetching messages. Please try again. ' + error.message;
  }
});

</script>

<template>
  <div class="chat-popup">
    <div class="chat-header">
      <h2>Chat Room</h2>
      <button @click="$emit('closeChat')">Close Chat</button>
    </div>
    <div v-if="errorMessage" class="error-message">
      {{ errorMessage }}
    </div>
    <div class="chat-messages" ref="chatMessagesRef">
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

<style scoped>
.chat-popup {
  position: fixed;
  bottom: 15px;
  right: 15px;
  border: 1px solid #ccc;
  background-color: #fff;
  padding: 15px;
  z-index: 1;
  width: 300px;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chat-messages {
  max-height: 200px;
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

.error-message {
  color: red;
  margin-bottom: 10px;
}

</style>
