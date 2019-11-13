# Running GO in browser with WebAssembly
<center>
<img height="100" style="margin:0 20px" src="https://webassembly.org/css/webassembly.svg">
<img height="100" style="margin:0 20px" src="https://camo.githubusercontent.com/98ed65187a84ecf897273d9fa18118ce45845057/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67">
</center>

## What is Web Assembly?
	
WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable target for compilation of high-level languages like C/C++/Rust, enabling deployment on the web for client and server applications.

1. **Binary Instruction Format** : The instruction are converted to a specific binary format.
<img src="https://i.ytimg.com/vi/ukGrEGO_eNA/maxresdefault.jpg">

2. **Stack Based Virtual Machine** : Virtual machines that uses stacks to execute the instructions. Basically, it uses two Pointer a Stack Pointer and an Instruction Pointer. The instruction pointer fetches the instructions and according values get pushed or popped on the stack pointer. 

<img src="https://images2015.cnblogs.com/blog/995377/201702/995377-20170211224627151-676617258.png">



Simply, it let's you run machine code inside the browser. Since we are running machine code we can compile languages like go, rust , c , c++ etc to the machine code and execute them inside the browser. 

Let's look at how we can achieve this.

# Hands On
Create a file called as  `main.go`
```go  
package main

import "fmt"

func main() {
	fmt.Println("Hello wasm")
}

```

Running our program manually.
```sh
go run main.go   
```

Creating a Wasm Exectutable
```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

copy wasm_exec.js to your working directory

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```

4. `index.html`
```html
<html>
    <head>
      <meta charset="utf-8">
      <script src="wasm_exec.js"></script>
      <script>
            async function init() {
                const go = new Go();
                let result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
                go.run(result.instance);

            }
            init();
       </script>
    </head>
    <body></body>
</html>
```

