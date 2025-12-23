

# gRPC with Golang

## Introduction

This project demonstrates how to use **gRPC with Golang** to implement all four types of gRPC communication patterns:

1. Unary RPC
2. Server Streaming RPC
3. Client Streaming RPC
4. Bidirectional Streaming RPC

The project acts as a **gRPC client** that communicates with a `GreetService` defined using Protocol Buffers. It sends greeting requests and receives responses using different streaming models.

---

## Table of Contents

* [Introduction](#introduction)
* [Project Structure](#project-structure)
* [gRPC Communication Types](#grpc-communication-types)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
* [Features](#features)
* [Dependencies](#dependencies)
* [Configuration](#configuration)
* [Troubleshooting](#troubleshooting)
* [License](#license)

---

## Project Structure

```text
.
├── client/
│   └── bi_stream.go
│   └── client_stream.go
│   └── main.go
│   └── server_stream.go
│   └── unary.go
├── proto/
│   └── greet.pb
│   └── greet.proto
│   └── greet_grpc.pb
├── server/
│   └── bi_stream.go
│   └── client_stream.go
│   └── main.go
│   └── server_stream.go
│   └── unary.go
└── go.mod
└── go.sum
└── README.md
```

---

## gRPC Communication Types

### 1. Unary RPC

A single request sent by the client and a single response returned by the server.

Implemented in:

* `callSayHello` (`unary.go`)

---

### 2. Server Streaming RPC

The client sends one request and receives a stream of responses from the server.

Implemented in:

* `callSayHelloServerStream` (`server_stream.go`)

---

### 3. Client Streaming RPC

The client sends a stream of requests and receives a single response from the server.

Implemented in:

* `callSayHelloClientStream` (`client_stream.go`)

---

### 4. Bidirectional Streaming RPC

Both client and server send streams of messages independently.

Implemented in:

* `callHelloBidirectionalStream` (`bi_stream.go`)

---

## Prerequisites

* Go 1.18+
* Protocol Buffers compiler (`protoc`)
* gRPC Go plugins
* A running gRPC server implementing `GreetService`

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/gRPC-with-golang.git
   cd gRPC-with-golang
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Generate gRPC code from `.proto` files (if not already generated):

   ```bash
   protoc --go_out=. --go-grpc_out=. proto/greet.proto
   ```

---

## Usage

1. Ensure the gRPC server is running on:

   ```
   localhost:8080
   ```

2. Run the client:

   ```bash
   go run .
   ```

3. Enable the desired RPC call in `main.go` by uncommenting one of the following:

```go
// callSayHello(client)
callSayHelloServerStream(client, names)
callSayHelloClientStream(client, names)
callHelloBidirectionalStream(client, names)
```

---

## Features

* Demonstrates all major gRPC streaming patterns
* Clean separation of streaming logic
* Uses insecure credentials for local development
* Easy to extend with additional RPC methods

---

## Dependencies

Key dependencies used in this project:

* `google.golang.org/grpc`
* `google.golang.org/grpc/credentials/insecure`
* `google.golang.org/protobuf`

---

## Configuration

* **Port:** `:8080`
* **Transport Security:** Insecure (for local development only)

```go
grpc.WithTransportCredentials(insecure.NewCredentials())
```

⚠️ For production use, replace insecure credentials with TLS.

---

## Troubleshooting

**Connection refused**

* Ensure the gRPC server is running
* Verify port `8080` is open

**Proto-related errors**

* Regenerate protobuf files
* Ensure `proto` package path matches imports

**EOF errors**

* Normal behavior indicating stream completion

---

## License

This project is licensed under the MIT License.
You are free to use, modify, and distribute it.

--
