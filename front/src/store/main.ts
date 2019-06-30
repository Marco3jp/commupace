import Vue from "vue"
import Vuex, {StoreOptions} from "vuex"
import {views} from "@store/module/view";
import {chat} from "@store/module/chat";

export interface MainState {
    loadingCount: number;
}

Vue.use(Vuex);

const store: StoreOptions<MainState> = {
    modules: {
        view: views,
        chat: chat,
    },
};

export default new Vuex.Store<MainState>(store);