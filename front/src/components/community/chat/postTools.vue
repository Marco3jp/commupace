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
                    const newPost = {
                        "communityAccountId": parseInt(this.$store.getters.communityAccountId),
                        "threadId": 0,
                        "communityId": parseInt(this.$route.params.communityId),
                        "postText": newPostMessage,
                        "postPath": "/",
                    };
                    this.$emit("post", newPost);
                    this.$store.commit("resetPostText", this.$route.params.communityId);
                }
            }
        }
    }
</script>

<style scoped lang="scss">
    #post-tools {
        display: flex;

        #post-tools-text-box {
            width: 75%;
            font-size: 18px;
            padding: 5px;
            resize: none;
        }

        #post-tools-send-button {
            margin-left: 5%;
            width: 20%;
            font-size: 18px;
            padding: 5px;
        }
    }
</style>