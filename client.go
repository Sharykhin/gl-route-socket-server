package main

import (
	"log"
	"net/http"
)

var output = `
<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:1234";

        window.onload = function() {

            console.log("onload");

            _init();
        };

		function _init() {
			sock = new WebSocket(wsuri);

            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }

            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
				var a = document.querySelector('#d').innerHTML;
				a += e.data + '<br/>';
				document.querySelector('#d').innerHTML = a;
            }
		}

        function send() {
            var msg = document.getElementById('message').value;
			var m = JSON.stringify({action: "message", "user_id":"12", payload:{text:msg}})
            sock.send(m);
        }

		function ccc() {
			console.log('close the connection')
			sock.close();
		}

		function reconnect() {
			_init();
		}

    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
	<button type="button" onclick="ccc();">Close</button>
	<button type="button" onclick="reconnect();">Reconnect</button>
	<div id="d" style="border: 1px solid #ccc;"></div>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(output))
	})
	log.Fatal(http.ListenAndServe(":8002", nil))
}
