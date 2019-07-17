<template>
    <div id="community-list">
        <community-list-content :community="community" v-for="(community, index) in communityList"
                                :key="index"></community-list-content>
    </div>
</template>

<script lang="ts">
    import CommunityListContent from "@components/explore/communityListContent.vue";

    export default {
        name: "community-list",
        components: {CommunityListContent},
        props: ['position'],
        data() {
            return {
                communityList: [],
            }
        },
        created() {
            fetch(`/api/v1/community/search?lat=${this.position.coords.latitude}&long=${this.position.coords.longitude}&zoom=17`, {
                headers: {
                    "Content-Type": "application/json",
                    "access-token": this.$store.getters.accessToken,
                    "refresh-token": this.$store.getters.refreshToken,
                    "manager-account-id": this.$store.getters.managerAccountId,
                },
            }).then((response) => {
                return response.json();
            }).then((result) => {
                this.communityList = result.communityList;
            });
        }
    }
</script>

<style scoped lang="scss">
    #community-list {
        overflow-y: scroll;
    }
</style>