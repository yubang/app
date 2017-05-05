app.init({
    api: {
      url: "/admin/api/getContainerServer",
      success: function(data){
          return {
              "tableData": data['data']["nodes"],
              tabIndex: "2",
              command: data["data"]["command"]
          }
      },
        before_success: beforeHandleAjx,
        error:handleError
    },
    data: {
        "tableData": [],
        tabIndex: "2"
    },
    methods: {
        findJoin: function(){
            this.$alert(this.command, '请在容器服务器输入下面命令：', {
                    confirmButtonText: '确定'
            });
        },
        deleteNode: function(index){
            var nodeName = this.tableData[index]["ID"];
            var that = this;
            $.post("/admin/api/deleteNode", {"nodeName": nodeName}, function(data){
                if(data["code"] == 0){
                    that.$message({
                        message: '移除节点成功！',
                        type: 'success'
                    });
                    app.reload();
                }else{
                    that.$message.error(data["msg"]);
                }
            }).error(function(){
                that.$message.error('服务器发生不可预料的错误！');
            });
        }
    }
})