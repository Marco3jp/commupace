<template>
    <router-view></router-view>
</template>

<script lang="ts">
    export default {
        name: "community",
        computed: {
            communityId: function (): undefined | string {
                if (Object.keys(this.$route.params).length === 0) return undefined;
                return this.$route.params.communityId
            }
        },
        methods: {
            recordLastVisitCommunityId(communityId) {
                localStorage.setItem("lastVisitCommunityId", communityId);
            },
            restoreLastVisitCommunityId(): undefined | string {
                let communityId = localStorage.getItem("lastVisitCommunityId");
                if (communityId === null) {
                    return undefined;
                } else {
                    return communityId;
                }
            },
            moveLastVisitCommunity(communityId) {
                this.$router.push({
                    name: 'communityFront',
                    params: {communityId: communityId},
                });
            }
        },
        created(): void {
            if (this.communityId === undefined) {
                let lastVisitCommunityId = this.restoreLastVisitCommunityId();
                if (lastVisitCommunityId !== undefined) {
                    this.moveLastVisitCommunity(lastVisitCommunityId);
                }
            } else {
                this.recordLastVisitCommunityId(this.communityId);
            }
        },
        watch: {
            communityId: function (newVal, oldVal) {
                if (this.communityId !== undefined) {
                    this.recordLastVisitCommunityId(newVal);
                } else {
                    this.moveLastVisitCommunity(oldVal);
                }
            }
        },

    }
</script>

<style scoped>

</style>