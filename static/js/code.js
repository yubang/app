var codes = null;
function getAllCode(){
    $.get("/code/getAllCode", {}, function(data){
        codes = data;
        var html = template("t", {"lists": data});
        $("#content").html(html);
        var url = "http://" + window.location.host;
        $(".address").each(function(){
            var t = $(this).html();
            $(this).html(url+t);
        });
    }).error(function(){
        alert("无法拉取数据！");
    });
}

function addCode(){

    var title = $("#title").val();
    var status = $("#status").val();

    $.post("/code/addCode", {
        "title": title,
        "status": status
    }, function(data){
        if(data['code'] == 0){
            alert('添加成功！');
            location.reload();
        }else{
            alert('添加失败！');
        }
    }, "json").error(function(){
        alert("服务器无法响应！");
    });

}

function editCode(index){
    var obj = codes[index];
    $("#update-id").val(obj.id);
    $("#update-title").val(obj.title);
    $("#update-status").val(obj.status);
    $('#codeDiv').modal({});
}

function updateCode(){
    var title = $("#update-title").val();
    var status = $("#update-status").val();
    var id = $("#update-id").val();
    $.post("/code/updateCode", {
        "title": title,
        "status": status,
        "id": id
    }, function(data){
        if(data['code'] == 0){
            alert('修改成功！');
            location.reload();
        }else{
            alert('修改失败！');
        }
    }, "json").error(function(){
        alert("服务器无法响应！");
    });
}

function uploadZip(){
    var mypic = document.getElementById('file').files[0];
    var fd = new FormData();
    fd.append("file",mypic);
    fd.append("code_id", code_id);

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function(){
        if(xhr.readyState==4 && xhr.status==200){
            $("#tip").html("上传完成");
            alert("上传代码成功！");
        }else if(xhr.readyState==4 && xhr.status==500){
            $("#tip").html("上传失败");
            alert("上传代码失败！");
        }
    }

    //侦查当前附件上传情况
    xhr.upload.onprogress = function(evt){
        //侦查附件上传情况
        //通过事件对象侦查
        //该匿名函数表达式大概0.05-0.1秒执行一次
        //console.log(evt);
        //console.log(evt.loaded);  //已经上传大小情况
        //evt.total; 附件总大小
        var loaded = evt.loaded;
        var tot = evt.total;
        var per = Math.floor(100*loaded/tot);  //已经上传的百分比
        var son =  document.getElementById('son');
        $("#p").html(per+"%");
        $("#p").css("width",per+"%");
        $("#tip").html("上传中："+per+"%");
    }

    xhr.open("post","/code/uploadFile");
    xhr.send(fd);
}

var code_id = 0;
function updateFile(id){
    code_id = id;
    $("#file").trigger("click");
}

$(document).ready(function(){
    getAllCode();
});