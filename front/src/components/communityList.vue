<template>
    <div id="community-list">
        <community-list-content :community="community" v-for="(community, index) in communityList"
                                :key="index"></community-list-content>
    </div>
</template>

<script lang="ts">
    import CommunityListContent from "@components/communityListContent";
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