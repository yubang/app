app.init({
    data: {
        form: {}
    },
    methods: {
        login: function(){
            var username =this.form.username;
            var password =this.form.password;

            var that = this;

            $.post("/admin/api/login", {
                username: username,
                password: password
            }, function(data){
                if(data['code'] == 0){
                    that.$message({type: 'success', message: data["data"]});
                    app.goto("/admin/web/apps");
                }else{
                    that.$message({type: 'error', message: data["msg"]});
                }
            }).error(function(){
                that.$message({type: 'error', message: "服务器未知错误！"});
            });


        }
    }
})