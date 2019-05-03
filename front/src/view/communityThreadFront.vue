<template>
    <div>
        <p>ここはコミュニティ掲示板のフロントページです。</p>
        <p>掲示板の一覧を表示する予定です。</p>
        <router-link :to="{name: 'communityThread', params: {threadId: '6b8ba8ee-ec8c-4dec-ba0a-a6cc3fb209fc'}}">
            テストスレッド
        </router-link>
    </div>
</template>

<script lang="ts">
    export default {
        name: "communityThreadFront"
    }
</script>

<style scoped>

</style>