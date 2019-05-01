import Vue from "vue"
import Vuex, {StoreOptions} from "vuex"
import {views} from "@store/module/view";

export interface MainState {
    loadingCount: number;
}

Vue.use(Vuex);

const store: StoreOptions<MainState> = {
    modules: {
        view: views,
    },
};

export default new Vuex.Store<MainState>(store);