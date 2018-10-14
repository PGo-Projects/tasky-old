<template>
<v-app id="inspire" dark>
  <v-toolbar app fixed clipped-right>
    <v-toolbar-title>Tasky</v-toolbar-title>
    <v-spacer></v-spacer>
    
    <v-toolbar-items class="hidden-sm-and-down">
      <v-btn flat href="/register">
        <v-icon>fa fa-user-plus</v-icon>
        <span class="ml-2">Register</span>
      </v-btn>
      
      <v-btn flat href="/login">
        <v-icon>fa fa-sign-in-alt</v-icon>
        <span class="ml-2">Login</span>
      </v-btn>
    </v-toolbar-items>
    
    <v-toolbar-side-icon class="hidden-md-and-up" @click.stop="drawer = !drawer;"></v-toolbar-side-icon>
  </v-toolbar>
  
  <v-navigation-drawer right fixed mobile-break-point=102400 v-model="drawer" app>
    <v-list>
      <v-list-tile>
        <v-list-tile-content>
          Tasky
        </v-list-tile-content>
      </v-list-tile>
      
      <v-list-tile href="/register">
        <v-list-tile-action>
          <v-icon>fa fa-user-plus</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Register</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      
      <v-list-tile href="/login">
        <v-list-tile-action>
          <v-icon>fa fa-sign-in-alt</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Login</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
  
  <v-content>
    <v-container fluid fill-height>
      <v-layout column align-center justify-center>
        <span class="display-3">Welcome to Tasky!</span>
        <span text-xs-center>The best To-Do app you will ever use!</span>
      </v-layout>
    </v-container>
  </v-content>
  
  <footer-component></footer-component>
</v-app>
</template>

<script>
import FooterComponent from 'Footer/Footer.vue';

export default {
    name: 'Index',
    components: {FooterComponent},
    data: () => ({
        drawer: false,
    }),
    methods: {
    },
    props: [],
}
</script>
