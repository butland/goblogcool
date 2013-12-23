<!DOCTYPE html>
<html>
<head>
    <title>编辑内容</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
    <link rel="stylesheet" href="/assets/css/preview.css" />
    <link rel="stylesheet" href="/assets/css/editor.css" />
    <link rel="stylesheet" href="/assets/css/dialog.css" />
</head>
<body>
    <div id="bar">
        <div id="title">
			<ul id="quickbtn">
				<li><a href="javascript:">加粗</a></li>
				<li><a href="javascript:">斜体</a></li>
				<li><a href="javascript:">标题2</a></li>
				<li><a href="javascript:">列表</a></li>
				<li><a href="javascript:">链接</a></li>
				<li><a href="javascript:">图片</a></li>
				<li><a href="/_sys/admin/help" target="_blank">更多..</a></li>
			</ul>
		</div>
        <div id="control">
            <a href="javascript:" class="button" id="upload_button">
                <span class="label">上传文件</span>
            </a>
            <a href="javascript:" class="button" id="save_button">
                <span class="label">保存</span>
            </a>
        </div>
    </div>
    <div id="container">
        <div class="pane" id="input" placeholder="type some markdown code or drag & drop a .md file here"></div>
        <div id="preview_pane" class="pane">
            <div id="preview"></div>
        </div>
    </div>
	<div id="form" style="display:none;">
		<h2>新建页面</h2>
		<p><input id="postTitle" type="text" value="请在这里键入标题"></p>
		<p><input id="postPath" type="text" value="友好URL(不填写将自动生成)"></p>
	</div>
    <script src="/assets/src/ace.js"></script>
    <script src="http://js.xcai.net/seajs/seajs/1.2.1/sea.js"></script>
	<script>
    seajs.use("/assets/router.js",function(r){r.load("made")});
	</script>	
</body>
</html>

