app.init({
    api:{
      url: "/admin/api/appList",
        data: {page: parseInt(app.get_args("page")||1)},
        success: function(data){
            return {
                apps: data["data"].apps,
                currentPage: parseInt(app.get_args("page") || 1),
                totalPage: data["data"].nums,
                tabIndex: "1"
            }
        },
        before_success: beforeHandleAjx,
        error:handleError
    },
    data: {
        apps: [1, 2, 3, 4, 5],
        currentPage: parseInt(app.get_args("page") || 1),
    },
    methods: {
        currentChange: function(page){
            app.goto("/admin/web/apps?page=" + page);
        },
        optionApp: function(appId){
            app.goto("/admin/web/option-app?appId="+appId);
        },
        deleteApp: function(appId){

            this.$confirm('此操作将永久删除应用, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {

                    var that = this;

                    $.post("/admin/api/deleteApp", {"appId": appId}, function(data){

                        if(data["code"]==0){
                            that.$message({
                                type: 'success',
                                message: data["data"]
                            });
                            app.goto("/admin/web/apps");
                        }else{
                            that.$message({
                                type: 'error',
                                message: data["msg"]
                            });
                        }

                    }).error(function(){
                        that.$message({
                            type: 'error',
                            message: data["msg"]
                        });
                    });
            })

        }
    }
})