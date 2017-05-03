app.init({
    api:{
      url: "/admin/api/appInfo",
      data:{"appId": app.get_args("appId")},
      method: "POST",
      success: function (data) {
          var obj = data["data"];

          var selectImages = [];
          for(var i in obj["images"]){
              var tmpObj = $.parseJSON(obj["images"][i]);
              selectImages.push({value: i, label: tmpObj["imageAbout"], time: tmpObj["imageCreateTime"], "image": tmpObj["imageName"]})
          }

          return {
              container: {"nums": obj["appInfo"]["nums"], cpu: obj["appInfo"]["cpu"]+"核", "memory": obj["appInfo"]["memory"]+"M"},
              options: selectImages,
              image: "",
              app: obj["appInfo"],
              appId: obj["appId"],
              logs: obj["logs"],
              imageMessage: "",
              tabIndex: "1"
          }
      },
        before_success: beforeHandleAjx,
        error:handleError
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
        },
        buildImage: function(){
            var that = this;
            this.$prompt('请输入镜像说明', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPattern: /\S/,
                inputErrorMessage: '说明不能为空！'
            }).then(({ value }) => {
                buildImage(that, value)
            });

        },
        selectImage: function(i){
            // 修改提示
            this.imageMessage = "镜像打包于：" + this.options[i]["time"];
        },
        useImage: function(){
            var index = this.image;
            if(index == ""){
                this.$message({
                    message: "请选择一个镜像！",
                    type: 'warning'
                });
                return;
            }
            var data = this.options[index];
            var that = this;

            $.post("/admin/api/useImage", {
                "appId": app.get_args("appId"),
                imageName:data["image"],
                imageTime:data["time"],
                imageAbout: data["label"]
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


function buildImage(that, v){
    $.post("/admin/api/buildImage", {"appId": app.get_args("appId"), "imageAbout": v}, function(data){
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