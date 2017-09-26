Vue.use(VueHtml5Editor, {
    showModuleName: true,
    image: {
        sizeLimit: 512 * 1024,
        upload: {
            url: '/upload'
        },
        compress: {
            width: 2000,
            height: 2000,
            quality: 80
        }
    }
});

new Vue({
    el: "#app",
    data: {
        UserId: '',
        UserName: '',
        Password: '',
        Subject: '',
        Body: '',
        UID: '',
        PWD: '',
        showModuleName: false
    },
    methods: {
        updateData: function(data) {
            this.Body = data
        },
        fullScreen: function() {
            this.$refs.editor.enableFullScreen()
        },
        focus: function() {
            this.$refs.editor.focus()
        },
        post: function() {
            var self = this;
            if (!$.trim(self.Subject) || !$.trim(self.Body))
                return;

            var url = '/post';
            $.ajax({
                url: url,
                type: 'POST',
                dataType: 'json',
                timeout: 60000,
                data: {
                    Subject: self.Subject,
                    Body: self.Body
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

            var url = '/loginpost';
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