
// https://github.com/dbravender/go_mapreduce
// http://dan.bravender.us/2009/11/24/MapReduce_for_Go.html

// http://gcc.gnu.org/svn/gcc/branches/cilkplus/libgo
// http://stackoverflow.com/questions/14142667/reflect-value-mapindex-returns-a-value-different-from-reflect-valueof

package main

import (
    "fmt"
    "wordcount"
    "reflect"
)


// https://github.com/QLeelulu/goku/blob/master/utils/utils.go



// https://github.com/astaxie/gopkg
// http://golang.org/pkg/reflect/#Value
func printWordCount(input interface{}) {

    inputValue := reflect.ValueOf(input)
    if !inputValue.CanSet() && inputValue.Kind() == reflect.Ptr {
        inputValue = inputValue.Elem()
    }

    var keys []reflect.Value = inputValue.MapKeys()

    for keyIndex := 0; keyIndex <len(keys); keyIndex++ {
        key := keys[keyIndex]
        value := inputValue.MapIndex(key)
        fmt.Println(key.String(), value.Int())
    }
}

func main() {

    // interface{} --> map[string]int
    printWordCount(wordcount.Count());
}
