# Revoltgo

Revoltgo is a go package for writing bots / self-bots in revolt easily.

**NOTE**: This package is still under development and not finished. Create an issue if you found a bug.

## Features

- Can listen an event multiple times
- Easy to use
- Supports self-bots
- Simple cache system

## Installation

- Create a new project and init go.mod file. `go mod init example`
- Install the package using `go get github.com/5elenay/revoltgo`

## API Reference

Click [here](https://pkg.go.dev/github.com/5elenay/revoltgo) for api reference.

## Ping Pong Example (Bot)

```go
package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/5elenay/revoltgo"
)

func main() {
    // Init a new client.
    client := revoltgo.Client{
        Token: "bot token",
    }

    // Listen a on message event.
    client.OnMessage(func(m *revoltgo.Message) {
        if m.Content == "!ping" {
            sendMsg := &revoltgo.SendMessage{}
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

```

## Ping Pong Example (Self-Bot)

```go
package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/5elenay/revoltgo"
)

func main() {
    // Init a new client.
    client := revoltgo.Client{
        SelfBot: &revoltgo.SelfBot{
            Id:           "session id",
            SessionToken: "session token",
            UserId:       "user id",
        },
    }

    // Listen a on message event.
    client.OnMessage(func(m *revoltgo.Message) {
        if m.Content == "!ping" {
            sendMsg := &revoltgo.SendMessage{}
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

```
