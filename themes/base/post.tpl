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
<script src="http://js.xcai.net/lib/jquery/1.8.2/jquery.js"></script>
<script src="/assets/scripts/toc.js"></script>
