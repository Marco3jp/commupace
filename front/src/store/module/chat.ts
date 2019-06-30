import {Module} from "vuex";
import {MainState} from "../main";
import Vue from "vue";

export interface chatData {
    chatId: string,
    text: string,
}

export interface chatInterface {
    postSaver: Object
}

const state: chatInterface = {
    postSaver: {}
};

export const chat: Module<chatInterface, MainState> = {
    state: state,
    getters: {
        getPostText: (state) => (chatId) => {
            return state.postSaver[chatId];
        }
    },
    mutations: {
        savePostText(state, chatData: chatData) {
            Vue.set(state.postSaver, chatData.chatId, chatData.text);
        },
        resetPostText(state, chatId) {
            Vue.set(state.postSaver, chatId, "");
        }
    },
};