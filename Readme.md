# Running GO in browser with WebAssembly

## Intro

`main.go`
```go  
package main

import "fmt"

func main() {
	fmt.Println("Hello wasm")
}

```

1. Running our program manually.
```sh
go run main.go   
```

2. Creating a Wasm Exectutable
```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

3. 

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```

4. `index.html`