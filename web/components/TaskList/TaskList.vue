<template>
  <div>
    <task v-for="task in this.tasks" :task="task">
    </task>
  </div>
</template>

<script>
import Task from 'Task/Task.vue';

export default {
    name: 'TaskList',
    components: {Task},
    data: () => ({
    }),
    methods: {
    },
    props: ['tasks'],
    watch: {
        tasks(newValue, oldValue) {
            console.log(newValue);
        },
    },
}
</script>
