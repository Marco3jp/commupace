<template>
    <button @click="$emit('click')">送信</button>
</template>

<script lang="ts">
    export default {
        name: "sendButton"
    }
</script>

<style scoped>

</style>