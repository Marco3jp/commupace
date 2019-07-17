<template>
    <div id="speech-balloon">
        <div id="speech-balloon-meta">
            <div id="speech-user-name">{{postData.CommunityAccount.DisplayName}}</div>
            <!--div id="speech-timestamp">
                <span v-if="!isToday" class="speech-timestamp-date">{{date}}</span>
                <span class="speech-timestamp-time">{{time}}</span>
            </div-->
        </div>
        <div id="speech-body">
            <div id="speech-body-text">{{postData.PostText}}</div>
        </div>
    </div>
</template>

<script lang="ts">
    export default {
        name: "speechBalloon",
        props: ['postData'],
        data: function () {
            return {
                postTime: {
                    year: 0,
                    month: 0,
                    day: 0,
                    hour: "0",
                    minute: "0",
                }
            }
        },
        computed: {
            isToday: function () {
                const now = new Date();
                return this.postTime.year === now.getFullYear()
                    && this.postTime.month === now.getMonth()
                    && this.postTime.day === now.getDay();
            },
            date: function () {
                return `${this.postTime.month}/${this.postTime.day}`
            },
            time: function () {
                return `${this.postTime.hour.substring(this.postTime.hour.length - 2)}:${this.postTime.minute.substring(this.postTime.minute.length - 2)}`
            }
        },
        created(): void {
            const date = new Date(parseInt(this.postData.postTime));
            this.postTime.year = date.getFullYear();
            this.postTime.month = date.getMonth();
            this.postTime.day = date.getDay();
            this.postTime.hour += date.getHours().toString();
            this.postTime.minute += date.getMinutes().toString();
        }
    }
</script>

<style scoped lang="scss">
    #speech-balloon {
        &:not(:first-of-type) {
            margin-top: 10px;
        }

        #speech-balloon-meta {
            display: flex;
            margin-bottom: 3px;

            #speech-user-name {
                color: $secondary-dark-color;
                flex-grow: 1;
            }

            #speech-timestamp {
                color: rgba($secondary-dark-color, .75);
            }
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