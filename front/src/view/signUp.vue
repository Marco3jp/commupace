<template>
    <div>
        <div @click="signUp">アカウントを作る</div>
        <h1>本サービスのルール</h1>
        <p>データが消えても泣かない！</p>
    </div>
</template>

<script lang="ts">
    import {tokenPairs} from "@store/module/authData";

    export default {
        name: "signUp",
        data: () => {
            return {
                signUpFlag: false
            }
        },
        methods: {
            signUp() {
                if (!this.signUpFlag) {
                    this.signUpFlag = true;
                    fetch("/api/v1/auth/sign_up").then((response) => {
                        if (response.ok && response.status === 200) {
                            return response.json()
                        }
                    }).then((result) => {
                        let tokenPairs: tokenPairs = {
                            accessToken: result.accessToken,
                            refreshToken: result.refreshToken,
                        };
                        this.$store.commit("saveTokenPairs", tokenPairs);
                        this.$store.commit("saveManagerAccountId", result.managerAccountId);
                        this.$router.push({name: "community"});
                    }).catch((e) => {
                        // console.log(e);
                    })
                }
            }
        }
    }
</script>

<style scoped>

</style>