import Vue from 'vue';
import Vuex from 'vuex'

Vue.use(Vuex);

const state={
    count: 1
};

const mutations={
    increment(state){
        state.count++
    },
    decrement(state){
        state.count--
    }
};

const actions={
    increment:({commit})=> {// 复制的意思，es6的语法
        // 提交increment，也就是上面的increment
        commit('increment')
    },
    decrement:({commit}) => {
        commit('decrement')
    }
};

export default new Vuex.Store({
    // 作为模块导出
    state,
    mutations,
    actions
})