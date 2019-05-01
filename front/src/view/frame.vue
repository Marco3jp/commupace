<template>
    <div>
        <div id="main">
            <router-view></router-view>
        </div>
        <navigation-bar></navigation-bar>
        <!--loading-overlay></loading-overlay-->
    </div>
</template>

<script lang="ts">
    import navigationBar from "@view/navigationBar"

    export default {
        name: "frame",
        components: {navigationBar}
    }
</script>

<style scoped lang="scss">
    #main {
        width: 100vw;
        height: calc(100vh - #{$navigation-bar-height});
    }
</style>