<template>
    <div>
        <p>{{communityInfoMessage}}</p>
        <div id="community-content">
            <div id="community-content-chat">
                <router-link :to="{name: 'communityChat'}">チャット</router-link>
            </div>
            <div id="community-content-thread">
                <router-link :to="{name: 'communityThreadFront'}">掲示板</router-link>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import {communityListMock} from "@mock/communityList";

    export default {
        name: "community",
        computed: {
            communityInfoMessage: function () {
                let thisCommunityId = parseInt(this.$route.params.communityId);
                let thisCommunity = communityListMock.filter(community => {
                    return thisCommunityId === community.id;
                });

                return `ここはid${thisCommunity[0].id}の『${thisCommunity[0].name}』コミュニティフロントです。`
            }
        }
    }
</script>

<style scoped>

</style>