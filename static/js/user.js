function addApp(){
    $("#add-app-submit").trigger("click");
}

function updateApp(){
    $("#update-app-submit").trigger("click");
}

function changeValue(obj, forSign){
    $("label[for='"+forSign+"'] b").html(obj.value);
}

function showUpdateDiv(i){
    var obj = objs[i];

    $("input[name='app_id']").val(obj.id);
    $("#update-title").val(obj.title);
    $("#update-description").val(obj.description);
    $("#update-min_number").val(obj.min_container_number).trigger("change");
    $("#update-max_number").val(obj.max_container_number).trigger("change");
    $("#update-memory").val(obj.memory).trigger("change");
    $('#updateAppDiv').modal({});
}

var objs = null;
function showAllApp(){
    $.get('/user/getApps', {}, function(data){
        objs = data;
        var html = template("t", {"list": data});
        $("#content").html(html);
    }, "json")
}

function login(){
    var username = $("input[name='username']").val();
    var password = $("input[name='password']").val();
    $.post("/user/login", {"username": username, "password": password}, function(data){
        if(data['code'] == 0){
            alert("登录成功");
            location.href="/user";
        }else{
            alert("用户名或密码错误！");
        }
    }, "json").error(function(){
        alert("服务器无法响应！");
    });
}

function deployment(app_id){
    $.post("/user/deploymentApp", {"app_id": app_id}, function(data){
        if(data['code'] == 0){
            alert("发布应用成功，可能有数分钟的延迟！");
        }else{
            alert("你已经重新发布应用了，请耐心等候！");
        }
    }).error(function(){
        alert("服务器无法响应！");
    });
}
