# WebSocket Mock Server


## Purpose

---

This project is a runs a simple HTTP server that allows to test client applications that needs to consume WebSocket.

For the current version you can customize it through 2 environment variables :

- **$MOCK_MESSAGE** : define the message that the server will send to the client over WebSocket protocol _it has a default value of "Default Message from Go Websocket Mock"_
- **$MESSAGE_INTERVAL** : define the interval (in seconds) that the server will send to the client. _its default value is 5 seconds_

---


## Install

---

You can easily run the dockerfile and run it. It will expose the port 8085

First build the image :

```shell
docker build -t jossefaz/web-socket-mock .   
```

Then run it 

```shell
docker run -p 8089:8085 jossefaz/web-socket-mock:latest   
```

If you want to customize the interval and the message

```shell
docker run -p 8089:8085 --env MESSAGE_INTERVAL=2 --MOCK_MESSAGE=[{"test": "ok"}] yossefaz/web-socket-mock:latest 
```