<template>
<v-app id="inspire" dark>
  <v-navigation-drawer clipped fixed v-model="drawer" app>
    <v-list>
      <v-list-tile @click="newTask();">
        <v-list-tile-action>
          <v-icon>add</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>New</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
    <v-divider></v-divider>
	<v-list dense>
      <v-list-tile @click="">
      	<v-list-tile-action>
          <v-icon>fa fa-star</v-icon>
      	</v-list-tile-action>
      	<v-list-tile-content>
          <v-list-tile-title>Today</v-list-tile-title>
      	</v-list-tile-content>
      </v-list-tile>
      
      <v-list-tile @click="">
      	<v-list-tile-action>
          <v-icon>fa fa-calendar-alt</v-icon>
      	</v-list-tile-action>
      	<v-list-tile-content>
          <v-list-tile-title>Upcoming</v-list-tile-title>
      	</v-list-tile-content>
      </v-list-tile>
      
   	  <v-list-tile @click="">
      	<v-list-tile-action>
          <v-icon>fa fa-cloud</v-icon>
      	</v-list-tile-action>
      	<v-list-tile-content>
          <v-list-tile-title>Thought Cloud</v-list-tile-title>
      	</v-list-tile-content>
      </v-list-tile>
      
      <v-list-tile @click="">
        <v-list-tile-action>
          <v-icon>fa fa-calendar-check</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Done</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
  	</v-list>
  </v-navigation-drawer>
  
  <v-toolbar app fixed clipped-left>
  	<v-toolbar-side-icon class="hidden-md-and-up" @click.stop="drawer = !drawer;"></v-toolbar-side-icon>
  	<v-toolbar-title>Tasky</v-toolbar-title>
    
    <v-spacer></v-spacer>
    
    <v-toolbar-items>
        <v-btn flat href="/auth/logout">
          <v-icon>fa fa-sign-out-alt</v-icon>
          <span class="ml-2">Logout</span>
        </v-btn>
      </v-toolbar-items>
	</v-toolbar>

    <v-content>
        <task-list :tasks="this.tasks" :csrfToken="this.csrfToken" :username="this.username"></task-list>
    </v-content>

    <footer-component></footer-component>
  </v-app>
</template>

<script>
import {CREATE_ACTION as TASK_CREATE_ACTION, DISPLAY_ACTION as TASK_DISPLAY_ACTION} from 'Task/Task.vue';
import TaskList from 'TaskList/TaskList.vue';
import FooterComponent from 'Footer/Footer.vue';

export default {
    name: 'Tasky',
    components: {TaskList, FooterComponent},
    data: () => ({
        drawer: null,
        tasks: [],
    }),
    async mounted() {
        const response = await fetch('/tasky/get_tasks');
        if (response.ok) {
            this.tasks = await response.json();
        } else {
            alert('There was a problem communicating with the server, please try again later.');
            return
        }
        
        if (this.tasks === null) {
            this.tasks = [];
        }
    },
    methods: {
        newTask() {
            var emptyTask = {
                index: -1,
                title: '',
                time: '',
                description: '',
                action: TASK_CREATE_ACTION,
                firstTime: true
            };
            this.tasks.unshift(emptyTask);
        },
    },
    props: ['csrfToken', 'username'],
}
</script>
