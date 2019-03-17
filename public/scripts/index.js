new Vue({
    el: "#dataformuser",
    data: {
        UserId: '',
        UserName: '',
        Password: '',
        UID: '',
        PWD: ''
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
        },
        loginpost: function() {
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
                        location.href = '/user/index';
                    } else {
                        alert('请确定账号是否存在与激活，密码是否正确，默认没有激活！');
                    }
                }
            });
        }
    }
});