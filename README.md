# Golang WebSocket Chat App

This is a simple chat application written in Golang using WebSockets. 
It allows multiple users to chat with each other in real-time.

## Features

- Real-time messaging using WebSockets
- Clean and simple user interface
- Support for multiple users

---

#### Chat Web
![Alt text](.github/img/chat-web-websocket.png?raw=true "Chat Web")

#### Code Log
![Alt text](.github/img/chat-web-websocket-2.png?raw=true "Code Log")

---

## Getting Started
### To run the server:

```bash
go run main.go
```

### To package the server:

```bash
go build ./cmd/web/*.go
```

### Using docker
```bash
docker build --tag go-chat:1 .
docker run -p 5000:5000 go-chat:1
```

