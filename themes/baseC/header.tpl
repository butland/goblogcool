<header>
    <div class="nav">
        <div class="link">
        {{if .CanBack}}
            <a href="javascript:" onclick="history.go(-1)" title="返回" class="active"><span>&lt;</span>返回</a>
            <a href="/" title="首页">首页</a>
        {{else}}
            <a href="/" title="首页" class="active">首页</a>
        {{end}}
        {{if .Logined}}
            <a href="/_sys/admin" title="管理">管理</a>
            <a href="{{.LogoutUrl}}" title="注销 {{.User}}">注销</a>
        {{else}}
            <a href="{{.LoginUrl}}" title="登录">登录</a>
        {{end}}
        </div>
        <div class="page">
            <a href="/about" title="关于我们">关于我们</a>
        </div>
    </div>
    <div>
         <p class="title"> docooler's Blog</p>
         <p>万物之始，大道至简，衍化至繁。</p>
    </div>
</header>
