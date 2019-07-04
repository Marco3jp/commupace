<template>
    <div id="speech-balloon">
        <div id="speech-user-name">{{postData.postUser.userName}}</div>
        <div id="speech-body">
            <div id="speech-body-text">{{postData.postBody.text}}</div>
        </div>
    </div>
</template>

<script lang="ts">
    export default {
        name: "speechBalloon",
        props: ['postData'],
    }
</script>

<style scoped lang="scss">
    #speech-balloon {
        &:not(:first-of-type) {
            margin-top: 10px;
        }

        #speech-user-name {
            color: $secondary-dark-color;
        }

        #speech-body {
            padding: 10px 5px;
            border-radius: 5px;
            background: rgba($primary-color, .5);

            #speech-body-text {
                overflow-wrap: break-word;
            }
        }
    }
</style>