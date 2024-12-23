/*
*  An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
Note: There is an error in the example code on line 22. Vertex (the value type) 
doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type). 
* */

/*
* Encapsulation of behavior: Interfaces allow you to define how something should behave.
* I.e which methods are available wihtout sepcifying how these methods are implemented.
*
* Polymorphism. Same as in Java, GO's type satisfy an interface if the required mehtods
* are found on the type, meaning any type that has the methods <any in {}> can be treated 
* as Abser
*
* You can decouple against an interface. See example below.
*
* The examples below and the explaination above are inspired from :
* https://go.dev/tour/methods/9
*/

// In Go, interfaces can hold a pointer to the actual struct implementing the interface.
package interface_testing

import (
	"fmt"
	"math"
)

// Interfaces in Go define what methods a type must have to “satisfy” that interface.
// Once you have something in your code that behaves just like Abser, you can pass it 
// and it will just work.
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//
type Vertex struct {
  X, Y float64
}

func main() {

	var a Abser

	f := MyFloat(-math.Sqrt2)
  // every MyFloat is an Abser, because MyFloat is receiver of Abs method, 
  // which is the only method in Abser
  a = f

  // -------------------------------------------------------------------------------

	v := Vertex{3, 4}
  // every *Vertex is an Abser, because *Vertex is receiver of Abs method
  // a is implicitly of type Abser, because Abser decouples the type from 
  // the implementation, which is the Abs method in this case.
  // This can appear in any package without prearrangement.(not knowing at compile time)
  // 
  // compile time
  // It's the process of converting the code into machine code.
  // Code is run at runtime. Runtime is after the code is executed.

	a = &v
  
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}


