import {Module} from "vuex";
import {MainState} from "../main";

export interface tokenPairs {
    accessToken: string,
    refreshToken: string,
}

export interface authDataInterface {
    tokens: tokenPairs,
    managerAccountId: string,
    communityAccountId: number,
}

const state: authDataInterface = {
    tokens: {
        accessToken: "",
        refreshToken: "",
    },
    managerAccountId: "",
    communityAccountId: 0
};

export const auth: Module<authDataInterface, MainState> = {
    state: state,
    getters: {
        accessToken: (state) => {
            return state.tokens.accessToken
        },
        refreshToken: (state) => {
            return state.tokens.refreshToken
        },
        tokenPairs: (state) => {
            return state.tokens
        },
        managerAccountId: (state) => {
            return state.managerAccountId
        },
        hasManagerAccount: (state) => {
            return state.tokens.accessToken !== "" && state.tokens.refreshToken !== "" && state.managerAccountId !== ""
        },
        communityAccountId: (state) => {
            return state.communityAccountId
        }
    },
    mutations: {
        saveTokenPairs(state, pairs: tokenPairs) {
            state.tokens = pairs;
            localStorage.setItem("access-token", pairs.accessToken);
            localStorage.setItem("refresh-token", pairs.refreshToken);
        },
        saveManagerAccountId(state, managerAccountId) {
            state.managerAccountId = managerAccountId;
            localStorage.setItem("manager-account-id", managerAccountId);
        },
        saveCommunityAccountId(state, communityAccountId: number) {
            state.communityAccountId = communityAccountId;
            localStorage.setItem("community-account-id", communityAccountId.toString());
        }
    },
};