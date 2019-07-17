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
            this.fetchNewPosts();
        },
        methods: {
            post(newPost) {
                fetch(`/api/v1/community/chat/post`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "access-token": this.$store.getters.accessToken,
                        "refresh-token": this.$store.getters.refreshToken,
                        "manager-account-id": this.$store.getters.managerAccountId,
                    },
                    body: JSON.stringify(newPost),
                }).then(response => {
                    return response.json()
                }).then(result => {
                    this.fetchNewPosts();
                })
            },
            fetchNewPosts() {
                fetch(`/api/v1/community/chat/post?communityId=${parseInt(this.$route.params.communityId)}&count=10`, {
                    headers: {
                        "Content-Type": "application/json",
                        "access-token": this.$store.getters.accessToken,
                        "refresh-token": this.$store.getters.refreshToken,
                        "manager-account-id": this.$store.getters.managerAccountId,
                    },
                }).then(response => {
                    return response.json();
                }).then(result => {
                    this.storePosts(result);
                })
            },
            storePosts(result) {
                for (let i = 0; i < result.posts.length; i++) {
                    let dupFlag = false;
                    for (let j = 0; j < this.postData.length; j++) {
                        if (result.posts[i].PostNumber === this.postData[j].PostNumber) {
                            dupFlag = true;
                            break;
                        }
                    }
                    if (!dupFlag) {
                        this.postData.push(result.posts[i]);
                    }
                }
                this.postData.sort((a, b) => {
                    if (a.PostNumber > b.PostNumber) {
                        return 1
                    } else {
                        return -1
                    }
                })
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