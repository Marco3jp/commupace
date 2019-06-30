<template>
    <div>
        <div>名前：{{postData.postUser.userName}}</div>
        <div>{{postData.postBody.text}}</div>
    </div>
</template>

<script lang="ts">
    export default {
        name: "speechBalloon",
        props: ['postData'],
    }
</script>

<style scoped>

</style>