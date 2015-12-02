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

var servers = null;
function getServers(){
    $.get('/admin/getServers', {}, function(data){
        servers = data;
        var html = template("t", {"lists": data});
        $("#content").html(html);
    }, "json").error(function(){alert('拉取服务器列表失败！');});
}

function addServer(){
    var title = $("#title").val();
    var server_host = $("#server_host").val();
    var server_port = $("#server_port").val();
    var status = $("#status").val();
    var max_container_number = $("#max_container_number").val();
    var max_memory = $("#max_memory").val();
    var sort = $("#sort").val();

    $.post("/admin/addServer", {
        "title": title,
        "server_host": server_host,
        "server_port": server_port,
        "status": status,
        "max_container_number": max_container_number,
        "max_memory": max_memory,
        "sort": sort
    }, function(data){
        if(data['code'] == 0){
            alert("添加成功！");
            location.reload();
        }else{
            alert("添加失败！");
        }
    }, "json").error(function(){
        alert("服务器出错！");
    });

}

function editServer(index){
    var obj = servers[index];
    $("#update-title").val(obj.title);
    $("#server_id").val(obj.id);
    $("#update-server_host").val(obj.server_host);
    $("#update-server_port").val(obj.server_port);
    $("#update-status").val(obj.status);
    $("#update-max_container_number").val(obj.max_container_number);
    $("#update-max_memory").val(obj.max_memory);
    $("#update-sort").val(obj.sort);
    $('#updateServerDiv').modal({});
}

function updateServer(){
    var id = $("#server_id").val();
    var title = $("#update-title").val();
    var server_host = $("#update-server_host").val();
    var server_port = $("#update-server_port").val();
    var status = $("#update-status").val();
    var max_container_number = $("#update-max_container_number").val();
    var max_memory = $("#update-max_memory").val();
    var sort = $("#update-sort").val();

    $.post("/admin/updateServer", {
        "id": id,
        "title": title,
        "server_host": server_host,
        "server_port": server_port,
        "status": status,
        "max_container_number": max_container_number,
        "max_memory": max_memory,
        "sort": sort
    }, function(data){
        if(data['code'] == 0){
            alert("修改成功！");
            location.reload();
        }else{
            alert("修改失败！");
        }
    }, "json").error(function(){
        alert("服务器出错！");
    });

}

function login(){
    var username = $("input[name='username']").val();
    var password = $("input[name='password']").val();
    $.post("/admin/login", {"username": username, "password": password}, function(data){
        if(data['code'] == 0){
            alert("登录成功");
            location.href="/admin/user";
        }else{
            alert("用户名或密码错误！");
        }
    }, "json").error(function(){
        alert("服务器无法响应！");
    });
}