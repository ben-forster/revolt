[![Go Reference](https://pkg.go.dev/badge/github.com/ben-forster/revolt.svg)](https://pkg.go.dev/github.com/ben-forster/revolt)
[![Go Report](https://goreportcard.com/badge/github.com/ben-forster/revolt)](https://goreportcard.com/report/github.com/ben-forster/revolt)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ben-forster/revolt?filename=go.mod)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-MIT%202.0-blue.svg)](https://github.com/ben-forster/revolt/blob/master/LICENSE)
[![Revolt.go Version](https://img.shields.io/github/v/release/ben-forster/revolt?label=release)](https://github.com/ben-forster/revolt/releases/latest)

# Revolt.go

Revolt.go is a Go package for writing bots and self-bots for Revolt.

## Features
- Event listener.
- Easy to use.
- Support for self bots.
- Simple caching system.

## Getting Started

### Installation
- Create a new project and initiate the go.mod file. `go mod init example`
- Install the package using `go get github.com/ben-forster/revolt`
- Create your main bot file. `touch main.go`

## API Reference
Click [here](https://pkg.go.dev/github.com/ben-forster/revolt@v0.0.1) for the library's API reference.

Official documentation will come in the near future.

## Notice

Please note that you will need **Go 1.21** to use this library.

This package is still under development and while you can create a working bot, the library is not finished. You can see a development roadmap [here](https://github.com/users/ben-forster/projects/8). Please create an issue if you would like to contribute.

## Ping Pong Example (Bot)

```go
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

```

## Ping Pong Example (Self-Bot)

```go
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
        SelfBot: &revolt.SelfBot{
            Id:           "session id",
            SessionToken: "session token",
            UserId:       "user id",
        },
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

```

## Credit

This project is a mantained and re-worked version of 5elenay's library [revoltgo](https://github.com/5elenay/revoltgo).