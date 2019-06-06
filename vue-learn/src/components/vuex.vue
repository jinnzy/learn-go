<template>
    <div class="vuex">
        vuex {{$store.state.count}}
        <!--store 就是main中引入的-->
        <button type="button" name="button" @click="increment">增加</button>
        <button type="button" name="button" @click="decrement">删减</button>
    </div>
</template>

<script>
    import { mapActions } from 'vuex'
    // mapActions进行关联操作的
    export default {
        name: "vuex",
        methods: mapActions([
            // 关联vuex定义的操作
            'increment',
            'decrement'
        ])
    }
</script>

<style scoped>

</style>