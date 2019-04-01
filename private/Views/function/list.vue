<template>
    <div id="page-container" style="position: relative;">
        <Menu tagIndex="10"></Menu>
        <div class="rightMain">
            <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
                <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
            </div>
            <div id="list">
                <div class="search form-horizontal" style="padding:10px 20px 0 10px;overflow:auto;">
                    <button class="hd-button" v-on:click="add">添加权限</button>
                </div>
                <div class="row">
                    <div class="col-md-12">
                        <!-- BEGIN EXAMPLE TABLE PORTLET-->
                        <div class="portlet light">
                            <div class="portlet-body">
                                <table class="table table-striped table-bordered table-hover" id="mytable">
                                    <thead>
                                        <tr>
                                            <th width="150">序号</th>
                                            <th width="*">权限</th>
                                            <th width="*">分组</th>
                                            <th width="*">控制器</th>
                                            <th width="150">创建时间</th>
                                            <th width="180">操作</th>
                                        </tr>
                                    </thead>
                                    <tbody id="resultTable">
                                        <tr v-for="data in datas" :key="data.ID">
                                            <td>{{data.ID}}</td>
                                            <td>{{data.Funname}}</td>
                                            <td>{{data.FunType}}</td>
                                            <td>{{data.Controller}}</td>
                                            <td>{{data.CreateTime | formatDate}}</td>
                                            <td align="center">
                                                <a v-on:click="edit(data.ID)" class="btn btn-sm btn-outline filter-submit blue">
                                                    <i class="fa fa-edit"></i> 修  改
                                                </a>
                                                <a v-on:click="remove(data.ID)" class="btn btn-sm btn-outline filter-submit dark" style="margin-right:0;">
                                                    <i class="fa fa-lock"></i> 删  除
                                                </a>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                            <Pager :total="total" :current='current' :display='display' @pagechange="pagechange"></Pager>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import Menu from '../components/menu';
    import Pager from '../components/pager';
    import httper from '../../util/httper';
    import { formatDate } from '../../util/date.js';
    import router from '../../router';

    export default {
        data: function () {
            return {
                datas: [],
                total: 5,
                display: 15,
                current: 1
            };
        },
        components: {
            Menu,
            Pager
        },
        created: function () {
            var self = this;
            self.current = parseInt(self.$route.params.pageno);
            self.display = parseInt(self.$route.params.size);
            self.load();
        },
        filters: {
            formatDate(time) {
                var date = new Date(time);
                return formatDate(date, "yyyy-MM-dd hh:mm:ss");
            }
        },
        methods: {
            pagechange: function (currentPage) {
                var self = this;
                self.current = currentPage;
                router.push({ name: 'function', params: { size: self.display, pageno: self.current } });
                self.load();
            },
            load: function () {
                var self = this;
                httper.post('/function/list/' + self.display + '/' + self.current).then(function (response) {
                    self.datas = response.data.Items;
                    self.total = response.data.PageArgs.TotalCount;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            add: function () {
                router.push({ name: 'FunctionAdd' });
            },
            edit: function (id) {
                router.push({ name: 'FunctionEdit', params: {id: id} });
            }
        }
    };
</script>