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
                            <div class="portlet-body form" style="margin:0 auto;width:80%;">
                                <div class="form-horizontal mg0" role="form">
                                    <div class="form-body" style="padding:0 0 0 0;">
                                        <div class="form-group">
                                            <div class="col-md-12">
                                                <div style="margin:0 auto auto;font-size:20px;text-align:center;color:#36c6d3">角色权限配置</div>
                                            </div>
                                        </div>
                                        <div v-for="(items,type) of datas" class="form-group" :key="type">
                                            <div style="font-weight:bold;border-bottom: 1px solid #eeeeee;padding-bottom:5px;">{{type}}</div>
                                            <div style="overflow:hidden;padding:5px;">
                                                <div style="float:left;display:inline;width:150px;" class="checkbox-custom" v-for="item in items" :key="item.ID">
                                                    <input :id="item.ID" type="checkbox" v-model="functions" :value="item.ID" @click="mappingset(item.ID)" />
                                                    <label :for="item.ID">{{item.Funname}}</label>
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
    </div>
</template>

<script>
    import Menu from '../components/menu';
    import httper from '../../util/httper';

    export default {
        data: function () {
            return {
                Id: 0,
                RoleName: '',
                Status: 0,
                datas: [{ FunType: '基础数据', items: [{ Id: 1, Funname: '日数据' }] }],
                functions: []
            }
        },
        components: {
            Menu
        },
        created: function () {
            var self = this;
            if (self.$route.params.id) {
                var url = '/function/group/' + self.$route.params.id;
                httper.get(url).then(function (response) {
                    self.datas = response.data.functions;
                    self.functions = response.data.selectedids;
                }).catch(function (error) {
                    console.log(error);
                });
            }
        },
        methods: {
            mappingset: function (id) {
                var self = this;
                var index = $.inArray(id, self.functions);
                console.log(index);
                console.log(self.$route.params.id);
                console.log(self.functions);

                if (id > 0 && self.$route.params.id > 0) {
                    httper.post('/function/mapping/post', {
                        FunId: id,
                        RoleId: self.$route.params.id,
                        Toggle: index >= 0
                    }).then(function (response) {
                        console.log(response.data.result);
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            }
        }
    };
</script>

<style scoped>
    .checkbox-custom {
        position: relative;
        padding: 0 15px 0 25px;
        margin-bottom: 7px;
        margin-top: 0;
        display: inline-block;
    }
        /*
将初始的checkbox的样式改变
*/
        .checkbox-custom input[type="checkbox"] {
            opacity: 0; /*将初始的checkbox隐藏起来*/
            position: absolute;
            cursor: pointer;
            z-index: 2;
            margin: -6px 0 0 0;
            top: 50%;
            left: 3px;
        }
        /*
设计新的checkbox，位置
*/
        .checkbox-custom label:before {
            content: '';
            position: absolute;
            top: 50%;
            left: 0;
            margin-top: -9px;
            width: 19px;
            height: 18px;
            display: inline-block;
            border-radius: 2px;
            border: 1px solid #bbb;
            background: #fff;
        }
        /*
点击初始的checkbox，将新的checkbox关联起来
*/
        .checkbox-custom input[type="checkbox"]:checked + label:after {
            position: absolute;
            display: inline-block;
            font-family: 'Glyphicons Halflings';
            content: "\e013";
            top: 42%;
            left: 3px;
            margin-top: -5px;
            font-size: 11px;
            line-height: 1;
            width: 16px;
            height: 16px;
            color: #333;
        }

        .checkbox-custom label {
            cursor: pointer;
            line-height: 1.2;
            font-weight: normal; /*改变了rememberme的字体*/
            margin-bottom: 0;
            text-align: left;
        }
</style>