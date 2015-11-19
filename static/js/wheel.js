/**
 * 轮子js框架
 * 为了前后端分离，但又尽量简单，依赖jq
 * @author: yubang
 */

function Wheel(){

    /**
     * 拦截所有符合规则的请求
     *
     */
    this.init = function(){

        $("a[data-route='wheel']").on('click', function(){
            var func = $(this).attr("data-function");
            var data = $(this).attr("data-args");
            if(data == undefined){
                data = "";
            }
            if(func != undefined){
                eval(func+"("+data+");");
            }
            return false;
        });

    }

}