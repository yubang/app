app.init({
    api: {
        url: "/admin/api/getSsh",
        success: function (data) {
            return {
                "sshContent": data['data'],
                tabIndex: "4"
            }
        },
        before_success: beforeHandleAjx,
        error: handleError
    }
});