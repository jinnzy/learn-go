<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <pagea></pagea>
    <pageb></pageb>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import pagea from './components/a'
import pageb from './components/b'
export default {
  name: 'app',
  data() {
    return {
      ages: 184
    }
  },
  components: {
    HelloWorld,
    pagea,
    pageb
  },
  methods: {
    msg: function () {
      console.log(this.ages++);

    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
