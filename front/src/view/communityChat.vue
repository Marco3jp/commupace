<template>
    <div id="community-chat">
        <post-tools id="post-tools" @post="post"></post-tools>
        <speech-frame id="speech-frame" :postData="postData"></speech-frame>
    </div>
</template>

<script lang="ts">
    import SpeechFrame from "@components/community/chat/speechFrame.vue";
    import PostTools from "@components/community/chat/postTools.vue";
    import {chatPostListMock} from "../test/mock/chatPostList";

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

        #post-tools {
            width: 100%;
            height: 10%;
        }

        #speech-frame {
            margin-top: 5%;
            width: 100%;
            height: 90%;
        }
    }
</style>