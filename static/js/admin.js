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
        var html = template("t", {"lists": data});
        $("#content").html(html);
    }, "json").error(function(){alert('拉取用户列表失败！');});
}