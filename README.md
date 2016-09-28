# Weblinks

Create pagination links for REST endpoints in accordance with RFC5988 for GO.
If you are writing a REST API then use this library to create your "Link" headers.

## Usage

### `New(url string, page, pageSize, totalCount int)`

```go
package main

import "github.com/blockloop/weblinks-go"

func main() {
    wl, err := weblinks.New("http://www.google.com", 2, 10, 100)
    wl.Self.Query().Get("page") // 2
    wl.Next.Query().Get("page") // 3
    wl.Prev.Query().Get("page") // 1
    wl.Last.Query().Get("page") // 10
    wl.First.Query().Get("page") // 10
  
    wl, err := weblinks.New("http://www.google.com", 1, 10, 100)
    wl.Prev // nil
}
```

### `Parse(header string)`

Use Parse to parse a Link header on the client side.

```go
	wl, err := weblinks.Parse(req.Header.Get("Link"))
```
