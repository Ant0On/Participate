<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <ExamplePage msg="Welcome to Your Vue.js App"/>
</template>

<script>
import ExamplePage from './components/ExamplePage.vue'
import './assets/css/app.css'

export default {
  name: 'App',
  components: {
    ExamplePage
  }
}
</script>
