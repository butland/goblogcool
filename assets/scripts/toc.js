/*!
  * jquery.toc.js - A jQuery plugin that will automatically generate a table of contents.
  * v0.0.3
  * https://github.com/jgallen23/toc
  * copyright JGA 2012
  * MIT License
  */
!function(e){e.fn.toc=function(t){var n=this,r=e.extend({},jQuery.fn.toc.defaults,t),i=e(r.container),s=e(r.selectors,i),o=[],u=r.prefix+"-active",a=function(t){for(var n=0,r=arguments.length;n<r;n++){var i=arguments[n],s=e(i);if(s.scrollTop()>0)return s;s.scrollTop(1);var o=s.scrollTop()>0;s.scrollTop(0);if(o)return s}return[]},f=a(r.container,"body","html"),l=function(t){if(r.smoothScrolling){t.preventDefault();var i=e(t.target).attr("href"),s=e(i);f.animate({scrollTop:s.offset().top},400,"swing",function(){location.hash=i})}e("li",n).removeClass(u),e(t.target).parent().addClass(u)},c,h=function(t){c&&clearTimeout(c),c=setTimeout(function(){var t=e(window).scrollTop(),i;for(var s=0,a=o.length;s<a;s++)if(o[s]>=t){e("li",n).removeClass(u),i=e("li:eq("+(s-1)+")",n).addClass(u),r.onHighlight(i);break}},50)};return r.highlightOnScroll&&(e(window).bind("scroll",h),h()),this.each(function(){var t=e(this),n=e("<ul/>");s.each(function(i,s){var u=e(s);o.push(u.offset().top-r.highlightOffset);var a=e("<span/>").attr("id",r.anchorName(i,s,r.prefix)).insertBefore(u),f=e("<a/>").text(u.text()).attr("href","#"+r.anchorName(i,s,r.prefix)).bind("click",function(n){l(n),t.trigger("selected",e(this).attr("href"))}),c=e("<li/>").addClass(r.prefix+"-"+u[0].tagName.toLowerCase()).append(f);n.append(c)}),t.html(n)})},jQuery.fn.toc.defaults={container:"body",selectors:"h1,h2,h3",smoothScrolling:!0,prefix:"toc",onHighlighted:function(){},highlightOnScroll:!0,highlightOffset:100,anchorName:function(e,t,n){return n+e}}}(jQuery)

$(document).ready(function(){
    $('#toc').toc({
        'selectors': 'h1,h2,h3', //elements to use as headings
        'container': 'body', //element to find all selectors in
        'smoothScrolling': true, //enable or disable smooth scrolling on click
        'prefix': 'toc', //prefix for anchor tags and class names
        'onHighlight': function (el) { }, //called when a new section is highlighted
        'highlightOnScroll': true, //add class to heading that is currently in focus
        'highlightOffset': 100, //offset to trigger the next headline
        'anchorName': function (i, heading, prefix) { //custom function for anchor name
            return prefix + i;
        },
        'headerText': function (i, heading, $heading) { //custom function building the header-item text
            return $heading.text();
        }
    });
    document.title = $("h1").text();
})