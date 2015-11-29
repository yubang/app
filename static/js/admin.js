var users = null;
function addUser(){

    var username = $("#username").val();
    var password = $("#password").val();
    var nickname = $("#nickname").val();
    var status = $("#status").val();

    $.post("/admin/addUser", {
        "username": username,
        "password": password,
        "nickname": nickname,
        "status": status
    }, function(data){
        if(data['code'] == 0){
            alert("添加用户成功！");
            location.reload();
        }else{
            alert("用户名已经存在！");
        }
    }, "json").error(function(){
        alert('服务器出错！');
    });

}


function getUsers(){
    $.get('/admin/getUsers', {}, function(data){
        users = data;
        var html = template("t", {"lists": data});
        $("#content").html(html);
    }, "json").error(function(){alert('拉取用户列表失败！');});
}


function editUser(index){
    var obj = users[index];
    $("#update-username").val(obj.username);
    $("#update-nickname").val(obj.nickname);
    $("#update-status").val(obj.status);
    $("#user_id").val(obj.id);
    $('#updateUserDiv').modal({});
}

function updateUser(){

    var username = $("#update-username").val();
    var password = $("#update-password").val();
    var nickname = $("#update-nickname").val();
    var status = $("#update-status").val();
    var user_id = $("#user_id").val();

    $.post("/admin/updateUser", {
        "username": username,
        "password": password,
        "nickname": nickname,
        "status": status,
        "user_id": user_id
    }, function(data){
        if(data['code'] == 0){
            alert("修改用户成功！");
            location.reload();
        }else{
            alert("修改用户失败！");
        }
    }, "json").error(function(){
        alert('服务器出错！');
    });

}