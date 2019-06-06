
const state = {
    money: 10
};
const mutations={
    add(state,parpm) {
        console.log(parpm);
        state.money+=parpm
    },
    reduce(state) {
        state.money--
    }
};

const actions = {
    add: ({commit},parpm) => {
        // parpm 传参，由模板那里传来
        commit('add',parpm)
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