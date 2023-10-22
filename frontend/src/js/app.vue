<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showMessage">This is test</button>
    <div>{{message}}</div>
  </div>
</template>
<script>
import axios from "axios";
const instance = axios.create({
  baseURL: 'http://localhost:3000/api/v1/',
});
const appData = {
  message: ""
};
export default {
  data() {
    return {
      message: ""
    };
  },
  methods: {
    showMessage: function () {
      instance.get("hello").then((res) => {
        console.log(res);
        this.message = res.data.message;
      });
    }
  }
};
</script>
<style>
</style>