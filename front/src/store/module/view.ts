import {Module} from "vuex";
import {MainState} from "../main";

export interface viewInterface {
    loadingCount: number;
}

const state: viewInterface = {
    loadingCount: 0,
};

export const views: Module<viewInterface, MainState> = {
    state: state,
    getters: {
        isLoading: state => state.loadingCount > 0,
    },
    mutations: {
        incrementLoadingCount(state) {
            ++state.loadingCount;
        },
        decrementLoadingCount(state) {
            if (state.loadingCount > 0) {
                --state.loadingCount;
            }
        },
    },
};