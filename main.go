package main

import "fmt"

func main() {
	//pointer()

	// Initializing struct
	//fmt.Println(Vertex{1, 2})

	//Access & modify struct
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)
	//fmt.Println(v.Y)

	arrayExperiment()
}

type Vertex struct {
	X int
	Y int
}

// Array has a fixed size
func arrayExperiment() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	preloadedstring := [2]string{"Init", "Value"}
	fmt.Println(preloadedstring)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func pointerV2() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //(*p).X --> can be write as p.X
	fmt.Println(v)
}

func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	fmt.Println(*p) // see the new value of *p (pointer)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
