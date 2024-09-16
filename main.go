package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
)

func main() {

	//////////////////////////////////////////////////// First Strike
	// var n int
	// fmt.Scanf("%d", &n)
	// m := make(map[string]string)

	// reader := bufio.NewReader(os.Stdin)

	// for i := 0; i < n; i++ {
	// 	// Read the entire line
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)

	// 	// Split into parts: operation, isbn, and title
	// 	parts := strings.SplitN(input, " ", 3)
	// 	q := parts[0]    // Operation (ADD or REMOVE)
	// 	isbn := parts[1] // ISBN
	// 	var title string
	// 	if q == "ADD" {
	// 		title = parts[2] // Title as a single string
	// 	}
	// 	// 		fmt.Printf("title: %v\n", title)

	// 	if q == "ADD" {
	// 		m[isbn] = title
	// 	} else if q == "REMOVE" {
	// 		delete(m, isbn)
	// 	} else {
	// 		fmt.Println("Invalid operation:", q)
	// 		return
	// 	}
	// }

	// // Create a slice of keys (ISBNs)
	// isbns := make([]string, 0, len(m))
	// for isbn := range m {
	// 	isbns = append(isbns, isbn)
	// }

	// // Sort by title, then by ISBN as a number if titles are equal
	// sort.Slice(isbns, func(i, j int) bool {
	// 	// Sort by title first
	// 	if m[isbns[i]] != m[isbns[j]] {
	// 		return m[isbns[i]] < m[isbns[j]]
	// 	}

	// 	// If titles are the same, sort by ISBN as a number
	// 	numI, _ := strconv.Atoi(isbns[i])
	// 	numJ, _ := strconv.Atoi(isbns[j])
	// 	return numI < numJ
	// })

	// // Print the sorted ISBNs
	// for _, isbn := range isbns {
	// 	fmt.Println(isbn)
	// }

	//////////////////////////////////////////////////// Second Strike

	// res := introduce("ERfan", "20")
	// fmt.Printf("%s\n", res)

	// num1 := 5
	// num2 := 10

	// num1, _ = swap(num1, num2)

	// fmt.Printf("num1 : %d, num2 : %d\n", num1, num2)

	// name, age := getPersonInfo()
	// fmt.Printf("name : %s, age : %d\n", name, age)

	// fmt.Printf("%d\n", sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// //function as a vaiable
	// printName := func(name string) string {
	// 	return fmt.Sprintf("Your Name : %s\n", name)
	// }

	// fmt.Println(printName("Yuri Gagarin"))

	// //ananymous function
	// func() {
	// 	fmt.Println("Anonymous Function")
	// }()
	// func(a, b int) {
	// 	fmt.Println(a + b)
	// }(4, 5)
	// //ananymous function as a variable
	// subtraction := func(a, b int) {
	// 	defer fmt.Println("Forth")
	// 	defer fmt.Println("Third")
	// 	defer fmt.Println("Second")
	// 	fmt.Println("First")
	// 	fmt.Println(a - b)
	// }
	// subtraction(6, 8)

	// //higher order function
	// fmt.Println(applyFunction(square, 5))
	// fmt.Println(applyFunction(cube, 5))

	// multiply := getMathFunction("multiply")
	// fmt.Println(multiply(5, 25))

	//////////////////////////////////////////////////// Third Strike
	var a int = 24
	var p *int = &a
	fmt.Println(p)  // Referencing
	fmt.Println(*p) //  Dereferencing

	*p = 27 // Changing the value of where the pointer points at
	fmt.Println(exampleFunction(p))

	//new variable
	p_new := new(string)
	*p_new = "example"
	fmt.Println(*p_new)
	fmt.Println(p_new)

	// Call by Value
	b := 250
	ChangeVariable(b)
	fmt.Printf("b : %v\n", b)

	// Call by Reference
	c := 270
	ChangeVariableByAddress(&c)
	fmt.Printf("c : %v\n", c)

	var o *[]int
	if o == nil {
		fmt.Println("Gorgali")
	}
}

// func applyFunction(f func(int) int, num int) int {
// 	return f(num)
// }
// func square(n int) int {
// 	return n * n
// }
// func cube(n int) int {
// 	return n * n * n
// }

// func getMathFunction(operation string) func(int, int) int {
// 	switch operation {
// 	case "multiply":
// 		return func(a, b int) int {
// 			return a * b
// 		}
// 	default:
// 		return func(a, b int) int {
// 			return a + b
// 		}
// 	}
// }

// func introduce(name string, age string) string {
// 	return "My Name is " + name + " and I'm " + age + " years old"
// }

// func swap(num1 int, num2 int) (int, int) {
// 	return num2, num1
// }

// // naked return
// func getPersonInfo() (name string, age int) {
// 	name = "Alice"
// 	age = 30
// 	return
// }

// // variadic functions
// func sum(numbers ...int) int {
// 	total := 0

// 	for _, num := range numbers {
// 		total += num
// 	}

// 	return total
// }

// func IsPalindrome(x int) bool {
// 	b := x
// 	var r int
// 	v := 0
// 	for x != 0 {
// 		r = x % 10
// 		v = v*10 + r
// 		x /= 10
// 	}
// 	return v == b
// }

func exampleFunction(p *int) int {
	return *p
}

func ChangeVariable(myVariable int) {
	myVariable += 1
}

func ChangeVariableByAddress(address *int) {
	*address += 1
}
