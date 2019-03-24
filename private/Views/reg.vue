<template>
    <div>
        <Header/>
        <div id="dataformuser" style="max-width:1200px;height:600px;margin:10px auto 0 auto;background-color: #fff;padding-top: 10px;" class="row">
            <div class="col-md-12">
                <div class="portlet light" style="margin-bottom: 0;">
                    <div class="portlet-body form" style="max-width: 1200px;margin:0 auto;">
                        <div class="form-horizontal mg0" role="form">
                            <div class="form-body" style="padding:0 0 0 0;">
                                <div class="form-group">
                                    <label class="col-md-3 control-label">账 号</label>
                                    <div class="col-md-6">
                                        <input type="text" v-model="UserId" class="form-control" placeholder="账号">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 control-label">昵 称</label>
                                    <div class="col-md-6">
                                        <input type="text" v-model="UserName" class="form-control" placeholder="昵称">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 control-label">密 码</label>
                                    <div class="col-md-6">
                                        <input type="password" v-model="Password" class="form-control" placeholder="密码">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 control-label"></label>
                                    <div class="col-md-6">
                                        <button style="color:#ffffff;padding:5px 30px 5px 30px;background-color: #36c6d3;font-size:16px;border-radius:10px;border: 1px solid #2bb8c4;" class="right" v-on:click="regpost">注 册</button>
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
    import httper from '../util/httper';
    import Header from './components/header';

    export default {
        data: function () {
            return {
                UserId: '',
                UserName: '',
                Password: '',
                UID: '',
                PWD: ''
            }
        },
        methods: {
            regpost: function() {
                var self = this;
                if (!$.trim(self.UserId) || !$.trim(self.UserName) || !$.trim(self.Password))
                    return;

                var url = '/regpost';
                $.ajax({
                    url: url,
                    type: 'POST',
                    dataType: 'json',
                    timeout: 60000,
                    data: {
                        UserId: self.UserId,
                        UserName: self.UserName,
                        Password: self.Password
                    },
                    error: function() { alert('Error loading'); },
                    beforeSend: function() {
                        //$("#resultTable").html('<img src="/Images/loading.gif" />');
                    },
                    success: function(result) {
                        location.href = '/';
                    }
                });
            }
        },
        components:{
            Header
        }
    }
</script>