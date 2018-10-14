<template>
<v-form v-model="valid">
  <v-text-field
    v-model="username"
    prepend-icon="account_circle"
    label="Username"
    required
    ></v-text-field>
  <v-text-field
    v-model="password"
    :append-icon="showPassword ? 'visibility_off' : 'visibility'"
    :type="showPassword ? 'text' : 'password'"
    prepend-icon="lock"
    label="Password"
    required
    @click:append="showPassword = !showPassword"
    ></v-text-field>
  
  <v-layout row>
    <v-btn flat small dark href="/register">Create Account</v-btn>
    <v-spacer></v-spacer>
    <v-btn dark @click="submit">Login</v-btn>
  </v-layout>
</v-form>
</template>

<script>
export default {
    name: 'LoginForm',
    components: {},
    data: () => ({
        username: "",
        password: "",
        showPassword: false,
    }),
    methods: {
        submit() {
            console.log('Submitted!');
        },
    },
    props: [],
}
</script>
