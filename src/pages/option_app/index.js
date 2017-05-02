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
    data:{
        container: {"nums": 1, cpu: "1核", "memory": "64M"},
        options: [{"value": "adhuehfee", "label": "erfre"}],
        image: ""
    }
})