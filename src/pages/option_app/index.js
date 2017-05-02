app.init({
    api:{
      url: "/admin/api/appInfo",
      data:{"appId": app.get_args("appId")},
      method: "POST",
      success: function (data) {
          var obj = data["data"];
          return {
              container: {"nums": obj["appInfo"]["nums"], cpu: obj["appInfo"]["cpu"]+"核", "memory": obj["appInfo"]["memory"]+"M"},
              options: [{"value": "adhuehfee", "label": "erfre"}],
              image: obj["appInfo"]["image"],
              app: obj["appInfo"],
              appId: obj["appId"],
              logs: obj["logs"]
          }
      }
    },
    methods: {
        optioneContauiner: function(){
            var that = this;
            $.post("/admin/api/updateAppContainerInfo", {
                "appId": app.get_args("appId"),
                "memory": that.container.memory.substring(0, that.container.memory.length-1),
                "cpu": that.container.cpu.substring(0, that.container.cpu.length-1),
                "nums": that.container.nums
            }, function(data){
                if(data["code"]==0){
                    that.$message({
                        message: data["data"],
                        type: 'success'
                    });
                    app.reload();
                }else{
                    that.$message({
                        message: data['msg'],
                        type: 'error'
                    });
                }
            }).error(function(){
                that.$message({
                    message: "服务器未知错误！",
                    type: 'error'
                });
            });
        }
    }
})