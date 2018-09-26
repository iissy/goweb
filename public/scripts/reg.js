new Vue({
    el: "#header",
    data: {},
    created: function() {
        var id = $.cookie('id');
        var userid = $.cookie('userid');
        var username = $.cookie('username');
        var token = $.cookie('token');
        if (id && userid && username && token) {
            $("#logininfo").html("<a href='/user/index'><span style='color:#ffffff;padding-right:20px;padding-left:20px;'>" + username + "</span></a><span style='color: #ffffff; padding: 5px 20px 5px 20px; background-color: #ff6a00; cursor: pointer;' v-on:click='logout'>注 销</span>")
        }
    },
    methods: {
        reg: function() {
            location.href = '/reg';
        },
        logout: function() {
            $.ajax({
                url: '/logout',
                type: 'GET',
                dataType: 'json',
                timeout: 6000,
                error: function() { alert('Error loading'); },
                success: function(result) {
                    location.href = '/';
                }
            });
        },
        login: function() {
            location.href = '/login';
        }
    }
});

var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?32cee789010835e47ba7585de7418cc5";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();