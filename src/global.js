function gotoApps(){
    app.goto('/admin/web/apps');
}

function gotoContainerServer() {
    app.goto('/admin/web/server');
}

function exitAccount() {
    $.get("/admin/api/exit", {}, function(data){
        app.goto("/admin/web/login");
    });
}

function beforeHandleAjx(data){
    if(data["code"] == 10005){
        app.goto("/admin/web/login");
        return false;
    }
    return true;
}

function handleError(data){
    app.goto("/web/html/500");
}