function gotoApps(){
    app.goto('/admin/web/apps');
}

function exitAccount() {
    $.get("/admin/api/exit", {}, function(data){
        app.goto("/admin/web/login");
    });
}