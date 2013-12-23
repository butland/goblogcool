(function () {
    var version = '1201005';
    var debug = false;
    var alias = {
        'mustache': 'http://js.xcai.net/seajs/mustache/0.5.0/mustache.js',
        'jquery': 'http://js.xcai.net/seajs/jquery/1.8.1/jquery.js',
        'querystring': 'http://js.xcai.net/seajs/querystring/1.0.2/querystring.js'
    };
    var base = 'http://' + location.host + '/assets/src/';
    var map_debug = [[/^(.*\/assets\/.*\.(?:css|js))(?:.*)$/i, "$1?" + version]];
    seajs.config({ alias: alias, preload: ['jquery'], base: base, debug: debug, map: map_debug });

    define(function (require, exports) {
        exports.load = function (file) {
            var path = './src/' + file;
            require.async(path, function (mod) {
                if (mod && mod.init) { mod.init(); }
            });
        };
    });
})();