define(function (require, exports, module) {

    var $ = require('jquery');

    var ma = require('mustache');
    var url = require('querystring');
    var markdown = require('markdown');
    var convertor = new markdown.Converter();

    window.$ = $;

    var FORCE = 1;
    var MODE_FULL = 1;
    var canUpdate = true;
    var emptyTpl = "\n# 请在这里输入文章的标题 \n\n";

    var preview = $("#preview");

    var update = function (mode) {
        if (!mode) mode = 0;
        if (canUpdate || mode === FORCE) {
            var source = $.trim(editor.getSession().getValue());
            if (source.lenght != 0) {
                var html = convertor.makeHtml(source);
                preview.html(html);
            }
        }
        setTimeout(update, 1000);
    }

    var resize = function () {
        var view_h = $(this).height();
        var view_w = $(this).width();
        $('#container').height(view_h - $('#bar').height() - 1);
        $('#container').children('.pane').height(view_h - $('#bar').height - 5);
        $('#input').width(parseInt(view_w / 2) + 10)
        $('#preview_pane').width(parseInt(view_w / 2) - 20);
    }

    exports.init = function () {

        editor.getSession().setValue(emptyTpl);
        editor.getSession().setTabSize(4);
        editor.getSession().setUseSoftTabs(true);
        editor.getSession().setUseWrapMode(true);
        editor.setShowPrintMargin(true);
        document.getElementById('input').style.fontSize = '14px';

        resize();
        $(window).resize(resize);
        update(FORCE);

        $('#preview_pane').hover(function () {
            canUpdate = false;
        }, function () {
            canUpdate = true;
        });

        var href = document.location.href;
        var q = url.parse(href.substring(href.indexOf("?") + 1), sep = '&', eq = '=')
        if (q.path) {

            $.ajax({
              url: "/_sys/admin/page/get",
              data: { path: q.path },
              success: function (data) {
                  if (data.Id) {
                      editor.getSession().setValue(data.Content);
                  }
              },
              dataType: "json",
              cache: false
            });
        }

        $("#save_button").click(function () {
            var text = editor.getSession().getValue();
            var title = $.trim($("#preview").find("h1").text());
            if (title == "") {
                alert("文档必须包含标题，标题以格式 '# 标题' 行来定义")
                return;
            }
            q.title = title;
            q.text = text;
            $.post("/_sys/admin/page/new", q, function (data, x, y) {
                if (data.Path) {
                    alert("发布成功");
                    if (window.opener) {
                        window.opener.location.reload();
                        window.close();
                    } else {
                        document.location.href = data.Path;
                    }
                }
            }, "json");
        });

        require.async(["uploadify.js", "artdialog.js"], function () {
            $("#upload_button").uploadify({
                fileObjName: "file",
                height: 20,
                swf: '/assets/css/uploadify.swf',
                uploader: '/_sys/admin/upload',
                width: 82,
                buttonText: "上传文件",
                debug: false,
                onUploadSuccess: function (file, data, response) {
                    var jsn = eval('(' + data + ')');
                    if (jsn.Id > 0) {
                        if (jsn.Mime.indexOf("image/") >= 0) {
                            editor.insert('![图片名](/file/' + jsn.Id + ')');
                        }
                        else {
                            editor.insert('[附件名](/file/' + jsn.Id + ')');
                        }
                    } else {
                        alert("上传失败");
                    }
                },
                onUploadError: function (file, errorCode, errorMsg, errorString) {
                    alert('文件 ' + file.name + ' 上传失败: ' + errorString);
                }
            });
            /*
            if (!q.path) {
                $.dialog({
                    content: $("#form").html(),
                    lock: true,
                    okValue: '确认',
                    esc: false,
                    ok: function () {
                        this.clost()
                        return false;
                    }
                });
            }*/
        });

        $("#quickbtn").find("a").each(function (idx) {
            $(this).click(function () {
                if (idx == 0) {
                    editor.insert('**加粗的内容**');
                } else if (idx == 1) {
                    editor.insert('_斜体内容_');
                } else if (idx == 2) {
                    editor.insert('\n\n## 标题\n');
                } else if (idx == 3) {
                    editor.insert('\n\n- 列表项目1\n- 列表项目2\n- 列表项目3\n');
                } else if (idx == 4) {
                    editor.insert('[链接文字](http://url.com/)');
                } else if (idx == 5) {
                    editor.insert('![图片名字](http://url.com/)');
                }
            });
        });
    };
});
