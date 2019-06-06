<template>
    <div class="event">
        <button type="button" name="button" @click="msg">test</button>
    </div>
</template>

<script>
    export default {
        name: "events",
        methods: {
            msg:function () {
                console.log(Math.random())
            }
        }
    }
</script>

<style scoped>

</style>