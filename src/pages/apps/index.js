app.init({
    api:{
      url: "/admin/api/appList",
        data: {page: parseInt(app.get_args("page")||1)},
        success: function(data){
            return {
                apps: data["data"].apps,
                currentPage: parseInt(app.get_args("page") || 1),
                totalPage: data["data"].nums
            }
        }
    },
    data: {
        apps: [1, 2, 3, 4, 5],
        currentPage: parseInt(app.get_args("page") || 1),
    },
    methods: {
        currentChange: function(page){
            app.goto("/admin/web/apps?page=" + page);
        },
        optionApp: function(appId){
            app.goto("/admin/web/option-app?appId="+appId);
        }
    }
})