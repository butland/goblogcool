<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
    <title>Markdown 快速参考</title>
    <link rel="stylesheet" href="/assets/css/preview.css" />
</head>
<body>
    <div id="preview">
        <h1>快速参考:</h1>

        <h2>文字强调</h2>

        <pre><code>*斜体*   **粗体**
_斜体_   __粗体__
</code></pre>

        <h2>链接</h2>

        <p>内联：</p>

        <pre><code>An [example](http://url.com/ "Title")
</code></pre>

        <p>引用：</p>

        <pre><code>An [example] [id]. Then, anywhere else in the doc, define the link:
[id]: http://example.com/  "Title"
</code></pre>

        <p>上述链接中 <code>Title</code> 是可选项，可以不输入。</p>

        <h2>图片</h2>

        <p>内联:</p>

        <pre><code>![alt text](https://www.google.com/images/srpr/logo3w.png "Title")
</code></pre>

        <p>引用：</p>

        <pre><code>![alt text][id]
[id]: https://www.google.com/images/srpr/logo3w.png "Title"
</code></pre>

        <h2>标题</h2>

        <p>下划线方式:</p>

        <pre><code>Header 1
========

Header 2
--------
</code></pre>

        <p>混合方式:</p>

        <pre><code># Header 1 #

## Header 2 ##

###### Header 6
</code></pre>

        <h2>列表</h2>

        <p>有序列表, 无段落:</p>

        <pre><code>1.  Foo
2.  Bar
</code></pre>

        <p>你可以这样定义缩进:</p>

        <pre><code>*   Abacus
* answer
*   Bubbles
1.  bunk
2.  bupkis
* BELITTLER
3. burper
*   Cunning
</code></pre>

        <h2>横线</h2>

        <pre><code>三个或更多的破折号或星号:
---

* * *

- - - -
</code></pre>

        <h2>分段换行</h2>

        <p>在行尾输入2个以上的空格手动换行:</p>

        <pre><code>Roses are red,
Violets are blue.
</code></pre>

        <p>不输入空格直接换行，最终效果是不会换行，</p>

        <p>分段通过在两行直接插入空白行</p>

        <pre><code>Roses are red.

Violets are blue.
</code></pre>

        <p>换行输出的<code>html</code>代码为<code>&lt;br&gt;</code>，而分段输出为<code>&lt;p&gt;&lt;/p&gt;</code></p>

        <hr>

        <p>ps: 整理这个有点费脑细胞，如果你想更个性话，当然可以直接输入<code>html</code>代码。</p>

        <p>更多内容请参考: <a href="http://wowubuntu.com/markdown/">Markdown 语法说明 (简体中文版) </a>。</p>
    </div>
</body>
</html>
