## Routing Requests

- Adding a couple more routes so that the application starts to shape up as:

```
/ -> home -> display the home page
/snippet -> showSnippet -> display a specific snippet
/snippet/create -> createSnippet -> create a new snippet
```

### Fixed Path and Subtree Patterns

- Go's servemux supports 2 different types of URL patterns ; fixed paths and subtree paths. Fixed paths dont end with a trailing slash, whereas subtree paths do end with a trailing slash.In Go's servemux, fixed path patterns like these are only matched -> when the request URL path exactly matches the fixed path.

- / is an example of subtree path. Subtree path patterns are matched whenever the start of a request URL path matches the subtree path.
- / pattern is acting like a catch all.

### Restricting the Root URL Pattern

- If the application we are building we want the home page to be displayed if - and only if - the request URL path matches exactly / , else we want the user to receive a 404 page not found response.

- can add a condition in the home handler which have the same effect.

### Default Serve Mux

- http.Handle() and http.HandleFunc() functions allow us to register routes without declaring a servemux.
- behind the scenes , these functions will register their routes with DefaultServeMux. There is nothing special about this -> like a regular servemux like we have aready been using, but which is initialized by default and stored in a net/http global variable.

- Issue is here if we use defaultservemux is a global variable, any package can access it and register a route -> including any third-party packages that our application imports. If any one of these third-party packages is compromised, they could use DefaultServerMux to expose a malicious handler to the web.

- In Go's servemux longer URL patterns always take precedence over the shorter ones. So if a servemux contains mutliple patterns which match a request, it will always dispatch the request to the handler corresponding to the longest pattern. This has the nice side-effect that we can register patterns in any order and it won't change how the servemux behaves.

- Request URL paths are automatically sanitized. If the request path contains any . and .. elements or repeated slashes, it will automatically redirect the user to an equivalent clean URL.

- If a subtree path has been registered and a request is received for that subtree path without a trailing slash, then the user will automatically be sent a 301 permanent redirect to the subtree path with the slash added.
