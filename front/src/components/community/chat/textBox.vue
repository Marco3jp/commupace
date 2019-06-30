<template>
    <textarea v-model="postText" @keydown.ctrl.enter="$emit('ctrl-enter')"></textarea>
</template>

<script lang="ts">
    export default {
        name: "textBox",
        props: ['postText'],
        computed: {
            postText: {
                get() {
                    return this.$store.getters.getPostText(this.$route.params.communityId);
                },
                set(value) {
                    this.$store.commit("savePostText", {
                        chatId: this.$route.params.communityId,
                        text: value,
                    });
                }
            }
        }
    }
</script>

<style scoped>

</style>