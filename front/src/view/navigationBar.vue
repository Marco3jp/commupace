<template>
    <nav id="navigation-bar">
        <router-link to="myPage" :class="{selecting: isMyPage}" class="navigation-bar-content"
                     id="navigation-bar-users-cog">
            <i class="fas fa-users-cog"></i>
        </router-link>
        <router-link to="community" :class="{selecting: isCommunity}" class="navigation-bar-content"
                     id="navigation-bar-community">
            <i class="fas fa-comments"></i>
        </router-link>
        <router-link to="explore" :class="{selecting: isExplore}" class="navigation-bar-content"
                     id="navigation-bar-explore">
            <i class="fas fa-compass"></i>
        </router-link>
    </nav>
</template>

<script lang="ts">
    export default {
        name: "navigationBar",
        data() {
            return {
                isMyPage: this.currentTab === "myPage",
                isCommunity: this.currentTab === "community",
                isExplore: this.currentTab === "explore",
            }
        },
        computed: {
            currentTab: () => {
                console.log(this.$route.path.split("/")[1]);
                return this.$route.path.split("/")[1];
            },
        }
    }
</script>

<style scoped lang="scss">
    #navigation-bar {
        height: $navigation-bar-height;
        width: 100vw;
        bottom: 0;
        position: fixed;
        display: flex;
        padding-bottom: 1.5vh;
        padding-top: 1.5vh;
        border-top: 1px solid $secondary-light-color;

        .navigation-bar-content {
            color: $secondary-light-color;
            flex-grow: 1;
            font-size: 5vh;
            line-height: 6vh;
            text-align: center;
        }
    }
</style>

<style lang="scss">
    #navigation-bar {
        .router-link-active {
            color: $secondary-dark-color !important;
        }
    }
</style>