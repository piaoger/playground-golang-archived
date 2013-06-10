package main

import "fmt"

func main() {

    var isError = true
    var number = 100
    if isError == true && number == 100 {
        fmt.Printf("hello, world\n")
        fmt.Printf("number %d", number)
     }

    // Array
    var (
        arrary    = [5]int{18, 20, 15, 22, 16}
        // var lazyArray = []int{5, 6, 7, 8, 22}
        lazyArray = [...]int{5, 6, 7, 8, 22}
    )
    fmt.Printf("Length: array %d ; lazyArray %d\n", len(arrary), len(lazyArray))

    //var arrKeyValue = []string{3: "Chris", 4: "Ron"}
    var keyvalues = [5]string{3: "Chris", 4: "Ron"}
    for i := 0; i < len(keyvalues); i++ {
        fmt.Printf("Person at %d is %s\n", i, keyvalues[i])
    }

    namevalues := map[string]string{
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }

    for key, value := range namevalues {
        fmt.Printf("namevalues at %s is %s\n", key, value)
    }

}