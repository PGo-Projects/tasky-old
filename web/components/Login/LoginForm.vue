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
	<v-form method="post">
    	<v-text-field
    		v-model="username"
	    	prepend-icon="account_circle"
    		label="Username"
			name="username"
    		required
    		autofocus
    		></v-text-field>
		<v-text-field
    		v-model="password"
    		:append-icon="showPassword ? 'visibility_off' : 'visibility'"
    		:type="showPassword ? 'text' : 'password'"
    		prepend-icon="lock"
    		label="Password"
			name="password"
    		required
    		@click:append="showPassword = !showPassword"
    		></v-text-field>
 	 
		<input type="hidden" name="csrf_token" :value="this.csrfToken" />
 	 
		<v-layout row>
    		<v-btn flat small dark href="/auth/register">Create Account</v-btn>
    		<v-spacer></v-spacer>
    		<v-btn type="submit" dark>Login</v-btn>
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
    name: 'LoginForm',
    components: {},
    data: () => ({
        username: "",
        password: "",
        csrfToken: csrfToken,
        formattedErrorMessage: formattedErrorMessage,
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
