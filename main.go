package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	// kullu := person{
	// 	firstName: "kuldeep",
	// 	lastName:  "goyal",
	// 	contact: contactInfo{
	// 		email:   "kgkuldeepgoyal66@gmail.com",
	// 		zipCode: 400065,
	// 	},
	// }
	//fmt.Println(kullu)
	//kullu.firstName = "kulluji"
	//fmt.Println(kullu)
	// kullupointer := &kullu
	// kullupointer.updateName("kullu")
	// kullu.print()

	// colors := map[string]string{
	// 	"red":   "ff0000",
	// 	"green": "4bf745",
	// 	"white": "ffffff",
	// }
	// printMap(colors)

	// toys := make(map[string]string)

	// toys["car"] = "Ferrari"
	// fmt.Println(toys)
	// delete(toys, "car")
	// fmt.Println(toys)

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

// func printMap(c map[string]string) {
// 	for color, hex := range c {
// 		fmt.Println(color, hex)
// 	}
// }

// func (pointertoperson *person) updateName(newName string) {
// 	(*pointertoperson).firstName = newName
// }
// func (p person) print() {
// 	fmt.Println(p)
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
