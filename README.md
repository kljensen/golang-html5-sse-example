Golang HTML5 SSE Example
========================

This is an minimalistic example of how to do
[HTML5 Server Sent Events](http://en.wikipedia.org/wiki/Server-sent_events)
with [Go (golang)](http://golang.org/).  From the server's perspective,
SSE is nearly identical to long polling.  The client makes a GET request
that establishes a TCP connection.  The server keeps this connection open
and sends events to the client when they are available. In this example,
the server pushes a new event every five seconds, consting of a short
message with the current time.  Any number of clients can be
connected: they will all receive the same events if they're connected
concurrently.  (This is achived using Go's channels and a fan-out
pattern.  In other languages you may need to use some kind of pubsub
messaging, like Redis or Zeromq.)

The main advantage of HTML5 SSE over long polling is that there is a nice
API for it in modern browsers, so that you need not use iframes and such.
SSE is easier than Websockets in the sense that it communicates exclusively
over HTTP and therefore does not require a separate server.  Websockets,
however, supports two-way real-time communication between the client and
the server.

## Installing

Check out the repository from GitHub

	git clone https://github.com/kljensen/golang-html5-sse-example

## Running

To run the server, do 

	go run ./server.go

Then point your web browser to `http://localhost:8000`.
You should see output in the browser like the following:

	Yo Duder, here are some facinating messages about the current time:
	Message: 20 - the time is 2013-03-08 21:08:01.260967 -0500 EST
	Message: 21 - the time is 2013-03-08 21:08:06.262034 -0500 EST
	Message: 22 - the time is 2013-03-08 21:08:11.262608 -0500 EST
	Message: 23 - the time is 2013-03-08 21:08:16.263491 -0500 EST
	Message: 24 - the time is 2013-03-08 21:08:21.264218 -0500 EST
	Message: 25 - the time is 2013-03-08 21:08:26.264433 -0500 EST

And, you should see output in the terminal like the following

	2013/03/15 03:55:55 Sent message 0 
	2013/03/15 03:55:55 Broadcast message to 0 clients
	2013/03/15 03:55:55 Added new client
	2013/03/15 03:55:55 Added new client
	2013/03/15 03:55:58 Finished HTTP request at  /
	2013/03/15 03:55:58 Added new client
	2013/03/15 03:56:00 Sent message 1 
	2013/03/15 03:56:00 Broadcast message to 3 clients
	2013/03/15 03:56:05 Sent message 2 
	2013/03/15 03:56:05 Broadcast message to 3 clients
	2013/03/15 03:56:10 Sent message 3 
	2013/03/15 03:56:10 Broadcast message to 3 clients
	2013/03/15 03:56:15 Sent message 4 
	2013/03/15 03:56:15 Broadcast message to 3 clients
	2013/03/15 03:56:20 Sent message 5 
	2013/03/15 03:56:20 Broadcast message to 3 clients

## Thanks

This code is based off of a few sources, mostly

* [Leroy Campbell's SSE example in Go](https://gist.github.com/artisonian/3836281); and,
* the [HTML5Rocks SSE tutorial](http://www.html5rocks.com/en/tutorials/eventsource/basics/).

 
## License (the Unlicense)

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>

