<template>
<div>
	<v-navigation-drawer clipped fixed v-model="drawer" app>
		<v-list dense>
    		<v-list-tile @click="">
      			<v-list-tile-action>
        			<v-icon>calendar_today</v-icon>
      			</v-list-tile-action>
      			<v-list-tile-content>
        			<v-list-tile-title>Today</v-list-tile-title>
      			</v-list-tile-content>
    		</v-list-tile>

    		<v-list-tile @click="">
      			<v-list-tile-action>
        			<v-icon>inbox</v-icon>
      			</v-list-tile-action>
      			<v-list-tile-content>
        			<v-list-tile-title>Inbox</v-list-tile-title>
      			</v-list-tile-content>
    		</v-list-tile>
    
   	 		<v-list-tile @click="">
      			<v-list-tile-action>
        			<v-icon>cloud</v-icon>
      			</v-list-tile-action>
      			<v-list-tile-content>
        			<v-list-tile-title>Thought Cloud</v-list-tile-title>
      			</v-list-tile-content>
    		</v-list-tile>
  		</v-list>
	</v-navigation-drawer>

	<v-toolbar app fixed clipped-left>
  		<v-toolbar-side-icon @click.stop="drawer = !drawer;"></v-toolbar-side-icon>
  		<v-toolbar-title>Tasky</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-toolbar-items>
          <v-btn flat href="/auth/logout">
            <v-icon>fa fa-sign-out-alt</v-icon>
            <span class="ml-2">Logout</span>
          </v-btn>
        </v-toolbar-items>
	</v-toolbar>
</div>
</template>

<script>
export default {
    name: 'AuthenticatedNav',
    components: {},
    data: () => ({
        drawer: true,
    }),
    methods: {
    },
    props: [],
}
</script>
