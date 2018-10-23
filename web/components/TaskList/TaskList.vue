<template>
<draggable :list="this.tasks" :options="dragOptions" :component-data="getComponentData()">
  <task class="task" v-for="task in this.tasks" :task="task">
  </task>
</draggable>
</template>

<script>
import draggable from 'vuedraggable'
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
        getComponentData() {
            return {
                on: {
                    update(event, argument) {
                        console.log(`Moved from ${event.oldIndex} to ${event.newIndex}`);
                    },
                }
            }
        },
    },
    props: ['tasks'],
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
