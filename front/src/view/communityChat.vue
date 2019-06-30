<template>
    <div>
        <p>ここはコミュニティのチャットを表示する場所です。</p>
        <post-tools @post="post"></post-tools>
        <speech-frame :postData="postData"></speech-frame>
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

<style scoped>

</style>