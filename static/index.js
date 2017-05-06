// 这里填充每一个页面的数据（一个字典）
var html_and_css_and_js_data = {"/admin/web/login": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<el-form :model=\"form\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u7528\u6237\u540d\uff1a\">\n<el-input v-model=\"form.username\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5bc6\u7801\uff1a\">\n<el-input type=\"password\" v-model=\"form.password\"></el-input>\n</el-form-item>\n<el-form-item>\n<el-button @click=\"login\" type=\"primary\">\u767b\u5f55</el-button>\n</el-form-item>\n</el-form>", "css": "form{\n    width:500px;\n    margin-left:auto;\n    margin-right: auto;\n    margin-top: 50px;\n}", "js": "app.init({\n    data: {\n        form: {}\n    },\n    methods: {\n        login: function(){\n            var username =this.form.username;\n            var password =this.form.password;\n\n            var that = this;\n\n            $.post(\"/admin/api/login\", {\n                username: username,\n                password: password\n            }, function(data){\n                if(data['code'] == 0){\n                    that.$message({type: 'success', message: data[\"data\"]});\n                    app.goto(\"/admin/web/apps\");\n                }else{\n                    that.$message({type: 'error', message: data[\"msg\"]});\n                }\n            }).error(function(){\n                that.$message({type: 'error', message: \"\u670d\u52a1\u5668\u672a\u77e5\u9519\u8bef\uff01\"});\n            });\n\n\n        }\n    }\n})"}, "/web/html/500": {"title": "paas\u5e73\u53f0", "html": "<el-alert :closable=\"false\" description=\"\u670d\u52a1\u5668\u51fa\u73b0\u4e86\u95ee\u9898\uff0c\u8bf7\u8054\u7cfb\u7ba1\u7406\u5458\uff01\" show-icon=\"\" title=\"\u6765\u81eapaas\u5e73\u53f0\u7684\u63d0\u793a\" type=\"error\">\n</el-alert>", "css": "", "js": "app.init({})"}, "/admin/web/apps": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu :default-active=\"tabIndex\" class=\"el-menu-vertical-demo\">\n<el-menu-item @click.native=\"gotoApps\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n<el-menu-item @click.native=\"gotoContainerServer\" index=\"2\"><i class=\"el-icon-menu\"></i>\u96c6\u7fa4</el-menu-item>\n<el-menu-item @click.native=\"gotoSsh\" index=\"4\"><i class=\"el-icon-menu\"></i>ssh</el-menu-item>\n<el-menu-item @click.native=\"gotoShell\" index=\"5\"><i class=\"el-icon-menu\"></i>shell</el-menu-item>\n<el-menu-item @click.native=\"exitAccount\" index=\"3\"><i class=\"el-icon-menu\"></i>\u9000\u51fa\u767b\u5f55</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<p>\n<el-button @click=\"app.goto('/admin/web/new-app');\" type=\"success\">\u6dfb\u52a0\u5e94\u7528</el-button>\n</p>\n<el-card class=\"box-card app_info\" v-for=\"app in apps\">\n<div class=\"text item\">\n<p><strong>\u5e94\u7528Id\uff1a</strong> {{app.appId}}</p>\n<p><strong>\u5e94\u7528\u540d\u5b57\uff1a</strong> {{app.name}}</p>\n<p><strong>\u5e94\u7528\u63cf\u8ff0\uff1a</strong> {{app.desc}}</p>\n<p><strong>\u5e94\u7528\u73af\u5883\uff1a</strong> {{app.image}}</p>\n<p><strong>\u5bb9\u5668\u4e2a\u6570\uff1a</strong> {{app.nums}}</p>\n<p><strong>\u5e94\u7528\u955c\u50cf\uff1a</strong>\n<template v-if=\"app.nowImageName != ''\">{{app.nowImageName}}</template>\n<template v-else=\"\">\u65e0</template>\n</p>\n<p><strong>\u4ee3\u7801\u5730\u5740\uff1a</strong> {{app.git}}</p>\n<p><strong>\u5e94\u7528\u57df\u540d\uff1a</strong>\n<a :href=\"'http://'+app.domain\" target=\"_blank\">http://{{app.domain}}</a>\n</p>\n<p><strong>CPU\uff1a</strong> {{app.cpu}}\u6838</p>\n<p><strong>\u5185\u5b58\uff1a</strong> {{app.memory}}M</p>\n<p>\n<el-button @click=\"optionApp(app.appId);\" type=\"info\">\u64cd\u4f5c\u5e94\u7528</el-button>\n                \u00a0\n                <el-button @click=\"deleteApp(app.appId);\" type=\"danger\">\u5220\u9664\u5e94\u7528</el-button>\n</p>\n</div>\n</el-card>\n<div class=\"block\">\n<el-pagination :current-page=\"currentPage\" :page-size=\"5\" :total=\"totalPage\" @current-change=\"currentChange\" layout=\"prev, pager, next\">\n</el-pagination>\n</div>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": ".app_info{\n    margin-bottom: 20px;\n}", "js": "app.init({\n    api:{\n      url: \"/admin/api/appList\",\n        data: {page: parseInt(app.get_args(\"page\")||1)},\n        success: function(data){\n            return {\n                apps: data[\"data\"].apps,\n                currentPage: parseInt(app.get_args(\"page\") || 1),\n                totalPage: data[\"data\"].nums,\n                tabIndex: \"1\"\n            }\n        },\n        before_success: beforeHandleAjx,\n        error:handleError\n    },\n    data: {\n        apps: [1, 2, 3, 4, 5],\n        currentPage: parseInt(app.get_args(\"page\") || 1),\n    },\n    methods: {\n        currentChange: function(page){\n            app.goto(\"/admin/web/apps?page=\" + page);\n        },\n        optionApp: function(appId){\n            app.goto(\"/admin/web/option-app?appId=\"+appId);\n        },\n        deleteApp: function(appId){\n\n            this.$confirm('\u6b64\u64cd\u4f5c\u5c06\u6c38\u4e45\u5220\u9664\u5e94\u7528, \u662f\u5426\u7ee7\u7eed?', '\u63d0\u793a', {\n                confirmButtonText: '\u786e\u5b9a',\n                cancelButtonText: '\u53d6\u6d88',\n                type: 'warning'\n            }).then(() => {\n\n                    var that = this;\n\n                    $.post(\"/admin/api/deleteApp\", {\"appId\": appId}, function(data){\n\n                        if(data[\"code\"]==0){\n                            that.$message({\n                                type: 'success',\n                                message: data[\"data\"]\n                            });\n                            app.goto(\"/admin/web/apps\");\n                        }else{\n                            that.$message({\n                                type: 'error',\n                                message: data[\"msg\"]\n                            });\n                        }\n\n                    }).error(function(){\n                        that.$message({\n                            type: 'error',\n                            message: data[\"msg\"]\n                        });\n                    });\n            })\n\n        }\n    }\n})"}, "/admin/web/ssh": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu :default-active=\"tabIndex\" class=\"el-menu-vertical-demo\">\n<el-menu-item @click.native=\"gotoApps\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n<el-menu-item @click.native=\"gotoContainerServer\" index=\"2\"><i class=\"el-icon-menu\"></i>\u96c6\u7fa4</el-menu-item>\n<el-menu-item @click.native=\"gotoSsh\" index=\"4\"><i class=\"el-icon-menu\"></i>ssh</el-menu-item>\n<el-menu-item @click.native=\"gotoShell\" index=\"5\"><i class=\"el-icon-menu\"></i>shell</el-menu-item>\n<el-menu-item @click.native=\"exitAccount\" index=\"3\"><i class=\"el-icon-menu\"></i>\u9000\u51fa\u767b\u5f55</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u8bf7\u628a\u4e0b\u9762\u5185\u5bb9\u52a0\u5165git\u670d\u52a1\u5546ssh\u914d\u7f6e</span>\n</div>\n<div class=\"text item\">\n            {{sshContent}}\n        </div>\n</el-card>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": "", "js": "app.init({\n    api: {\n        url: \"/admin/api/getSsh\",\n        success: function (data) {\n            return {\n                \"sshContent\": data['data'],\n                tabIndex: \"4\"\n            }\n        },\n        before_success: beforeHandleAjx,\n        error: handleError\n    }\n});"}, "/admin/web/new-app": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu :default-active=\"tabIndex\" class=\"el-menu-vertical-demo\">\n<el-menu-item @click.native=\"gotoApps\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n<el-menu-item @click.native=\"gotoContainerServer\" index=\"2\"><i class=\"el-icon-menu\"></i>\u96c6\u7fa4</el-menu-item>\n<el-menu-item @click.native=\"gotoSsh\" index=\"4\"><i class=\"el-icon-menu\"></i>ssh</el-menu-item>\n<el-menu-item @click.native=\"gotoShell\" index=\"5\"><i class=\"el-icon-menu\"></i>shell</el-menu-item>\n<el-menu-item @click.native=\"exitAccount\" index=\"3\"><i class=\"el-icon-menu\"></i>\u9000\u51fa\u767b\u5f55</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-form :model=\"form\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u5e94\u7528\u540d\u5b57\">\n<el-input v-model=\"form.name\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u63cf\u8ff0\">\n<el-input type=\"textarea\" v-model=\"form.desc\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u57df\u540d\">\n<el-input v-model=\"form.domain\"></el-input>\n</el-form-item>\n<el-form-item label=\"git\u5730\u5740\">\n<el-input v-model=\"form.git\"></el-input>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668CPU\">\n<el-radio-group v-model=\"form.cpu\">\n<el-radio-button label=\"1\u6838\"></el-radio-button>\n<el-radio-button label=\"2\u6838\"></el-radio-button>\n<el-radio-button label=\"3\u6838\"></el-radio-button>\n<el-radio-button label=\"4\u6838\"></el-radio-button>\n<el-radio-button label=\"5\u6838\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u5185\u5b58\">\n<el-radio-group v-model=\"form.memory\">\n<el-radio-button label=\"64M\"></el-radio-button>\n<el-radio-button label=\"128M\"></el-radio-button>\n<el-radio-button label=\"256M\"></el-radio-button>\n<el-radio-button label=\"512M\"></el-radio-button>\n<el-radio-button label=\"1024M\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u6570\u91cf\">\n<el-input-number :max=\"10\" :min=\"1\" v-model=\"form.nums\"></el-input-number>\n</el-form-item>\n<el-form-item label=\"\u5e94\u7528\u73af\u5883\">\n<el-select placeholder=\"\u8bf7\u9009\u62e9\u5e94\u7528\u73af\u5883\" v-model=\"form.image\">\n<el-option :key=\"item.value\" :label=\"item.label\" :value=\"item.value\" v-for=\"item in options\">\n</el-option>\n</el-select>\n</el-form-item>\n<el-form-item>\n<el-button @click=\"createApp();\" type=\"primary\">\u7acb\u5373\u521b\u5efa</el-button>\n</el-form-item>\n</el-form>\n\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": "", "js": "app.init({\n    api: {\n      url: \"/admin/api/getAddMessage\",\n      success: function(data){\n          return {\n              form: {\n                  image: \"\",\n                  memory: \"64M\",\n                  cpu: \"1\u6838\",\n                  nums: 1\n              },\n              options: data['data'],\n              tabIndex: \"1\"\n          };\n      },\n        before_success: beforeHandleAjx,\n        error:handleError\n    },\n    methods: {\n        createApp: function(){\n            var name = this.form.name;\n            var desc = this.form.desc;\n            var domain = this.form.domain;\n            var git = this.form.git;\n            var cpu = this.form.cpu;\n            var memory = this.form.memory;\n            var nums = parseInt(this.form.nums);\n            var image = this.form.image;\n\n            if(!$.trim(name)){\n                this.$message({\n                    message: '\u5e94\u7528\u540d\u5b57\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(desc)){\n                this.$message({\n                    message: '\u5e94\u7528\u63cf\u8ff0\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(domain)){\n                this.$message({\n                    message: '\u5e94\u7528\u57df\u540d\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(git)){\n                this.$message({\n                    message: '\u4ee3\u7801\u4ed3\u5e93\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            if(!$.trim(image)){\n                this.$message({\n                    message: '\u5e94\u7528\u73af\u5883\u4e0d\u80fd\u4e3a\u7a7a\uff01',\n                    type: 'warning'\n                });\n                return ;\n            }\n\n            var that = this;\n            $.post(\"/admin/api/createApp\", {\n                name: name,\n                desc: desc,\n                domain: domain,\n                git: git,\n                cpu: cpu,\n                memory: memory,\n                nums: nums,\n                image: image\n            }, function(data){\n                if(data[\"code\"] == 0){\n                    that.$message({\n                        message: '\u521b\u5efa\u5e94\u7528\u6210\u529f\uff01',\n                        type: 'success'\n                    });\n                    setTimeout(function(){app.goto(\"/admin/web/apps\");}, 200);\n                }else{\n                    that.$message.error(data[\"msg\"]);\n                }\n            }).error(function(){\n                that.$message.error('\u670d\u52a1\u5668\u53d1\u751f\u4e0d\u53ef\u9884\u6599\u7684\u9519\u8bef\uff01');\n            });\n\n        }\n    }\n})"}, "/admin/web/server": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu :default-active=\"tabIndex\" class=\"el-menu-vertical-demo\">\n<el-menu-item @click.native=\"gotoApps\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n<el-menu-item @click.native=\"gotoContainerServer\" index=\"2\"><i class=\"el-icon-menu\"></i>\u96c6\u7fa4</el-menu-item>\n<el-menu-item @click.native=\"gotoSsh\" index=\"4\"><i class=\"el-icon-menu\"></i>ssh</el-menu-item>\n<el-menu-item @click.native=\"gotoShell\" index=\"5\"><i class=\"el-icon-menu\"></i>shell</el-menu-item>\n<el-menu-item @click.native=\"exitAccount\" index=\"3\"><i class=\"el-icon-menu\"></i>\u9000\u51fa\u767b\u5f55</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-button @click=\"findJoin\" type=\"info\">\u67e5\u770b\u83b7\u53d6\u52a0\u5165\u96c6\u7fa4\u547d\u4ee4</el-button>\n<br><br>\n<template>\n<el-table :data=\"tableData\" border=\"\" style=\"width: 100%\">\n<el-table-column label=\"ID\" prop=\"ID\" width=\"180\">\n</el-table-column>\n<el-table-column label=\"\u4e3b\u673a\u540d\" prop=\"Description.Hostname\" width=\"180\">\n</el-table-column>\n<el-table-column label=\"CPU\">\n<template scope=\"scope\">\n                    {{parseInt(scope.row.Description.Resources.NanoCPUs / 1024/ 1024)}}\n                </template>\n</el-table-column>\n<el-table-column label=\"\u5185\u5b58\">\n<template scope=\"scope\">\n                    {{parseInt(scope.row.Description.Resources.MemoryBytes / 1024/ 1024) + 'M'}}\n                </template>\n</el-table-column>\n<el-table-column label=\"docker\u7248\u672c\" prop=\"Description.Engine.EngineVersion\">\n</el-table-column>\n<el-table-column label=\"\u7c7b\u578b\" prop=\"Spec.Role\">\n</el-table-column>\n<el-table-column label=\"\u72b6\u6001\" prop=\"Spec.Availability\">\n</el-table-column>\n<el-table-column label=\"\u64cd\u4f5c\">\n<template scope=\"scope\">\n<el-button @click=\"deleteNode(scope.$index)\" type=\"danger\" v-if=\"scope.row.Spec.Role != 'manager'\">\u79fb\u9664</el-button>\n<el-button disabled=\"true\" type=\"danger\" v-else=\"\">\u79fb\u9664</el-button>\n</template>\n</el-table-column>\n</el-table>\n</template>\n</br></br>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": "", "js": "app.init({\n    api: {\n      url: \"/admin/api/getContainerServer\",\n      success: function(data){\n          return {\n              \"tableData\": data['data'][\"nodes\"],\n              tabIndex: \"2\",\n              command: data[\"data\"][\"command\"]\n          }\n      },\n        before_success: beforeHandleAjx,\n        error:handleError\n    },\n    data: {\n        \"tableData\": [],\n        tabIndex: \"2\"\n    },\n    methods: {\n        findJoin: function(){\n            this.$alert(this.command, '\u8bf7\u5728\u5bb9\u5668\u670d\u52a1\u5668\u8f93\u5165\u4e0b\u9762\u547d\u4ee4\uff1a', {\n                    confirmButtonText: '\u786e\u5b9a'\n            });\n        },\n        deleteNode: function(index){\n            var nodeName = this.tableData[index][\"ID\"];\n            var that = this;\n            $.post(\"/admin/api/deleteNode\", {\"nodeName\": nodeName}, function(data){\n                if(data[\"code\"] == 0){\n                    that.$message({\n                        message: '\u79fb\u9664\u8282\u70b9\u6210\u529f\uff01',\n                        type: 'success'\n                    });\n                    app.reload();\n                }else{\n                    that.$message.error(data[\"msg\"]);\n                }\n            }).error(function(){\n                that.$message.error('\u670d\u52a1\u5668\u53d1\u751f\u4e0d\u53ef\u9884\u6599\u7684\u9519\u8bef\uff01');\n            });\n        }\n    }\n})"}, "/web/html/404": {"title": "paas\u5e73\u53f0", "html": "<el-alert :closable=\"false\" description=\"\u4f60\u8bbf\u95ee\u7684\u9875\u9762\u4f3c\u4e4e\u662f\u4e0d\u5b58\u5728\u7684\uff01\" show-icon=\"\" title=\"\u6765\u81eapaas\u5e73\u53f0\u7684\u63d0\u793a\" type=\"warning\">\n</el-alert>", "css": "", "js": "app.init({})"}, "/admin/web/option-app": {"title": "paas\u5e73\u53f0", "html": "<el-menu :default-active=\"1\" class=\"el-menu-demo\" theme=\"dark\">\n<el-menu-item index=\"1\">Paas\u7ba1\u7406\u540e\u53f0</el-menu-item>\n</el-menu>\n<br>\n<el-row class=\"tac\">\n<el-col :span=\"4\">\n<el-menu :default-active=\"tabIndex\" class=\"el-menu-vertical-demo\">\n<el-menu-item @click.native=\"gotoApps\" index=\"1\"><i class=\"el-icon-menu\"></i>\u5e94\u7528\u7ba1\u7406</el-menu-item>\n<el-menu-item @click.native=\"gotoContainerServer\" index=\"2\"><i class=\"el-icon-menu\"></i>\u96c6\u7fa4</el-menu-item>\n<el-menu-item @click.native=\"gotoSsh\" index=\"4\"><i class=\"el-icon-menu\"></i>ssh</el-menu-item>\n<el-menu-item @click.native=\"gotoShell\" index=\"5\"><i class=\"el-icon-menu\"></i>shell</el-menu-item>\n<el-menu-item @click.native=\"exitAccount\" index=\"3\"><i class=\"el-icon-menu\"></i>\u9000\u51fa\u767b\u5f55</el-menu-item>\n</el-menu>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n<el-col :span=\"18\">\n\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5e94\u7528\u4fe1\u606f</span>\n</div>\n<div class=\"text item\">\n<p><strong>\u5e94\u7528Id\uff1a</strong> {{appId}}</p>\n<p><strong>\u5e94\u7528\u540d\u5b57\uff1a</strong> {{app.name}}</p>\n<p><strong>\u5e94\u7528\u63cf\u8ff0\uff1a</strong> {{app.desc}}</p>\n<p><strong>\u5e94\u7528\u73af\u5883\uff1a</strong> {{app.image}}</p>\n<p><strong>\u5e94\u7528\u955c\u50cf\uff1a</strong> <template v-if=\"app.nowImageName != ''\">{{app.nowImageName}}</template>\n<template v-else=\"\">\u65e0</template></p>\n<p><strong>\u4ee3\u7801\u5730\u5740\uff1a</strong> {{app.git}}</p>\n<p><strong>\u5e94\u7528\u57df\u540d\uff1a</strong>\n<a :href=\"'http://'+app.domain\" target=\"_blank\">http://{{app.domain}}</a>\n</p>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5bb9\u5668\u4fe1\u606f</span>\n<el-button @click=\"optioneContauiner();\" style=\"float: right;\" type=\"primary\">\u66f4\u65b0</el-button>\n</div>\n<div class=\"text item\">\n<el-form :model=\"container\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u5bb9\u5668\u6570\u91cf\">\n<el-input-number :max=\"10\" :min=\"0\" v-model=\"container.nums\"></el-input-number>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668CPU\">\n<el-radio-group v-model=\"container.cpu\">\n<el-radio-button label=\"1\u6838\"></el-radio-button>\n<el-radio-button label=\"2\u6838\"></el-radio-button>\n<el-radio-button label=\"3\u6838\"></el-radio-button>\n<el-radio-button label=\"4\u6838\"></el-radio-button>\n<el-radio-button label=\"5\u6838\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n<el-form-item label=\"\u5bb9\u5668\u5185\u5b58\">\n<el-radio-group v-model=\"container.memory\">\n<el-radio-button label=\"64M\"></el-radio-button>\n<el-radio-button label=\"128M\"></el-radio-button>\n<el-radio-button label=\"256M\"></el-radio-button>\n<el-radio-button label=\"512M\"></el-radio-button>\n<el-radio-button label=\"1024M\"></el-radio-button>\n</el-radio-group>\n</el-form-item>\n</el-form>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u955c\u50cf\u4fe1\u606f</span>\n</div>\n<div class=\"text item\">\n<div class=\"text item\">\n<el-form :model=\"container\" label-width=\"80px\" ref=\"form\">\n<el-form-item label=\"\u5f53\u524d\u955c\u50cf\">\n                        {{app.nowImageAbout}}\uff0c\u6253\u5305\u4e8e{{app.nowImageCreateTime}}\n                    </el-form-item>\n<el-form-item label=\"\u6700\u8fd1\u955c\u50cf\">\n<el-select @change=\"selectImage\" placeholder=\"\u8bf7\u9009\u62e9\u955c\u50cf\" v-model=\"image\">\n<el-option :key=\"item.value\" :label=\"item.label\" :value=\"item.value\" v-for=\"item in options\">\n</el-option>\n</el-select>\n<el-button @click=\"useImage\" type=\"info\">\u5e94\u7528\u955c\u50cf</el-button>\n<br>\n<template v-if=\"imageMessage != ''\">{{imageMessage}}</template>\n</br></el-form-item>\n<el-form-item label=\"\u6253\u5305\u955c\u50cf\">\n<template>\n<template v-if=\"app.nowImageStatus == 0\">\u672a\u6253\u5305</template>\n<template v-else-if=\"app.nowImageStatus == 1\">\u6253\u5305\u6210\u529f</template>\n<template v-else-if=\"app.nowImageStatus == 2\">\u6253\u5305\u4e2d</template>\n<template v-else-if=\"app.nowImageStatus == 3\">\u6253\u5305\u5931\u8d25</template>\n</template>\n<template>\n<el-button @click=\"buildImage();\" type=\"success\" v-if=\"app.nowImageStatus == 0\">\u4ee3\u7801\u6253\u5305</el-button>\n<el-button @click=\"buildImage();\" type=\"success\" v-else-if=\"app.nowImageStatus == 1\">\u4ee3\u7801\u6253\u5305</el-button>\n<el-button @click=\"buildImage();\" type=\"success\" v-else-if=\"app.nowImageStatus == 3\">\u4ee3\u7801\u6253\u5305</el-button>\n</template>\n</el-form-item>\n</el-form>\n</div>\n</div>\n</el-card>\n<br>\n<el-card class=\"box-card\">\n<div class=\"clearfix\" slot=\"header\">\n<span style=\"line-height: 36px;\">\u5e94\u7528\u65e5\u5fd7</span>\n</div>\n<div class=\"text item log\">\n<el-alert :closable=\"false\" :title=\"'\u3010' + log.time + '\u3011   ' + log.content\" :type=\"log.type\" style=\"margin-top:5px;\" v-for=\"log in logs\">\n</el-alert>\n</div>\n</el-card>\n</br></br></br>\n</el-col>\n<el-col :span=\"1\">\u00a0</el-col>\n</el-row></br>", "css": ".log{\n    max-height: 500px;\n    overflow-y: scroll;\n}", "js": "app.init({\n    api:{\n      url: \"/admin/api/appInfo\",\n      data:{\"appId\": app.get_args(\"appId\")},\n      method: \"POST\",\n      success: function (data) {\n          var obj = data[\"data\"];\n\n          var selectImages = [];\n          for(var i in obj[\"images\"]){\n              var tmpObj = $.parseJSON(obj[\"images\"][i]);\n              selectImages.push({value: i, label: tmpObj[\"imageAbout\"], time: tmpObj[\"imageCreateTime\"], \"image\": tmpObj[\"imageName\"]})\n          }\n\n          return {\n              container: {\"nums\": obj[\"appInfo\"][\"nums\"], cpu: obj[\"appInfo\"][\"cpu\"]+\"\u6838\", \"memory\": obj[\"appInfo\"][\"memory\"]+\"M\"},\n              options: selectImages,\n              image: \"\",\n              app: obj[\"appInfo\"],\n              appId: obj[\"appId\"],\n              logs: obj[\"logs\"],\n              imageMessage: \"\",\n              tabIndex: \"1\"\n          }\n      },\n        before_success: beforeHandleAjx,\n        error:handleError\n    },\n    methods: {\n        optioneContauiner: function(){\n            var that = this;\n            $.post(\"/admin/api/updateAppContainerInfo\", {\n                \"appId\": app.get_args(\"appId\"),\n                \"memory\": that.container.memory.substring(0, that.container.memory.length-1),\n                \"cpu\": that.container.cpu.substring(0, that.container.cpu.length-1),\n                \"nums\": that.container.nums\n            }, function(data){\n                if(data[\"code\"]==0){\n                    that.$message({\n                        message: data[\"data\"],\n                        type: 'success'\n                    });\n                    app.reload();\n                }else{\n                    that.$message({\n                        message: data['msg'],\n                        type: 'error'\n                    });\n                }\n            }).error(function(){\n                that.$message({\n                    message: \"\u670d\u52a1\u5668\u672a\u77e5\u9519\u8bef\uff01\",\n                    type: 'error'\n                });\n            });\n        },\n        buildImage: function(){\n            var that = this;\n            this.$prompt('\u8bf7\u8f93\u5165\u955c\u50cf\u8bf4\u660e', '\u63d0\u793a', {\n                confirmButtonText: '\u786e\u5b9a',\n                cancelButtonText: '\u53d6\u6d88',\n                inputPattern: /\\S/,\n                inputErrorMessage: '\u8bf4\u660e\u4e0d\u80fd\u4e3a\u7a7a\uff01'\n            }).then(({ value }) => {\n                buildImage(that, value)\n            });\n\n        },\n        selectImage: function(i){\n            // \u4fee\u6539\u63d0\u793a\n            this.imageMessage = \"\u955c\u50cf\u6253\u5305\u4e8e\uff1a\" + this.options[i][\"time\"];\n        },\n        useImage: function(){\n            var index = this.image;\n            if(index == \"\"){\n                this.$message({\n                    message: \"\u8bf7\u9009\u62e9\u4e00\u4e2a\u955c\u50cf\uff01\",\n                    type: 'warning'\n                });\n                return;\n            }\n            var data = this.options[index];\n            var that = this;\n\n            $.post(\"/admin/api/useImage\", {\n                \"appId\": app.get_args(\"appId\"),\n                imageName:data[\"image\"],\n                imageTime:data[\"time\"],\n                imageAbout: data[\"label\"]\n            }, function(data){\n                if(data[\"code\"]==0){\n                    that.$message({\n                        message: data[\"data\"],\n                        type: 'success'\n                    });\n                    app.reload();\n                }else{\n                    that.$message({\n                        message: data['msg'],\n                        type: 'error'\n                    });\n                }\n            }).error(function(){\n                that.$message({\n                    message: \"\u670d\u52a1\u5668\u672a\u77e5\u9519\u8bef\uff01\",\n                    type: 'error'\n                });\n            });\n\n        }\n\n    }\n})\n\n\nfunction buildImage(that, v){\n    $.post(\"/admin/api/buildImage\", {\"appId\": app.get_args(\"appId\"), \"imageAbout\": v}, function(data){\n        if(data[\"code\"]==0){\n            that.$message({\n                message: data[\"data\"],\n                type: 'success'\n            });\n            app.reload();\n        }else{\n            that.$message({\n                message: data['msg'],\n                type: 'error'\n            });\n        }\n    }).error(function(){\n        that.$message({\n            message: \"\u670d\u52a1\u5668\u672a\u77e5\u9519\u8bef\uff01\",\n            type: 'error'\n        });\n    });\n}"}};

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
        this.goto(window.location.href);
    }

}

var app = new Applet();


window.addEventListener('popstate', function(event) {
   goto_url(window.location.pathname);
});


// 全局js
function gotoApps(){
    app.goto('/admin/web/apps');
}

function gotoContainerServer() {
    app.goto('/admin/web/server');
}

function gotoSsh() {
    app.goto('/admin/web/ssh');
}

function gotoShell() {
    app.goto('/admin/web/Shell');
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


$(document).ready(function(){
    goto_url(window.location.pathname);
});