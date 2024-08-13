## Web App Basics

1. Handler -> like controllers. They are responsible for executing our application logic and for writing HTTP response headers and bodies.
2. Router -> servemux in Go . This stores a mapping bw the URL patterns in our application and the corresponding handlers. Usually we have 1 servemux for our application containing all of our routes.
3. Web Server -> Built into the language and can establish a web server and listen for incoming requests as part of the application itself. Dont need a external third-party server like Nginx or Apache.

- Go's servemux treats URL pattern like / like a catch all. All HTTP requests will be handled by the home function regardless of their URL Path.

### Network Addresses

- TCP Network address that we pass onto http.ListenAndServe() should be of the format host:port.
- If we omit the host , then the server will listen on all of our computer's available network interfaces.
- Can also use named ports like :http, :http-alt instead of a number. If we are using a named port then Go, will attempt to look up the relevant port number from the /etc/services file when starting the server , or will return a error if a match can't be found.

### Logging

- using Go's built in logger to output log messages.
- I.e prefixes messages with the local date and time and writes them to the standard error stream.
- log.Fatal() will also call os.Exit(1) after witing the message, causing the application to immediately exit.
-
