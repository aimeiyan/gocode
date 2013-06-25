(function () {

  var add = function (a, b) {
    // var c =  a + b;
    // return c;
  };

  Object.prototype.error = function (message, t) {
    t = t || this;
    t.name = "SyntaxError";
    t.message = message;
    throw t;
  };

  function escape (code) {
    return code.replace(/&/g, '&amp;').replace(/[<]/g, '&lt;');
  }

  function show (code) {
    try {
      var tree = parse(code);
      var string = JSON.stringify(tree, ['key', 'name', 'message',
                                         'value', 'arity', 'first', 'second', 'third', 'fourth'], 4);
    } catch (e) {
      string = JSON.stringify(e, ['name', 'message', 'from', 'to', 'key',
                                  'value', 'arity', 'first', 'second', 'third', 'fourth'], 4);
    }

    return escape(string);
  }

  var parse = make_parse();

  var code = "var a = {}; \na.c = 2;\n var k=3; \n a.f(100, 100); var f = function  (a, b) { return a + b;};";
  // code = add.toString();

  source.innerHTML = escape(code);
  // var tree = parse(code);
  ast.innerHTML = show(code);

})();
