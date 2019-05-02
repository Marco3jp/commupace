<template>
    <router-link :to="location" id="community-list-content">
        <div id="icon">
            <img id="icon-img" v-if="!isFontIcon" :src="this.community.icon" :alt=alt>
            <i v-if="isFontIcon" id="icon-i" class="fas" :class="this.community.icon"></i>
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
        height: $explore-community-list-content-height;
        padding-top: $explore-community-list-content-vertical-padding;
        padding-bottom: $explore-community-list-content-vertical-padding;
        border-bottom: 1px solid $secondary-light-color;

        text-decoration: none;
        font-size: $explore-community-list-inner-content-height * .75;
        line-height: $explore-community-list-inner-content-height;

        #icon {
            height: $explore-community-list-inner-content-height;
            width: $explore-community-list-inner-content-height;
            margin-left: 15px;

            #icon-img {
                width: 100%;
                height: 100%;
            }

            #icon-i {
                width: 100%;
                height: 100%;
                font-size: $explore-community-list-inner-content-height;
            }
        }

        #community-name {
            margin-left: 30px;
        }
    }
</style>