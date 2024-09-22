package main

// "bufio"
import (
	"fmt"
)

// "os"
// "sort"
// "strconv"
// "strings"

// type Human struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// type Cat struct {
// 	Color string
// }

// func (c Cat) ChangeColorByValue(color string) {
// 	c.Color = color
// }

// func (c *Cat) ChangeColorByReference(color string) {
// 	c.Color = color
// }

// // interface
// type WashingMachine interface {
// 	// Cleaning(dc []dirtyClothes) []cleanClothes
// 	Cleaning()
// 	Drying()
// }

// // Empty Interface
// var emptyInterface interface{}

// // The Definition of types
// type Bosch struct{}
// type GPlus struct{}

// // The below funcs implement WashingMachine interface
// func (machine Bosch) Cleaning() {
// 	fmt.Println("Bosch Cleaning")
// }

// func (machine Bosch) Drying() {
// 	fmt.Println("Bosch Drying")
// }

// func (machine GPlus) Cleaning() {
// 	fmt.Println("GPlus Cleaning")
// }

// func (machine GPlus) Drying() {
// 	fmt.Println("GPlus Drying")
// }

// func CleanAndDry(machine WashingMachine) {
// 	machine.Cleaning()
// 	machine.Drying()
// }

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
	// var a int = 24
	// var p *int = &a
	// fmt.Println(p)  // Referencing
	// fmt.Println(*p) //  Dereferencing

	// *p = 27 // Changing the value of where the pointer points at
	// fmt.Println(exampleFunction(p))

	// //new variable
	// p_new := new(string)
	// *p_new = "example"
	// fmt.Println(*p_new)
	// fmt.Println(p_new)

	// // Call by Value
	// b := 250
	// ChangeVariable(b)
	// fmt.Printf("b : %v\n", b)

	// // Call by Reference
	// c := 270
	// ChangeVariableByAddress(&c)
	// fmt.Printf("c : %v\n", c)

	// println(*intPointer())
	// println(intValue())

	//////////////////////////////////////////////////// Forth Strike
	// type Student struct {
	// 	FirstName, LastName string
	// 	Age                 int
	// 	Major               string
	// }

	// // without initialization
	// var s1 Student
	// fmt.Printf("%+v\n", s1)

	// // initialize all fields
	// var s2 Student = Student{FirstName: "Olivia", LastName: "Martin", Age: 18, Major: "Electrical Engineering"}
	// fmt.Printf("%+v\n", s2)

	// // initialize some fields
	// var s3 = Student{FirstName: "Jack", Age: 24}
	// fmt.Printf("%+v\n", s3)

	// // definition with ':=' can possible with quantification
	// s4 := Student{FirstName: "James", Major: "Industrial Engineering"}
	// fmt.Printf("%+v\n", s4)

	// // definition without mentioning feilds
	// var s5 = Student{"Sophia", "Smith", 19, "Computer Engineering"}
	// fmt.Printf("%+v\n", s5)

	// // Anonymous Struct (Struct Without Name)
	// var tmp = struct {
	// 	ProductName string
	// 	Price       int
	// }{Price: 40000, ProductName: "Chips"}

	// fmt.Printf("Anonymous Struct: %+v\n", tmp)

	// s := Student{
	// 	Age:   22,
	// 	Major: "Civil Engineering",
	// }

	// fmt.Printf("Major : %s\n", s.Major)
	// fmt.Printf("Age : %d\n", s.Age)

	// type Person struct {
	// 	Name      string
	// 	Age       int
	// 	Favorites []string
	// }

	// // Slice of Structs
	// people := make([]Person, 0)

	// people = append(people, Person{Name: "Emily", Age: 20, Favorites: []string{"running", "watch TV"}})

	// p2 := Person{
	// 	Name:      "Joe",
	// 	Age:       30,
	// 	Favorites: []string{"chess"},
	// }

	// people = append(people, p2)

	// for i, v := range people {
	// 	fmt.Printf("%d. Name : %s Age : %d Favorites : ", i+1, v.Name, v.Age)

	// 	for j, f := range v.Favorites {
	// 		fmt.Printf("%d. %s ", j+1, f)
	// 	}

	// 	fmt.Printf("\n")

	// }

	// // Building a struct of other structs =(Composition)=
	// // When We have "has a" relationship between structs
	// type CPU struct {
	// 	Brand string
	// 	Core  int
	// 	Speed float32
	// }
	// type Memory struct {
	// 	Capacity       int
	// 	FrequencyInMHz int
	// }
	// type Storage struct {
	// 	Capacity int
	// 	Type     string
	// }
	// type Computer struct {
	// 	cpu     CPU
	// 	memory  Memory
	// 	storage Storage
	// }

	// cpu := CPU{Brand: "Intel", Core: 8, Speed: 4.5}
	// memory := Memory{Capacity: 16, FrequencyInMHz: 3200}
	// storage := Storage{Capacity: 500, Type: "SSD"}
	// // Computer has a cpu
	// comp := Computer{cpu, memory, storage}

	// fmt.Printf("%+v\n", comp)
	// fmt.Println("Core : ", comp.cpu.Core)
	// fmt.Println("Memory Capacity : ", comp.memory.Capacity)

	// // Building a struct of other structs =(struct embedding)=
	// // When We have "is a" relationship between structs
	// type Product struct {
	// 	Name  string
	// 	Price int
	// }
	// type Electrical struct {
	// 	Product
	// 	WarrantyMonths int
	// }
	// type Clothing struct {
	// 	Product
	// 	Size     string
	// 	Material string
	// }

	// // Clothing is a Product
	// cl := Clothing{Product{"Shirt", 300}, "S", "Cotton"}
	// fmt.Printf("%+v\n", cl)

	// // Electrical is a Product
	// el := Electrical{Product: Product{Name: "Lamp", Price: 40}, WarrantyMonths: 12}
	// fmt.Printf("%+v\n", el)

	// // Pointer to struct

	// p_human := Human{FirstName: "John", LastName: "Doe", Age: 40}
	// ptr := &p_human
	// fmt.Println(ptr.FirstName)
	// ptr.FirstName = "Bob"
	// fmt.Println(p_human.FirstName)
	// fmt.Println(p_human.Age)
	// increamentAge(&p_human)
	// fmt.Println(p_human.Age)

	//////////////////////////////////////////////////// Fifth Strike
	// persianCat := Cat{Color: "white"}
	// fmt.Println(persianCat.Color)

	// // No Change Because copies the color string
	// persianCat.ChangeColorByValue("black")
	// fmt.Println(persianCat.Color)
	// // Changes the color
	// persianCat.ChangeColorByReference("black")
	// fmt.Println(persianCat.Color)

	// b := Bosch{}
	// gp := GPlus{}

	// CleanAndDry(b)
	// CleanAndDry(gp)

	/////// Concurrency
	// people := []string{"Rose", "Erich", "Amelia"}
	// for _, person := range people {
	// 	go sayHello(person)
	// }
	// time.Sleep(1 * time.Second)

	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("Goroutine:", i)
	// 		time.Sleep(time.Millisecond * 300)
	// 	}
	// }()
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Main:", i)
	// 	time.Sleep(time.Millisecond * 300)
	// }
	// time.Sleep(1 * time.Second)

	c := make(chan int, 2)
	go func() {
		c <- 20
	}()
	c <- 40
	fmt.Println(<-c)
	second := <-c
	fmt.Println(second)

	//deadlock occurance
	c1 := make(chan string, 2)
	c1 <- "hi"
	c1 <- "hey"
	c1 <- "hello"
	value := <-c1
	fmt.Println(value)

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

// func exampleFunction(p *int) int {
// 	return *p
// }

// func ChangeVariable(myVariable int) {
// 	myVariable += 1
// }

// func ChangeVariableByAddress(address *int) {
// 	*address += 1
// }

// func intPointer() *int {
// 	local1 := 1
// 	return &local1
// }

// func intValue() int {
// 	local2 := 2
// 	return local2
// }

// func increamentAge(p *Human) {
// 	p.Age++
// }

// func sayHello(name string) {
// 	fmt.Println("Hello " + name)
// }
