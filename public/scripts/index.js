Vue.use(VueHtml5Editor, {
    showModuleName: true,
    image: {
        sizeLimit: 512 * 1024,
        upload: {
            url: '/upload'
        },
        compress: {
            width: 1000,
            height: 1000,
            quality: 80
        }
    }
});

new Vue({
    el: "#app",
    data: {
        Id: '',
        Subject: '',
        Body: '',
        showModuleName: false
    },
    beforeCreate: function() {
        var self = this;
        self.Id = $('#Id').val();
        if (self.Id) {
            var url = '/get/' + self.Id;
            $.ajax({
                url: url,
                type: 'GET',
                dataType: 'json',
                timeout: 60000,
                error: function() { alert('Error loading'); },
                beforeSend: function() {
                    //$("#resultTable").html('<img src="/Images/loading.gif" />');
                },
                success: function(response) {
                    self.Id = response.id;
                    self.Subject = response.subject;
                    self.Body = response.body;
                }
            });
        }
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
            var url = '/newask';
            $.ajax({
                url: url,
                type: 'POST',
                dataType: 'json',
                timeout: 60000,
                data: {
                    Id: self.Id,
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
        }
    }
});