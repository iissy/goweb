<template>
    <div style="position: relative;">
        <Menu tagIndex="5"></Menu>
        <div class="rightMain">
            <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
                <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
            </div>
            <div id="list">
                <div class="search form-horizontal" style="padding:10px 20px 0 10px;overflow:auto;">
                    <div class="form-group">
                        <div class="col-md-2">
                            <label class="col-md-3 control-label">栏目：</label>
                            <div class="col-md-9">
                                <select v-model="Catalog" class="inputclass">
                                    <option value=""></option>
                                    <option value="技术">技术</option>
                                    <option value="科技">科技</option>
                                    <option value="新闻">新闻</option>
                                    <option value="故事">故事</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-md-3">
                            <label class="col-md-2 control-label">编号：</label>
                            <div class="col-md-10">
                                <input type="text" v-model="Id" class="inputclass" style="width:300px;" />
                            </div>
                        </div>
                        <div class="col-md-3">
                            <label class="col-md-2 control-label">标题：</label>
                            <div class="col-md-10">
                                <input type="text" v-model="Title" class="inputclass" style="width:300px;" />
                            </div>
                        </div>
                        <div class="col-md-1">
                            <button class="hd-button" v-on:click="search">搜  索</button>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-12">
                        <!-- BEGIN EXAMPLE TABLE PORTLET-->
                        <div class="portlet light">
                            <div class="portlet-body">
                                <table class="table table-striped table-bordered table-hover" id="mytable">
                                    <thead>
                                        <tr>
                                            <th width="400">编号</th>
                                            <th>标题</th>
                                            <th width="100">访问</th>
                                            <th width="150">来源</th>
                                            <th width="150">栏目</th>
                                            <th width="150">时间</th>
                                            <th width="150">操作</th>
                                        </tr>
                                    </thead>
                                    <tbody id="resultTable">
                                        <tr v-for="data in datas" :key="data.ID">
                                            <td style="max-width:340px;overflow:hidden;white-space:nowrap;">{{data.ID}}</td>
                                            <td><a :href="'https://www.hrefs.cn/article/'+data.id" target="_blank">{{data.Subject}}</a></td>
                                            <td align="center">{{data.Visited}}</td>
                                            <td align="center">{{data.Origin}}</td>
                                            <td align="center">
                                                <select name="public-choice" class="inputclass" v-model="data.PostType" @change="getCatalogSelected(data)">
                                                    <option value="技术">技术</option>
                                                    <option value="科技">科技</option>
                                                    <option value="新闻">新闻</option>
                                                    <option value="故事">故事</option>
                                                </select>
                                            </td>
                                            <td>{{data.AddDate | formatDate}}</td>
                                            <td align="center">
                                                <a v-on:click="editarticle('ArticleEdit', {id: data.ID})" class="btn btn-sm btn-outline filter-submit purple">
                                                    <i class="fa fa-edit"></i> 修改
                                                </a>
                                                <a v-on:click="delarticle(data.ID)" class="btn btn-sm btn-outline filter-submit dark" style="margin-right:0;">
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
    import { formatDate } from '../../util/date.js';
    import router from '../../router';

    export default {
        data: function () {
            return {
                datas: [],
                total: 5,
                display: 10,
                current: 1,
                Id: '',
                Catalog: '',
                Title: ''
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
                router.push({ name: 'ArticleList', params: { size: self.display, pageno: self.current } });
                self.load();
            },
            load: function () {
                var self = this;
                httper.post('/article/list/' + self.display + '/' + self.current, {
                    id: self.Id,
                    catalog: self.Catalog,
                    title: self.Title
                }).then(function (response) {
                    self.datas = response.data.Items;
                    self.total = response.data.PageArgs.TotalCount;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            editarticle: function (name, params) {
                router.push({ name: name, params: params });
            },
            delarticle: function (id) {
                var self = this;
                if (confirm("确认要删除？")) {
                    httper.get('/article/delete/' + id).then(function (response) {
                        if (response.data.ok) {
                            self.load();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            },
            getCatalogSelected: function (article) {
                var self = this;
                httper.post('/article/catalog/update', {
                    Id: article.id,
                    Catalog: article.catalog
                }).then(function (response) {
                    if (response.data.result === 1) {
                        router.push({ name: 'ArticleList', params: { size: self.display, pageno: self.current } });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            search: function () {
                var self = this;
                self.pagechange(1);
                self.load();
            }
        }
    };
</script>