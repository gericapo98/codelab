package functions_methods

import(
  "fmt"
 // "os"
 // "bufio"
 // "string"
 // "slices"
)
// the function needs derf for passing by ref but method doesn't
// functions are tied to a specific task and are not inherently tied
// to any object or data type
type Vertex struct {
  X, Y, float64
}

func (v *Vertex) Scale(f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){



  v := Vertex{3, 4}
  v.Scale(2)
  ScaleFunc(&v, 10)

  p:= &Vertex{4, 3}
  p.Scale(3)
  ScaleFunc(p,8)

  fmt.Println(v,p)

}
