<script setup>
import {defineProps, defineEmits, onMounted} from 'vue';
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

const props = defineProps({
  messages: [],
  newMessage: '',
  email: '',
  offerID: 0,
  userID: 0,
  chatID: 0
})

const emits = defineEmits(['closeChat'])

function sendMessage() {
  if (props.newMessage.trim() !== '') {
    props.messages.push(`${props.email}: ${props.newMessage}`);
    fetchWrapper.post(`/api/customer/${props.userID}/${props.offerID}/message/send`, {
      'customer_id': props.userID,
      'email': props.email,
      'content': props.newMessage,
      'chat_id': props.chatID
    }).catch()
    props.newMessage = '';
  }
}

onMounted(async()=> {
  await fetchWrapper.get(`/api/chat/${props.chatID}/messages`).then(response => {
    props.messages = response.data
  })
})

</script>
<template>
  <div class="chat-popup">
    <div class="chat-header">
      <h2>Chat Room</h2>
      <button @click="$emit('closeChat')">Close Chat</button>
    </div>
    <div class="chat-messages">
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
</style>
