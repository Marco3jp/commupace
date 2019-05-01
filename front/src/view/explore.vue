<template>
    <div>
        <div v-if="this.isErrorMessage" class="error-message" id="location-error"><p>{{errorMessage}}</p></div>
        <embed-map-component id="embed-map" v-if="completedPositioning" :position="location"></embed-map-component>
    </div>
</template>

<script lang="ts">
    import EmbedMapComponent from "@components/embedMapComponent.vue";

    export default {
        name: "explore",
        components: {EmbedMapComponent},
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
    #location-error {
        width: 90%;
        padding: 10px 5%;
    }
    #embed-map{
        height: $explore-map-height;
    }
</style>