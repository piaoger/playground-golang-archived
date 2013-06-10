
package main

import "fmt"

type Shaper interface {
   Area() int
}

type Rectangle struct {
   length, width int
}

func (r Rectangle) Area() int {
   return r.length * r.width
}

type Square struct {
   side int
}

func (sq Square) Area() int {
   return sq.side * sq.side
}

func main() {
   r := Rectangle{length:5, width:3}
   q := Square{side:5}
   shapesArr := [...]Shaper{r, q}

   fmt.Println("Looping through shapes for area ...")
   for n, _ := range shapesArr {
       fmt.Println("Shape details: ", shapesArr[n])
       fmt.Println("Area of this shape is: ", shapesArr[n].Area())
   }
}