<template>
    <div class="a">page a {{ $store.state.money.money}}
    <!--store是main.js里导入的，第一个money是模块，第二个是数值-->
    <button @click="add(2)">增加</button>
    <button @click="reduce">删减</button>
    </div>
</template>

<script>
    import { mapActions } from 'vuex'
    export default {
        name: "a",
        methods: mapActions('money',['add','reduce'])
        // 导入money中的add和reduce函数
    }
</script>

<style scoped>

</style>