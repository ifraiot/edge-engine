<template>
  <v-app>
    <nav>
      <v-toolbar>
        <v-toolbar-title>
          <img :src="require('@/assets/logo.png')" alt="Logo" height="40">
        </v-toolbar-title>
        <router-link to="/">
          <v-btn class="text-none" stacked>
            <v-icon>mdi-home</v-icon>
            Home
          </v-btn>
        </router-link>
        <router-link to="/connectors">
          <v-btn class="text-none" stacked>
            <v-icon>mdi-link-variant</v-icon>
            Connectors
          </v-btn>
        </router-link>
        <router-link to="/analyzers">
          <v-btn class="text-none" stacked>
            <v-icon>mdi-chart-bell-curve</v-icon>
            Analyzers
          </v-btn>
        </router-link>
        <router-link to="/integrators">
          <v-btn class="text-none" stacked>
            <v-icon>mdi-connection</v-icon>
            Integrators
          </v-btn>
        </router-link>

        <router-link to="/data">
          <v-btn class="text-none" stacked>
            <v-icon>mdi-database</v-icon>
            Data
          </v-btn>
        </router-link>

        <v-spacer></v-spacer>

        <v-btn text>
          <span>Sign Out</span>
          <v-icon right>exit_to_app</v-icon>
        </v-btn>
      </v-toolbar>
    </nav>

    <v-container>
      <router-view />
    </v-container>

  </v-app>
</template>

<script>
export default {
  name: 'App',
  components: {
    // HomePage
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
</style>
