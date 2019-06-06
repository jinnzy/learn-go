
const state = {
    count: 1
};
const mutations={
    add(state) {
        state.count++
    },
    reduce(state) {
        state.count--
    }
};

const actions = {
    add: ({commit}) => {
        commit('add')
    },
    reduce: ({commit}) => {
        commit('reduce')
    }
};

// 导出
export default {
    namespaced: true,
    state,
    mutations,
    actions
}