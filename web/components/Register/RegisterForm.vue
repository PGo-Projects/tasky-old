<template>
<div>
	<div class="mb-3">
		<v-alert v-if="formattedErrorMessage !== ''"
			:value="true"
			type="error"
			outline=true
			>
			{{ formattedErrorMessage }}
		</v-alert>
	</div>
	<v-form ref="form" method="post">
  		<v-text-field
    		v-model="username"
    		:rules="[rules.minusername]"
    		prepend-icon="account_circle"
    		label="Username"
			name="username"
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
			name="password"
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
			name="confirm_password"
		    hint="Enter the same password again"
		    required
		    @click:append="showConfirmationPassword = !showConfirmationPassword"
		    ></v-text-field>

		<input type="hidden" name="csrf_token" :value="this.csrfToken" />
  
		<v-layout row>
    		<v-btn flat small dark href="/auth/login">Sign in instead</v-btn>
    		<v-spacer></v-spacer>
    		<v-btn type="submit" dark @click.native="validateForm()">Register</v-btn>
  		</v-layout>
	</v-form>
</div>
</template>

<script>
var csrfToken = document.getElementById('csrf-token').value;
var errorMessageElement = document.getElementById('error-message');
var errorMessage = (errorMessageElement != null ? errorMessageElement.value : '');
var formattedErrorMessage = (errorMessage !== '' ? errorMessage.charAt(0).toUpperCase() + errorMessage.slice(1) + '.' : '');

export default {
    name: 'RegisterForm',
    components: {},
    data: () => ({
        username: '',
        password: '',
        confirmationPassword: '',
		csrfToken: csrfToken,
		formattedErrorMessage: formattedErrorMessage,
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
		validateForm() {
			return this.$refs.form.validate() && this.password == this.confirmationPassword;
		},
    },
}
</script>
