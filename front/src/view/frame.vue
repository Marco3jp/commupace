<template>
    <div id="commupace">
        <div id="main">
            <router-view></router-view>
        </div>
        <navigation-bar id="navigation-bar"></navigation-bar>
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
    #commupace {
        height: 100%;
        display: flex;
        flex-direction: column;

        #main {
            width: 100vw;
            height: 90%;
        }

        #navigation-bar {
            height: $navigation-bar-height;
        }
    }
</style>