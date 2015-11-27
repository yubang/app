function addApp(){
    $("#add-app-submit").trigger("click");
}

function changeValue(obj, forSign){
    $("label[for='"+forSign+"'] b").html(obj.value);
}

function showAllApp(){
    $.get('/user/getApps', {}, function(data){
        var html = template("t", {"list": data});
        $("#content").html(html);
    }, "json")
}

$(document).ready(function(){
    showAllApp();
});