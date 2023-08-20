<template>
  <v-container>
    <ApplicationManagement />
  </v-container>
</template>

<script>
import ApplicationManagement from './components/ApplicationManagement.vue'

export default {
  name: 'App',
  components: {
    ApplicationManagement
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  /* color: #2c3e50; */
}
</style>
