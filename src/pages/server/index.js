app.init({
    api: {
      url: "/admin/api/getContainerServer",
      success: function(data){
          return {
              "tableData": data['data'],
              tabIndex: "2"
          }
      },
        before_success: beforeHandleAjx,
        error:handleError
    },
    data: {
        "tableData": [],
        tabIndex: "2"
    }
})