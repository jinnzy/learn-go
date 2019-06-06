<template>
    <div>
    child components
    {{ namet }} {{ age }}
    <button type="button" name="button" @click="clickS">发送父组件</button>
    </div>
</template>

<script>
    export default {
        name: "com",
        props: ['age'],
        data() {
            return{
                namet: 'com',
            }
        },
        methods: {
            clickS: function () {
                console.log("给父组件传值");
                this.$emit("patch")
            }
        }


    }
</script>

<style scoped>

</style>