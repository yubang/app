// 这里填充每一个页面的数据（一个字典）
var html_and_css_and_js_data = {"/admin/web/apps": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu class=\"el-menu-vertical-demo\" default-active=\"1\">\n<el-menu-item index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<p>\n<el-button type=\"success\">\u6dfb\u52a0\u5e94\u7528</el-button>\n</p>\n<el-card class=\"box-card app_info\" v-for=\"app in apps\">\n<div class=\"text item\">\n<p><strong>\u5e94\u7528Id\uff1a</strong> abcdef</p>\n<p><strong>\u5e94\u7528\u540d\u5b57\uff1a</strong> \u6d4b\u8bd5\u5e94\u7528A</p>\n<p><strong>\u5e94\u7528\u63cf\u8ff0\uff1a</strong> \u8fd9\u662f\u4e00\u4e2a\u5c0f\u5e94\u7528\uff01</p>\n<p><strong>\u5e94\u7528\u73af\u5883\uff1a</strong> static</p>\n<p><strong>\u5bb9\u5668\u4e2a\u6570\uff1a</strong> 0</p>\n<p><strong>\u5e94\u7528\u955c\u50cf\uff1a</strong> \u65e0</p>\n<p><strong>\u4ee3\u7801\u5730\u5740\uff1a</strong> git@github.com:yubang/ilog.git</p>\n<p><strong>\u5e94\u7528\u57df\u540d\uff1a</strong> http://blog.yubangweb.com</p>\n<p><strong>CPU\uff1a</strong> 1\u6838</p>\n<p><strong>\u5185\u5b58\uff1a</strong> 128M</p>\n<p>\n<el-button type=\"info\">\u64cd\u4f5c\u5e94\u7528</el-button>\n                \u00a0\n                <el-button type=\"danger\">\u5220\u9664\u5e94\u7528</el-button>\n</p>\n</div>\n</el-card>\n<div class=\"block\">\n<el-pagination :current-page=\"currentPage\" :page-size=\"5\" :total=\"50\" @current-change=\"currentChange\" layout=\"prev, pager, next\">\n</el-pagination>\n</div>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": ".app_info{\n    margin-bottom: 20px;\n}", "js": "app.init({\n    data: {\n        apps: [1, 2, 3, 4, 5],\n        currentPage: parseInt(app.get_args(\"page\") || 1),\n    },\n    methods: {\n        currentChange: function(page){\n            app.goto(\"/admin/web/apps?page=\" + page);\n        }\n    }\n})"}, "/web/html/404": {"title": "paas\u5e73\u53f0", "html": "", "css": "", "js": ""}};

// 404页面地址
var not_found_path = "/web/html/404";

// js name
var js_label_name = "js_";
var js_label_index = 0;

// 路由处理
function goto_url(url){
    if(html_and_css_and_js_data[url]){
        $("#css").html(html_and_css_and_js_data[url]['css']);
        $("#html").html(html_and_css_and_js_data[url]['html']);

        $("#" + js_label_name + js_label_index).remove();
        js_label_index++;
        $("body").append('<script id="'+js_label_name + js_label_index+'">'+html_and_css_and_js_data[url]['js']+'</script>');
    }else{
        if(html_and_css_and_js_data[not_found_path]){
            $("#css").html(html_and_css_and_js_data[not_found_path]['css']);
            $("#html").html(html_and_css_and_js_data[not_found_path]['html']);
            $("#" + js_label_name + js_label_index).remove();
            js_label_index++;
            $("body").append('<script id="'+js_label_name + js_label_index+'">'+html_and_css_and_js_data[not_found_path]['js']+'</script>');
        }else{
            alert("你访问的页面已经被吃掉了！");
        }
    }
}

// 组件js


// app内置的方法
function Applet(){

    this.vm = null;
    this.data = null;

    // 是否显示loading
    this.show_loading = function(sign){
        if(sign){
            $("#html").hide();
            $("#loading").show();
        }else{
            $("#html").show();
            $("#loading").hide();
        }
    }

    this.render = function(d){
        var that = this;
        this.vm = new Vue({
            el: '#html',
            data: d['data'] || {},
            methods: d['methods'] || {},
            mounted: function(){
                that.show_loading(false);
            }
        });
    }

    // 渲染页面
    this.init = function(vue_data){
        // 显示loadin动画
        this.show_loading(true);

        var api = vue_data['api'] || {};
        // 获取api数据
        var api_url = api['url'] || null;
        var api_data = api['data'] || {};
        var api_method = api['method'] || "GET";
        var api_headers = api['headers'] || {};
        var success = api['success'] || function(d){return d;}
        var error = api['error'] || function(){}
        var before_success = api['before_success'] || function(d){return true;};

        if(api_url == null){
            this.render(vue_data);
        }else{
            var that = this;
            $.ajax({
                url: api_url,
                type: api_method,
                headers: api_headers,
                data: api_data,
                success: function(d){
                    if(before_success(d)){
                        vue_data['data'] = success(d);
                        that.render(vue_data);
                    }
                },
                error: function(){
                    error();
                    that.render(vue_data);
                }
            });

        }

    }

    // 无刷新跳转页面
    this.goto = function(url){
        window.history.pushState({}, null,url);
        goto_url(window.location.pathname);
    }

    // 获取get参数
    this.get_args = function(name){
      var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
      var r = window.location.search.substr(1).match(reg);
      if (r != null) {
        return unescape(r[2]);
      }
      return false;
    }

    this.reload = function(){
        this.goto(window.location.pathname);
    }

}

var app = new Applet();


window.addEventListener('popstate', function(event) {
   goto_url(window.location.pathname);
});


// 全局js



$(document).ready(function(){
    goto_url(window.location.pathname);
});