new Vue({
    el: "#header",
    data: {
        UID: '',
        PWD: ''
    },
    created: function() {
        var id = $.cookie('id');
        var name = $.cookie('username');
        if (id && name) {
            $("#logininfo").html("<a href='/user/" + id + "'><span style='color:#ffffff;padding-right:10px;'>" + name + "</span></a><span style='color: #ffffff; padding: 5px 20px 5px 20px; background-color: #ff6a00; cursor: pointer;' v-on:click='logout'>注 销</span>")
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
            var self = this;
            if (!$.trim(self.UID) || !$.trim(self.PWD))
                return;

            var url = '/login';
            $.ajax({
                url: url,
                type: 'POST',
                dataType: 'json',
                timeout: 60000,
                data: {
                    UID: self.UID,
                    PWD: self.PWD
                },
                error: function() { alert('Error loading'); },
                beforeSend: function() {
                    //$("#resultTable").html('<img src="/Images/loading.gif" />');
                },
                success: function(result) {
                    if (result.ok) {
                        location.href = '/';
                    }
                }
            });
        }
    }
});