package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/ben-forster/revolt"
)

func main() {
    // Init a new client.
    client := revolt.Client{
        Token: "bot token",
    }

    // Listen a on message event.
    client.OnMessage(func(m *revolt.Message) {
        if m.Content == "!ping" {
            sendMsg := &revolt.SendMessage{}
            sendMsg.SetContent("🏓 Pong!")

            m.Reply(true, sendMsg)
        }
    })

    // Start the client.
    client.Start()

    // Wait for close.
    sc := make(chan os.Signal, 1)

    signal.Notify(
        sc,
        syscall.SIGINT,
        syscall.SIGTERM,
        os.Interrupt,
    )
    <-sc

    // Destroy client.
    client.Destroy()
}
