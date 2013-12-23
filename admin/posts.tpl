<div id="toc">
  <ul>
    <li class="toc-h1">
      <a href="javascript:">管理后台</a>
    </li>
    <li class="toc-h2 toc-active">
      <a href="#show=posts" id="showPages">文章列表</a>
    </li>
    <li class="toc-h2">
      <a href="#show=latest" id="showLatest">最近更新</a>
    </li>
    <!--
    <li class="toc-h2">
      <a href="#show=users" id="showUsers">用户管理</a>
    </li>-->
  </ul>
</div>
<div id="admin">
  <h1>文章列表</h1>
  <div class="link">
    <a href="/_sys/admin/page" target="_blank">添加新文章</a>
  </div>
  <table id="posts">
    <thead>
      <tr>
        <th>标题</th>
        <th>作者</th>
        <th>编辑时间</th>
        <th>最新版本</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="7">加载中，请稍后...</td>
      </tr>
    </tbody>
  </table>
  <div class="page">
    <ul>
      <li><a href="javascript:" id="prev">上一页</a></li>
      <li><a href="javascript:" id="next">下一页</a></li>
    </ul>
  </div>
  <script>
    seajs.use("/assets/router.js",function(r){r.load("posts")});
  </script>
</div>