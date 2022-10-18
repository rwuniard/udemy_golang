package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Function syntax
// func (r receiver) identifier(parameters) (return(s)) { ... }
// This receiver will allow any data with type secretAgent to access speak function
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- the person")
}

type human interface {
	// So any type that is associated with speak method, will become a human type as well.
	// So secretAgent and person are also human, since they both have speak().
	// This is similar to polymorphism in Java, C# and C++.
	// except it java it uses class and interface. The interface defines the signature of the method, then
	// it is implemented in class for various different type, such as person and secretAgent in this case.
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
	// Here where you call speak(), and it will be based on the type, whether that's a person or secretAgent.
	// It will call the appropriate implementation of speak
	h.speak()

	// Showing assertion. In the switch statement you can check the type.
	// Then in the println, h.(person) -> assertion.
	switch h.(type) {
	case person:
		fmt.Println("I am person", h.(person).first)
	case secretAgent:
		fmt.Println("I am a secret agent with license to kill:", h.(secretAgent).ltk)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1)

	p1 := person{
		first: "Dr",
		last:  "Yes",
	}

	// This is an example of using the receiver
	// This is attaching a method to a struct basically.
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
