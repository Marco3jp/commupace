<template>
    <div class="speechFrame">
        <speech-balloon v-for="post in postData" :postData="post" :key="post.PostNumber"></speech-balloon>
    </div>
</template>

<script lang="ts">
    import SpeechBalloon from "./speechBalloon";

    export default {
        name: "speechFrame",
        components: {SpeechBalloon},
        props: ['postData'],
    }
</script>

<style scoped>

</style>