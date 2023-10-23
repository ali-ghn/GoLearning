package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// http.Get returns a pointer to a Response struct
	// Response struct contains the response body, headers, and status code
	// what we are intrested here is response body
	// Response body is actually an io.ReadCloser which is an interface that implements
	// Read and Close interfaces which are used to read and close the response body.
	// Here interfaces get a little complicated but the main reason we use nested interfaces
	// is to add flexibility to our code and also to make it more readable.
	// For instance Read interface implements a function called Read with this signature
	// func (r io.Reader) Read(b []byte) (n int, err error)
	// therefore this function can be used to read the response body with any kind of
	// data type from the serevr as it is a general data type.

	r, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// closes the response body after we are done with it
	defer r.Body.Close()

	// io.Copy is a function with this signature
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// we are familiar with Reader interface but we dont know what is the Writer interface
	// an implementation of Writer interface is any type that implements Write method
	// which has this signature func (w io.Writer) Write(p []byte) (n int, err error)
	// the io.Copy function copies whatever is in the response body to the os.Stdout

	// io.Copy(os.Stdout, r.Body)

	// After implementing our own implementation of writer we can use it in our own fashion way
	lw := logWriter{}
	io.Copy(lw, r.Body)

	// So with that in mind, you actially can implement any kind of interfaces to do what you want,
	// that is the whole purpose of interfaces.
	// For instance you can implement a writer interface to write to a file or to a database or to a socket.
}

// Using the felexibility of nested interfaces we can implement a type which implements
// Writer interface and then pass it to io.Copy function.
type logWriter struct{}

// Write method of Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	// we are using the fmt package to print the bytes to the console
	fmt.Println(string(bs))

	// just to make sure it is our own code
	fmt.Printf("Writing %d bytes to log\n", len(bs))

	// we are returning the number of bytes written and nil error
	return len(bs), nil
}
