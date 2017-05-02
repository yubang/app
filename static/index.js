// 这里填充每一个页面的数据（一个字典）
var html_and_css_and_js_data = {"/admin/web/apps": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu class=\"el-menu-vertical-demo\" default-active=\"1\">\n<el-menu-item @click.native=\"app.goto('/admin/web/apps');\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<p>\n<el-button @click=\"app.goto('/admin/web/new-app');\" type=\"success\">\u6dfb\u52a0\u5e94\u7528</el-button>\n</p>\n<el-card class=\"box-card app_info\" v-for=\"app in apps\">\n<div class=\"text item\">\n<p><strong>\u5e94\u7528Id\uff1a</strong> {{app.appId}}</p>\n<p><strong>\u5e94\u7528\u540d\u5b57\uff1a</strong> {{app.name}}</p>\n<p><strong>\u5e94\u7528\u63cf\u8ff0\uff1a</strong> {{app.desc}}</p>\n<p><strong>\u5e94\u7528\u73af\u5883\uff1a</strong> {{app.image}}</p>\n<p><strong>\u5bb9\u5668\u4e2a\u6570\uff1a</strong> {{app.nums}}</p>\n<p><strong>\u5e94\u7528\u955c\u50cf\uff1a</strong>\n<template v-if=\"app.nowImageName != ''\">{{app.nowImageName}}</template>\n<template v-else=\"\">\u65e0</template>\n</p>\n<p><strong>\u4ee3\u7801\u5730\u5740\uff1a</strong> {{app.git}}</p>\n<p><strong>\u5e94\u7528\u57df\u540d\uff1a</strong> {{app.domain}}</p>\n<p><strong>CPU\uff1a</strong> {{app.cpu}}</p>\n<p><strong>\u5185\u5b58\uff1a</strong> {{app.memory}}</p>\n<p>\n<el-button @click=\"optionApp(app.appId);\" type=\"info\">\u64cd\u4f5c\u5e94\u7528</el-button>\n                \u00a0\n                <el-button @click=\"deleteApp(app.appId);\" type=\"danger\">\u5220\u9664\u5e94\u7528</el-button>\n</p>\n</div>\n</el-card>\n<div class=\"block\">\n<el-pagination :current-page=\"currentPage\" :page-size=\"5\" :total=\"totalPage\" @current-change=\"currentChange\" layout=\"prev, pager, next\">\n</el-pagination>\n</div>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": ".app_info{\n    margin-bottom: 20px;\n}", "js": "app.init({\n    api:{\n      url: \"/admin/api/appList\",\n        data: {page: parseInt(app.get_args(\"page\")||1)},\n        success: function(data){\n            return {\n                apps: data[\"data\"].apps,\n                currentPage: parseInt(app.get_args(\"page\") || 1),\n                totalPage: data[\"data\"].nums\n            }\n        }\n    },\n    data: {\n        apps: [1, 2, 3, 4, 5],\n        currentPage: parseInt(app.get_args(\"page\") || 1),\n    },\n    methods: {\n        currentChange: function(page){\n            app.goto(\"/admin/web/apps?page=\" + page);\n        },\n        optionApp: function(appId){\n            app.goto(\"/admin/web/option-app?appId=\"+appId);\n        },\n        deleteApp: function(appId){\n\n            this.$confirm('\u6b64\u64cd\u4f5c\u5c06\u6c38\u4e45\u5220\u9664\u5e94\u7528, \u662f\u5426\u7ee7\u7eed?', '\u63d0\u793a', {\n                confirmButtonText: '\u786e\u5b9a',\n                cancelButtonText: '\u53d6\u6d88',\n                type: 'warning'\n            }).then(() => {\n\n                    var that = this;\n\n                    $.post(\"/admin/api/deleteApp\", {\"appId\": appId}, function(data){\n\n                        if(data[\"code\"]==0){\n                            that.$message({\n                                type: 'success',\n                                message: data[\"data\"]\n                            });\n                            app.goto(\"/admin/web/apps\");\n                        }else{\n                            that.$message({\n                                type: 'error',\n                                message: data[\"msg\"]\n                            });\n                        }\n\n                    }).error(function(){\n                        that.$message({\n                            type: 'error',\n                            message: data[\"msg\"]\n                        });\n                    });\n            })\n\n        }\n    }\n})"}, "/admin/web/new-app": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu class=\"el-menu-vertical-demo\" default-active=\"1\">\n<el-menu-item @click.native=\"app.goto('/admin/web/apps');\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-form :model=\"form\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u5e94\u7528\u540d\u5b57\">\n<el-input v-model=\"form.name\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u63cf\u8ff0\">\n<el-input type=\"textarea\" v-model=\"form.desc\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u57df\u540d\">\n<el-input v-model=\"form.domain\"></el-input>\n</el-form-item>\n<el-form-item label=\"git\u5730\u5740\">\n<el-input v-model=\"form.git\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668CPU\">\n<el-radio-group v-model=\"form.cpu\">\n<el-radio-button label=\"1\u6838\"></el-radio-button>\n<el-radio-button label=\"2\u6838\"></el-radio-button>\n<el-radio-button label=\"3\u6838\"></el-radio-button>\n<el-radio-button label=\"4\u6838\"></el-radio-button>\n<el-radio-button label=\"5\u6838\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u5185\u5b58\">\n<el-radio-group v-model=\"form.memory\">\n<el-radio-button label=\"64M\"></el-radio-button>\n<el-radio-button label=\"128M\"></el-radio-button>\n<el-radio-button label=\"256M\"></el-radio-button>\n<el-radio-button label=\"512M\"></el-radio-button>\n<el-radio-button label=\"1024M\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u6570\u91cf\">\n<el-input-number :max=\"10\" :min=\"1\" v-model=\"form.nums\"></el-input-number>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u73af\u5883\">\n<el-select placeholder=\"\u8bf7\u9009\u62e9\u5e94\u7528\u73af\u5883\" v-model=\"form.image\">\n<el-option :key=\"item.value\" :label=\"item.label\" :value=\"item.value\" v-for=\"item in options\">\n</el-option>\n</el-select>\n</el-form-item>\n<el-form-item>\n<el-button @click=\"createApp();\" type=\"primary\">\u7acb\u5373\u521b\u5efa</el-button>\n</el-form-item>\n</el-form>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": "", "js": "app.init({\n    data: {\n        form: {\n            image: \"static\",\n            memory: \"64M\",\n            cpu: \"1\u6838\",\n            nums: 1\n        },\n        options: [{label: \"\u9759\u6001\u8d44\u6e90\u73af\u5883\", value: \"static\"}]\n    },\n    methods: {\n        createApp: function(){\n            var name = this.form.name;\n            var desc = this.form.desc;\n            var domain = this.form.domain;\n            var git = this.form.git;\n            var cpu = this.form.cpu;\n            var memory = this.form.memory;\n            var nums = parseInt(this.form.nums);\n            var image = this.form.image;\n\n            if(!$.trim(name)){\n                this.$message({\n                    message: '\u5e94\u7528\u540d\u5b57\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(desc)){\n                this.$message({\n                    message: '\u5e94\u7528\u63cf\u8ff0\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(domain)){\n                this.$message({\n                    message: '\u5e94\u7528\u57df\u540d\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(git)){\n                this.$message({\n                    message: '\u4ee3\u7801\u4ed3\u5e93\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n            var that = this;\n            $.post(\"/admin/api/createApp\", {\n                name: name,\n                desc: desc,\n                domain: domain,\n                git: git,\n                cpu: cpu,\n                memory: memory,\n                nums: nums,\n                image: image\n            }, function(data){\n                if(data[\"code\"] == 0){\n                    that.$message({\n                        message: '\u521b\u5efa\u5e94\u7528\u6210\u529f\uff01',\n                        type: 'success'\n                    });\n                    setTimeout(function(){app.goto(\"/admin/web/apps\");}, 2000);\n                }else{\n                    that.$message.error(data[\"msg\"]);\n                }\n            }).error(function(){\n                that.$message.error('\u670d\u52a1\u5668\u53d1\u751f\u4e0d\u53ef\u9884\u6599\u7684\u9519\u8bef\uff01');\n            });\n\n        }\n    }\n})"}, "/web/html/404": {"title": "paas\u5e73\u53f0", "html": "<el-alert :closable=\"false\" description=\"\u4f60\u8bbf\u95ee\u7684\u9875\u9762\u4f3c\u4e4e\u662f\u4e0d\u5b58\u5728\u7684\uff01\" show-icon=\"\" title=\"\u6765\u81eapaas\u5e73\u53f0\u7684\u63d0\u793a\" type=\"warning\">\n</el-alert>", "css": "", "js": "app.init({})"}, "/admin/web/option-app": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu class=\"el-menu-vertical-demo\" default-active=\"1\">\n<el-menu-item @click.native=\"app.goto('/admin/web/apps');\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5e94\u7528\u4fe1\u606f</span>\n</div>\n<div class=\"text item\">\n<p><strong>\u5e94\u7528Id\uff1a</strong> abcdef</p>\n<p><strong>\u5e94\u7528\u540d\u5b57\uff1a</strong> \u6d4b\u8bd5\u5e94\u7528A</p>\n<p><strong>\u5e94\u7528\u63cf\u8ff0\uff1a</strong> \u8fd9\u662f\u4e00\u4e2a\u5c0f\u5e94\u7528\uff01</p>\n<p><strong>\u5e94\u7528\u73af\u5883\uff1a</strong> static</p>\n<p><strong>\u5e94\u7528\u955c\u50cf\uff1a</strong> \u65e0</p>\n<p><strong>\u4ee3\u7801\u5730\u5740\uff1a</strong> git@github.com:yubang/ilog.git</p>\n<p><strong>\u5e94\u7528\u57df\u540d\uff1a</strong> http://blog.yubangweb.com</p>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5bb9\u5668\u4fe1\u606f</span>\n<el-button style=\"float: right;\" type=\"primary\">\u66f4\u65b0</el-button>\n</div>\n<div class=\"text item\">\n<el-form :model=\"container\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u5bb9\u5668\u6570\u91cf\">\n<el-input-number :max=\"10\" :min=\"1\" v-model=\"container.nums\"></el-input-number>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668CPU\">\n<el-radio-group v-model=\"container.cpu\">\n<el-radio-button label=\"1\u6838\"></el-radio-button>\n<el-radio-button label=\"2\u6838\"></el-radio-button>\n<el-radio-button label=\"3\u6838\"></el-radio-button>\n<el-radio-button label=\"4\u6838\"></el-radio-button>\n<el-radio-button label=\"5\u6838\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u5185\u5b58\">\n<el-radio-group v-model=\"container.memory\">\n<el-radio-button label=\"64M\"></el-radio-button>\n<el-radio-button label=\"128M\"></el-radio-button>\n<el-radio-button label=\"256M\"></el-radio-button>\n<el-radio-button label=\"512M\"></el-radio-button>\n<el-radio-button label=\"1024M\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n</el-form>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u955c\u50cf\u4fe1\u606f</span>\n</div>\n<div class=\"text item\">\n<div class=\"text item\">\n<el-form :model=\"container\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u6700\u8fd1\u955c\u50cf\">\n<el-select placeholder=\"\u8bf7\u9009\u62e9\u955c\u50cf\" v-model=\"image\">\n<el-option :key=\"item.value\" :label=\"item.label\" :value=\"item.value\" v-for=\"item in options\">\n</el-option>\n</el-select>\n</el-form-item>\n<el-form-item label=\"\u5f53\u524d\u955c\u50cf\">\n                        \u65e0\uff0c\u6253\u5305\u4e8e2017\u5e745\u67081\u65e5\n                        \u00a0\u00a0\u00a0\n                        <el-button type=\"success\">\u4ee3\u7801\u6253\u5305</el-button>\n</el-form-item>\n</el-form>\n</div>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5e94\u7528\u65e5\u5fd7</span>\n</div>\n<div class=\"text item\">\n<p>\u955c\u50cf\u6253\u5305\uff01</p>\n</div>\n</el-card>\n</br></br></br>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": "", "js": "app.init({\n    data:{\n        container: {\"nums\": 1, cpu: \"1\u6838\", \"memory\": \"64M\"},\n        options: [{\"value\": \"adhuehfee\", \"label\": \"erfre\"}],\n        image: \"\"\n    }\n})"}};

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