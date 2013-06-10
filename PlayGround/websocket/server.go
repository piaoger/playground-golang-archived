
// Materials:
//     https://github.com/ukai/go-websocket-sample/blob/master/websocket_echo_sample.go
//     http://music.573114.com/Blog/Html/6A36/665685.html
//     http://www.websocket.org/echo.html

// Questions:
//  1. Why I have to place backend.go in %GOPATH%\src? If not, a "import "backend": cannot find package" will be prompted.
//     Go code must be kept inside a workspace. A workspace is a directory hierarchy with three directories at its root:
//       src contains Go source files organized into packages (one package per directory),
//       pkg contains package objects, and
//       bin contains executable commands.
//     Reference: [How to Write Go Code](http://golang.org/doc/code.html)

//  2. Why below error:
//         # command-line-arguments
//         cannot create server.exe
//     Reason: server.exe is read only.

package main

import (
    "fmt"
    "backend"
)

func main() {

    fmt.Println("begin")

    backend.Connect();

    fmt.Println("end")
}