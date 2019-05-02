<template>
    <div>
        <p>この端末でコミュニティにアクセスした履歴がありませんでした。</p>
        <p>いつも訪れているコミュニティに向かうなら
            <router-link :to="{name:'myPage'}">マイページ</router-link>
            に行きましょう。
        </p>
        <p>新しいコミュニティを探すなら
            <router-link :to="{name:'explore'}">探索ページ</router-link>
            に行きましょう。
        </p>
    </div>
</template>

<script lang="ts">
    export default {
        name: "communityFallBack"
    }
</script>

<style scoped>

</style>