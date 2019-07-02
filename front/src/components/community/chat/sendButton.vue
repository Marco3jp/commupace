<template>
    <div @click="$emit('click')" id="send-button" title="Ctrl+Enter">
        <span>送信</span>
    </div>
</template>

<script lang="ts">
    export default {
        name: "sendButton"
    }
</script>

<style scoped lang="scss">
    #send-button {
        border: 1px solid $secondary-light-color;
        text-align: center;
        cursor: pointer;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }
</style>