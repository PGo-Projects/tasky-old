<template>
<v-card dark v-if="task.action == DISPLAY_ACTION">
    <v-container fluid>
      <v-layout row>
        <v-flex column xs1>
          <v-checkbox></v-checkbox>
        </v-flex>

        <v-flex xs10>
          <v-card-title>
            <div>
              <div class="headline">{{ task.title }}</div>
              <span class="gray--text">{{ task.time }}</span>
            </div>
          </v-card-title>
          
          <v-card-text>{{ task.description }}</v-card-text>
          <input type="hidden" :value="task.index" />
        </v-flex>
        
        <v-flex column xs1 align-content-space-around>
            <div class="mt-2 mb-5">
              <v-icon color="red" @click="remove();">clear</v-icon>
            </div>
            <div class="mt-5">
              <v-icon @click="edit()">edit</v-icon>
            </div>
        </v-flex>
      </v-layout>
    </v-container>
</v-card>

<v-card dark v-else-if="task.action == SAVE_ACTION || task.action == CREATE_ACTION">
  <v-container fluid>
    <v-form ref="form" method="post">
      <v-text-field
        :error-messages="emptyField(task.title)"
        label="Title"
        name="title"
        v-model="task.title"
        required
        autofocus
        ></v-text-field>
      <v-text-field
        :error-messages="emptyField(task.time)"
        label="Time"
        name="time"
        v-model="task.time"
        required
        ></v-text-field>
      <v-textarea
        :error-messages="emptyField(task.description)"
        label="Description"
        name="description"
        v-model="task.description"
        auto-grow
        required
        ></v-textarea>
      <input type="hidden" :value="task.index" />

      <v-layout row>
        <v-btn flat dark @click.stop="createOrSave();">{{ task.action }}</v-btn>
        <v-spacer></v-spacer>
        <v-btn flat dark @click.stop="cancel();">Cancel</v-btn>
      </v-layout>
    </v-form>
  </v-container>
</v-card>

<div v-else>
</div>
</template>

<script>
export const CREATE_ACTION = 'create';
export const DELETE_ACTION = 'delete';
export const DISPLAY_ACTION = 'display';
export const SAVE_ACTION = 'save';

export default {
    name: 'Task',
    components: {},
    data: () => ({
        CREATE_ACTION: CREATE_ACTION,
        DELETE_ACTION: DELETE_ACTION,
        DISPLAY_ACTION: DISPLAY_ACTION,
        SAVE_ACTION: SAVE_ACTION,
    }),
    methods: {
        cancel() {
            if (this.task.action == CREATE_ACTION) {
                this.task.action = DELETE_ACTION;
            } else if (this.task.action == SAVE_ACTION) {
                this.task.action = DISPLAY_ACTION;
            }
        },
        createOrSave() {
            if (this.task.title !== '' && this.task.time !== '' && this.task.description !== '') {
                this.task.action = DISPLAY_ACTION;
            } else {
                this.task.firstTime = false;
                return this.$refs.form.validate();
            }

            if (!this.task.index >= 0) {
                if (this.missingTaskIndices.length > 0) {
                    this.task.index = this.missingTaskIndices.shift();
                } else {
                    this.task.index = this.totalNumOfTasks - 1;
                    this.totalNumOfTasks++;
                }
            }
        },
        edit() {
            // This would make the card a form the user can edit and then *save*
            this.task.action = SAVE_ACTION;
        },
        emptyField(value) {
            return (this.task.firstTime || value !== '') ? '' : 'Field cannot be empty';
        },
        remove() {
            this.sendTask('/tasky/remove_task');
            this.task.action = DELETE_ACTION;
        },
    },
    props: ['csrfToken', 'task', 'username'],
}
</script>
