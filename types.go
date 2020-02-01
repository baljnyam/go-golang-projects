// package main

// func main() {
// pointer
// i, j := 42, 2701

// p := &i  // point to i
// fmt.Println(*p) // read i through the pointer = 42
// *p = 21 // set i through the pointer
// fmt.Println(i) // see the new value of i = 21

// p = &j // point to j
// *p = *p / 37 // divide j through the pointer
// fmt.Println(j) // see the new value of j = 73

// }

//structs
// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	fmt.Println(Vertex{1, 2}) // {1 2}
// }

// Pointers to structs

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1e9
// 	fmt.Println(v) // {1000000000 2}
// }

//  Struct Literals

// type Vertex struct {
// 	X, Y int
// }

// var (
// 	v1 = Vertex{1, 2}  // has type Vertex
// 	v2 = Vertex{X: 1}  // Y:0 is implicit
// 	v3 = Vertex{}      // X:0 and Y:0
// 	p  = &Vertex{1, 2} // has type *Vertex
// )

// func main() {
// 	fmt.Println(v1, p, v2, v3)  // {1 2} &{1 2} {1 0} {0 0}
// }

// Arrays

// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a[0], a[1]) // Hello World
// 	fmt.Println(a) // [Hello World]

// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes) // [2 3 5 7 11 13]
// }

// Slices

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

// 	var s []int = primes[1:4]
// 	fmt.Println(s) //[3 5 7]
// }

// Slices are like references to arrays

// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names) // [John Paul George Ringo]

// 	a := names[0:2]
// 	b := names[1:3]
// 	fmt.Println(a, b) // [John Paul] [Paul George]

// 	b[0] = "XXX"
// 	fmt.Println(a, b) // [John XXX] [XXX George]
// 	fmt.Println(names) // [John XXX George Ringo]
// }

// Slice literals

// func main() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q) // [2 3 5 7 11 13]

// 	r := []bool{true, false, true, true, false, true}
// 	fmt.Println(r) // [true false true true false true]

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
// }

//Slice defaults

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}

// 	s = s[1:4]
// 	fmt.Println(s) // [3 5 7]

// 	s = s[:2]
// 	fmt.Println(s) // [3 5]

// 	s = s[1:]
// 	fmt.Println(s) // [5]
// }

// Slice length and capacity

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)

// 	// Slice the slice to give it zero length.
// 	s = s[:0]
// 	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

// 	// Extend its length.
// 	s = s[:4]
// 	printSlice(s) // len=0 cap=6 []

// 	// Drop its first two values.
// 	s = s[2:]
// 	printSlice(s) // len=4 cap=6 [2 3 5 7]
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=2 cap=4 [5 7]
// }

// Nil slices

// func main() {
// 	var s []int
// 	fmt.Println(s, len(s), cap(s)) // [] 0 0
// 	if s == nil {
// 		fmt.Println("nil!") //  nil!
// 	}
// }

// Creating a slice with make

// func main() {
// 	a := make([]int, 5)
// 	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

// 	b := make([]int, 0, 5)
// 	printSlice("b", b) // b len=0 cap=5 []

// 	c := b[:2]
// 	printSlice("c", c) // c len=2 cap=5 [0 0]

// 	d := c[2:5]
// 	printSlice("d", d) // d len=3 cap=3 [0 0 0]
// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// Slices of slices

// func main() {
// 	// Create a tic-tac-toe board.
// 	board := [][]string{
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 	}

// The players take turns.
// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "O"
// 	board[0][2] = "X"

// 	for i := 0; i < len(board); i++ {
// 		fmt.Printf("%s\n", strings.Join(board[i], " "))
// 	}
// 	// X _ X
// 	// O _ X
// 	// _ _ O
// }

// Appending to a slice

// func main() {
// 	var s []int
// 	printSlice(s)

// 	// append works on nil slices.
// 	s = append(s, 0)
// 	printSlice(s)

// 	// The slice grows as needed.
// 	s = append(s, 1)
// 	printSlice(s) //len=0 cap=0 []

// 	// We can add more than one element at a time.
// 	s = append(s, 2, 3, 4)
// 	printSlice(s) //len=1 cap=1 [0]
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=2 cap=2 [0 1] // len=5 cap=6 [0 1 2 3 4]
// }

// range

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func main() {
// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }

/*
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
*/

// Range continued
// func main() {
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}
// }

/*
1
2
4
8
16
32
64
128
256
512
*/

// Maps

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
// }

// Map literals

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// func main() {
// 	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
// }

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

// func main() {
// 	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
// }

// Mutating Maps

// func main() {
// 	m := make(map[string]int)

// 	m["Answer"] = 42
// 	fmt.Println("The value:", m["Answer"])

// 	m["Answer"] = 48
// 	fmt.Println("The value:", m["Answer"])

// 	delete(m, "Answer")
// 	fmt.Println("The value:", m["Answer"])

// 	v, ok := m["Answer"]
// 	fmt.Println("The value:", v, "Present?", ok)
// }

/*
The value: 42
The value: 48
The value: 0
The value: 0 Present? false
*/

// Function values

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func main() {
// 	hypot := func(x, y float64) float64 {
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	fmt.Println(hypot(5, 12))

// 	fmt.Println(compute(hypot))
// 	fmt.Println(compute(math.Pow))
// }

/*
13
5
81
*/

// Function closures

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }

/*
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
*/
