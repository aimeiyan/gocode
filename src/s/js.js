(function () {
  var conn = new WebSocket("ws://127.0.0.1:9898/ws");

  conn.onopen = function (e) {
    console.log("open", e);
    conn.send("good, this is send message from client");
  };

  conn.onclose = function (e) {
    console.log("close", e);
  };
  conn.onmessage = function (e) {
    console.log("message", e.data, e);
    add_msg(e.data);
  };

  var $i = $('#i'),
      $history = $('#history');

  var id = 1;

  function add_msg (msg) {
    if(msg) {
      $history.append('<li>' + msg + '</li>');
      $history.find('li:last')[0].scrollIntoView();
    }
  }

  function send_to_server () {
    var msg = $.trim($i.val());
    conn.send(msg);
    $i.val('');
  }

  $('#send').click(send_to_server);

  $i.keyup(function (e) {
    if(e.which === 13) {        // enter
      send_to_server();
    }
  });

  $i.focus();
})();