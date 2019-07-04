<template>
    <div id="post-tools">
        <text-box id="post-tools-text-box" @ctrl-enter="post"></text-box>
        <send-button id="post-tools-send-button" @click="post"></send-button>
    </div>
</template>

<script lang="ts">
    import TextBox from "./textBox";
    import SendButton from "./sendButton";
    import {postChat} from "@test/postChat";

    export default {
        name: "postTools",
        components: {SendButton, TextBox},
        methods: {
            post() {
                const newPostMessage = this.$store.getters.getPostText(this.$route.params.communityId);
                if (newPostMessage !== undefined && newPostMessage !== "") {
                    const newPost = postChat(newPostMessage);
                    this.$emit("post", newPost); // for test. if connect server, reload automatically.
                    this.$store.commit("resetPostText", this.$route.params.communityId);
                }
            }
        }
    }
</script>

<style scoped lang="scss">
    #post-tools {
        display: flex;

            width: 70%;
        #post-tools-text-box {
        }

        #post-tools-send-button {
            margin-left: 5%;
            width: 30%;
            font-size: 250%;
        }
    }
</style>