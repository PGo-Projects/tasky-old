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
            console.log(`Moved from ${event.moved.oldIndex} to ${event.moved.newIndex}`);
            console.log(event.moved.element);
        },
    },
    props: ['csrfToken', 'tasks', 'username'],
    watch: {
        tasks(newValue, oldValue) {
            console.log(newValue);
        },
    },
}
</script>

<style>
.task {
    cursor: move;
}
</style>
