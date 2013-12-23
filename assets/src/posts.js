define(function (require, exports, module) {

    var $ = require('jquery');
    var ma = require('mustache');
    var url = require('querystring');

    var tmpl = '{{#posts}}\
     <tr>\
		<td><a href="/_sys/go/show/{{Id}}" target="_blank">{{Title}}</a></td>\
		<td>{{Author}}</td>\
		<td>{{Date}}</td>\
		<td>{{Version}}</td>\
		<td><a href="/_sys/go/edit/{{Id}}" target="_blank">编辑</a>\
            <!--<a href="/_sys/admin/page#path={{Path}}" target="_blank">修改地址</a>\
            <a href="/_sys/admin/page#path={{Path}}" target="_blank">历史</a>-->\
        </td>\
    </tr>{{/posts}}';

    var loading = '<tr><td colspan="6">加载中，请稍后...</td></tr>';
    var empty = '<tr><td colspan="6">没有任何内容</td></tr>';

    var postsTbl = $("#posts").find("tbody")
    var btnPrev = $("#prev");
    var btnNext = $("#next");

    var limit = 20;
    var offset = 0;
    var maxOffset = 0;
    var order = "-Id";

    var dateFormater = function (date, format) {
        var o = {
            "M+": date.getMonth() + 1, //month 
            "d+": date.getDate(),    //day 
            "h+": date.getHours(),   //hour 
            "m+": date.getMinutes(), //minute 
            "s+": date.getSeconds(), //second 
            "q+": Math.floor((date.getMonth() + 3) / 3),  //quarter 
            "S": date.getMilliseconds() //millisecond 
        }
        if (/(y+)/.test(format)) format = format.replace(RegExp.$1,
          (date.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o) if (new RegExp("(" + k + ")").test(format))
            format = format.replace(RegExp.$1,
              RegExp.$1.length == 1 ? o[k] :
                ("00" + o[k]).substr(("" + o[k]).length));
        return format;
    }

    var renderPosts = function (posts) {
        var len = posts.length;
        for (var i = 0; i < len; i++) {
            var date = new Date(posts[i].Date);
            posts[i].Date = dateFormater(date, "yyyy-MM-dd hh:mm:ss");
        }

        var obj = {};
        obj.posts = posts;
        var html = ma.to_html(tmpl, obj);
        postsTbl.html(html);
    }

    var loadPosts = function () {
        postsTbl.html(loading);

        $.ajax({
            url: "/_sys/admin/pages/get",
            data: { offset: offset, limit: limit, order: order },
            dataType: "json",
            cache: false,
            success: function (json) {
                if (json) {
                    renderPosts(json);
                } else {
                    postsTbl.html(empty);
                }
                renderPage();
            }
        });
    }

    var renderPage = function () {
        if (offset <= 0) {
            btnPrev.addClass("disabled");
        } else {
            btnPrev.removeClass("disabled");
        }
        if (offset >= maxOffset - limit) {
            btnNext.addClass("disabled");
        } else {
            btnNext.removeClass("disabled");
        }
    }

    exports.init = function () {
        var href = document.location.href;
        var q = url.parse(href.substring(href.indexOf("#") + 1), sep = '&', eq = '=')
        if (q.offset) {
            offset = parseInt(q.offset);
        }
        if (q.show && q.show == "posts") {
            order = "-Id";
        }

        $.ajax({
            url: "/_sys/admin/paths/max",
            dataType: "json",
            cache: false,
            success: function (json) {
                maxOffset = json.Id - 1;
                loadPosts();

                btnPrev.click(function () {
                    if (offset > 0) {
                        offset = offset - limit;
                        if (offset < 0) {
                            offset = 0;
                        }
                        loadPosts();
                    }
                });
                btnNext.click(function () {
                    if (offset < maxOffset - limit) {
                        offset = offset + limit;
                        loadPosts();
                    }
                });
            }
        });


        var linkPages = $("#showPages");
        var linkLatests = $("#showLatest");
        var linkUsers = $("#showUsers");

        linkPages.click(function () {
            offset = 0;
            order = "-Id";
            loadPosts();
        });
        linkLatests.click(function () {
            offset = 0;
            order = "-Date";
            loadPosts();
        });

        var toc = $("#toc").find("li");
        toc.each(function () {
            var a = $(this);
            a.click(function () {
                toc.removeClass("toc-active");
                a.addClass("toc-active");
            });
        });

    };
});