package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// Connected clients
var clients = make(map[*websocket.Conn]bool)

// Channel for broadcasting messages
var broadcast = make(chan string)

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket Upgrade Error:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            delete(clients, conn)
            break
        }
        broadcast <- string(msg)
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        fmt.Println("Broadcasting:", msg)

        for client := range clients {
            err := client.WriteMessage(websocket.TextMessage, []byte(msg))
            if err != nil {
                log.Println("Write Error:", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()

    log.Println("Chat server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe Error:", err)
    }
}
