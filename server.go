// Golang HTML5 Server Side Events Example
//
// Run this code like:
//  > go run server.go
//
// Then open up your browser to http://localhost:8000
// Your browser must support HTML5 SSE, of course.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Create a map `attachedClients`, which is a global variable.
// The keys of the map are the channels over which we can 
// push messages to attached clients.  (The values are just
// booleans and are meaningless.)
//
var attachedClients map[chan string]bool = make(map[chan string]bool)

// Handler for events.
//
func EventHandler(w http.ResponseWriter, r *http.Request) {

	// Make sure that the writer support flushing.
	//
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Add this client to the map of those that should
	// receive updates
	messageChan := make(chan string)
	attachedClients[messageChan] = true
	log.Println("New client attached.")

	// Remove ourselves from the list of attached clients
	// when `EventHandler` exits.
	defer delete(attachedClients, messageChan)

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Don't close the connection, instead loop 100 times,
	// sending messages and flushing the response each time
	// there is a new message to send along.
	// NOTE: we could loop endlessly; however, then you 
	// could not easily detect clients that dettach and the
	// server would continue to send them messages log after
	// they're gone due to the "keep-alive" header.  One of
	// the nifty aspects of SSE is that clients automatically
	// reconnect when they lose their connection.
	//
	for i := 0; i < 10; i++ {

		// Read from our messageChan.
		msg := <-messageChan

		// Write to the ResponseWriter, `w`.
		fmt.Fprintf(w, "data: Message: %s\n\n", msg)

		// Flush the response.  This is only possible if
		// the repsonse supports streaming.
		f.Flush()
	}

	// Log disconnect.
	log.Println("Client disconnected.")
}

// Handler for the main page, which we wire up to the 
// route at "/" below in `main`.
//
func MainPageHandler(c http.ResponseWriter, req *http.Request) {

	// Read out HTML into it
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	// Render the template, writing to `c`
	t.Execute(c, nil)

	// Log that we're done, just for kicks.
	log.Println("Done with MainPageHandler.")
}

// Main routine
//
func main() {

	// Start our goroutine that will send the current time
	// into the messageChan, which is a global variable.
	go func() {

		// Start an infinite loop (no end condition) that
		// sends messages into the messageChan.
		for i := 0; ; i++ {
			message := fmt.Sprintf("%d - the time is %v", i, time.Now())

			// For each attached client, push the new message
			// into the client's message channel.
			for messageChan, _ := range attachedClients {
				messageChan <- message
			}

			// Print a nice log message and sleep for 5s.
			log.Printf("Sent message %d to %d attached clients", i, len(attachedClients))
			time.Sleep(5 * 1e9)

		}
	}()

	// When we get a request at "/", call `MainPageHandler`
	// in a goroutine.
	http.Handle("/", http.HandlerFunc(MainPageHandler))

	// When we get a request at "/event/", call `EventHandler`
	// in a goroutine.
	http.Handle("/events/", http.HandlerFunc(EventHandler))

	// Start the server and listen forever on port 8000.
	http.ListenAndServe(":8000", nil)
}
