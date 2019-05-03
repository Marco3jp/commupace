<template>
    <div>
        <p>ここはスレッドを表示します</p>
        <p>スレッドIDは："<span>{{threadId}}</span>"です</p>
    </div>
</template>

<script lang="ts">
    export default {
        name: "communityThread",
        computed: {
            threadId: function () {
                return this.$route.params.threadId;
            }
        }
    }
</script>

<style scoped>

</style>