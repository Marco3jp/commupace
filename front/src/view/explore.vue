<template>
    <div id="explore">
        <div v-if="this.isErrorMessage" class="error-message" id="location-error"><p>{{errorMessage}}</p></div>
        <embed-map-component id="embed-map" v-if="completedPositioning" :position="location"></embed-map-component>
        <community-list id="community-list" v-if="completedPositioning" :position="location"></community-list>
    </div>
</template>

<script lang="ts">
    import EmbedMapComponent from "@components/explore/embedMapComponent.vue";
    import CommunityList from "@components/explore/communityList.vue";

    export default {
        name: "explore",
        components: {CommunityList, EmbedMapComponent},
        data() {
            return {
                location: {},
                completedPositioning: false,
                errorMessage: ""
            };
        },
        computed: {
            isErrorMessage: function () {
                return this.errorMessage !== "";
            },
        },
        created(): void {
            navigator.geolocation.getCurrentPosition(result => {
                this.location = result;
                this.completedPositioning = true;
            }, error => {
                if (error.code === 1) {
                    this.errorMessage = "権限が足りていないため場所を確認することができません。本サイトに位置情報の利用を許可することで解消されることがあります。"
                } else if (error.code === 2) {
                    this.errorMessage = "エラーが発生しました。ブラウザやスマートフォンを再起動することで解消されることがあります。"
                } else if (error.code === 3) {
                    this.errorMessage = "うまく測位できません。地下などのGPSが働きにくい場所にいる場合、移動することで解消されることがあります。"
                }
            })
        }
    }
</script>

<style scoped lang="scss">
    #explore {
        height: 100%;
    }

    #location-error {
        width: 90%;
        padding: 10px 5%;
        margin: auto;
    }

    #embed-map {
        height: $explore-map-height;
        width: 100%;
    }

    #community-list {
        height: $explore-community-list-height;
        width: 100%
    }
</style>