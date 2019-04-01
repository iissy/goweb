<template>
    <div id="page-container" style="position: relative;">
        <Menu tagIndex="9"></Menu>
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
                                            <label class="col-md-2 control-label">角色：</label>
                                            <div class="col-md-10">
                                                <input type="text" v-model="RoleName" class="form-control" placeholder="请在此填写角色" />
                                            </div>
                                        </div>
                                        <div class="form-group">
                                            <label class="col-md-2 control-label">状态：</label>
                                            <div class="col-md-10">
                                                <select v-model="Status" class="form-control">
                                                    <option value="1">启用</option>
                                                    <option value="0">停用</option>
                                                </select>
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
                RoleName: '',
                Status: 0
            };
        },
        components: {
            Menu
        },
        created: function () {
            var self = this;
            if (self.$route.params.id) {
                var url = '/role/get/' + self.$route.params.id;
                httper.get(url).then(function (response) {
                    self.Id = response.data.ID;
                    self.RoleName = response.data.RoleName;
                    self.Status = response.data.Status;
                }).catch(function (error) {
                    console.log(error);
                });
            }
        },
        methods: {
            post: function () {
                var self = this;
                if (!!$.trim(self.RoleName)) {
                    httper.post('/role/post', {
                        Id: self.Id,
                        RoleName: self.RoleName,
                        Status: self.Status,
                    }).then(function (response) {
                        if (response.data.ok) {
                            router.push({ name: 'RoleList', params: { size: 15, pageno: 1 } });
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            }
        }
    };
</script>