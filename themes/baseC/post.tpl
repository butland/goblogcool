<div id="toc"></div>
<article>
    {{.Content}}
</article>
<hr>
<div>
    <p><span>作者：</span>{{.Author}} <span>时间：</span>{{.Date.Format "2006-01-02 15:04"}} <span>该页版本：</span>{{.Version}} </p>
</div>
<div>
    <a href="/_sys/admin/page?path={{.Path}}" target="_blank" rel="nofollow">编辑该页</a> |
    <a href="{{.Path}}?raw=yes&ver={{.Version}}" target="_blank" rel="nofollow">源文件</a>
</div>
<div class="ds-thread" data-thread-key="{{.Path}}" data-title="{{.Title}}" data-author-key="{{.Author}}" data-url=""></div>
<script src="http://js.xcai.net/lib/jquery/1.8.2/jquery.js"></script>
<script src="/assets/scripts/toc.js"></script>
<script type="text/javascript">
	var duoshuoQuery = {short_name:"YOUR-SITE-NAME"};
	(function() {
		var ds = document.createElement('script');
		ds.type = 'text/javascript';ds.async = true;
		ds.src = 'http://static.duoshuo.com/embed.js';
		ds.charset = 'UTF-8';
		(document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(ds);
	})();
</script>