<template>
<draggable :list="this.tasks" :options="dragOptions" @change="change">
  <task class="task" v-for="task in this.tasks"
        :task="task"
        :csrf-token="csrfToken"
        :username="username">
  </task>
</draggable>
</template>

<script>
import draggable from 'vuedraggable';
import qs from 'qs';
import Task from 'Task/Task.vue';

export default {
    name: 'TaskList',
    components: {draggable, Task},
    computed: {
        dragOptions() {
            return {
                animation: 0,
            }
        },
    },
    data: () => ({
    }),
    methods: {
        change(event) {
            const data = {
                index: event.moved.element.index,
                from: event.moved.oldIndex,
                to: event.moved.newIndex,
                csrf_token: this.csrfToken,
                username: this.username,
            };

            fetch('/tasky/update_task_pos', {
                method: 'POST',
                mode: 'cors',
                credentials: 'same-origin',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: qs.stringify(data),
            });
        },
    },
    props: ['csrfToken', 'tasks', 'username'],
}
</script>

<style>
.task {
    cursor: move;
}
</style>
