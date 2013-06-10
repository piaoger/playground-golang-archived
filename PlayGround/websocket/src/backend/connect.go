
package backend

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "log"
    "net/http"
)

func echo(ws *websocket.Conn) {

    var err error

    for {
        var reply string

        // Receive Message
        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }

        // Print message
        fmt.Println("Received back from client: " + reply)
        msg := "Received:  " + reply
        fmt.Println("Sending to client: " + msg)

        // Send message
        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
}

func Connect() {

    // Register our handlers with the http package.

    // Root handler: Routing Url to ./web/index.html.
    http.Handle("/", http.FileServer(http.Dir("./web")))

    // ECHO Handler:
    http.Handle("/socket", websocket.Handler(echo))

    // Listening to 1234 Port.
    if err := http.ListenAndServe(":1234", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}