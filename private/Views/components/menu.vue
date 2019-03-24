<template>
    <div class="leftBar">
        <div id="header">
            <div style="margin:0 auto 0 auto;text-align:left;height:60px;">
                <div style="overflow:hidden;background-color: #17C4BB;">
                    <div style="margin-left:10px;padding:5px 5px 5px 5px;">
                        <a href="/"><img src="/assets/images/logo.png" height="50" width="200"/></a>
                    </div>
                </div>
            </div>
        </div>
        <ul>
            <li :class="{active: (currentIndex == 1)}"><router-link to="/main/index">个人首页</router-link></li>
            <li><span></span></li>
            <li :class="{active: (currentIndex == 4)}"><router-link to="/main/article/add">添加文章</router-link></li>
            <li :class="{active: (currentIndex == 5)}"><router-link to="/main/article/14/1">文章管理</router-link></li>
            <li><span></span></li>
            <li :class="{active: (currentIndex == 8)}"><router-link to="/main/account/14/1">用户列表</router-link></li>
        </ul>
    </div>
</template>

<script>
    export default {
        data: function () {
            return {
                currentIndex: parseInt(this.tagIndex)
            };
        },
        props: ['tagIndex']
    };
</script>