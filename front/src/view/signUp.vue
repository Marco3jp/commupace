<template>
    <div>
        <label>
            DisplayID:
            <input v-model="communityAccount.displayId" ref="displayId">
        </label>
        <label>
            DisplayName:
            <input v-model="communityAccount.displayName" ref="displayName">
        </label>
        <div @click="signUp">アカウントを作る</div>
        <p>{{errorMessage}}</p>
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
                signUpFlag: false,
                communityAccount: {
                    displayId: "",
                    displayName: "",
                    icon: "no",
                    status: "commupace alpha version"
                },
                errorMessage: ""
            }
        },
        methods: {
            async signUp() {
                if (this.communityAccount.displayId !== "" && this.communityAccount.displayName !== "") {
                    if (!this.signUpFlag) {
                        this.signUpFlag = true;
                        this.$refs.displayId.disabled = true;
                        this.$refs.displayName.disabled = true;
                        await this.createManagerAccount();
                        let result = await this.createCommunityAccount();
                        console.log(result);
                        if (result) {
                            this.$router.push({name: "community"});
                        } else {
                            this.errorMessage = "idが重複しているか、サーバー側でエラーが起こりました。@Marco_utauを叩くとなんか出ます。"
                        }
                    }
                } else {
                    this.errorMessage = "idかNameが空です"
                }
            },
            createManagerAccount() {
                return new Promise(((resolve, reject) => {
                    fetch("/api/v1/auth/sign_up").then((response) => {
                        if (response.ok && response.status === 200) {
                            return response.json()
                        }
                    }).then((result) => {
                        this.$store.commit("saveTokenPairs", {
                            accessToken: result.accessToken,
                            refreshToken: result.refreshToken,
                        });
                        this.$store.commit("saveManagerAccountId", result.managerAccountId);
                        resolve();
                    }).catch((e) => {
                        reject();
                    })
                }))
            },
            createCommunityAccount() {
                return new Promise(((resolve, reject) => {
                    fetch("/api/v1/manager_account/add_community_account", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                            "access-token": this.$store.getters.accessToken,
                            "refresh-token": this.$store.getters.refreshToken,
                            "manager-account-id": this.$store.getters.managerAccountId,
                        },
                        body: JSON.stringify(this.communityAccount)
                    }).then((response) => {
                        return response.json();
                    }).then((result) => {
                        console.log(result.communityAccount.ID);
                        this.$store.commit("saveCommunityAccountId", result.communityAccount.ID);
                        resolve(true);
                    }).catch((e) => {
                        reject(false);
                    })
                }))
            }
        }
    }
</script>

<style scoped>

</style>