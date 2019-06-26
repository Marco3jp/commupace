<template>
    <div>
        <iframe width="100%" height="100%" frameborder="0" style="border:0" id="embed-map"
                :src=mapUrl allowfullscreen></iframe>
    </div>
</template>

<script lang="ts">
    export default {
        name: "embedMapComponent",
        props: ['position'],
        computed: {
            mapUrl: function () {
                return `https://www.google.com/maps/embed/v1/view?key=AIzaSyCtUVwCfCKqz-qz_GOKDramQAz889JbA0Q&center=${this.position.coords.latitude},${this.position.coords.longitude}&zoom=17`
            }
        }
    }
</script>

<style scoped lang="scss">
    #embed-map {
        height: 100%;
    }
</style>