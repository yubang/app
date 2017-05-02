app.init({
    data: {
        form: {
            image: "static",
            memory: "64M",
            cpu: "1核",
            nums: 1
        },
        options: [{label: "静态资源环境", value: "static"}]
    },
    methods: {
        createApp: function(){
            var name = this.form.name;
            var desc = this.form.desc;
            var domain = this.form.domain;
            var git = this.form.git;
            var cpu = this.form.cpu;
            var memory = this.form.memory;
            var nums = parseInt(this.form.nums);
            var image = this.form.image;

            if(!$.trim(name)){
                this.$message({
                    message: '应用名字不能为空！',
                    type: 'warning'
                });
                return ;
            }

            if(!$.trim(desc)){
                this.$message({
                    message: '应用描述不能为空！',
                    type: 'warning'
                });
                return ;
            }

            if(!$.trim(domain)){
                this.$message({
                    message: '应用域名不能为空！',
                    type: 'warning'
                });
                return ;
            }

            if(!$.trim(git)){
                this.$message({
                    message: '代码仓库不能为空！',
                    type: 'warning'
                });
                return ;
            }
            var that = this;
            $.post("/admin/api/createApp", {
                name: name,
                desc: desc,
                domain: domain,
                git: git,
                cpu: cpu,
                memory: memory,
                nums: nums,
                image: image
            }, function(data){
                if(data["code"] == 0){
                    that.$message({
                        message: '创建应用成功！',
                        type: 'success'
                    });
                    setTimeout(function(){app.goto("/admin/web/apps");}, 200);
                }else{
                    that.$message.error(data["msg"]);
                }
            }).error(function(){
                that.$message.error('服务器发生不可预料的错误！');
            });

        }
    }
})