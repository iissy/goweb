<template>
    <div id="page-container" style="position: relative;">
        <Menu tagIndex="10"></Menu>
        <div class="rightMain">
            <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
                <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
            </div>
            <div id="list">
                <div class="row">
                    <div class="col-md-12">
                        <div class="portlet light" style="margin-bottom: 0;">
                            <div class="portlet-body form" style="margin:0 auto;width:600px;">
                                <div class="form-horizontal mg0" role="form">
                                    <div class="form-body" style="padding:0 0 0 0;">
                                        <div class="form-group">
                                            <label class="col-md-2 control-label">功能：</label>
                                            <div class="col-md-10">
                                                <input type="text" v-model="Funname" class="form-control" placeholder="请在此填写功能" />
                                            </div>
                                        </div>
                                        <div class="form-group">
                                            <label class="col-md-2 control-label">分组：</label>
                                            <div class="col-md-10">
                                                <select v-model="FunType" class="form-control">
                                                    <option value="内容管理">内容管理</option>
                                                    <option value="用户管理">用户管理</option>
                                                    <option value="系统管理">系统管理</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group">
                                            <label class="col-md-2 control-label">控制器：</label>
                                            <div class="col-md-10">
                                                <input type="text" v-model="Controller" class="form-control" placeholder="请在此填写控制器" />
                                            </div>
                                        </div>
                                        <div class="form-group">
                                            <label class="col-md-1 control-label"></label>
                                            <div class="col-md-11">
                                                <button class="hd-button" v-on:click="post">添  加</button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import Menu from '../components/menu';
    import httper from '../../util/httper';
    import router from '../../router';

    export default {
        data: function () {
            return {
                Id: 0,
                Funname: '',
                FunType: '',
                Controller: ''
            };
        },
        components: {
            Menu
        },
        created: function () {
            var self = this;
            if (self.$route.params.id) {
                var url = '/function/get/' + self.$route.params.id;
                httper.get(url).then(function (response) {
                    self.Id = response.data.ID;
                    self.Funname = response.data.Funname;
                    self.FunType = response.data.FunType;
                    self.Controller = response.data.Controller;
                }).catch(function (error) {
                    console.log(error);
                });
            }
        },
        methods: {
            post: function () {
                var self = this;
                if (!!$.trim(self.Funname) && !!$.trim(self.FunType) && !!$.trim(self.Controller)) {
                    httper.post('/function/post', {
                        Id: self.Id,
                        Funname: self.Funname,
                        FunType: self.FunType,
                        Controller: self.Controller
                    }).then(function (response) {
                        if (response.data.ok) {
                            router.push({ name: 'FunctionList', params: { size: 15, pageno: 1 } });
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            }
        }
    };
</script>