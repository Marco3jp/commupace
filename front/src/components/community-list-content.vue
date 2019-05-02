<template>
    <router-link :to="location" id="community-list-content">
        <div id="icon">
            <img v-if="!isFontIcon" :src="this.community.icon" :alt=alt>
            <i v-if="isFontIcon" class="fas" :class="this.community.icon"></i>
        </div>
        <div id="community-name">
            <span>{{this.community.name}}</span>
        </div>
    </router-link>
</template>

<script lang="ts">
    export default {
        name: "community-list-content",
        props: ['community'],
        data() {
            return {
                iconName: "",
                iconUrl: "",
                location: {
                    name: 'communityFront',
                    params: {
                        communityId: this.community.id,
                    }
                },
                alt: `${this.community.name}のアイコン`
            }
        },
        computed: {
            isFontIcon: function () {
                return this.community.isFontIcon
            }
        },
    }
</script>

<style scoped lang="scss">
    #community-list-content {
        display: flex;
    }
</style>