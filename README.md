# revgo

### This is a fork of the original to fix some bugs and add features.

revgo is a go package for writing bots / self-bots in revolt easily.

**NOTE**: This package is still under development and not finished. Create an issue if you found a bug.

**NOTE 2**: This package requires go 1.17.

## Features

- Multiple event listen
- Easy to use
- Supports self-bots
- Simple cache system

## Installation

- Create a new project and init go.mod file. `go mod init example`
- Install the package using `go get github.com/itzTheMeow/revolt-go`

## API Reference

Click [here](https://pkg.go.dev/github.com/5elenay/revgo@v0.3.1) for api reference.

## Ping Pong Example (Bot)

```go
package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/5elenay/revgo"
)

func main() {
    // Init a new client.
    client := revgo.Client{
        Token: "bot token",
    }

    // Listen a on message event.
    client.OnMessage(func(m *revgo.Message) {
        if m.Content == "!ping" {
            sendMsg := &revgo.SendMessage{}
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

    "github.com/5elenay/revgo"
)

func main() {
    // Init a new client.
    client := revgo.Client{
        SelfBot: &revgo.SelfBot{
            Id:           "session id",
            SessionToken: "session token",
            UserId:       "user id",
        },
    }

    // Listen a on message event.
    client.OnMessage(func(m *revgo.Message) {
        if m.Content == "!ping" {
            sendMsg := &revgo.SendMessage{}
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
