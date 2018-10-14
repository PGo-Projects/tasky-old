<template>
  <v-app id="inspire" dark>
    <authenticated-nav></authenticated-nav>
  </v-app>
</template>

<script>
import AuthenticatedNav from 'AuthenticatedNav/AuthenticatedNav.vue';

export default {
    name: 'Tasky',
    components: {AuthenticatedNav},
    data: () => ({
    }),
    methods: {
    },
    props: [],
}
</script>
