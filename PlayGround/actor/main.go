
// Actor-oriented programming
// http://www.scribd.com/doc/76337990/Concurrency-in-Go

package main

import (
    "fmt"
)

type ReadReq struct {
    key string
    ack chan<- string
}

type WriteReq struct {
    key, val string
}

func main() {

    c := make(chan interface{})

    go func() {
        m := make(map[string]string)
        for {

            switch r := (<-c).(type) {
                case ReadReq:
                    // Read value and write to another channel
                    r.ack <- m[r.key]

                case WriteReq:
                    // m["foo"] = "bar"
                    m[r.key] = r.val
            }
        }
    }()

    // Send WriteReq Object to channel
    // m["foo"] = "bar"
    c <-WriteReq{"foo", "bar"}

    // Send ReadReq Object to channel
    // ack = m["foo"]
    ack := make(chan string)
    c <-ReadReq{"foo", ack}

    fmt.Printf("Got %s", <-ack)
}