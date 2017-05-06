app.init({
    api: {
        url: "/admin/api/getShellLog",
        method: "POST",
        data: {
            page: app.get_args("page") || 1
        },
        success: function (data) {
            return {
                "objs": data['data']['objs'],
                "nums": data['data']['nums'],
                tabIndex: "5",
                currentPage: parseInt(app.get_args("page") || 1)
            }
        },
        before_success: beforeHandleAjx,
        error: handleError
    },
    methods: {
        currentChange: function(page){
            app.goto("/admin/web/shell?page="+page);
        }
    }
});