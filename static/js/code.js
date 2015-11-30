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

$(document).ready(function(){
    getAllCode();
});