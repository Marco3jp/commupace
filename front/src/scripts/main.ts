import "@scss/main.scss";
import Vue from "vue";
import Vuex from "vuex";
import Main from "../view/frame.vue";
import router from "../routes/main";
import store from "../store/main";

Vue.use(Vuex);

const accessToken = localStorage.getItem("access-token");
const refreshToken = localStorage.getItem("refresh-token");
const managerAccountId = localStorage.getItem("manager-account-id");
const communityAccountId = localStorage.getItem("community-account-id");

if (accessToken !== null || refreshToken !== null || managerAccountId !== null || communityAccountId !== null) {
    store.commit("saveTokenPairs", {accessToken: accessToken, refreshToken: refreshToken});
    store.commit("saveManagerAccountId", managerAccountId);
    store.commit("saveCommunityAccountId", communityAccountId);
}

// main vue の初期化
new Vue({
    router,
    store,
    render: h => h(Main)
}).$mount("#app");