<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showMessage">This is test</button>
    <div>{{message}}</div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      message: ""
    };
  },
  methods: {
    showMessage: function () {
      axios.get("http://localhost:3000/api/v1/hello").then((res) => { // TODO ogarnąć adres
        console.log(res);
        this.message = res.data.message;
      });
    }
  }
};
</script>
<style>
</style>