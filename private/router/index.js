import Vue from 'vue';
import Router from 'vue-router';

import Main from '../views/main';
import Login from '../views/login';
import Reg from '../views/reg';
import ArticleAdd from '../views/article/add';
import ArticleList from '../views/article/list';
import AccountList from '../views/account/list';
import FunctionAdd from '../views/function/add';
import FunctionList from '../views/function/list';
import RoleAdd from '../views/role/add';
import RoleList from '../views/role/list';
import Mapping from '../views/role/mapping';

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
}, {
    path: '/main/function/add',
    meta: { title: "添加权限" },
    name: 'FunctionAdd',
    component: FunctionAdd
}, {
    path: '/main/function/edit/:id',
    meta: { title: "编辑权限" },
    name: 'FunctionEdit',
    component: FunctionAdd
}, {
    path: '/main/function/:size/:pageno',
    meta: { title: "权限管理" },
    name: 'FunctionList',
    component: FunctionList
}, {
    path: '/main/role/add',
    meta: { title: "添加角色" },
    name: 'RoleAdd',
    component: RoleAdd
}, {
    path: '/main/role/edit/:id',
    meta: { title: "编辑角色" },
    name: 'RoleEdit',
    component: RoleAdd
}, {
    path: '/main/role/:size/:pageno',
    meta: { title: "角色管理" },
    name: 'RoleList',
    component: RoleList
}, {
    path: '/main/mapping/:id',
    meta: { title: "权限配置" },
    name: 'Mapping',
    component: Mapping
}];

export default new Router({
    mode: 'history',
    routes: routes,
    scrollBehavior(to, from, savedPosition) {
        return { x: 0, y: 0 };
    }
});