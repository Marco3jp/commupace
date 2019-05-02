<template>
    <div id="community-list">
        <community-list-content :community="community" v-for="community in communityList"></community-list-content>
    </div>
</template>

<script lang="ts">
    import CommunityListContent from "@components/community-list-content";
    import {communityListMock} from "@mock/communityList"

    export default {
        name: "community-list",
        components: {CommunityListContent},
        data() {
            return {
                communityList: [],
            }
        },
        created() {
            this.communityList = communityListMock;
        }
    }
</script>

<style scoped lang="scss">
    #community-list {
        overflow: scroll;
    }
</style>