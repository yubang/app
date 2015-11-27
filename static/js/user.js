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

$(document).ready(function(){
    showAllApp();
});