<template>
<v-app id="inspire" dark>
  <v-content>
    <v-container fluid fill-height>
      <v-layout align-center>
        <v-flex xs12 sm8 offset-sm2 md6 offset-md3 lg6 offset-lg3>
          <span class="display-4">
            <v-img src="/public/assets/logo.png" max-height=175 alt="Tasky" contain class="mb-5"></v-img>
          </span>
          <register-form :csrf-token="this.csrfToken" :error-message="this.errorMessage"></register-form>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>

  <footer-component></footer-component>
</v-app>
</template>

<script>
import RegisterForm from 'Register/RegisterForm.vue';
import FooterComponent from 'Footer/Footer.vue';

export default {
    name: 'Register',
    components: {RegisterForm, FooterComponent},
    data: () => ({
    }),
    methods: {
    },
    props: ['csrfToken', 'errorMessage'],
}
</script>
