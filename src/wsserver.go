package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var (
	js string = `
<script>
(function () {
  if(!WebSocket) {return; }
  var TOP = '_r_s_top', conn = new WebSocket("ws://localhost%s/ws");
  conn.onmessage = function (e) {
    localStorage.setItem(TOP, JSON.stringify([window.scrollX, window.scrollY]));
    location.reload(true);
  };

  window.onload = function () {
    if(localStorage.getItem(TOP)) {
      var d = JSON.parse(localStorage.getItem(TOP));
      window.scrollTo(d[0], d[1]);
    }
    localStorage.removeItem(TOP);
   };
   conn.onopen = function (e) {
     console.log("reload connected");
   };
})();
</script>
`
	activeClients map[*websocket.Conn]int = make(map[*websocket.Conn]int)
	mu            sync.Mutex              = sync.Mutex{}
)

var addr = flag.String("addr", ":3456", "The addr to listen (':3456')")

func changed() {
	mu.Lock()
	fmt.Println(len(activeClients))
	for c, _ := range activeClients {
		c.Write([]byte("changed"))
		delete(activeClients, c)
		c.Close()
	}
	defer mu.Unlock()
}

func main() {

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		changed()
		fmt.Fprintf(w, "ok")
	})
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println("connected")
		mu.Lock()
		defer mu.Unlock()
		activeClients[ws] = 0
	}))

	fmt.Printf(js, *addr)
	err := http.ListenAndServe(*addr, nil)
	fmt.Println(err)
}
