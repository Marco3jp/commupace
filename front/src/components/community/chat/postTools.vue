<template>
    <div>
        <text-box @ctrl-enter="post"></text-box>
        <send-button @click="post"></send-button>
    </div>
</template>

<script lang="ts">
    import TextBox from "./textBox";
    import SendButton from "./sendButton";
    import {postChat} from "@mock/postChat";

    export default {
        name: "postTools",
        components: {SendButton, TextBox},
        methods: {
            post() {
                const newPost = postChat(this.$store.getters.getPostText(this.$route.params.communityId));
                this.$emit("post", newPost); // for test. if connect server, reload automatically.
                this.$store.commit("resetPostText", this.$route.params.communityId);
            }
        }
    }
</script>

<style scoped>

</style>