import Vue from 'vue';
import Router from 'vue-router';

import Main from '../views/main';
import Login from '../views/login';
import Reg from '../views/reg';
import ArticleAdd from '../views/article/add';
import ArticleList from '../views/article/list';
import AccountList from '../views/account/list';

Vue.use(Router);

const routes = [{
    path: '/main/index',
    meta: { title: "首页" },
    component: Main
}, {
    path: '/login',
    meta: { title: "登录" },
    component: Login
}, {
    path: '/reg',
    meta: { title: "注册" },
    component: Reg
}, {
    path: '/main/article/add',
    meta: { title: "添加文章" },
    component: ArticleAdd
}, {
    path: '/main/article/edit/:id',
    meta: { title: "编辑文章" },
    name: 'ArticleEdit',
    component: ArticleAdd
}, {
    path: '/main/article/:size/:pageno',
    meta: { title: "文章管理" },
    name: 'ArticleList',
    component: ArticleList
}, {
    path: '/main/account/:size/:pageno',
    meta: { title: "用户列表" },
    name: 'AccountList',
    component: AccountList
}];

export default new Router({
    mode: 'history',
    routes: routes,
    scrollBehavior(to, from, savedPosition) {
        return { x: 0, y: 0 };
    }
});