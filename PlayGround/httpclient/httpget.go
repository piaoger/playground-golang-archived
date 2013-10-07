package main

import (
    "os"
    "net/http"
    "log"
    "fmt"
)

func main() {

    // Request
    resp, err := http.Get("http://curl.haxdx.se/docs/thxanks.html")
    if err != nil {
        log.Fatal(err)
    }
    if resp.StatusCode == http.StatusOK {
        fmt.Println(resp.StatusCode)
    }
    defer resp.Body.Close()

    // Open a File
    fd, err := os.OpenFile("result.html", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }
    defer fd.Close()

    // Write request body to "result.html"
    var buf [512]byte
    for {
            nr, _ := resp.Body.Read(buf[:])
            if nr == 0 {
                break
            }

            if nw, err := fd.Write(buf[0:nr]); nw != nr {
                log.Fatal(err)
            }
    }


    // https://github.com/lifelog/lg/blob/master/request.go
}
