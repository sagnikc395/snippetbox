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
