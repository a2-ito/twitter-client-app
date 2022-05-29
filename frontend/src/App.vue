<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar color="primary" dark app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Twitter Client Application</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text>Login</v-btn>
    </v-app-bar>
    <v-main>
      <router-link to="/">home</router-link>
      <router-link to="/top">top</router-link>
      <router-link to="/about">about</router-link>
      <router-view></router-view>
    </v-main>
    <v-footer color="primary" dark app>
      Vuetify
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    data: () => ({ drawer: null }),
  }
</script>
