Vue.use(VueHtml5Editor, {
    showModuleName: true,
    image: {
        sizeLimit: 512 * 1024,
        upload: {
            url: '/upload'
        }
    }
});

new Vue({
    el: "#app",
    data: {
        PostType: '技术',
        Id: '',
        Subject: '',
        Body: '',
        Picture: '',
        Origin: '爱施缘',
        Adding: true,
        showModuleName: false
    },
    created: function() {
        var self = this;
        if (!!id) {
            var url = '/article/get/' + id;
            $.ajax({
                url: url,
                type: 'GET',
                dataType: 'json',
                timeout: 60000,
                data: null,
                error: function() { alert('Error loading'); },
                beforeSend: function() {
                    //$("#resultTable").html('<img src="/Images/loading.gif" />');
                },
                success: function(result) {
                    self.Id = result.ID;
                    self.PostType = result.PostType;
                    self.Subject = result.Subject;
                    self.Picture = result.Picture;
                    self.Body = result.Body;
                    self.Origin = result.Origin;
                    if (!!result.ID) {
                        self.Adding = false;
                        $("#Id").attr("disabled", true);
                    }
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
        change: function() {
            var self = this;
            $("#file").upload('/upload', function(resp) {
                self.Picture = resp.data;
            });
        },
        post: function() {
            var self = this;
            if (!$.trim(self.Subject) || !$.trim(self.Body) || !$.trim(self.Picture))
                return;

            var url = '/article/post';
            $.ajax({
                url: url,
                type: 'POST',
                dataType: 'json',
                timeout: 60000,
                data: {
                    Adding: self.Adding,
                    Id: self.Id,
                    PostType: self.PostType,
                    Subject: self.Subject,
                    Picture: self.Picture,
                    Body: self.Body,
                    Origin: self.Origin
                },
                error: function() { alert('Error loading'); },
                beforeSend: function() {
                    //$("#resultTable").html('<img src="/Images/loading.gif" />');
                },
                success: function(result) {
                    if (result.ok)
                        location.href = '/article/list/1';
                }
            });
        }
    }
});