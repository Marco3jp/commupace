import {Module} from "vuex";
import {MainState} from "../main";

export interface tokenPairs {
    accessToken: string,
    refreshToken: string,
}

export interface authDataInterface {
    tokens: tokenPairs,
    managerAccountId: string,
}

const state: authDataInterface = {
    tokens: {
        accessToken: "",
        refreshToken: "",
    },
    managerAccountId: "",
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
        }
    },
};