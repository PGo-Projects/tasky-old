<template>
<v-form ref="form">
  <v-text-field
    v-model="username"
    :rules="[rules.minusername]"
    prepend-icon="account_circle"
    label="Username"
    required
    autofocus
    ></v-text-field>
  <v-text-field
    v-model="password"
    :append-icon="showPassword ? 'visibility_off' : 'visibility'"
    :rules="[rules.minpassword]"
    :type="showPassword ? 'text' : 'password'"
    prepend-icon="lock"
    label="Password"
    hint="At least 12 characters"
    required
    @click:append="showPassword = !showPassword"
    ></v-text-field>
  <v-text-field
    v-model="confirmationPassword"
    :append-icon="showConfirmationPassword ? 'visibility_off' : 'visibility'"
    :type="showConfirmationPassword ? 'text' : 'password'"
    :error-messages='confirmPasswordMatch()'
    prepend-icon="lock"
    label="Confirmation Password"
    hint="Enter the same password again"
    required
    @click:append="showConfirmationPassword = !showConfirmationPassword"
    ></v-text-field>
  
  <v-layout row>
    <v-btn flat small dark href="/login">Sign in instead</v-btn>
    <v-spacer></v-spacer>
    <v-btn dark @click="submit">Register</v-btn>
  </v-layout>
</v-form>
</template>

<script>
export default {
    name: 'RegisterForm',
    components: {},
    data: () => ({
        username: "",
        password: "",
        confirmationPassword: "",
        showPassword: false,
        showConfirmationPassword: false,
        rules: {
            minusername: v => v.length >= 4 || 'Needs to be at least 4 characters',
            minpassword: v => v.length >= 12 || 'Needs to be at least 12 characters',
        }
    }),
    methods: {
        confirmPasswordMatch() {
            return (this.password == this.confirmationPassword) ? '' : 'Password does not match';
        },
        submit() {
            var isValid = this.$refs.form.validate() && this.password == this.confirmationPassword;
            if (isValid) {
                console.log('Submitted!');
            } else {
                console.log('No good');
            }
        },
    },
    props: [],
}
</script>
