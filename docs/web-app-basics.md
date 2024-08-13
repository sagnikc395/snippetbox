## Web App Basics

1. Handler -> like controllers. They are responsible for executing our application logic and for writing HTTP response headers and bodies.
2. Router -> servemux in Go . This stores a mapping bw the URL patterns in our application and the corresponding handlers. Usually we have 1 servemux for our application containing all of our routes.
3. Web Server -> Built into the language and can establish a web server and listen for incoming requests as part of the application itself. Dont need a external third-party server like Nginx or Apache.
