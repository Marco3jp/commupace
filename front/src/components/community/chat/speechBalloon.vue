<template>
    <div id="speech-balloon">
        <div id="speech-user-name">名前：{{postData.postUser.userName}}</div>
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
        #speech-body {
            #speech-body-text {
                overflow-wrap: break-word;
            }
        }
    }
</style>