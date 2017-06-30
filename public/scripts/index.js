jQuery(function() {
    var $ = jQuery;
    new Vue({
        el: "#app",
        data: {
            blogs: []
        },
        beforeCreate() {
            var self = this;
            $.ajax({
                url: "/blogs",
                type: 'GET',
                dataType: 'json',
                timeout: 5000,
                error: function() { alert('Error loading'); },
                beforeSend: function() {
                    //$("#app").append('<img src="/images/timg.gif" />');
                },
                success: function(result) {
                    $("#app").children("img:last-child").remove();
                    self.blogs = result;
                }
            });
        }
    });
});