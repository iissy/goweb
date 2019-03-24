<template>
    <div style="position: relative;">
        <Menu tagIndex="8"></Menu>
        <div class="rightMain">
            <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
                <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
            </div>
            <div id="list">
                <div class="row">
                    <div class="col-md-12">
                        <!-- BEGIN EXAMPLE TABLE PORTLET-->
                        <div class="portlet light">
                            <div class="portlet-body">
                                <table class="table table-striped table-bordered table-hover" id="mytable">
                                    <thead>
                                        <tr>
                                            <th width="100">序号</th>
                                            <th width="200">登陆号</th>
                                            <th width="*">名字</th>
                                            <th width="100">状态</th>
                                            <th width="150">注册时间</th>
                                            <th width="150">最后登录</th>
                                            <th width="150">操作</th>
                                        </tr>
                                    </thead>
                                    <tbody id="resultTable">
                                        <tr v-for="data in datas" :key="data.ID">
                                            <td align="center">{{data.ID}}</td>
                                            <td align="center">{{data.UserID}}</td>
                                            <td><div style="max-width:785px;overflow:hidden;white-space:nowrap">{{data.UserName}}</div></td>
                                            <td align="center">{{data.Status}}</td>
                                            <td align="center">{{data.RegDate | formatDate}}</td>
                                            <td>{{data.LastLoginDate | formatDate}}</td>
                                            <td align="center">
                                                <a v-on:click="editlink('LinkEdit', {id: data.ID})" class="btn btn-sm btn-outline filter-submit purple">
                                                    <i class="fa fa-edit"></i> 修改
                                                </a>
                                                <a v-on:click="dellink(data.ID)" class="btn btn-sm btn-outline filter-submit dark" style="margin-right:0;">
                                                    <i class="fa fa-lock"></i> 删除
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
    import { formatDate } from '../../util/date';
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
                router.push({ name: 'AccountList', params: { size: self.display, pageno: self.current } });
                self.load();
            },
            load: function () {
                var self = this;
                httper.post('/account/list/' + self.display + '/' + self.current, {
                }).then(function (response) {
                    self.datas = response.data.Items;
                    self.total = response.data.PageArgs.TotalCount;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            editlink: function (name, params) {
                router.push({ name: name, params: params });
            },
            dellink: function (id) {
                var self = this;
                if (confirm("确认要删除？")) {
                    httper.get('/account/delete/' + id).then(function (response) {
                        if (response.data.result === "1") {
                            self.load();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            },
            search: function () {
                var self = this;
                self.pagechange(1);
            }
        }
    };
</script>