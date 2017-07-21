new Vue({
    el: "#app",
    data: {
        UserId: '',
        UserName: '',
        Password: ''
    },
    methods: {
        post: function() {
            var self = this;
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
    }
});