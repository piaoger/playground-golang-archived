

package main

import (
    "fmt"
    "time"
)

func main() {

    // Create a new synchronous chan of string
    ch := make(chan string)

    // Create two goroutines.
    go sendData(ch)
    go receiveData(ch)

    // Mimics thread.join
    // sleep works with a Duration in nanoseconds (ns) !
    time.Sleep(1e9)
}

func sendData(ch chan string) {
    // Sending data to a channel
    ch <- "One"
    ch <- "Two"
    ch <- "Three"
    ch <- "Four"
    ch <- "Five"
}

func receiveData(ch chan string) {
    // Receiving data from a channel
    var input string
    for {
        input = <-ch
        fmt.Printf("%s ", input)
    }
}