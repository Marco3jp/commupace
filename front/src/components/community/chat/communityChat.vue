<template>
    <div id="community-chat">
        <speech-frame id="speech-frame" :postData="postData"></speech-frame>
        <post-tools id="post-tools" @post="post"></post-tools>
    </div>
</template>

<script lang="ts">
    import SpeechFrame from "@components/community/chat/speechFrame.vue";
    import PostTools from "@components/community/chat/postTools.vue";
    import {chatPostListMock} from "@mock/chatPostList";

    export default {
        name: "communityChat",
        components: {PostTools, SpeechFrame},
        data: function () {
            return {
                postData: [],
            }
        },
        created(): void {
            this.postData = chatPostListMock;
        },
        methods: {
            post(newPost) {
                this.postData.push(newPost);
            }
        }
    }
</script>

<style scoped lang="scss">
    #community-chat {
        width: 100%;
        height: 100%;
        padding: 5%;

        #speech-frame {
            width: 100%;
            height: calc(100% - #{$chat-post-tools-height * 1.75});
            overflow-y: auto;
            overflow-x: hidden;
            margin-bottom: $chat-post-tools-height * .75;
        }

        #post-tools {
            //margin-top: 5%;
            width: 100%;
            height: $chat-post-tools-height;
        }
    }
</style>