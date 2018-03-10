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
            }
		}

        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        }

		function close() {
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
	<button type="button" onclick="close();">Close</button>
	<button type="button" onclick="reconnect();">Reconnect</button>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(output))
	})
	log.Fatal(http.ListenAndServe(":8002", nil))
}
